package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/blockstack/go-blockstack/blockstack"
	"github.com/blockstack/go-blockstack/indexer"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	port      = 3000
	logPrefix = "[main]"
)

func main() {
	configs := blockstack.ServerConfigs{
		blockstack.ServerConfig{Address: "52.175.238.175", Port: 6264, TLS: false},
		blockstack.ServerConfig{Address: "52.175.238.175", Port: 6265, TLS: false},
		blockstack.ServerConfig{Address: "52.175.238.175", Port: 6266, TLS: false},
		blockstack.ServerConfig{Address: "52.175.238.175", Port: 6267, TLS: false},
		blockstack.ServerConfig{Address: "13.64.158.139", Port: 6264, TLS: false},
		blockstack.ServerConfig{Address: "13.64.158.139", Port: 6265, TLS: false},
		blockstack.ServerConfig{Address: "13.64.158.139", Port: 6266, TLS: false},
		blockstack.ServerConfig{Address: "13.64.158.139", Port: 6267, TLS: false},
	}

	go indexer.StartIndexer(configs)
	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Serving the prometheus metrics for the indexing service on port :%v...", 3000)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
