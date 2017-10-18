package main

import (
	"log"
	"net/http"

	"github.com/jackzampolin/go-blockstack/api"
	"github.com/jackzampolin/go-blockstack/blockstack"
)

func main() {
	router := api.NewRouter(blockstack.ServerConfig{Address: "52.175.238.175", Port: 6264, TLS: false})

	log.Fatal(http.ListenAndServe(":8080", router))

}
