package indexer

import (
	"log"
	"sync"
	"time"

	"github.com/blockstack/go-blockstack/blockstack"
)

// startBlock is the first block on the main bitcoin network
const (
	logPrefix      = "[indexer]"
	resolveTimeout = 30

	concurrentPageFetch = 1000
	domainChanWorkers   = 100
)

// StartIndexer returns a Indexer
func StartIndexer(clients []*blockstack.Client) {
	i := &Indexer{
		StartBlock: blockstack.StartBlock,
		Domains:    make([]*Domain, 0),
		// NameZonefileHash: make(map[string]string),

		currentClient: 0,
		clients:       clients,
		domainChan:    make(chan []*Domain),
		stats:         newIndexerStats(),
	}

	// Get the expected number of names in all namespaces
	go i.setExpectedNames()

	i.Add(i.domainChanWorkers)
	// Fetch the full list of names
	go i.getAllNames()

	i.Wait()

	for _, d := range i.Domains {
		log.Println(d.JSON())
	}
	// Resolve Domains
	// log.Printf("%v Resolv	ing domains", logPrefix)
	// i.resolveDomains()
}

// The Indexer talks to blockstack-core and resolves all
// the domains and subdomains - hopefully
type Indexer struct {
	StartBlock    int
	CurrentBlock  int
	Domains       []*Domain
	ExpectedNames int

	// stats holds the prometheus statss
	stats *indexerStats

	// clients is an array of blockstack-core nodes
	clients []*blockstack.Client

	// currentClient tracks the last used client out of a list
	currentClient int

	// nameChan handles the names coming back from fetching the full list of names
	domainChan chan []*Domain

	// names holds the list of names from the network
	names []string

	concurrentPageFetch int
	domainChanWorkers   int

	sync.Mutex
	sync.WaitGroup
}

// GetAllNames retrieves all the names from all namespaces
func (i *Indexer) getAllNames() {
	// Kick off i.domainChanWorkers worker channels
	for iter := 0; iter < i.domainChanWorkers; iter++ {
		go i.handleDomainChan()
	}
	ns, err := i.Client().GetAllNamespaces()
	if err != nil {
		log.Fatal(err)
	}
	i.CurrentBlock = ns.Lastblock
	for _, n := range ns.Namespaces {
		go i.getAllNamesInNamespace(n)
	}
}

// getAllNamesInNamespace gets all the names in a namespace
func (i *Indexer) getAllNamesInNamespace(ns string) {
	numNames, err := i.Client().GetNumNamesInNamespace(ns)
	if err != nil {
		log.Fatal(err)
	}
	iter := (numNames.Count/100 + 1)
	sem := make(chan struct{}, i.concurrentPageFetch)
	for page := 0; page <= iter; page++ {
		sem <- struct{}{}
		go i.getNamePageAsync(page, iter, ns, sem)
	}
}

// A goroutine safe method for fetching the list of names from blockstack-core
func (i *Indexer) getNamePageAsync(page int, iter int, ns string, sem chan struct{}) {
	namePage, err := i.Client().GetNamesInNamespace(ns, page*100, 100)
	if err != nil {
		log.Fatal(err)
	}
	var domains []*Domain
	for _, name := range namePage.Names {
		// Fetch the name details here as well
		dom := NewDomain(name)
		res, err := i.Client().GetNameAt(name, i.CurrentBlock)
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
		res, err := i.Client().GetZonefiles(zfhs)
		if err != nil {
			panic(err)
		}
		i.stats.zonefilesFetched.Add(float64(len(zfhs)))
		zfs := res.Decode()
		i.Lock()
		for _, dom := range n {
			dom.AddZonefile(zfs[n[0].getNameAtRes.Records[0].ValueHash])
			i.Domains = append(i.Domains, dom)
		}
		i.Unlock()
		if len(i.names) == i.ExpectedNames {
			close(i.domainChan)
		}
	}
	i.Done()
}

