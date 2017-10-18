package indexer

import (
	// "log"
	// "time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	promNameSpace     = "indexer"
	reportingInterval = 2
)

type indexerStats struct {
	callsMade          prometheus.Counter
	namePagesFetched   prometheus.Counter
	nameDetailsFetched prometheus.Counter
	zonefilesFetched   prometheus.Counter
	namesResolved      prometheus.Counter
	namesOnNetwork     prometheus.Gauge
	timeSinceStart     prometheus.Gauge
}

func newIndexerStats() *indexerStats {
	s := &indexerStats{
		callsMade: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: promNameSpace,
			Subsystem: "core_calls",
			Name:      "num_made",
			Help:      "the number of core RPC calls made",
		}),
		namePagesFetched: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: promNameSpace,
			Subsystem: "name",
			Name:      "pages_fetched",
			Help:      "the number of pages of 100 names fetched",
		}),
		nameDetailsFetched: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: promNameSpace,
			Subsystem: "name",
			Name:      "details_fetched",
			Help:      "the number names where details have been fetched",
		}),
		zonefilesFetched: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: promNameSpace,
			Subsystem: "zonefiles",
			Name:      "num_fetched",
			Help:      "the number zonefiles for given names that have been fetched",
		}),
		namesResolved: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: promNameSpace,
			Subsystem: "name",
			Name:      "resolved",
			Help:      "the number names that have been resolved",
		}),
		namesOnNetwork: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: promNameSpace,
			Subsystem: "name",
			Name:      "on_network",
			Help:      "the number names on the blockstack network",
		}),
	}
	prometheus.MustRegister(s.callsMade)
	prometheus.MustRegister(s.namePagesFetched)
	prometheus.MustRegister(s.nameDetailsFetched)
	prometheus.MustRegister(s.zonefilesFetched)
	prometheus.MustRegister(s.namesResolved)
	prometheus.MustRegister(s.namesOnNetwork)
	return s
}
