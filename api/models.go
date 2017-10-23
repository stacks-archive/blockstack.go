package api

import (
	"encoding/json"
	// "fmt"
	"log"

	"github.com/blockstack/go-blockstack/indexer"
)

// V1GetNameResponse is the response for the /v1/names/:name
type V1GetNameResponse struct {
	Address      string `json:"address"`
	Blockchain   string `json:"blockchain"`
	ExpireBlock  int    `json:"expire_block"`
	LastTxid     string `json:"last_txid"`
	Status       string `json:"status"`
	Zonefile     string `json:"zonefile"`
	ZonefileHash string `json:"zonefile_hash"`
}

// V1GetNameResponse is the response for the /v1/names/:name
type V1GetNameNoZResponse struct {
	Address      string            `json:"address"`
	Blockchain   string            `json:"blockchain"`
	ExpireBlock  int               `json:"expire_block"`
	LastTxid     string            `json:"last_txid"`
	Status       string            `json:"status"`
	Zonefile     map[string]string `json:"zonefile"`
	ZonefileHash string            `json:"zonefile_hash"`
}

// JSON proves a JSON output for ResponseWritert for ResponseWriter
func (r V1GetNameResponse) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return byt
}

// JSON proves a JSON output for ResponseWritert for ResponseWriter
func (r V1GetNameNoZResponse) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return byt
}

// Transaction models a bitcoin transaction
type Transaction struct {
	Address              string  `json:"address"`
	BlockNumber          int     `json:"block_number"`
	ConsensusHash        string  `json:"consensus_hash"`
	FirstRegistered      int     `json:"first_registered"`
	Importer             string  `json:"importer"`
	ImporterAddress      string  `json:"importer_address"`
	LastCreationOp       string  `json:"last_creation_op"`
	LastRenewed          int     `json:"last_renewed"`
	Name                 string  `json:"name"`
	BurnAddress          string  `json:"burn_address"`
	HistorySnapshot      bool    `json:"history_snapshot"`
	NameHash128          string  `json:"name_hash128"`
	NamespaceBlockNumber int     `json:"namespace_block_number"`
	NamespaceID          string  `json:"namespace_id"`
	Op                   string  `json:"op"`
	OpFee                float64 `json:"op_fee"`
	Opcode               string  `json:"opcode"`
	PreorderBlockNumber  int     `json:"preorder_block_number"`
	PreorderHash         string  `json:"preorder_hash"`
	Revoked              bool    `json:"revoked"`
	Sender               string  `json:"sender"`
	SenderPubkey         string  `json:"sender_pubkey"`
	TransferSendBlockID  int     `json:"transfer_send_block_id"`
	Txid                 string  `json:"txid"`
	ValueHash            string  `json:"value_hash"`
	Vtxindex             int     `json:"vtxindex"`
}

// JSON proves a JSON output for ResponseWriter
func (r Transaction) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return byt
}

// V1GetNameHistoryResponse holds the response for the /v1/names/{name}/history route
// NOTE: result of the get_name_blockchain_record rpc call
type V1GetNameHistoryResponse map[int][]Transaction

// JSON proves a JSON output for ResponseWriter
func (r V1GetNameHistoryResponse) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return byt
}

// V1GetNamesInNamespaceResponse holds the response for the /v1/namespaces/{namespace}/names?page={page} route
// NOTE: result of the get_names_in_namespace rpc call
type V1GetNamesInNamespaceResponse []string

// JSON proves a JSON output for ResponseWriter
func (r V1GetNamesInNamespaceResponse) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return byt
}

// V2GetUserProfileResponse holds the response for the /v2/users/{name} route
// NOTE: This is the big one
// type V2GetUserProfileResponse struct{}

// JSON proves a JSON output for ResponseWriter
func (r V2GetUserProfileResponse) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return byt
}

type V2GetUserProfileResponse map[string]V2GetUserProfile

type V2GetUserProfile struct {
	Expired       string          `json:"expired"`
	Profile       indexer.Profile `json:"profile"`
	Verifications []struct {
		Identifier string `json:"identifier"`
		ProofURL   string `json:"proof_url"`
		Service    string `json:"service"`
		Valid      bool   `json:"valid"`
	} `json:"verifications"`
	ZoneFile indexer.Zonefile `json:"zone_file"`
}

