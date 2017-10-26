package indexer

import (
	"log"
	"sync"

	"github.com/blockstack/go-blockstack/blockstack"
)

const (
	logPrefix = "[indexer]"
)

// The Indexer talks to blockstack-core and resolves all
// the domains and subdomains - hopefully
type Indexer struct {
	CurrentBlock  int
	ExpectedNames int
	Domains       []*Domain

	// Config holds all the config vars
	Config *Config

	// stats holds the prometheus stats
	stats *indexerStats

	// currentBlock holds the currenBlock number
	current *current

	// nameChan handles the names coming back from fetching the full list of names
	// the namePageWorkers then process them and add the zonefiles to the domains
	namePageChan chan []*Domain
	namePageWait sync.WaitGroup

	// once the zonefiles are added then names travel down the resolve chan
	// the workers then resolve them.
	resolveChan chan *Domain
	resolveWait sync.WaitGroup

	// once the *Domains are resolved they are sent to the database batcher
	// for insert/update of the MongoDB instance
	dbChan chan *Domain
	dbWait sync.WaitGroup

	// TODO: Remove? This might still be helpful for the []*Domains object,
	// but that might not be needed either once database is incorporated
	sync.Mutex
}

// NewIndexer returns a new *Indexer
func NewIndexer(conf *Config) *Indexer {
	return &Indexer{
		Domains: make([]*Domain, 0),
		Config:  conf,

		namePageChan: make(chan []*Domain),
		resolveChan:  make(chan *Domain),
		stats:        newIndexerStats(),
	}
}

// Start runs the Indexer
func (i *Indexer) Start() {

	// Kick off the client updater
	go i.Config.runClientUpdater()

	// Get the expected number of names in all namespaces
	log.Println(logPrefix, "Fetching expected number of names...")
	i.setExpectedNames()
	log.Println(logPrefix, i.ExpectedNames, "found on the Blockstack Network, fetching...")

	if i.Config.IndexMethod == "byName" {
		log.Println(logPrefix, "Resolving all names...")
		// TODO: Add waitGroup handling here and wait until the index is Done
		// then exit. This will allow for looping!
		go i.startByNames()
	} else {
		log.Printf("%s Invalid indexMethod '%s', byName and byBlock supported", logPrefix, i.Config.IndexMethod)
	}
}

// client provides a convinent interface to loop through multiple
// blockstack-core clients in a goroute safe manner
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

	// Increment stats in a goroutine to avoid blocking
	// Maybe add functions for this to indexerStats?
	go func(s *indexerStats) {
		i.stats.Lock()
		i.stats.callsMade.Add(1)
		i.stats.Unlock()
	}(i.stats)

	return client
}

// Gets the expected number of names from blockstack-core
func (i *Indexer) setExpectedNames() {
	// First get the list of Namespaces
	res, err := i.client().GetAllNamespaces()
	if err != nil {
		log.Fatal(err)
	}

	// Then find the number of names in each Namespace
	for _, ns := range res.Namespaces {
		res, err := i.client().GetNumNamesInNamespace(ns)
		if err != nil {
			log.Fatal(err)
		}
		i.ExpectedNames += res.Count
	}

	// This is the only time this stat is set so no need to lock
	i.stats.namesOnNetwork.Set(float64(i.ExpectedNames))
}

// setCB is a goroutine safe way to set the current block
func (i *Indexer) setCB(block int) {
	i.current.Lock()
	i.current.block = block
	i.current.Unlock()
}

// GetCB reads the current block in a goroutine safe manner
func (i *Indexer) GetCB() int {
	var block int
	i.current.Lock()
	block = i.current.block
	i.current.Unlock()
	return block
}

// current holds the value of the current block as well as a mutex to prevent contention
type current struct {
	block int

	sync.Mutex
}
