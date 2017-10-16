package indexer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/miekg/dns"
)

// Domain models a Blockstack domain name
type Domain struct {
	Name          string         `json:"name"`
	Address       string         `json:"address"`
	Zonefile      *Zonefile      `json:"zonefile"`
	Profile       *Profile       `json:"profile"`
	LegacyProfile *LegacyProfile `json:"legacy_profile"`

	resolved bool
}

// JSON returns the JSON representation of Domain
func (d *Domain) JSON() string {
	byt, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// Domains is a collection of Domain
type Domains []*Domain

// NewDomain returns an initialized Domain
func NewDomain(name string, zonefile string) *Domain {
	out := &Domain{
		Name:     name,
		resolved: false,
	}
	out.AddZonefile(zonefile)
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
			d.resolved = true
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
	<-sem
	d.resolved = true
}

// Zonefile models a Zonefile
type Zonefile struct {
	Raw string   `json:"raw"`
	RRs []dns.RR `json:"RRs"`

	compliant bool
}

// AddZonefile takes a string representation of a Zonefile and parses out some info
func (d *Domain) AddZonefile(zonefile string) {
	d.Zonefile = &Zonefile{
		Raw:       zonefile,
		RRs:       make([]dns.RR, 0),
		compliant: true,
	}
	for x := range dns.ParseZone(strings.NewReader(zonefile), "", "") {
		if x.Error != nil {
			d.Zonefile.compliant = false
			var legacyProfile LegacyProfile
			// NOTE: Squash error here. We don't care about it
			json.Unmarshal([]byte(zonefile), &legacyProfile)
			if legacyProfile.Account == nil {
				d.LegacyProfile = nil
			} else {
				d.LegacyProfile = &legacyProfile
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

// Profile models a blockstack profile
type Profile struct {
	Token           string       `json:"token"`
	ParentPublicKey string       `json:"parentPublicKey"`
	Encrypted       bool         `json:"encrypted"`
	DecodedToken    DecodedToken `json:"decodedToken"`
}

// Account models a social media proof
type Account struct {
	Type       string `json:"@type"`
	Service    string `json:"service"`
	ProofType  string `json:"proofType"`
	Identifier string `json:"identifier"`
	ProofURL   string `json:"proofUrl"`
}

// Image models a Profile Image
type Image struct {
	Type       string `json:"@type"`
	ContentURL string `json:"contentUrl"`
	Name       string `json:"name"`
}

// Claim contains social proofs and images
type Claim struct {
	Type    string    `json:"@type"`
	Image   []Image   `json:"image"`
	Account []Account `json:"account"`
}

// PublicKey models {publicKey: "030ec5101181a8e528b70141b0cde18fda231ab1be5f166e49f813c63914f4ebc8"}
type PublicKey struct {
	PublicKey string `json:"publicKey"`
}

// Payload contains social media claim info and other
type Payload struct {
	Claim     Claim     `json:"claim"`
	IssuedAt  string    `json:"issuedAt"`
	Subject   PublicKey `json:"subject"`
	Issuer    PublicKey `json:"issuer"`
	ExpiresAt string    `json:"expiresAt"`
}

// Header describes the encryption types
type Header struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

// DecodedToken contains most of the profile information
type DecodedToken struct {
	Payload   Payload `json:"payload"`
	Signature string  `json:"signature"`
	Header    Header  `json:"header"`
}

// LegacyProfile models a zonefile that holds the full profile format
type LegacyProfile struct {
	Account []struct {
		Type       string `json:"@type"`
		Identifier string `json:"identifier"`
		Service    string `json:"service"`
	} `json:"account"`
	Avatar struct {
		URL string `json:"url"`
	} `json:"avatar"`
	Bio     string `json:"bio"`
	Bitcoin struct {
		Address string `json:"address"`
	} `json:"bitcoin"`
	Cover struct {
		URL string `json:"url"`
	} `json:"cover"`
	Facebook struct {
		Proof struct {
			URL string `json:"url"`
		} `json:"proof"`
		Username string `json:"username"`
	} `json:"facebook"`
	Github struct {
		Proof struct {
			URL string `json:"url"`
		} `json:"proof"`
		Username string `json:"username"`
	} `json:"github"`
	Linkedin struct {
		URL string `json:"url"`
	} `json:"linkedin"`
	Location struct {
		Formatted string `json:"formatted"`
	} `json:"location"`
	Name struct {
		Formatted string `json:"formatted"`
	} `json:"name"`
	Pgp struct {
		Fingerprint string `json:"fingerprint"`
		URL         string `json:"url"`
	} `json:"pgp"`
	Twitter struct {
		Proof struct {
			URL string `json:"url"`
		} `json:"proof"`
		Username string `json:"username"`
	} `json:"twitter"`
	V       string `json:"v"`
	Website string `json:"website"`
}
