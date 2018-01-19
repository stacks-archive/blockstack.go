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

	"github.com/blockstack/blockstack.go/indexer"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the indexer and serves the metrics server",
	Run: func(cmd *cobra.Command, args []string) {
		prt := viper.GetString("port")

		cfg := &indexer.Config{
			IndexMethod:          viper.GetString("indexMethod"),
			NamePageWorkers:      viper.GetInt("namePageWorkers"),
			ResolveWorkers:       viper.GetInt("resolveWorkers"),
			ConcurrentPageFetch:  viper.GetInt("pageFetchConc"),
			DBBatchSize:          viper.GetInt("dbBatchSize"),
			DBWorkers:            viper.GetInt("dbWorkers"),
			URLs:                 viper.GetStringSlice("hosts"),
			ClientUpdateInterval: viper.GetInt("updateInterval"),
			MongoConnection:      viper.GetString("mongoConn"),
		}

		log.Println(serveLog, cfg)
		log.Println(serveLog, "Setting valid clients...")
		cfg.SetClients()
		idx := indexer.NewIndexer(cfg)

		go idx.Start()

		http.Handle("/metrics", promhttp.Handler())
		log.Printf("%v Serving the prometheus metrics for the indexing service on port :%v...", serveLog, prt)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", prt), nil))
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
