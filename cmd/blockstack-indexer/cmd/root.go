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

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile         string
	pageFetchConc   int
	namePageWorkers int
	resolveWorkers  int
	dbBatchSize     int
	dbWorkers       int
	indexMethod     string
	hosts           []string
	port            int
	rootLog         = "[main]"
	serveLog        = "[server]"
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

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.blockstack-indexer.yaml)")
	RootCmd.PersistentFlags().StringSlice("hosts", []string{"https://node.blockstack.org:6263"}, "blockstack-core nodes to run indexer against")
	RootCmd.PersistentFlags().IntVar(&port, "port", 3000, "port to run the prometheus metrics server on")
	RootCmd.PersistentFlags().IntVar(&pageFetchConc, "pageFetchConc", 100, "number of namePages to fetch concurrently")
	RootCmd.PersistentFlags().IntVar(&namePageWorkers, "namePageWorkers", 10, "number of workers to process namePages")
	RootCmd.PersistentFlags().IntVar(&resolveWorkers, "resolveWorkers", 50, "number of workers to resolve names")
	RootCmd.PersistentFlags().StringVar(&indexMethod, "indexMethod", "byName", "indexing method to employ")
	RootCmd.PersistentFlags().IntVar(&dbBatchSize, "dbBatchSize", 20, "number of names to insert/update at same time")
	RootCmd.PersistentFlags().IntVar(&dbWorkers, "dbWorkers", 4, "number of workers to manage inserts into database")
	viper.BindPFlag("port", RootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("hosts", RootCmd.PersistentFlags().Lookup("hosts"))
	viper.BindPFlag("pageFetchConc", RootCmd.PersistentFlags().Lookup("pageFetchConc"))
	viper.BindPFlag("namePageWorkers", RootCmd.PersistentFlags().Lookup("namePageWorkers"))
	viper.BindPFlag("resolveWorkers", RootCmd.PersistentFlags().Lookup("resolveWorkers"))
	viper.BindPFlag("indexMethod", RootCmd.PersistentFlags().Lookup("indexMethod"))
	viper.BindPFlag("dbBatchSize", RootCmd.PersistentFlags().Lookup("dbBatchSize"))
	viper.BindPFlag("dbWorkers", RootCmd.PersistentFlags().Lookup("dbWorkers"))
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
