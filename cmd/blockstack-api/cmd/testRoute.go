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
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// "/v2/users/muneeb.id", *

// "/v1/addresses/bitcoin/{address}",
// "/v1/namespaces",
// "/v1/names/muneeb.id/zonefile",
// "/v1/names/muneeb.id",
// "/v1/names/muneeb.id/history",
// "/v1/namespaces/{namespace}/names?page={number}",
// "/v1/blockchains/{blockchain}/operations/{blockHeight}",
// "/v1/namespaces/{namespace}",
// "/v1/blockchains/{blockchain}/name_count",

// testRouteCmd represents the testRoute command
var testRouteCmd = &cobra.Command{
	Use:   "test-route",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get(fmt.Sprintf("http://localhost:%v/%v", viper.GetInt("port"), args[0]))
		if err != nil {
			panic(err)
		}
		bdy, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bdy))
	},
}

func init() {
	RootCmd.AddCommand(testRouteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testRouteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testRouteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
