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

	"github.com/blockstack/go-blockstack/indexer"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the indexer and serves the metrics server",
	Run: func(cmd *cobra.Command, args []string) {
		cfgs := viper.GetStringSlice("hosts")
		prt := viper.GetString("port")
		im := viper.GetString("indexMethod")
		pfc := viper.GetInt("pageFetchConc")
		npw := viper.GetInt("namePageWorkers")
		rw := viper.GetInt("resolveWorkers")
		bs := viper.GetInt("dbBatchSize")
		dbw := viper.GetInt("dbWorkers")

		log.Println(serveLog, "indexMethod", im)
		log.Println(serveLog, "hosts", len(cfgs))
		log.Println(serveLog, "port", prt)
		log.Println(serveLog, "pageFetchConc", pfc)
		log.Println(serveLog, "namePageWorkers", npw)
		log.Println(serveLog, "resolveWorkers", rw)
		log.Println(serveLog, "dbBatchSize", bs)
		log.Println(serveLog, "dbWorkers", dbw)

		// Create the indexer object with config
		// (clients []string, pageFetchConc, namePageWorkers, resolveWorkers, dbBatchSize, dbWorkers int, indexMethod string)
		idx := indexer.NewIndexer(indexer.NewConfig(cfgs, pfc, npw, rw, bs, dbw, im))

		// Kick off the indexer in a goroutine. Currently it just runs once.
		go idx.Start()

		// Expose the registered metrics via HTTP.
		http.Handle("/metrics", promhttp.Handler())
		log.Printf("%v Serving the prometheus metrics for the indexing service on port :%v...", serveLog, prt)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", prt), nil))
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
