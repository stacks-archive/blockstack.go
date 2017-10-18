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

	"github.com/spf13/cobra"
)

// getZonefilesByBlockCmd represents the getZonefilesByBlock command
var getZonefilesByBlockCmd = &cobra.Command{
	Use:   "get_zonefiles_by_block",
	Short: "Calls the get_zonefiles_by_block rpc method on the configured host",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		client := getClient()
		res, err := client.GetZonefilesByBlock(
			validateIntArg(args[0], "startBlock"),
			validateIntArg(args[1], "endBlock"),
			validateIntArg(args[2], "offset"),
			validateIntArg(args[3], "count"),
		)
		if err != nil {
			fmt.Println("Error, TODO: Make some real JSON errors for the blockstack library", err)
		}
		fmt.Println(res.JSON())
	},
}

func init() {
	RootCmd.AddCommand(getZonefilesByBlockCmd)
}
