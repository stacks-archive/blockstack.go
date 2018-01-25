// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import "github.com/spf13/cobra"

const namePageSize = 100
const concurrency = 10
const fetchNamesDetails = true
const fetchZonefiles = true
const reportingSeconds = 10

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// nameScanCmd represents the nameScan command
var nameScanCmd = &cobra.Command{
	Use:   "nameScan",
	Short: "Scan all names from all namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		ns := newNameScan(cfg.TestServer)
		go ns.report()
		ns.runNameScan()
	},
}

func (ns nameScan) runNameScan() {
	// Get the list of namespaces
	res := ns.gans()

	// Add them to the testing struct
	for _, name := range res.Namespaces {
		ns.namespaces = append(ns.namespaces, namespace{name: name})
	}

	// Loop through namespaces getting all the names
	for _, namespace := range ns.namespaces {
		// Find the number of names in the namespace
		res := ns.gnnin(namespace.name)

		namespace.numNames = res.Count
		iter := (res.Count/namePageSize + 1)
		sem := make(chan struct{}, concurrency)
		for page := 0; page <= iter; page++ {
			sem <- struct{}{}
			go ns.getNamePageAsync(page, namespace, sem)
		}
	}

	// if zonefiles need fetching, also fetch the zonefiles
	if fetchZonefiles {
		zfhs := make([][]string, 0)
		for _, namespace := range ns.namespaces {
			zfha := make([]string, 0)
			for _, zfh := range namespace.zfhs {
				if len(zfha) == 100 {
					zfhs = append(zfhs, zfha)
					zfha = make([]string, 0)
				}
				zfha = append(zfha, *zfh)
			}
		}
		for _, query := range zfhs {
			ns.gz(query)
		}
	}
}

// A goroutine safe method for fetching the list of names from blockstack-core
func (ns nameScan) getNamePageAsync(page int, namespace namespace, sem chan struct{}) {
	namePage := ns.gnin(namespace.name, page, namePageSize)

	for _, name := range namePage.Names {
		namespace.names = append(namespace.names, &name)
		if fetchNamesDetails {
			res := ns.gnbr(name)
			namespace.zfhs = append(namespace.zfhs, &res.Record.ValueHash)
		}
	}

	<-sem
}

func init() {
	rootCmd.AddCommand(nameScanCmd)
}
