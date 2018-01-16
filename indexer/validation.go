package indexer

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"time"

	"gopkg.in/square/go-jose.v1"
)

func NewDecodedProfileToken(pt string) *DecodedProfileToken {
	object, err := jose.ParseSigned(pt)
	if err != nil {
		log.Println(err)
	}
	dpt := &DecodedProfileToken{}
	err = json.Unmarshal([]byte(object.FullSerialize()), dpt)
	if err != nil {
		log.Println(err)
	}
	return dpt
}

type DecodedProfileToken struct {
	Payload   string `json:"payload"`
	Protected string `json:"protected"`
	Signature string `json:"signature"`
}

func (dpt *DecodedProfileToken) DecodedPayload() *DecodedProfileTokenPayload {
	data, err := base64.RawURLEncoding.DecodeString(dpt.Payload)
	if err != nil {
		log.Println(err)
	}
	prof := &DecodedProfileTokenPayload{}
	err = json.Unmarshal(data, prof)
	if err != nil {
		log.Println(err)
	}
	return prof
}

func (so SOProfile) Validate() bool {
	dptp := NewDecodedProfileToken(so.Token).DecodedPayload()
	if dptp.Issuer.PublicKey != so.DecodedToken.Payload.Issuer.PublicKey {
		log.Println("[validation] New Profile failed validation for issuer", dptp.Issuer.PublicKey, so.DecodedToken.Payload.Issuer.PublicKey)
		return false
	}
	if dptp.Subject.PublicKey != so.DecodedToken.Payload.Subject.PublicKey {
		log.Println("[validation] New Profile failed validation for subject", dptp.Subject.PublicKey, so.DecodedToken.Payload.Subject.PublicKey)
		return false
	}
	return true
}

func (lp LegacyProfile) Validate() bool {
	return true
}

type DecodedProfileTokenPayload struct {
	IssuedAt  string `json:"issuedAt,omitempty"`
	Claim     Claim
	ExpiresAt string `json:"expiresAt,omitempty"`
	Issuer    PublicKey
	Subject   PublicKey
	Jti       string    `json:"jti,omitempty"`
	Iat       time.Time `json:"iat,omitempty"`
	Exp       time.Time `json:"exp,omitempty"`
}
