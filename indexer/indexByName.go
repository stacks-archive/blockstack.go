package indexer

import (
	"log"
	"time"
)

// StartByNames kicks off an indexing by name
// Run in a goroutine
func (i *Indexer) StartByNames() {

	// Kick off the client updater
	go i.Config.runClientUpdater()

	// Get the expected number of names in all namespaces
	log.Println(logPrefix, "Fetching expected number of names...")
	i.setExpectedNames()

	// Increment the WaitGroup with number of workers
	i.Add(i.Config.DomainChanWorkers)

	// Fetch the full list of names
	log.Println(logPrefix, "Resolving all names...")
	go i.getAllNames()

	// Wait for all the names to be resolved
	i.Wait()

	// Print them out when done
	for _, d := range i.Domains {
		log.Println(d.JSON())
	}
}

// GetAllNames retrieves all the names from all namespaces
func (i *Indexer) getAllNames() {
	// Kick off i.domainChanWorkers worker channels
	for iter := 0; iter < i.Config.DomainChanWorkers; iter++ {
		go i.handleDomainChan()
	}

	// First fetch the namespaces
	ns, err := i.client().GetAllNamespaces()
	if err != nil {
		log.Fatal(err)
	}

	// Set the current block
	i.CurrentBlock = ns.Lastblock
	for _, n := range ns.Namespaces {
		go i.getAllNamesInNamespace(n)
	}
}

// getAllNamesInNamespace gets all the names in a namespace
func (i *Indexer) getAllNamesInNamespace(ns string) {
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
	i.domainChan <- domains
	<-sem
}

// handleDomainChan handles the names coming back from blockstack core
// then it fetches the zonefiles and appends them to i.Domains
func (i *Indexer) handleDomainChan() {
	for n := range i.domainChan {
		zfhs := Domains(n).getZonefiles()
		res, err := i.client().GetZonefiles(zfhs)
		if err != nil {
			panic(err)
		}
		i.stats.zonefilesFetched.Add(float64(len(zfhs)))
		zfs := res.Decode()
		// i.Lock()
		for _, dom := range n {
			dom.AddZonefile(zfs[n[0].getNameAtRes.Records[0].ValueHash])
			i.resolveChan <- dom
			i.stats.Lock()
			i.stats.sentDownResolveChan.Inc()
			i.stats.Unlock()
			// i.Domains = append(i.Domains, dom)
		}
		// i.Unlock()
		if len(i.Domains) == i.ExpectedNames {
			close(i.domainChan)
		}
	}
	i.Done()
}

// resolveDomains takes []*Domains and resolves the URI records
func (i *Indexer) resolveDomains() {
	sem := make(chan struct{}, i.Config.ConcurrentPageFetch)
	t0 := time.Time{}
	for _, domain := range i.Domains {
		if domain.lastResolved == t0 || time.Now().Sub(domain.lastResolved) > (resolveTimeout*time.Minute) {
			sem <- struct{}{}
			go domain.ResolveProfile(sem)
			i.stats.namesResolved.Add(1)
		}
	}
}
