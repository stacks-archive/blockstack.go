package indexer

import (
	"fmt"

	"github.com/miekg/dns"
)

// Zonefile models a Zonefile
type Zonefile struct {
	Raw string   `json:"raw"`
	RRs []dns.RR `json:"RRs"`

	Compliant bool
}

// This is a JSON representation of a zonefile for out output
// TODO: Do I need this?
type zonefileOut struct {
	Origin string   `json:"$origin"`
	TTL    string   `json:"$ttl"`
	URI    []dns.RR `json:"URI"`
	TXT    []dns.RR `json:"TXT"`
}

// GetURI returns the first URI with a Target starting with http
func (zf *Zonefile) GetURI() *dns.URI {
	var URI *dns.URI
	for _, rr := range zf.RRs {
		if rr.Header().Rrtype == 256 {
			uri := rr.(*dns.URI)
			if goodTarget(uri.Target) {
				URI = uri
				break
			}
		}
	}
	return URI
}

// JSON representation of a Zonefile
func (zf *Zonefile) JSON() []byte {
	var byt []byte
	for _, rr := range zf.RRs {
		fmt.Printf("%#v\n", rr)
	}
	return byt
}
