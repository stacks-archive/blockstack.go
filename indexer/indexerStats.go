package indexer

import (
	// "log"
	// "time"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	promNameSpace     = "indexer"
	reportingInterval = 2
)

type indexerStats struct {
	callsMade           prometheus.Gauge
	namePagesFetched    prometheus.Gauge
	nameDetailsFetched  prometheus.Gauge
	zonefilesFetched    prometheus.Gauge
	namesResolved       prometheus.Gauge
	namesOnNetwork      prometheus.Gauge
	timeSinceStart      prometheus.Gauge
	sentDownResolveChan prometheus.Gauge
	writtenToDatabase   prometheus.Gauge

	sync.Mutex
}

func newIndexerStats() *indexerStats {
	s := &indexerStats{
		callsMade: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: promNameSpace,
			Subsystem: "core_calls",
			Name:      "num_made",
			Help:      "the number of core RPC calls made",
		}),
		namePagesFetched: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: promNameSpace,
			Subsystem: "name",
			Name:      "pages_fetched",
			Help:      "the number of pages of 100 names fetched",
		}),
		nameDetailsFetched: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: promNameSpace,
			Subsystem: "name",
			Name:      "details_fetched",
			Help:      "the number names where details have been fetched",
		}),
		zonefilesFetched: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: promNameSpace,
			Subsystem: "zonefiles",
			Name:      "num_fetched",
			Help:      "the number zonefiles for given names that have been fetched",
		}),
		namesResolved: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: promNameSpace,
			Subsystem: "resolve",
			Name:      "num_resolved",
			Help:      "the number names that have been resolved",
		}),
		namesOnNetwork: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: promNameSpace,
			Subsystem: "name",
			Name:      "on_network",
			Help:      "the number names on the blockstack network",
		}),
		sentDownResolveChan: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: promNameSpace,
			Subsystem: "resolve",
			Name:      "down_chan",
			Help:      "the number names sent down the resolve channel",
		}),
		writtenToDatabase: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: promNameSpace,
			Subsystem: "database",
			Name:      "written",
			Help:      "the number names inserted/updated on the database",
		}),
	}
	prometheus.MustRegister(s.callsMade)
	prometheus.MustRegister(s.namePagesFetched)
	prometheus.MustRegister(s.nameDetailsFetched)
	prometheus.MustRegister(s.zonefilesFetched)
	prometheus.MustRegister(s.namesResolved)
	prometheus.MustRegister(s.namesOnNetwork)
	prometheus.MustRegister(s.sentDownResolveChan)
	prometheus.MustRegister(s.writtenToDatabase)
	return s
}