// V1GetNameOpsAtHeightResponse holds the response for the /v1/blockchains/bitcoin/operations/:blockHeight route
// NOTE: List of all transactions affected at a blockHeight
type V1GetNameOpsAtHeightResponse []struct {
	Address              string `json:"address"`
	BlockNumber          int    `json:"block_number"`
	ConsensusHash        string `json:"consensus_hash"`
	FirstRegistered      int    `json:"first_registered"`
	History              map[int][]Transaction
	Importer             interface{} `json:"importer"`
	ImporterAddress      interface{} `json:"importer_address"`
	KeepData             bool        `json:"keep_data"`
	LastCreationOp       string      `json:"last_creation_op"`
	LastRenewed          int         `json:"last_renewed"`
	Name                 string      `json:"name"`
	NameHash128          string      `json:"name_hash128"`
	NamespaceBlockNumber int         `json:"namespace_block_number"`
	NamespaceID          string      `json:"namespace_id"`
	Op                   string      `json:"op"`
	OpFee                int         `json:"op_fee"`
	Opcode               string      `json:"opcode"`
	PreorderBlockNumber  int         `json:"preorder_block_number"`
	PreorderHash         string      `json:"preorder_hash"`
	Recipient            string      `json:"recipient"`
	RecipientAddress     string      `json:"recipient_address"`
	Revoked              bool        `json:"revoked"`
	Sender               string      `json:"sender"`
	SenderPubkey         interface{} `json:"sender_pubkey"`
	TransferSendBlockID  int         `json:"transfer_send_block_id"`
	Txid                 string      `json:"txid"`
	ValueHash            string      `json:"value_hash"`
	Vtxindex             int         `json:"vtxindex"`
}

// JSON proves a JSON output for ResponseWriter
func (r V1GetNameOpsAtHeightResponse) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return byt
}

// V1GetNamesOwnedByAddressResponse holds the response for the /v1/addresses/bitcoin/:address route
// NOTE: result of get_names_owned_by_address rpc call with a slightly different format
type V1GetNamesOwnedByAddressResponse struct {
	Names []string `json:"names"`
}

// JSON proves a JSON output for ResponseWriter
func (r V1GetNamesOwnedByAddressResponse) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return byt
}

// V1GetZonefileResponse holds the response for the /v1/names/{name}/zonefile route
// NOTE: returns just the plaintext zonefile
type V1GetZonefileResponse struct {
	Zonefile string `json:"zonefile"`
}

// JSON proves a JSON output for ResponseWriter
func (r V1GetZonefileResponse) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return byt
}

// NamespaceTransaction models and individual namespace transaction
type NamespaceTransaction struct {
	Address          string `json:"address"`
	Base             int    `json:"base"`
	BlockNumber      int    `json:"block_number"`
	Buckets          []int  `json:"buckets"`
	BurnAddress      string `json:"burn_address"`
	Coeff            int    `json:"coeff"`
	ConsensusHash    string `json:"consensus_hash"`
	HistorySnapshot  bool   `json:"history_snapshot"`
	Lifetime         int    `json:"lifetime"`
	NamespaceID      string `json:"namespace_id"`
	NonalphaDiscount int    `json:"nonalpha_discount"`
	NoVowelDiscount  int    `json:"no_vowel_discount"`
	Op               string `json:"op"`
	Opcode           string `json:"opcode"`
	OpFee            int64  `json:"op_fee"`
	PreorderHash     string `json:"preorder_hash"`
	Recipient        string `json:"recipient"`
	RecipientAddress string `json:"recipient_address"`
	RevealBlock      int    `json:"reveal_block"`
	Sender           string `json:"sender"`
	SenderPubkey     string `json:"sender_pubkey"`
	Txid             string `json:"txid"`
	Version          int    `json:"version"`
	Vtxindex         int    `json:"vtxindex"`
}

// V1GetNamespaceBlockchainRecordResponse holds the response for the /v1/namespaces/{namespace} route
// NOTE: returns result of get_namespace_blockchain_record rpc call
type V1GetNamespaceBlockchainRecordResponse struct {
	Address          string                       `json:"address"`
	Base             int                          `json:"base"`
	BlockNumber      int                          `json:"block_number"`
	Buckets          []int                        `json:"buckets"`
	Coeff            int                          `json:"coeff"`
	History          map[int]NamespaceTransaction `json:"history"`
	Lifetime         int                          `json:"lifetime"`
	NamespaceID      string                       `json:"namespace_id"`
	NoVowelDiscount  int                          `json:"no_vowel_discount"`
	NonalphaDiscount int                          `json:"nonalpha_discount"`
	Op               string                       `json:"op"`
	OpFee            int64                        `json:"op_fee"`
	PreorderHash     string                       `json:"preorder_hash"`
	Ready            bool                         `json:"ready"`
	ReadyBlock       int                          `json:"ready_block"`
	Recipient        string                       `json:"recipient"`
	RecipientAddress string                       `json:"recipient_address"`
	RevealBlock      int                          `json:"reveal_block"`
	Sender           string                       `json:"sender"`
	SenderPubkey     string                       `json:"sender_pubkey"`
	Txid             string                       `json:"txid"`
	Version          int                          `json:"version"`
	Vtxindex         int                          `json:"vtxindex"`
}

// JSON proves a JSON output for ResponseWriter
func (r V1GetNamespaceBlockchainRecordResponse) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return byt
}

// V1GetNamespacesResponse holds the response for the /v1/namespaces route
// NOTE: returns result of get_all_namespaces rpc call
type V1GetNamespacesResponse []string

// JSON proves a JSON output for ResponseWriter
func (r V1GetNamespacesResponse) JSON() []byte {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return byt
}
