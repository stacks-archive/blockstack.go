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
	"encoding/json"
	"fmt"
	"sort"
	"sync"
	"time"
)

type byCall map[string][]*latency

func (bc byCall) JSON(start time.Time) string {
	var out = make(map[string]map[string]string, 0)
	var totalCalls = 0
	for k, v := range bc {
		out[k] = ls(v).summary()
		totalCalls += len(v)
	}
	testTime := time.Now().Sub(start)
	out["totals"] = make(map[string]string)
	out["totals"]["time"] = fmt.Sprint(testTime)
	out["totals"]["calls"] = fmt.Sprint(totalCalls)
	out["totals"]["callsPerSec"] = fmt.Sprint(float64(totalCalls) / testTime.Seconds())
	o, err := json.Marshal(out)
	check(err)
	return string(o)
}

type latencies struct {
	l     []*latency
	start time.Time

	sync.Mutex
}

func newLatencies() *latencies {
	return &latencies{
		l:     make([]*latency, 0),
		start: time.Now(),
	}
}

func (l ls) summary() map[string]string {
	return map[string]string{
		"numCalls": fmt.Sprint(len(l)),
		"mean":     fmt.Sprint(l.mean()),
		"max":      fmt.Sprint(l.max()),
		"min":      fmt.Sprint(l.min()),
		"p50":      fmt.Sprint(l.perc(.5)),
		"p90":      fmt.Sprint(l.perc(.9)),
		"p95":      fmt.Sprint(l.perc(.95)),
	}
}

func (l ls) perc(p float32) time.Duration {
	var index = len(l) - int(float32(len(l))*p)
	if sort.IsSorted(l) {
		if len(l) == 1 {
			return l[0].t()
		}
		return l[index].t()
	}
	sort.Sort(l)
	if len(l) == 1 {
		return l[0].t()
	}
	return l[index].t()
}

func (l ls) Len() int {
	return len(l)
}

func (l ls) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ls) Less(i, j int) bool {
	return l[i].t() > l[j].t()
}

func (l ls) max() time.Duration {
	if sort.IsSorted(l) {
		return l[0].t()
	}
	sort.Sort(l)
	return l[0].t()
}

// TODO: Percentiles

func (l ls) min() time.Duration {
	if sort.IsSorted(l) {
		return l[len(l)-1].t()
	}
	sort.Sort(l)
	return l[len(l)-1].t()
}

func (l ls) mean() time.Duration {
	var out time.Duration
	for _, lat := range l {
		out += lat.t()
	}
	return time.Duration(int(out) / len(l))
}

func (l latencies) byCallSummary() string {
	ret := make(byCall, 0)
	for _, lat := range l.l {
		if _, ok := ret[lat.call]; ok {
			ret[lat.call] = append(ret[lat.call], lat)
		} else {
			ret[lat.call] = []*latency{lat}
		}
	}
	return ret.JSON(l.start)
}

type ls []*latency

type latency struct {
	call      string
	startTime time.Time
	endTime   time.Time
}

func (l *latency) String() string {
	return fmt.Sprintf("{call: %v, start: %v, end: %v}", l.call, l.startTime, l.endTime)
}

func (l *latency) t() time.Duration {
	return l.endTime.Sub(l.startTime)
}

func newLatency(call string) *latency {
	return &latency{
		call:      call,
		startTime: time.Now(),
	}
}
