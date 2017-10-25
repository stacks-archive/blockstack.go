package indexer

import (
	"log"
	"sync"

	"github.com/blockstack/go-blockstack/blockstack"
)

// startBlock is the first block on the main bitcoin network
const (
	logPrefix      = "[indexer]"
	resolveTimeout = 30
)

// The Indexer talks to blockstack-core and resolves all
// the domains and subdomains - hopefully
type Indexer struct {
	CurrentBlock  int
	ExpectedNames int
	Domains       []*Domain

	// Config holds all the config vars
	Config *Config

	// stats holds the prometheus statss
	stats *indexerStats

	// nameChan handles the names coming back from fetching the full list of names
	// the workers then process them and add the zonefiles to the domains
	domainChan chan []*Domain
	domainWait sync.WaitGroup

	// once the zonefiles are added then names travel down the resolve chan
	// the workers then resolve them.
	resolveChan chan *Domain
	resolveWait sync.WaitGroup

	sync.Mutex
	sync.WaitGroup
}

// NewIndexer returns a new *Indexer
func NewIndexer(conf *Config) *Indexer {
	return &Indexer{
		Domains: make([]*Domain, 0),
		// Config:  NewConfig(clients, 100, 1000),
		Config: conf,

		domainChan:  make(chan []*Domain),
		resolveChan: make(chan *Domain),
		stats:       newIndexerStats(),
	}
}

// client provides a convinent interface to loop through provided multiple clients
func (i *Indexer) client() *blockstack.Client {
	var client *blockstack.Client

	// Lock the config object to prevent concurrent access
	i.Config.Lock()

	// If there is only one client, return it quickly
	if len(i.Config.clients) == 1 {
		client = i.Config.clients[0]
		i.Config.Unlock()
		return client
	}

	// Reset the counter if currentClient is greater than the number of clients
	i.Config.currentClient++
	if len(i.Config.clients) <= i.Config.currentClient {
		i.Config.currentClient = 0
	}

	// Return the client at index currentClient
	client = i.Config.clients[i.Config.currentClient]

	i.Config.Unlock()

	// Increment the number of calls made in a goroutine
	go func(s *indexerStats) {
		// Lock the stats object to prevent concurrent access
		i.stats.Lock()
		i.stats.callsMade.Add(1)
		i.stats.Unlock()
	}(i.stats)

	return client
}

// Gets the expected number of names from blockstack-core
func (i *Indexer) setExpectedNames() {
	res, err := i.client().GetAllNamespaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, ns := range res.Namespaces {
		res, err := i.client().GetNumNamesInNamespace(ns)
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
// 		res, err := i.client().GetZonefilesByBlock(startBlock, endBlock, (iter * 100), 100)
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
// 		res, err := i.client().GetZonefiles(batch.Zonefiles())
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
