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

// getClient returns the client for the configured node
func getClient() blockstack.Client {
	scheme := "http"
	url, err := url.Parse(viper.GetString("node"))
	if err != nil {
		fmt.Printf("Unable to parse node address: %v\n", err)
		os.Exit(1)
	}
	if url.Scheme != "" {
		scheme = url.Scheme
	}
	conf := blockstack.ServerConfig{Address: url.Hostname(), Port: url.Port(), Scheme: scheme}
	if conf.Address == "" {
		fmt.Printf("Unable to parse node address: %#v\n", conf)
		os.Exit(1)
	}
	return *blockstack.NewClient(conf)
}

// handleResult prints results from the RPC calls
func handleResult(res blockstack.Response, err blockstack.Error) {
	if !viper.GetBool("pretty") {
		if err != nil {
			fmt.Println(err.PrettyJSON())
		} else {
			fmt.Println(res.PrettyJSON())
		}
	} else {
		if err != nil {
			fmt.Println(err.JSON())
		} else {
			fmt.Println(res.JSON())
		}
	}
}

// validateIntArg takes the string arg from the commandline and converts it to an int exiting on error
func validateIntArg(i, argName string) int {
	out, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		fmt.Println("Error parsing int for", argName)
		os.Exit(1)
	}
	return int(out)
}

func init() {
	RootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.blockstack.yaml)")
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringP("node", "n", "http://localhost:6264", "blockstack-core node to run rpc comands against")
	RootCmd.PersistentFlags().BoolP("pretty", "p", false, "toggle to turn off pretty print JSON responses")
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose mode")
	viper.BindPFlag("node", RootCmd.PersistentFlags().Lookup("node"))
	viper.BindPFlag("pretty", RootCmd.PersistentFlags().Lookup("pretty"))
	viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))
}

// initConfig makes sure the config file is properly initialized
func initConfig() {
	if viper.GetString("config") != "" {
		viper.SetConfigFile(viper.GetString("config"))
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".blockstackd-cli")
	}
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if viper.GetBool("verbose") {
		if err != nil {
			fmt.Println("Using config at", viper.ConfigFileUsed())
		} else {
			fmt.Println("Config Error:", err)
		}
	}
}
