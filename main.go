package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jackzampolin/blockstack-indexer/api"
	"github.com/jackzampolin/blockstack-indexer/blockstack"
)

func main() {
	port := 3000
	conf := blockstack.ServerConfig{
		Address: "node.blockstack.org",
		Port:    6263,
		TLS:     true,
	}
	router := api.NewRouter(conf)

	log.Printf("Serving the Blockstack API on port %v...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}

// import (
// 	"fmt"
//
// 	"github.com/jackzampolin/blockstack-indexer/blockstack"
// )
//
// func main() {
// 	conf := blockstack.ServerConfig{
// 		Address: "node.blockstack.org",
// 		Port:    6263,
// 		TLS:     true,
// 	}
// 	bsk := blockstack.NewClient(conf)
//
// 	foo := bsk.TestMethod("get_namespace_blockchain_record", []interface{}{"id"})
// 	fmt.Println(foo)
// }
