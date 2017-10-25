package indexer

import (
	"log"
	"sync"
	"time"

	"github.com/blockstack/go-blockstack/blockstack"
)

// Config is a config object for the Indexer
type Config struct {
	DomainChanWorkers   int
	ConcurrentPageFetch int
	URLs                []string

	clients       []*blockstack.Client
	currentClient int
	sync.Mutex
}

// NewConfig returns a new config object
func NewConfig(clients []string, domainChanW int, pageFetchConc int) *Config {
	conf := &Config{
		DomainChanWorkers:   domainChanW,
		ConcurrentPageFetch: pageFetchConc,
		URLs:                clients,
		currentClient:       0,
	}
	log.Println(logPrefix, "Setting valid clients...")
	conf.setClients()
	return conf
}

// setClients takes the configured URLs and returns only
// the blockstack-core nodes that are in consensus
func (c *Config) setClients() {
	clients, errs := blockstack.ValidClients(c.URLs)
	for _, err := range errs {
		if err != nil {
			log.Println(logPrefix, err.JSON())
		}
	}
	c.Lock()
	c.clients = clients
	c.Unlock()
}

// Run as a goroutine to continually update clients
func (c *Config) runClientUpdater() {
	log.Println(logPrefix, "Kicking off client update routine...")
	for {
		time.Sleep(5 * time.Minute)
		log.Println(logPrefix, "Updating blockstack-core clients...")
		c.setClients()
	}
}
