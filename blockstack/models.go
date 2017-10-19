package blockstack

import (
	"encoding/base64"
	"encoding/json"
	"log"
)

// StartBlock is the first block on the bitcoin blockchain with blockstack transactions
const StartBlock = 373601

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

// PrettyJSON returns the Pretty Printed JSON representation of GetInfoResult
func (r GetInfoResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of Transaction
func (r Transaction) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of GetNameBlockchainRecordResult
func (r GetNameBlockchainRecordResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// LastTx returns the last transcation from the history
func (r GetNameBlockchainRecordResult) LastTx() Transaction {
	var tx int
	for block := range r.Record.History {
		if block > tx {
			tx = block
		}
	}
	return r.Record.History[tx][0]
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

// PrettyJSON returns the Pretty Printed JSON representation of PingResult
func (r PingResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of GetNameHistoryBlocksResult
func (r GetNameHistoryBlocksResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of GetNameAtResult
func (r GetNameAtResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of GetNamesOwnedByAddressResult
func (r GetNamesOwnedByAddressResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of GetNameCostResult
func (r GetNameCostResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of GetNamespaceCostResult
func (r GetNamespaceCostResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of GetAllNamesResult
func (r GetAllNamesResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of GetAllNamespacesResult
func (r GetAllNamespacesResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of GetNamesInNamespaceResult
func (r GetNamesInNamespaceResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of GetConsensusAtResult
func (r GetConsensusAtResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of GetBlockFromConsensusResult
func (r GetBlockFromConsensusResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// ZonefileHashResult is a go represenation of a zonefile_hash_result
type ZonefileHashResult struct {
	Txid         string `json:"txid"`
	Name         string `json:"name"`
	ZonefileHash string `json:"zonefile_hash"`
	BlockHeight  int    `json:"block_height"`
}

// ZonefileHashResults is a collection of ZonefileHashResult
type ZonefileHashResults []ZonefileHashResult

// Zonefiles returns zonefiles from a ZonefileHashResults
func (zfhr ZonefileHashResults) Zonefiles() []string {
	var out []string
	for _, zfh := range zfhr {
		out = append(out, zfh.ZonefileHash)
	}
	return out
}

// LatestZonefileHash returns the latest zonefile from a batch for a given zonefileHash
func (zfhr ZonefileHashResults) LatestZonefileHash(zonefileHash string) ZonefileHashResult {
	var out ZonefileHashResult
	for _, zfh := range zfhr {
		if zfh.ZonefileHash == zonefileHash && zfh.BlockHeight > out.BlockHeight {
			out = zfh
		}
	}
	return out
}

// GetZonefilesByBlockResult is the go represenation of the get_zonefiles_by_block rpc method
type GetZonefilesByBlockResult struct {
	Status       bool                 `json:"status"`
	Lastblock    int                  `json:"lastblock"`
	Indexing     bool                 `json:"indexing"`
	ZonefileInfo []ZonefileHashResult `json:"zonefile_info"`
}

// JSON returns the JSON representation of GetZonefilesByBlockResult
func (r GetZonefilesByBlockResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// PrettyJSON returns the Pretty Printed JSON representation of GetZonefilesByBlockResult
func (r GetZonefilesByBlockResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// Zonefiles is an affordance to return the zonefile hashes in []string
func (r GetZonefilesByBlockResult) Zonefiles() []string {
	var out []string
	for _, zfh := range r.ZonefileInfo {
		out = append(out, zfh.ZonefileHash)
	}
	return out
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

// PrettyJSON returns the Pretty Printed JSON representation of GetAtlasPeersResult
func (r GetAtlasPeersResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of GetZonefileInventoryResult
func (r GetZonefileInventoryResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of GetNameOpsHashAtResult
func (r GetNameOpsHashAtResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// PrettyJSON returns the Pretty Printed JSON representation of Transaction
func (r NamespaceTransaction) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
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

// JSON returns the JSON representation of GetNamespaceBlockchainRecordResult
func (r GetNamespaceBlockchainRecordResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// PrettyJSON returns the Pretty Printed JSON representation of GetNamespaceBlockchainRecordResult
func (r GetNamespaceBlockchainRecordResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetZonefilesResult is the go represenation of the get_zonfiles rpc method
type GetZonefilesResult struct {
	Status    bool              `json:"status"`
	Lastblock int               `json:"lastblock"`
	Indexing  bool              `json:"indexing"`
	Zonefiles map[string]string `json:"zonefiles"`
}

// Decode is an affordance that returns the results in map[string]string
func (r GetZonefilesResult) Decode() map[string]string {
	out := make(map[string]string)
	for k := range r.Zonefiles {
		dec, err := base64.StdEncoding.DecodeString(r.Zonefiles[k])
		if err != nil {
			log.Fatal(err)
		}
		out[k] = string(dec)
	}
	return out
}

// JSON returns the JSON representation of GetZonefilesResult
func (r GetZonefilesResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// PrettyJSON returns the Pretty Printed JSON representation of GetZonefilesResult
func (r GetZonefilesResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetOpHistoryRowsResult is the go represenation of the get_zonfiles rpc method
// NOTE: The HistoryData field is a Transaction, but the JSON is returned from blockstack-core escaped
// TODO: Get this fixed server-side
type GetOpHistoryRowsResult struct {
	Status      bool `json:"status"`
	HistoryRows []struct {
		BlockID     int    `json:"block_id"`
		Op          string `json:"op"`
		HistoryID   string `json:"history_id"`
		HistoryData string `json:"history_data"`
		Vtxindex    int    `json:"vtxindex"`
		Txid        string `json:"txid"`
	} `json:"history_rows"`
	Lastblock int  `json:"lastblock"`
	Indexing  bool `json:"indexing"`
}

// JSON returns the JSON representation of GetOpHistoryRowsResult
func (r GetOpHistoryRowsResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// PrettyJSON returns the Pretty Printed JSON representation of GetOpHistoryRowsResult
func (r GetOpHistoryRowsResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// CountResult is the go represenation of the
// get_num_names, get_num_names_in_namespace, get_num_nameops_affected_at, get_num_op_history_rows
// rpc methods
type CountResult struct {
	Status    bool `json:"status"`
	Count     int  `json:"count"`
	Lastblock int  `json:"lastblock"`
	Indexing  bool `json:"indexing"`
}

// JSON returns the JSON representation of CountResult
func (r CountResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// PrettyJSON returns the Pretty Printed JSON representation of CountResult
func (r CountResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetNameOpsAffectedAtResult is the go represenation of the get_nameops_affected_at rpc method
type GetNameOpsAffectedAtResult struct {
	Status    bool          `json:"status"`
	Nameops   []Transaction `json:"nameops"`
	Lastblock int           `json:"lastblock"`
	Indexing  bool          `json:"indexing"`
}

// JSON returns the JSON representation of GetNameOpsAffectedAtResult
func (r GetNameOpsAffectedAtResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// PrettyJSON returns the Pretty Printed JSON representation of GetNameOpsAffectedAtResult
func (r GetNameOpsAffectedAtResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// GetConsensusHashesResult is the go representation of the get_consensus_hashes rpc method
type GetConsensusHashesResult struct {
	Status          bool           `json:"status"`
	ConsensusHashes map[int]string `json:"consensus_hashes"`
	Lastblock       int            `json:"lastblock"`
	Indexing        bool           `json:"indexing"`
}

// JSON returns the JSON representation of GetConsensusHashesResult
func (r GetConsensusHashesResult) JSON() string {
	byt, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}

// PrettyJSON returns the Pretty Printed JSON representation of GetConsensusHashesResult
func (r GetConsensusHashesResult) PrettyJSON() string {
	byt, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(byt)
}
