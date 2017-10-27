package indexer

import (
	"log"
)

// startByNames retrieves all the names from all namespaces
func (i *Indexer) startByNames() {
	i.startWorkers()

	ns, err := i.client().GetAllNamespaces()
	if err != nil {
		// TODO: Better error handling here
		panic(err)
	}

	go i.setCB(ns.Lastblock)
	for _, n := range ns.Namespaces {
		go i.getAllNamePagesInNamespace(n)
	}
}

// Starts all the worker routines
func (i *Indexer) startWorkers() {
	go i.startNamePageWorkers()
	go i.startResolveWorkers()
	go i.startDBWorkers()
}

// startNamePageWorkers kicks off i.Config.NamePageWorkers workers
// to handle the GetNamesInNamespace returns and Zonefile fetching
func (i *Indexer) startNamePageWorkers() {
	for iter := 0; iter < i.Config.NamePageWorkers; iter++ {
		go i.handleNamePageChan()
	}
}

// startResolveWorkers kicks off i.Config.ResolveWorkers workers
// to handle the *Domains that have zonefiles
func (i *Indexer) startResolveWorkers() {
	for iter := 0; iter < i.Config.ResolveWorkers; iter++ {
		go i.handleResolveChan()
	}
}

// startDBWorkers kicks off i.Config.DBWorkers workers
// to handle batching and insertion/update of the database
func (i *Indexer) startDBWorkers() {
	for iter := 0; iter < i.Config.DBWorkers; iter++ {
		go i.handleDBChan()
	}
}

// getAllNamePagesInNamespace gets all the NamePages in a namespace
func (i *Indexer) getAllNamePagesInNamespace(ns string) {
	numNames, err := i.client().GetNumNamesInNamespace(ns)
	if err != nil {
		// TODO: Better error handling here
		panic(err)
	}

	iter := (numNames.Count/100 + 1)
	sem := make(chan struct{}, i.Config.ConcurrentPageFetch)
	for page := 0; page <= iter; page++ {
		sem <- struct{}{}
		go i.getNamePageAsync(page, iter, ns, sem)
	}
}

// A goroutine safe method for fetching the list of names from blockstack-core
func (i *Indexer) getNamePageAsync(page int, iter int, ns string, sem chan struct{}) {
	namePage, err := i.client().GetNamesInNamespace(ns, page*100, 100)
	if err != nil {
		// TODO: Better error handling here
		panic(err)
	}

	go i.setCB(namePage.Lastblock)

	var domains []*Domain
	for _, name := range namePage.Names {
		dom := NewDomain(name)
		res, err := i.client().GetNameAt(name, i.CurrentBlock)
		if err != nil {
			// TODO: Better error handling here
			log.Println("Error fetching name details", err)
		}
		dom.getNameAtRes = res
		domains = append(domains, dom)
		i.stats.nameDetailsFetched.Inc()
	}
	i.stats.namePagesFetched.Inc()
	i.namePageChan <- domains
	<-sem
}

// handleNamePageChan handles namePages coming back from blockstack core
// It fectches zonfiles and adds them to the *Domains, sending them for resolution
func (i *Indexer) handleNamePageChan() {
	for doms := range i.namePageChan {

		// Get zonefileHashes from Domains and get zonefiles
		res, err := i.client().GetZonefiles(doms.getZonefileHashes())
		if err != nil {
			// TODO: Better error handling here
			log.Fatal(logPrefix, err)
		}

		go i.setCB(res.Lastblock)
		i.stats.zonefilesFetched.Add(float64(len(res.Zonefiles)))

		// TODO: Double check behavior here. Make sure this is doing what you think it is
		zonefiles := res.Decode()
		for _, dom := range doms {
			if zonefileHash := dom.zonefileHash(); zonefileHash != "" {
				dom.AddZonefile(zonefiles[zonefileHash])
			}
			i.resolveChan <- dom
			i.stats.sentDownResolveChan.Inc()
		}

		// TODO: find a way to close this chan so that the waitgroup finishes
	}
	// TODO: Double check these WaitGroups
	// Once the for loop exits decrement the WaitGroup
	i.namePageWait.Done()
}

// handleResolveChan handles *Domain after they have zonefiles
func (i *Indexer) handleResolveChan() {
	for d := range i.resolveChan {
		if d.Profile != nil {
			d.ResolveProfile()
		}
		i.dbChan <- d
		i.stats.namesResolved.Inc()
	}
}

// handleDBChan batches *Domain for insert/update of the MongoDB instance
func (i *Indexer) handleDBChan() {
	var doms Domains
	for d := range i.dbChan {
		if len(doms) >= i.Config.DBBatchSize {
			// TODO: replace with actual DB calls
			log.Println(logPrefix, "Sending", len(doms), "domains to the database...")
			i.stats.writtenToDatabase.Inc()
			doms = Domains{}
		}
		doms = append(doms, d)
	}
}
