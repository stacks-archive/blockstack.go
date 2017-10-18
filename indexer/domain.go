package indexer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jackzampolin/go-blockstack/blockstack"
	"github.com/miekg/dns"
)

// Domain models a Blockstack domain name
type Domain struct {
	Name          string         `json:"name"`
	Address       string         `json:"{add}ress"`
	Zonefile      *Zonefile      `json:"zonefile"`
	Profile       *Profile       `json:"profile"`
	LegacyProfile *LegacyProfile `json:"legacy_profile"`

	getNameAtRes blockstack.GetNameAtResult
	lastResolved time.Time
}

type Domains []*Domain

func (d Domains) getZonefiles() (out []string) {
	for _, dom := range d {
		recs := dom.getNameAtRes.Records
		if len(recs) > 0 && recs[0].ValueHash != "" {
			out = append(out, recs[0].ValueHash)
		}
	}
	return out
}

// JSON returns the JSON representation of Domain
func (d *Domain) JSON() string {
	byt, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// NewDomain returns an initialized Domain
func NewDomain(name string) *Domain {
	out := &Domain{
		Name:         name,
		lastResolved: time.Time{},
	}
	return out
}

func goodTarget(uri string) bool {
	// Check for URLs, if not a url return false
	h := strings.TrimPrefix(uri, "http")
	w := strings.TrimPrefix(uri, "www")
	if uri != "" && (h != uri || w != uri) {
		return true
	}
	return false
}

// GetURI returns the first URI with a Target starting with http
func (d *Domain) GetURI() *dns.URI {
	var URI *dns.URI
	for _, rr := range d.Zonefile.RRs {
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

// ResolveProfile takes an initialized domain and fetches the resulting profile for that domain
func (d *Domain) ResolveProfile(sem chan struct{}) {
	// fmt.Println(d.Name)
	URI := d.GetURI()
	// Check that a URI was returned and there is a Target
	if URI != nil {
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
		var p1 []Profile
		err = json.Unmarshal(body, &p1)
		if err != nil && strings.Contains(err.Error(), "cannot unmarshal object") {
			var p2 Profile
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
	<-sem
}
