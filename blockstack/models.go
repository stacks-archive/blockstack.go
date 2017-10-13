package blockstack

import (
	"encoding/json"
	"log"
)

// GetInfoResult is the go represenation of the getinfo rpc method
type GetInfoResult struct {
	ServerAlive        bool   `json:"server_alive"`
	ZonefileCount      int    `json:"zonefile_count"`
	LastBlockProcessed int    `json:"last_block_processed"`
	Indexing           bool   `json:"indexing"`
	Consensus          string `json:"consensus"`
	LastBlockSeen      int    `json:"last_block_seen"`
	ServerVersion      string `json:"server_version"`
}

// JSON returns the JSON representation of GetInfoResult
func (r GetInfoResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// Transaction models a Bitcoin Transaction for various structs here
type Transaction struct {
	ValueHash            string  `json:"value_hash"`
	LastRenewed          int     `json:"last_renewed"`
	LastCreationOp       string  `json:"last_creation_op"`
	Revoked              bool    `json:"revoked"`
	SenderPubkey         string  `json:"sender_pubkey"`
	BlockNumber          int     `json:"block_number"`
	ConsensusHash        string  `json:"consensus_hash"`
	NamespaceBlockNumber int     `json:"namespace_block_number"`
	PreorderBlockNumber  int     `json:"preorder_block_number"`
	Vtxindex             int     `json:"vtxindex"`
	Op                   string  `json:"op"`
	Txid                 string  `json:"txid"`
	Importer             string  `json:"importer"`
	Opcode               string  `json:"opcode"`
	OpFee                float64 `json:"op_fee"`
	Address              string  `json:"address"`
	PreorderHash         string  `json:"preorder_hash"`
	ImporterAddress      string  `json:"importer_address"`
	FirstRegistered      int     `json:"first_registered"`
	TransferSendBlockID  int     `json:"transfer_send_block_id"`
	Sender               string  `json:"sender"`
}

// JSON returns the JSON representation of Transaction
func (r Transaction) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetNameBlockchainRecordResult needs testing rpc method get_name_blockchain_record
type GetNameBlockchainRecordResult struct {
	Status bool `json:"status"`
	Record struct {
		RenewalDeadline      int                   `json:"renewal_deadline"`
		BlockNumber          int                   `json:"block_number"`
		LastCreationOp       string                `json:"last_creation_op"`
		NamespaceID          string                `json:"namespace_id"`
		NamespaceBlockNumber int                   `json:"namespace_block_number"`
		ExpireBlock          int                   `json:"expire_block"`
		Address              string                `json:"address"`
		ImporterAddress      string                `json:"importer_address"`
		Expired              bool                  `json:"expired"`
		TransferSendBlockID  interface{}           `json:"transfer_send_block_id"`
		Sender               string                `json:"sender"`
		ValueHash            string                `json:"value_hash"`
		LastRenewed          int                   `json:"last_renewed"`
		Name                 string                `json:"name"`
		Revoked              bool                  `json:"revoked"`
		ConsensusHash        string                `json:"consensus_hash"`
		PreorderBlockNumber  int                   `json:"preorder_block_number"`
		Txid                 string                `json:"txid"`
		Importer             string                `json:"importer"`
		NameHash128          string                `json:"name_hash128"`
		Opcode               string                `json:"opcode"`
		OpFee                int                   `json:"op_fee"`
		SenderPubkey         string                `json:"sender_pubkey"`
		PreorderHash         string                `json:"preorder_hash"`
		History              map[int][]Transaction `json:"history"`
	}
	Lastblock int  `json:"lastblock"`
	Indexing  bool `json:"indexing"`
}

// JSON returns the JSON representation of GetNameBlockchainRecordResult
func (r GetNameBlockchainRecordResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// PingResult is the go represenation of the ping rpc method
type PingResult struct {
	Status string `json:"status"`
}

// JSON returns the JSON representation of PingResult
func (r PingResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetNameHistoryBlocksResult is the go represenation of the get_name_history_blocks method
type GetNameHistoryBlocksResult struct {
	Status        bool  `json:"status"`
	Lastblock     int   `json:"lastblock"`
	Indexing      bool  `json:"indexing"`
	HistoryBlocks []int `json:"history_blocks"`
}

// JSON returns the JSON representation of GetNameHistoryBlocksResult
func (r GetNameHistoryBlocksResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetNameAtResult is the go represenation of the get_name_at method
type GetNameAtResult struct {
	Status  bool `json:"status"`
	Records []struct {
		BlockNumber          int         `json:"block_number"`
		NamespaceID          string      `json:"namespace_id"`
		ImporterAddress      string      `json:"importer_address"`
		ValueHash            string      `json:"value_hash"`
		ConsensusHash        string      `json:"consensus_hash"`
		Txid                 string      `json:"txid"`
		Importer             string      `json:"importer"`
		NameHash128          string      `json:"name_hash128"`
		TransferSendBlockID  interface{} `json:"transfer_send_block_id"`
		PreorderHash         string      `json:"preorder_hash"`
		FirstRegistered      int         `json:"first_registered"`
		LastCreationOp       string      `json:"last_creation_op"`
		Name                 string      `json:"name"`
		NamespaceBlockNumber int         `json:"namespace_block_number"`
		Address              string      `json:"address"`
		OpFee                int         `json:"op_fee"`
		Revoked              bool        `json:"revoked"`
		LastRenewed          int         `json:"last_renewed"`
		Sender               string      `json:"sender"`
		SenderPubkey         string      `json:"sender_pubkey"`
		PreorderBlockNumber  int         `json:"preorder_block_number"`
		Opcode               string      `json:"opcode"`
		Op                   string      `json:"op"`
		Vtxindex             int         `json:"vtxindex"`
	} `json:"records"`
	Lastblock int  `json:"lastblock"`
	Indexing  bool `json:"indexing"`
}

// JSON returns the JSON representation of GetNameAtResult
func (r GetNameAtResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetNamesOwnedByAddressResult is the go represenation of the get_names_owned_by_address method
type GetNamesOwnedByAddressResult struct {
	Status    bool     `json:"status"`
	Lastblock int      `json:"lastblock"`
	Indexing  bool     `json:"indexing"`
	Names     []string `json:"names"`
}

// JSON returns the JSON representation of GetNamesOwnedByAddressResult
func (r GetNamesOwnedByAddressResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetNameCostResult is the go represenation of the get_name_cost method
type GetNameCostResult struct {
	Status    bool `json:"status"`
	Lastblock int  `json:"lastblock"`
	Indexing  bool `json:"indexing"`
	Satoshis  int  `json:"satoshis"`
}

// JSON returns the JSON representation of GetNameCostResult
func (r GetNameCostResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetNamespaceCostResult is the go represenation of the get_name_cost method
type GetNamespaceCostResult struct {
	Status    bool `json:"status"`
	Lastblock int  `json:"lastblock"`
	Indexing  bool `json:"indexing"`
	Satoshis  int  `json:"satoshis"`
}

// JSON returns the JSON representation of GetNamespaceCostResult
func (r GetNamespaceCostResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetNumNamesResult is the go represenation of the get_num_names method
type GetNumNamesResult struct {
	Status    bool `json:"status"`
	Count     int  `json:"count"`
	Lastblock int  `json:"lastblock"`
	Indexing  bool `json:"indexing"`
}

// JSON returns the JSON representation of GetNumNamesResult
func (r GetNumNamesResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetAllNamesResult is the go represenation of the get_all_names method
type GetAllNamesResult struct {
	Status    bool     `json:"status"`
	Lastblock int      `json:"lastblock"`
	Indexing  bool     `json:"indexing"`
	Names     []string `json:"names"`
}

// JSON returns the JSON representation of GetAllNamesResult
func (r GetAllNamesResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetAllNamespacesResult is the go represenation of the get_all_namespaces method
type GetAllNamespacesResult struct {
	Status     bool     `json:"status"`
	Lastblock  int      `json:"lastblock"`
	Indexing   bool     `json:"indexing"`
	Namespaces []string `json:"namespaces"`
}

// JSON returns the JSON representation of GetAllNamespacesResult
func (r GetAllNamespacesResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetNumNamesInNamespaceResult is the go represenation fo the get_num_names_in_namespace method
type GetNumNamesInNamespaceResult struct {
	Status    bool `json:"status"`
	Count     int  `json:"count"`
	Lastblock int  `json:"lastblock"`
	Indexing  bool `json:"indexing"`
}

// JSON returns the JSON representation of GetNumNamesInNamespaceResult
func (r GetNumNamesInNamespaceResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetNamesInNamespaceResult is the go represenation of the get_names_in_namespace rpc method
type GetNamesInNamespaceResult struct {
	Status    bool     `json:"status"`
	Lastblock int      `json:"lastblock"`
	Indexing  bool     `json:"indexing"`
	Names     []string `json:"names"`
}

// JSON returns the JSON representation of GetNamesInNamespaceResult
func (r GetNamesInNamespaceResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetConsensusAtResult is the go represenation of the get_consensus_at method
type GetConsensusAtResult struct {
	Status    bool   `json:"status"`
	Consensus string `json:"consensus"`
	Lastblock int    `json:"lastblock"`
	Indexing  bool   `json:"indexing"`
}

// JSON returns the JSON representation of GetConsensusAtResult
func (r GetConsensusAtResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetBlockFromConsensusResult is the go represenation of the get_block_from_consensus method
type GetBlockFromConsensusResult struct {
	Status    bool `json:"status"`
	Lastblock int  `json:"lastblock"`
	Indexing  bool `json:"indexing"`
	BlockID   int  `json:"block_id"`
}

// JSON returns the JSON representation of GetBlockFromConsensusResult
func (r GetBlockFromConsensusResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetZonefilesByBlockResult is the go represenation of the get_zonefiles_by_block rpc method
type GetZonefilesByBlockResult struct {
	Status       bool `json:"status"`
	Lastblock    int  `json:"lastblock"`
	Indexing     bool `json:"indexing"`
	ZonefileInfo []struct {
		Txid         string `json:"txid"`
		Name         string `json:"name"`
		ZonefileHash string `json:"zonefile_hash"`
		BlockHeight  int    `json:"block_height"`
	} `json:"zonefile_info"`
}

// JSON returns the JSON representation of GetZonefilesByBlockResult
func (r GetZonefilesByBlockResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetAtlasPeersResult is the go represenation of the get_atlas_peers rpc method
type GetAtlasPeersResult struct {
	Status    bool     `json:"status"`
	Lastblock int      `json:"lastblock"`
	Indexing  bool     `json:"indexing"`
	Peers     []string `json:"peers"`
}

// JSON returns the JSON representation of GetAtlasPeersResult
func (r GetAtlasPeersResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetZonefileInventoryResult is the go represenation of the get_zonefile_inventory rpc method
type GetZonefileInventoryResult struct {
	Status    bool   `json:"status"`
	Lastblock int    `json:"lastblock"`
	Indexing  bool   `json:"indexing"`
	Inv       string `json:"inv"`
}

// JSON returns the JSON representation of GetZonefileInventoryResult
func (r GetZonefileInventoryResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetNameOpsHashAtResult is the go represenation of the get_nameops_hash_at rpc method
type GetNameOpsHashAtResult struct {
	Status    bool   `json:"status"`
	Lastblock int    `json:"lastblock"`
	Indexing  bool   `json:"indexing"`
	OpsHash   string `json:"ops_hash"`
}

// JSON returns the JSON representation of GetNameOpsHashAtResult
func (r GetNameOpsHashAtResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// NamespaceTransaction is used to decode the get_namespace_blockchain_record return
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

// JSON returns the JSON representation of Transaction
func (r NamespaceTransaction) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetNamespaceBlockchainRecordResult is the go represenation of the get_namespace_blockchain_record rpc method
type GetNamespaceBlockchainRecordResult struct {
	Status bool `json:"status"`
	Record struct {
		BlockNumber      int                            `json:"block_number"`
		NonalphaDiscount int                            `json:"nonalpha_discount"`
		NamespaceID      string                         `json:"namespace_id"`
		RevealBlock      int                            `json:"reveal_block"`
		Buckets          []int                          `json:"buckets"`
		Base             int                            `json:"base"`
		Address          string                         `json:"address"`
		Ready            bool                           `json:"ready"`
		Lifetime         int                            `json:"lifetime"`
		Recipient        string                         `json:"recipient"`
		OpFee            int64                          `json:"op_fee"`
		Sender           string                         `json:"sender"`
		RecipientAddress string                         `json:"recipient_address"`
		SenderPubkey     string                         `json:"sender_pubkey"`
		ReadyBlock       int                            `json:"ready_block"`
		Coeff            int                            `json:"coeff"`
		Txid             string                         `json:"txid"`
		Version          int                            `json:"version"`
		Opcode           string                         `json:"opcode"`
		NoVowelDiscount  int                            `json:"no_vowel_discount"`
		PreorderHash     string                         `json:"preorder_hash"`
		History          map[int][]NamespaceTransaction `json:"history"`
		Vtxindex         int                            `json:"vtxindex"`
		Op               string                         `json:"op"`
	} `json:"record"`
	Lastblock int  `json:"lastblock"`
	Indexing  bool `json:"indexing"`
}

// JSON returns the JSON representation of GetNameOpsHashAtResult
func (r GetNamespaceBlockchainRecordResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}
