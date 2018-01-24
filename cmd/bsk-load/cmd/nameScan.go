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

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/blockstack/blockstack.go/blockstack"
	"github.com/spf13/cobra"
)

const namePageSize = 100
const concurrency = 5
const fetchNamesDetails = true
const fetchZonefiles = true

func newNameScan(cfg blockstack.ServerConfig) nameScan {
	return nameScan{
		c:          blockstack.NewClient(cfg),
		namespaces: make([]namespace, 0),
		latencies:  make([]*latency, 0),
	}
}

type nameScan struct {
	c          *blockstack.Client
	namespaces []namespace
	latencies  []*latency

	namesScanned int

	*sync.Mutex
}

func (ns nameScan) addNamespaces(n []string) {
	for _, name := range n {
		ns.namespaces = append(ns.namespaces, namespace{name: name})
	}
}

func (ns nameScan) addLatency(l *latency) {
	ns.Lock()
	l.endTime = time.Now()
	ns.latencies = append(ns.latencies, l)
	ns.Unlock()
}

type namespace struct {
	name     string
	names    []*string
	zfhs     []*string
	numNames int

	*sync.WaitGroup
	*sync.Mutex
}

func (ns namespace) addName(name string) {
	ns.Lock()
	ns.names = append(ns.names, &name)
	ns.Unlock()
}

func (ns namespace) addZFH(zfh string) {
	ns.Lock()
	ns.zfhs = append(ns.zfhs, &zfh)
	ns.Unlock()
}

type byCall map[string]latencies

func (bc byCall) JSON() string {

}

type latencies []*latency

func (l latencies) Len() int {
	return len(l)
}

func (l latencies) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l latencies) Less(i, j int) bool {
	return l[i].t() < l[j].t()
}

func (l latencies) max() time.Duration {
	if sort.IsSorted(l) {
		return l[0].t()
	}
	sort.Sort(l)
	return l[0].t()
}

// TODO: Percentiles

func (l latencies) min() time.Duration {
	if sort.IsSorted(l) {
		return l[len(l)-1].t()
	}
	sort.Sort(l)
	return l[len(l)-1].t()
}

func (l latencies) mean() float64 {
	var out time.Duration
	for _, lat := range l {
		out += lat.t()
	}
	return float64(out) / float64(len(l))
}

func (l latencies) meanResponseByCall() byCall {
	ret := make(byCall, 0)
	for _, lat := range l {
		if val, ok := ret[lat.call]; ok {
			ret[lat.call] = append(val, lat)
		} else {
			ret[lat.call] = latencies{lat}
		}
	}
	return ret
}

type latency struct {
	call      string
	startTime time.Time
	endTime   time.Time
}

func (l *latency) t() time.Duration {
	return l.startTime.Sub(l.endTime)
}

func newLatency(call string) *latency {
	return &latency{
		call:      call,
		startTime: time.Now(),
	}
}

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
		fmt.Println("Running test against", cfg.TestServer.String())
		// Initialize the testing struct and client, and run the test
		ns := newNameScan(cfg.TestServer)
		ns.runNameScan()
	},
}

func (ns nameScan) runNameScan() {
	// Get the list of namespaces
	l := newLatency("get_all_namespaces")
	res, err := ns.c.GetAllNamespaces()
	check(err)
	ns.addLatency(l)

	// Add them to the testing struct
	ns.addNamespaces(res.Namespaces)

	// Loop through namespaces getting all the names
	for _, namespace := range ns.namespaces {
		l := newLatency("get_num_names_in_namespace")
		res, err := ns.c.GetNumNamesInNamespace(namespace.name)
		check(err)
		ns.addLatency(l)
		namespace.numNames = res.Count
		iter := (res.Count/namePageSize + 1)
		sem := make(chan struct{}, concurrency)
		for page := 0; page <= iter; page++ {
			sem <- struct{}{}
			go ns.getNamePageAsync(page, namespace, sem)
			namespace.Add(1)
		}
		namespace.Wait()
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
			l := newLatency("get_zonfiles")
			_, err := ns.c.GetZonefiles(query)
			check(err)
			ns.addLatency(l)
		}
	}
}

// A goroutine safe method for fetching the list of names from blockstack-core
func (ns nameScan) getNamePageAsync(page int, namespace namespace, sem chan struct{}) {
	l := newLatency("get_names_in_namespace")
	namePage, err := ns.c.GetNamesInNamespace(namespace.name, page*namePageSize, namePageSize)
	check(err)
	ns.addLatency(l)

	for _, name := range namePage.Names {
		namespace.addName(name)
		if fetchNamesDetails {
			l := newLatency("get_blockchain_name_record")
			res, err := ns.c.GetNameBlockchainRecord(name)
			check(err)
			ns.addLatency(l)

			namespace.addZFH(res.Record.ValueHash)
		}
	}
	<-sem
	namespace.Done()
}

func init() {
	rootCmd.AddCommand(nameScanCmd)
}
