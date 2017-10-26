package indexer

import (
	"log"
	// "time"
)

// StartByNames retrieves all the names from all namespaces
func (i *Indexer) startByNames() {
	// Start workers
	go i.startNamePageWorkers()
	go i.startResolveWorkers()
	go i.startDBWorkers()

	// Fetch the list of namespaces
	ns, err := i.client().GetAllNamespaces()
	if err != nil {
		log.Fatal(err)
	}

	// Set the current block
	go i.setCB(ns.Lastblock)
	for _, n := range ns.Namespaces {
		go i.getAllNamePagesInNamespace(n)
	}
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

// getAllNamePagesInNamespace gets all the names in a namespace
func (i *Indexer) getAllNamePagesInNamespace(ns string) {
	numNames, err := i.client().GetNumNamesInNamespace(ns)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	go i.setCB(namePage.Lastblock)
	var domains []*Domain
	for _, name := range namePage.Names {
		// Fetch the name details here as well
		dom := NewDomain(name)
		res, err := i.client().GetNameAt(name, i.CurrentBlock)
		if err != nil {
			log.Println("Error fetching name details", err)
		}
		i.stats.nameDetailsFetched.Add(1)
		dom.getNameAtRes = res
		domains = append(domains, dom)
	}
	i.stats.namePagesFetched.Add(1)
	i.namePageChan <- domains
	<-sem
}

// handleNamePageChan handles the names coming back from blockstack core
// then it fetches the zonefiles and appends them to i.Domains
func (i *Indexer) handleNamePageChan() {
	// range over the namePageChan
	for doms := range i.namePageChan {

		// gather the zonefileHashes for the Domains
		// and get the zonefiles for them
		zfhs := Domains(doms).getZonefiles()
		res, err := i.client().GetZonefiles(zfhs)
		if err != nil {
			log.Fatal(logPrefix, err)
		}

		// set the current block and inc stats
		go i.setCB(res.Lastblock)
		i.stats.zonefilesFetched.Add(float64(len(zfhs)))

		// Decode the base64 encoded zonefiles into map[zonefileHash]zonefile
		zfs := res.Decode()

		// TODO: Double check behavior here. Make sure this is doing what you think it is
		for _, dom := range doms {
			// If the *Domain has a zonefile associated add the zonefile to the domain
			if len(dom.getNameAtRes.Records) > 0 && dom.getNameAtRes.Records[0].ValueHash != "" {
				dom.AddZonefile(zfs[dom.getNameAtRes.Records[0].ValueHash])
			}

			// Send the name down the resolve chan
			i.resolveChan <- dom

			// Increment stats in a goroutine to avoid blocking
			// Maybe add functions for this to indexerStats?
			go func(s *indexerStats) {
				s.Lock()
				s.sentDownResolveChan.Inc()
				s.Unlock()
			}(i.stats)
		}

		// Currently checking for it being the last domain,
		// This could leave some running routines
		// Maybe kick off a namePageChanCloser at start?
		if len(i.Domains) == i.ExpectedNames {
			close(i.namePageChan)
		}
	}

	// Once the for loop exits decrement the WaitGroup
	i.namePageWait.Done()
}

// handleResolveChan handles *Domain after they have zonefiles
func (i *Indexer) handleResolveChan() {
	for d := range i.resolveChan {
		// Resolve the profile
		d.ResolveProfile()

		// Send the domain for persistence
		i.dbChan <- d

		// Increment the number of resolved names
		go func(s *indexerStats) {
			s.Lock()
			s.namesResolved.Inc()
			s.Unlock()
		}(i.stats)
	}
}

// handleDBChan batches *Domain for insert/update of the MongoDB instance
func (i *Indexer) handleDBChan() {
	var doms []*Domain
	for d := range i.dbChan {
		// if the batch size has been reached then persist the domains and reset doms
		if len(doms) >= i.Config.DBBatchSize {

			// TODO: replace with actual DB calls
			log.Println(logPrefix, "Sending", len(doms), "domains to the database...")

			// Increment the number of names writtenToDatabase
			go func(s *indexerStats, num int) {
				s.Lock()
				s.writtenToDatabase.Inc()
				s.Unlock()
			}(i.stats, len(doms))

			// reset doms
			doms = []*Domain{}
		}

		doms = append(doms, d)
	}
}
