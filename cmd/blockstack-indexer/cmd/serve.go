// Copyright © 2017 Jack Zampolin <jack.zampolin@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/blockstack/go-blockstack/blockstack"
	"github.com/blockstack/go-blockstack/indexer"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// register the clients before passing them to the indexer
func registerClient(conf blockstack.ServerConfig) *blockstack.Client {
	client := blockstack.NewClient(conf)
	_, err := client.GetInfo()
	if err != nil {
		log.Printf("%s Failed to contact %s", serveLog, conf)
		return nil
	}
	return client
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the indexer and serves the metrics server",
	Run: func(cmd *cobra.Command, args []string) {
		cfgs := viper.GetStringSlice("hosts")
		var clients []*blockstack.Client

		// Make sure the urls are parsable and the hosts are reachable
		for _, cfg := range cfgs {
			url, err := url.Parse(cfg)
			if err != nil {
				log.Printf("Unable to parse URL '%v' not adding node to rotation", cfg)
				continue
			}
			conf := blockstack.ServerConfig{Address: url.Hostname(), Port: url.Port(), Scheme: url.Scheme}
			if c := registerClient(conf); c != nil {
				clients = append(clients, c)
			}
		}

		// Exit if none of the servers are reachable
		if len(clients) < 1 {
			log.Println(serveLog, "no reachable blockstack-nodes configured")
			os.Exit(1)
		}

		// Kick off the indexing process in a goroutine for now
		// TODO: make this run periodically and save the results somewhere
		go indexer.StartIndexer(clients)

		// Expose the registered metrics via HTTP.
		http.Handle("/metrics", promhttp.Handler())
		log.Printf("%v Serving the prometheus metrics for the indexing service on port :%v...", serveLog, viper.Get("port"))
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", viper.Get("port")), nil))
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
