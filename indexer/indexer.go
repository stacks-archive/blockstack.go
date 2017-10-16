package indexer

import (
	"log"
	"sync"
	"time"

	"github.com/jackzampolin/go-blockstack/blockstack"
)

// startBlock is the first block on the main bitcoin network
const (
	startBlock = blockstack.StartBlock
	logPrefix  = "[indexer]"
)

// StartIndexer returns a Indexer
func StartIndexer(confs []blockstack.ServerConfig) {
	i := &Indexer{
		StartBlock:       startBlock,
		Names:            make([]*Domain, 0),
		NameZonefileHash: make(map[string]string),

		currentClient: 0,
		clients:       make([]*blockstack.Client, 0),
		nameChan:      make(chan string),
		stats:         newIndexerStats(),
	}

	// Register Clients
	for _, conf := range confs {
		i.registerClient(conf)
		log.Printf("%s Added core node %s:%v to rotation", logPrefix, conf.Address, conf.Port)
	}
	// Kick off metrics goroutine
	go i.setMetrics()

	// Fetch the full list of names
	log.Printf("%v Fetching full list of domains...", logPrefix)
	i.getAllNames()

	// Fetch Zonefiles
	log.Printf("%v Indexing Blockstack Network from %v to %v", logPrefix, startBlock, i.CurrentBlock)
	i.fetchAllZonefiles()

	// Create Domains
	log.Printf("%v Create []*Domain from zonefiles", logPrefix)
	i.processNameZonefileMap()

	// Resolve Domains
	log.Printf("%v Resolving domains", logPrefix)
	i.resolveDomains()
}

// The Indexer talks to blockstack-core and resolves all
// the domains and subdomains - hopefully
type Indexer struct {
	StartBlock       int
	CurrentBlock     int
	Names            []*Domain
	ExpectedNames    int
	NameZonefileHash map[string]string

	// stats holds the prometheus statss
	stats *indexerStats

	// clients is an array of blockstack-core nodes
	clients []*blockstack.Client

	// currentClient tracks the last used client out of a list
	currentClient int

	// nameChan handles the names coming back from fetching the full list of names
	nameChan chan string

	sync.Mutex
}

func (i *Indexer) setMetrics() {
	for {
		i.Lock()
		i.stats.NumZonefiles.Set(float64(len(i.NameZonefileHash)))
		i.stats.NumDomains.Set(float64(len(i.Names)))
		i.stats.NumResolved.Set(float64(i.resolved()))
		i.Unlock()
		time.Sleep(2 * time.Second)
	}
}

// Resolved returns the number of resolved domains
func (i *Indexer) resolved() int {
	var out int
	for _, dom := range i.Names {
		if dom.resolved {
			out++
		}
	}
	return out
}

// resolveDomains takes []*Domains and resolves the URI records
func (i *Indexer) resolveDomains() {
	numCalls := 200
	sem := make(chan struct{}, numCalls)
	for _, domain := range i.Names {
		sem <- struct{}{}
		go domain.ResolveProfile(sem)
	}
}

// processNameZonefileMap takes a map[name]zonefile and returns the go representation
// TODO: Check this
func (i *Indexer) processNameZonefileMap() {
	for name := range i.NameZonefileHash {
		i.Lock()
		zfh := i.NameZonefileHash[name]
		i.Unlock()
		if zfh != "" {
			i.Names = append(i.Names, NewDomain(name, zfh))
		}
	}
}

// fetchAllZonefiles fetches all the zonefiles from the startBlock to the CurrentBlock
func (i *Indexer) fetchAllZonefiles() {
	numBlocks := i.CurrentBlock - startBlock
	fetchBlocks := 100
	iter := (numBlocks/fetchBlocks + 1)
	for page := 0; page <= iter; page++ {
		st := startBlock + (page * fetchBlocks)
		end := st + 100
		log.Printf("%v Fetching zonefiles from block %v to block %v", logPrefix, st, end)
		i.GetZonefilesByBlock(st, end)
		i.stats.ProcessedTo.Set(float64(end))
	}
}

// GetZonefilesByBlock returns a map[name]zonefile
// TODO: Parallelize - difficult, need to keep blocks in order. Maybe hold off on this
func (i *Indexer) GetZonefilesByBlock(startBlock, endBlock int) {
	zfhrs := make([]blockstack.ZonefileHashResults, 0)

	iter := 0
	for iter < 100000 {
		// Fetch batch of Zonefiles by block
		res, err := i.Client().GetZonefilesByBlock(startBlock, endBlock, (iter * 100), 100)
		if err != nil {
			log.Fatal(err)
		}
		iter++

		// Make a batch that maps to a get_zonefiles rpc call
		batch := make([]blockstack.ZonefileHashResult, 0)
		for _, zfhrs := range res.ZonefileInfo {
			batch = append(batch, zfhrs)
		}

		// Save the batch
		zfhrs = append(zfhrs, blockstack.ZonefileHashResults(batch))

		// If the batch doesn't have 100 records then stop
		if len(res.ZonefileInfo) != 100 {
			iter = 100000
		}
	}

	// Loop over the batches and call get_zonefile for each
	for _, batch := range zfhrs {
		res, err := i.Client().GetZonefiles(batch.Zonefiles())
		if err != nil {
			log.Fatal(err)
		}

		// Decode the base64 encoded zonefiles
		dec := res.Decode()

		// Range over the decoded zonefiles and associate them with names
		for zfh := range dec {
			l := batch.LatestZonefileHash(zfh)
			i.NameZonefileHash[l.Name] = dec[zfh]
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
	i.Unlock()
	return client
}

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
		i.Lock()
		i.ExpectedNames += res.Count
		i.Unlock()
	}
}

// GetAllNames retrieves all the names from all namespaces
func (i *Indexer) getAllNames() {
	ns, err := i.Client().GetAllNamespaces()
	if err != nil {
		log.Fatal(err)
	}
	i.CurrentBlock = ns.Lastblock
	for _, n := range ns.Namespaces {
		go i.getAllNamesInNamespace(n)
	}
}

func (i *Indexer) getNamePageAsync(page int, iter int, ns string, sem chan struct{}) {
	namePage, err := i.Client().GetNamesInNamespace(ns, page*100, 100)
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range namePage.Names {
		i.nameChan <- name
	}
	<-sem
}

func (i *Indexer) handleNameChan() {
	for n := range i.nameChan {
		i.Lock()
		i.NameZonefileHash[n] = ""
		i.Unlock()
	}
}

// getAllNamesInNamespace gets all the names in a namespace
func (i *Indexer) getAllNamesInNamespace(ns string) {
	numNames, err := i.Client().GetNumNamesInNamespace(ns)
	if err != nil {
		log.Fatal(err)
	}
	iter := (numNames.Count/100 + 1)
	numCalls := 50
	sem := make(chan struct{}, numCalls)
	for page := 0; page <= iter; page++ {
		sem <- struct{}{}
		go i.getNamePageAsync(page, iter, ns, sem)
	}
}

// registerClient takes a blockstack.ServerConfig and tries to contact that Server
// if it is successful it is added to the rotation if not it is excluded
func (i *Indexer) registerClient(conf blockstack.ServerConfig) {
	client := blockstack.NewClient(conf)
	res, err := client.GetInfo()
	if err != nil {
		log.Printf("%s Failed to contact %s:%v, excluding from rotation", logPrefix, conf.Address, conf.Port)
		return
	}
	i.clients = append(i.clients, client)
	i.CurrentBlock = res.LastBlockProcessed
}
