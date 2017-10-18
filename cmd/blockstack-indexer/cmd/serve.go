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

	"github.com/blockstack/go-blockstack/blockstack"
	"github.com/blockstack/go-blockstack/indexer"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port      = 3000
	logPrefix = "[main]"
)

func serverConfFromString(s string) {

}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfgs := viper.GetStringSlice("hosts")
		var configs blockstack.ServerConfigs
		for _, conf := range cfgs {
			url, err := url.Parse(conf)
			if err != nil {
				log.Printf("Unable to parse URL %v, not adding to rotation", conf)
				continue
			}
			configs = append(configs, blockstack.ServerConfig{Address: url.Hostname(), Port: url.Port(), Scheme: url.Scheme})
		}

		go indexer.StartIndexer(configs)
		// Expose the registered metrics via HTTP.
		http.Handle("/metrics", promhttp.Handler())
		log.Printf("Serving the prometheus metrics for the indexing service on port :%v...", 3000)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
