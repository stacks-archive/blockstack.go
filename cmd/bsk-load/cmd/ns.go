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
	"sync"
	"time"

	"github.com/blockstack/blockstack.go/blockstack"
)

func newNameScan(cfg blockstack.ServerConfig) nameScan {
	ns := nameScan{
		c:          blockstack.NewClient(cfg),
		namespaces: make([]namespace, 0),
		l:          newLatencies(),
		lchan:      make(chan *latency, 0),
	}
	go ns.handleLatencies()
	return ns
}

type nameScan struct {
	c          *blockstack.Client
	namespaces []namespace
	l          *latencies

	namesScanned int

	*sync.Mutex
	lchan      chan *latency
	totalCalls int
}

type l struct {
	latencies []*latency

	sync.Mutex
}

func newL() *l {
	return &l{
		latencies: make([]*latency, 0),
	}
}

func (ns nameScan) handleLatencies() {
	for l := range ns.lchan {
		ns.l.Lock()
		ns.l.l = append(ns.l.l, l)
		ns.l.Unlock()
	}
}

func (ns nameScan) gans() blockstack.GetAllNamespacesResult {
	call := "get_all_namespaces"
	l := newLatency(call)
	res, err := ns.c.GetAllNamespaces()
	check(err)
	l.endTime = time.Now()
	ns.lchan <- l
	return res
}

func (ns nameScan) gnnin(namespace string) blockstack.CountResult {
	call := "get_num_names_in_namespace"
	l := newLatency(call)
	res, err := ns.c.GetNumNamesInNamespace(namespace)
	check(err)
	l.endTime = time.Now()
	ns.lchan <- l
	return res
}

func (ns nameScan) gnin(namespace string, page int, namePageSize int) blockstack.GetNamesInNamespaceResult {
	call := "get_names_in_namespace"
	l := newLatency(call)
	res, err := ns.c.GetNamesInNamespace(namespace, page*namePageSize, namePageSize)
	check(err)
	l.endTime = time.Now()
	ns.lchan <- l
	return res
}

func (ns nameScan) gnbr(name string) blockstack.GetNameBlockchainRecordResult {
	call := "get_blockchain_name_record"
	l := newLatency(call)
	res, err := ns.c.GetNameBlockchainRecord(name)
	check(err)
	l.endTime = time.Now()
	ns.lchan <- l
	return res
}

func (ns nameScan) gz(zonefiles []string) blockstack.GetZonefilesResult {
	call := "get_zonefiles"
	l := newLatency(call)
	res, err := ns.c.GetZonefiles(zonefiles)
	check(err)
	l.endTime = time.Now()
	ns.lchan <- l
	return res
}

func (ns nameScan) report() {
	ticker := time.NewTicker(time.Second * reportingSeconds)
	for _ = range ticker.C {
		ns.l.Lock()
		fmt.Println(ns.l.byCallSummary())
		ns.l.Unlock()
	}
}

type namespace struct {
	name     string
	names    []*string
	zfhs     []*string
	numNames int

	*sync.WaitGroup
	*sync.Mutex
}
