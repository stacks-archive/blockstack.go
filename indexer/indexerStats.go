package indexer

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	mns = "indexer"
)

type indexerStats struct {
	ProcessedTo  prometheus.Gauge
	NumZonefiles prometheus.Gauge
	NumDomains   prometheus.Gauge
	NumResolved  prometheus.Gauge
}

func newIndexerStats() *indexerStats {
	s := &indexerStats{
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
	}
	prometheus.MustRegister(s.NumDomains)
	prometheus.MustRegister(s.NumResolved)
	prometheus.MustRegister(s.NumZonefiles)
	prometheus.MustRegister(s.ProcessedTo)
	return s
}
