package main

import (
	// "encoding/json"
	"fmt"
	// "log"

	"github.com/blockstack/go-blockstack/blockstack"
)

func main() {
	conf := blockstack.ServerConfig{
		Address: "52.175.238.175",
		Port:    6265,
		TLS:     false,
	}
	client := blockstack.NewClient(conf)
	one, _ := client.GetNameBlockchainRecord("muneeb.id")
	fmt.Println(one.JSON())
	two, _ := client.GetZonefiles([]string{one.Record.ValueHash})
	fmt.Println(two.JSON())

}
