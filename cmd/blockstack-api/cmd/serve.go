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

	"github.com/blockstack/blockstack.go/api"
	"github.com/blockstack/blockstack.go/blockstack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "This serves the blockstack api",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement the same multiclient thing here that I did over in the indexer
		router := api.NewRouter(blockstack.ServerConfig{Address: "node.blockstack.org", Port: "6263", Scheme: "https"})
		log.Println("Serving the blockstack-api on port", viper.GetInt("port"))
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("port")), router))

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
