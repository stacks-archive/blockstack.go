package main

import (
	// "encoding/json"
	"fmt"
	// "log"

	"github.com/jackzampolin/go-blockstack/blockstack"
)

func main() {
	conf := blockstack.ServerConfig{
		Address: "52.175.238.175",
		Port:    6265,
		TLS:     false,
	}
	client := blockstack.NewClient(conf)
	one, _ := client.GetOpHistoryRows("id", 0, 10)
	fmt.Println(one.JSON())
	two, _ := client.GetNumOpHistoryRows("id")
	fmt.Println(two.JSON())
	three, _ := client.GetNumNameOpsAffectedAt(480003)
	fmt.Println(three.JSON())
	four, _ := client.GetNameOpsAffectedAt(480003, 0, 10)
	fmt.Println(four.JSON())
	five, _ := client.GetConsensusHashes([]int{480003, 480005})
	fmt.Println(five.JSON())
}
