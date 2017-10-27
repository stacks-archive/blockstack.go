package indexer

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/blockstack/go-blockstack/blockstack"
)

// Config is a config object for the Indexer
type Config struct {
	URLs                 []string
	IndexMethod          string
	NamePageWorkers      int
	ResolveWorkers       int
	ConcurrentPageFetch  int
	ClientUpdateInterval int
	DBBatchSize          int
	DBWorkers            int

	clients       []*blockstack.Client
	currentClient int

	sync.Mutex
}

func (c *Config) String() string {
	return fmt.Sprintf(`Configuration Settings:
  Number of Clients:            %v
  Number of Name Page Workers:  %v
  Number of Resolve Workers:    %v
  Name Page Concurrency:        %v
  Client Update Interval:       %v
  Database Batch Size:          %v
  Database Insert Workers:      %v`,
		len(c.URLs),
		c.NamePageWorkers,
		c.ResolveWorkers,
		c.ConcurrentPageFetch,
		c.ClientUpdateInterval,
		c.DBBatchSize,
		c.DBWorkers,
	)
}

// SetClients takes the configured URLs and returns only
// the blockstack-core nodes that are in consensus
func (c *Config) SetClients() {
	clients, errs := blockstack.ValidClients(c.URLs)
	for _, err := range errs {
		er := err.(blockstack.ClientRegistrationError)
		if err != nil {
			log.Println(logPrefix, er.URL, er.Err)
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
		time.Sleep(time.Duration(c.ClientUpdateInterval) * time.Minute)
		log.Println(logPrefix, "Updating blockstack-core clients...")
		c.SetClients()
	}
}
