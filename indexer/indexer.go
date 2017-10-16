package indexer

import (
	"log"
	"sync"
	"time"

	"github.com/jackzampolin/go-blockstack/blockstack"
	"github.com/prometheus/client_golang/prometheus"
)

// StartBlock is the first block on the main bitcoin network
const (
	StartBlock = 373601
	prefix     = "[indexer]"
	mns        = "indexer"
)

// StartIndexer returns a Indexer
func StartIndexer(confs []blockstack.ServerConfig) {
	i := &Indexer{
		StartBlock:    StartBlock,
		currentClient: 0,
		clients:       make([]*blockstack.Client, 0),
		Names:         make([]*Domain, 0),
		nameChan:      make(chan *Domain),
		ProcessedTo: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: mns,
			Subsystem: "blocks",
			Name:      "processed",
			Help:      "The block number that zonefiles have been fetched to",
		}),
		NumZonefiles: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: mns,
			Subsystem: "zonefiles",
			Name:      "number",
			Help:      "Number of zonefiles in inventory",
		}),
		NumDomains: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: mns,
			Subsystem: "domains",
			Name:      "number",
			Help:      "Number of domains initialized",
		}),
		NumResolved: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: mns,
			Subsystem: "domains",
			Name:      "resolved",
			Help:      "Number of domains resolved",
		}),
		enrichChan:       make(chan *Domain),
		NameZonefileHash: make(map[string]string),
	}

	// Register Metrics with Prom handler
	prometheus.MustRegister(i.NumZonefiles)
	prometheus.MustRegister(i.ProcessedTo)
	prometheus.MustRegister(i.NumResolved)
	prometheus.MustRegister(i.NumDomains)

	// Register Clients
	for _, conf := range confs {
		i.registerClient(conf)
		log.Printf("%s Added core node %s:%v to rotation", prefix, conf.Address, conf.Port)
	}
	// Kick off metrics goroutine
	go i.setMetrics()

	// Fetch Zonefiles
	log.Printf("%v Indexing Blockstack Network from %v to %v", prefix, StartBlock, i.CurrentBlock)
	i.FetchAllZonefiles()

	// Create Domains
	log.Printf("%v Create []*Domain from zonefiles", prefix)
	i.ProcessNameZonefileMap()

	// Resolve Domains
	log.Printf("%v Resolving domains", prefix)
	i.ResolveDomains()
}

// The Indexer talks to blockstack-core and resolves all
// the domains and subdomains - hopefully
type Indexer struct {
	StartBlock    int
	CurrentBlock  int
	Names         []*Domain
	ExpectedNames int

	// Metrics for debugging
	ProcessedTo  prometheus.Gauge
	NumZonefiles prometheus.Gauge
	NumDomains   prometheus.Gauge
	NumResolved  prometheus.Gauge

	currentClient    int
	clients          []*blockstack.Client
	NameZonefileHash map[string]string

	nameChan   chan *Domain
	enrichChan chan *Domain

	sync.Mutex
}

func (i *Indexer) setMetrics() {
	for {
		i.Lock()
		i.NumZonefiles.Set(float64(len(i.NameZonefileHash)))
		i.NumDomains.Set(float64(len(i.Names)))
		i.NumResolved.Set(float64(i.resolved()))
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

// ResolveDomains takes []*Domains and resolves the URI records
func (i *Indexer) ResolveDomains() {
	numCalls := 200
	sem := make(chan struct{}, numCalls)
	for _, domain := range i.Names {
		sem <- struct{}{}
		domain.ResolveProfile(sem)
	}
}

// ProcessNameZonefileMap takes a map[name]zonefile and returns the go representation
func (i *Indexer) ProcessNameZonefileMap() {
	for name := range i.NameZonefileHash {
		i.Names = append(i.Names, NewDomain(name, i.NameZonefileHash[name]))
	}
}

// FetchAllZonefiles fetches all the zonefiles from the StartBlock to the CurrentBlock
func (i *Indexer) FetchAllZonefiles() {
	numBlocks := i.CurrentBlock - StartBlock
	fetchBlocks := 100
	iter := (numBlocks/fetchBlocks + 1)
	for page := 0; page <= iter; page++ {
		st := StartBlock + (page * fetchBlocks)
		end := st + 100
		log.Printf("%v Fetching zonefiles from block %v to block %v", prefix, st, end)
		i.GetZonefilesByBlock(st, end)
		i.ProcessedTo.Set(float64(end))
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
		go i.GetAllNamesInNamespace(n)
	}
}

func (i *Indexer) getNamePageAsync(page int, iter int, ns string) {
	namePage, err := i.Client().GetNamesInNamespace(ns, page*100, 100)
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range namePage.Names {
		i.enrichChan <- NewDomain(name, "ioapsdjfoaisdjf")
	}
}

func (i *Indexer) handleNameChan() {
	for n := range i.nameChan {
		i.Lock()
		i.Names = append(i.Names, n)
		i.Unlock()
	}
}

func (i *Indexer) handleEnrichChan() {
	for n := range i.enrichChan {
		go i.enrichName(n)
	}
}

func (i *Indexer) enrichName(name *Domain) {
	// i.Lock()
	// cb := i.CurrentBlock
	// i.Unlock()
	res, err := i.Client().GetNameBlockchainRecord(name.Name)
	if err != nil {
		log.Fatal(err)
	}
	name.Address = res.Record.Address
	log.Println(res.JSON())
	i.nameChan <- name
}

// GetAllNamesInNamespace gets all the names in a namespace
func (i *Indexer) GetAllNamesInNamespace(ns string) {
	numNames, err := i.Client().GetNumNamesInNamespace(ns)
	if err != nil {
		log.Fatal(err)
	}
	iter := (numNames.Count/100 + 1)
	for page := 0; page <= iter; page++ {
		go i.getNamePageAsync(page, iter, ns)
	}
}

// registerClient takes a blockstack.ServerConfig and tries to contact that Server
// if it is successful it is added to the rotation if not it is excluded
func (i *Indexer) registerClient(conf blockstack.ServerConfig) {
	client := blockstack.NewClient(conf)
	res, err := client.GetInfo()
	if err != nil {
		log.Printf("%s Failed to contact %s:%v, excluding from rotation", prefix, conf.Address, conf.Port)
		return
	}
	i.clients = append(i.clients, client)
	i.CurrentBlock = res.LastBlockProcessed
}
