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
	"net/url"
	"os"
	"strconv"

	"github.com/blockstack/go-blockstack/blockstack"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	host    string
)

// RootCmd is the root command of the CLI
var RootCmd = &cobra.Command{
	Use:   "blockstack",
	Short: "An RPC call runner for blockstack",
}

// Execute runs the rest of the commands
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getClient() blockstack.Client {
	url, err := url.Parse(viper.GetString("host"))
	if err != nil {
		fmt.Printf("Unable to parse URL not adding node to rotation: %v", err)
	}
	conf := blockstack.ServerConfig{Address: url.Hostname(), Port: url.Port(), Scheme: url.Scheme}
	return *blockstack.NewClient(conf)
}

func validateIntArg(i, argName string) int {
	out, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		fmt.Println("Error parsing int for", argName)
	}
	return int(out)
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.blockstack.yaml)")
	RootCmd.PersistentFlags().StringVar(&host, "host", "http://localhost:6264", "blockstack-core node to run rpc comands against")
	viper.BindPFlag("host", RootCmd.PersistentFlags().Lookup("host"))
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".blockstack")
	}
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Failed to parse config file", viper.ConfigFileUsed(), ":", err)
	}
}
