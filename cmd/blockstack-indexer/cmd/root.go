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
	"log"
	"os"

	"github.com/blockstack/go-blockstack/blockstack"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	hosts    []string
	port     int
	rootLog  = "[main]"
	serveLog = "[server]"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "blockstack-indexer",
	Short: "The blockstack-indexer is a program that indexes the blockstack network",
}

// Execute registers all of the other commands
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

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

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.blockstack-indexer.yaml)")
	RootCmd.PersistentFlags().StringSlice("hosts", []string{"https://node.blockstack.org:6263"}, "blockstack-core nodes to run indexer against")
	RootCmd.PersistentFlags().IntVar(&port, "port", 3000, "port to run the prometheus metrics server on")
	viper.BindPFlag("port", RootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("hosts", RootCmd.PersistentFlags().Lookup("hosts"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".blockstack-indexer")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Println(rootLog, "Using config file:", viper.ConfigFileUsed())
	}
}
