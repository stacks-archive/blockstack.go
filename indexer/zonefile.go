package indexer

import (
	"encoding/json"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
	"strings"

	"github.com/miekg/dns"
)

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
