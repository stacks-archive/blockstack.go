package indexer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/blockstack/go-blockstack/blockstack"
	"github.com/miekg/dns"
)

// NewDomain returns an initialized Domain
func NewDomain(name string) *Domain {
	out := &Domain{
		Name:         name,
		lastResolved: time.Time{},
	}
	return out
}

// Domain models a Blockstack domain name
type Domain struct {
	Name string `json:"name"`
	// Address  string    `json:"address"`
	Zonefile *Zonefile `json:"zonefile"`
	Profile  Profile   `json:"profile"`

	getNameAtRes blockstack.GetNameAtResult
	lastResolved time.Time
}

func (d *Domain) zonefileHash() string {
	if len(d.getNameAtRes.Records) > 0 && d.getNameAtRes.Records[0].ValueHash != "" {
		return d.getNameAtRes.Records[0].ValueHash
	}
	return ""
}

// JSON returns the JSON representation of Domain
func (d *Domain) JSON() string {
	byt, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(byt)
}

// GetURI returns the first URI with a Target starting with http
func (d *Domain) GetURI() *dns.URI {
	var URI *dns.URI
	for _, rr := range d.Zonefile.RRs {
		if rr.Header().Rrtype == 256 && goodTarget(rr.(*dns.URI).Target) {
			URI = rr.(*dns.URI)
			break
		}
	}
	return URI
}

// AddZonefile takes a string representation of a Zonefile and parses out some info
func (d *Domain) AddZonefile(zonefile string) {
	d.Zonefile = &Zonefile{
		Raw:       zonefile,
		RRs:       make([]dns.RR, 0),
		Compliant: true,
	}
	for x := range dns.ParseZone(strings.NewReader(zonefile), "", "") {
		if x.Error != nil {
			d.Zonefile.Compliant = false
			var legacyProfile LegacyProfile
			// NOTE: Squash error here. We don't care about it
			json.Unmarshal([]byte(zonefile), &legacyProfile)
			if legacyProfile.Account == nil {
				d.Profile = nil
			} else {
				d.Profile = legacyProfile
			}
		} else {
			// TODO: Handle fetching subdomains here...
			// Might need to lock the []*Domain in a bunch of places
			// if x.RR.Header().Rrtype == 16 {
			// 	fmt.Println(x.RR)
			// }
			// created_equal.self_evident_truth.id.	3600	IN	TXT	"owner=1AYddAnfHbw6bPNvnsQFFrEuUdhMhf2XG9" "seqn=0" "parts=1" "zf0=JE9SSUdJTiBjcmVhdGVkX2VxdWFsCiRUVEwgMzYwMApfaHR0cHMuX3RjcCBVUkkgMTAgMSAiaHR0cHM6Ly93d3cuY3MucHJpbmNldG9uLmVkdS9+YWJsYW5rc3QvY3JlYXRlZF9lcXVhbC5qc29uIgpfZmlsZSBVUkkgMTAgMSAiZmlsZTovLy90bXAvY3JlYXRlZF9lcXVhbC5qc29uIgo="
			d.Zonefile.RRs = append(d.Zonefile.RRs, x.RR)
		}
	}
}

// ResolveProfile takes an initialized domain and fetches the resulting profile for that domain
// TODO: Make this fail early and often to prevent bottleneck also this basically doesn't work
func (d *Domain) ResolveProfile() {
	// fmt.Println(d.Name)
	URI := d.GetURI()
	// Check that a URI was returned and there is a Target
	// Also check that the profile was not in the legacy format
	if URI != nil && d.Profile != nil {
		target := URI.Target

		// Handle dropbox urls with no http prefix
		if strings.TrimPrefix(target, "www") != target {
			target = fmt.Sprintf("http://%s", target)
		}
		res, err := http.Get(target)
		if err != nil || res.StatusCode == 404 {
			log.Printf("Error fetching profile for %v: %v\n", d.Name, err)
			d.lastResolved = time.Now()
			return
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("Error reading body for %v: %v", d.Name, err)
		}
		var p1 []SOProfile
		err = json.Unmarshal(body, &p1)
		if err != nil && strings.Contains(err.Error(), "cannot unmarshal object") {
			var p2 SOProfile
			json.Unmarshal(body, &p2)
			if err != nil {
				fmt.Printf("Error unmarshalling %v\nERROR: %v\nRES: %v\nTARGET: %v\n\n", d.Name, err, string(body), target)
			}
		} else {
			if len(p1) > 0 {
				d.Profile = &p1[0]
			}
		}
		res.Body.Close()
	}
	d.lastResolved = time.Now()
}

// Domains is a collection of *Domain
type Domains []*Domain

// getZonefileHashes returns an array of
func (d Domains) getZonefileHashes() (out []string) {
	for _, dom := range d {
		recs := dom.getNameAtRes.Records
		// TODO: turn this into a method on domain
		if len(recs) > 0 && recs[0].ValueHash != "" {
			out = append(out, recs[0].ValueHash)
		}
	}
	return out
}
