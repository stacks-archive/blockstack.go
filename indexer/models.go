package indexer

import (
	"encoding/json"
	"log"
)

// Profile is an interface to wrap Schema.org profiles and
// LegacyProfiles
type Profile interface {
	JSON() string
	Validate() bool
}

// SOProfile models a Schema.org profile
type SOProfile struct {
	Token           string       `json:"token"`
	ParentPublicKey string       `json:"parentPublicKey"`
	Encrypted       bool         `json:"encrypted"`
	DecodedToken    DecodedToken `json:"decodedToken"`
}

// JSON statisfiles the Profile interface
func (p SOProfile) JSON() string {
	byt, err := json.Marshal(p)
	if err != nil {
		// We Unmarshal this JSON so there should be no errors Marshaling
		log.Fatal(logPrefix, "error marshalling profile", err)
	}
	return string(byt)
}

// DecodedToken contains most of the profile information
type DecodedToken struct {
	Payload   Payload `json:"payload"`
	Signature string  `json:"signature"`
	Header    Header  `json:"header"`
}

// Payload contains social media claim info and other
type Payload struct {
	Claim     Claim     `json:"claim"`
	IssuedAt  string    `json:"issuedAt"`
	Subject   PublicKey `json:"subject"`
	Issuer    PublicKey `json:"issuer"`
	ExpiresAt string    `json:"expiresAt"`
}

// Account models a social media proof
// TODO: Write method on Account to check proof
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

// Header describes the encryption types
type Header struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

// LegacyProfile models a zonefile that holds the full profile format
type LegacyProfile struct {
	Account  []map[string]string `json:"account"`
	Avatar   map[string]string   `json:"avatar"`
	Bio      string              `json:"bio"`
	Bitcoin  map[string]string   `json:"bitcoin"`
	Cover    map[string]string   `json:"cover"`
	Facebook LProof              `json:"facebook"`
	Github   LProof              `json:"github"`
	Linkedin map[string]string   `json:"linkedin"`
	Location map[string]string   `json:"location"`
	Name     map[string]string   `json:"name"`
	Pgp      map[string]string   `json:"pgp"`
	Twitter  LProof              `json:"twitter"`
	V        string              `json:"v"`
	Website  string              `json:"website"`
}

// LProof is a legacy proof
type LProof struct {
	Proof    map[string]string `json:"proof"`
	Username string            `json:"username"`
}

// JSON statisfiles the Profile interface
func (p LegacyProfile) JSON() string {
	byt, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(byt)
}

// type LegacyProfile struct {
// 	Account []struct {
// 		Type       string `json:"@type"`
// 		Identifier string `json:"identifier"`
// 		Service    string `json:"service"`
// 	} `json:"account"`
// 	Avatar struct {
// 		URL string `json:"url"`
// 	} `json:"avatar"`
// 	Bio     string `json:"bio"`
// 	Bitcoin struct {
// 		Address string `json:"address"`
// 	} `json:"bitcoin"`
// 	Cover struct {
// 		URL string `json:"url"`
// 	} `json:"cover"`
// 	Facebook struct {
// 		Proof struct {
// 			URL string `json:"url"`
// 		} `json:"proof"`
// 		Username string `json:"username"`
// 	} `json:"facebook"`
// 	Github struct {
// 		Proof struct {
// 			URL string `json:"url"`
// 		} `json:"proof"`
// 		Username string `json:"username"`
// 	} `json:"github"`
// 	Linkedin struct {
// 		URL string `json:"url"`
// 	} `json:"linkedin"`
// 	Location struct {
// 		Formatted string `json:"formatted"`
// 	} `json:"location"`
// 	Name struct {
// 		Formatted string `json:"formatted"`
// 	} `json:"name"`
// 	Pgp struct {
// 		Fingerprint string `json:"fingerprint"`
// 		URL         string `json:"url"`
// 	} `json:"pgp"`
// 	Twitter struct {
// 		Proof struct {
// 			URL string `json:"url"`
// 		} `json:"proof"`
// 		Username string `json:"username"`
// 	} `json:"twitter"`
// 	V       string `json:"v"`
// 	Website string `json:"website"`
// }
