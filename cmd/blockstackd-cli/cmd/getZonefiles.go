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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getZonefilesCmd represents the getZonefiles command
var getZonefilesCmd = &cobra.Command{
	Use:   "get_zonefiles [zonefileHashes...]",
	Short: "[zonefileHashes...]",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := getClient()
		var zfhs []string
		for _, zfh := range args {
			zfhs = append(zfhs, zfh)
		}
		res, err := client.GetZonefiles(zfhs)
		if viper.GetBool("decode") {
			zfs := res.Decode()
			for zfh := range zfs {
				res.Zonefiles[zfh] = zfs[zfh]
			}
			handleResult(res, err)
		} else {
			handleResult(res, err)
		}
	},
}

func init() {
	RootCmd.AddCommand(getZonefilesCmd)
	getZonefilesCmd.Flags().BoolP("decode", "d", false, "toggle to decode the base64 encoded results")
	viper.BindPFlag("decode", getZonefilesCmd.Flags().Lookup("decode"))
}