// resolveDomains takes []*Domains and resolves the URI records
func (i *Indexer) resolveDomains() {
	sem := make(chan struct{}, i.concurrentPageFetch)
	t0 := time.Time{}
	for _, domain := range i.Domains {
		if domain.lastResolved == t0 || time.Now().Sub(domain.lastResolved) > (resolveTimeout*time.Minute) {
			sem <- struct{}{}
			go domain.ResolveProfile(sem)
			i.stats.namesResolved.Add(1)
		}
	}
}

// Client provides a convinent interface to loop through provided multiple clients
func (i *Indexer) Client() *blockstack.Client {
	var client *blockstack.Client
	i.Lock()
	if len(i.clients) == 1 {
		return i.clients[0]
	}
	client = i.clients[i.currentClient]
	i.currentClient++
	if len(i.clients) == i.currentClient {
		i.currentClient = 0
	}
	i.stats.callsMade.Add(1)
	i.Unlock()
	return client
}

// Gets the expected number of names from blockstack-core
func (i *Indexer) setExpectedNames() {
	res, err := i.Client().GetAllNamespaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, ns := range res.Namespaces {
		res, err := i.Client().GetNumNamesInNamespace(ns)
		if err != nil {
			log.Fatal(err)
		}
		i.ExpectedNames += res.Count
	}
	log.Println(logPrefix, "fetching all", i.ExpectedNames, "from the blockstack network")
	i.stats.namesOnNetwork.Set(float64(i.ExpectedNames))
}

// processNameZonefileMap takes a map[name]zonefile and returns the go representation
// TODO: Check this
// func (i *Indexer) processNameZonefileMap() {
// 	for name := range i.NameZonefileHash {
// 		i.Domains = append(i.Domains, NewDomain(name, i.NameZonefileHash[name]))
// 	}
// }

// fetchAllZonefiles fetches all the zonefiles from the startBlock to the CurrentBlock
// func (i *Indexer) fetchAllZonefiles() {
// 	numBlocks := i.CurrentBlock - i.StartBlock
// 	fetchBlocks := 100
// 	iter := (numBlocks/fetchBlocks + 1)
// 	for page := 0; page <= iter; page++ {
// 		st := i.StartBlock + (page * fetchBlocks)
// 		end := st + 100
// 		log.Printf("%v Fetching zonefiles from block %v to block %v", logPrefix, st, end)
// 		i.GetZonefilesByBlock(st, end)
// 	}
// }
//
// // GetZonefilesByBlock returns a map[name]zonefile
// // TODO: Parallelize - difficult, need to keep blocks in order. Maybe hold off on this
// func (i *Indexer) GetZonefilesByBlock(startBlock, endBlock int) {
// 	zfhrs := make([]blockstack.ZonefileHashResults, 0)
// 	iter := 0
// 	for iter < 100000 {
// 		// Fetch batch of Zonefiles by block
// 		res, err := i.Client().GetZonefilesByBlock(startBlock, endBlock, (iter * 100), 100)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		iter++
//
// 		// Make a batch that maps to a get_zonefiles rpc call
// 		batch := make([]blockstack.ZonefileHashResult, 0)
// 		for _, zfhrs := range res.ZonefileInfo {
// 			batch = append(batch, zfhrs)
// 		}
//
// 		// Save the batch
// 		zfhrs = append(zfhrs, blockstack.ZonefileHashResults(batch))
//
// 		// If the batch doesn't have 100 records then stop
// 		if len(res.ZonefileInfo) != 100 {
// 			iter = 100000
// 		}
// 	}
//
// 	// Loop over the batches and call get_zonefile for each
// 	for _, batch := range zfhrs {
// 		res, err := i.Client().GetZonefiles(batch.Zonefiles())
// 		if err != nil {
// 			log.Fatal(err)
// 		}
//
// 		// Decode the base64 encoded zonefiles
// 		dec := res.Decode()
//
// 		// Range over the decoded zonefiles and associate them with names
// 		for zfh := range dec {
// 			l := batch.LatestZonefileHash(zfh)
// 			i.NameZonefileHash[l.Name] = dec[zfh]
// 		}
// 	}
// 	i.stats.LastProcessed.Set(float64(endBlock))
// }
