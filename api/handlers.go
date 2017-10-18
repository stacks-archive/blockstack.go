package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/blockstack/go-blockstack/blockstack"
	"github.com/blockstack/go-blockstack/indexer"
)

const (
	logPrefix = "[api]"
)

// Handlers is a collection of Hanlder
type Handlers struct {
	Client  *blockstack.Client
	Indexer *indexer.Indexer

	lastBlock int
}

// NewHandlers creates the Handlers struct where all the handlers are defined.
// It is defined this way so database connections and other clients
// can be shared between handler methods easily
func NewHandlers(conf blockstack.ServerConfig) *Handlers {
	h := &Handlers{
		Client: blockstack.NewClient(conf),
	}
	res, err := h.Client.GetInfo()
	if err != nil {
		log.Fatalf("Failed to contact blockstack-core node: %v", err)
	}
	h.lastBlock = res.LastBlockSeen
	return h
}

func jsonKV(k, v string) []byte {
	ret, _ := json.Marshal(map[string]string{k: v})
	return ret
}

// V1GetNameHandler handles the /v1/names/{name} route
func (h *Handlers) V1GetNameHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	spl := strings.Split(name, ".")
	if len(spl) != 3 && len(spl) != 2 {
		w.Write(jsonKV("error", "invalid name"))
		return
	} else if spl[len(spl)-1] != "id" {
		w.Write(jsonKV("error", "invalid namespace"))
		return
	}
	nameDetails, err := h.Client.GetNameBlockchainRecord(name)
	if err != nil {
		er := strings.Split(err.Error(), ": ")[1]
		if er == "invalid name" {
			w.Write(jsonKV("error", er))
			return
		}
	}

	// if there are no name details then the name is available
	if !nameDetails.Status {
		w.Write(jsonKV("status", "available"))
		return
	}
	// divine the status of the name
	lastTx := nameDetails.LastTx()
	var status string
	if lastTx.Opcode == "NAME_PREORDER" {
		status = "pending"
	} else if nameDetails.Record.ExpireBlock > nameDetails.Lastblock {
		status = "expired"
	} else {
		status = "registered"
	}

	// If it is registered and there is a zonefile hash look that up
	if nameDetails.Status && nameDetails.Record.ValueHash != "" {
		zonefile, err := h.Client.GetZonefiles([]string{nameDetails.Record.ValueHash})
		if err != nil {
			log.Fatal(err)
		}

		out := V1GetNameResponse{
			Address:      nameDetails.Record.Address,
			Blockchain:   "bitcoin",
			ExpireBlock:  nameDetails.Record.ExpireBlock,
			LastTxid:     nameDetails.Record.Txid,
			Status:       status,
			ZonefileHash: nameDetails.Record.ValueHash,
			Zonefile:     zonefile.Decode()[nameDetails.Record.ValueHash],
		}
		w.Write(out.JSON())
		return
	} else if nameDetails.Status {
		out := V1GetNameNoZResponse{
			Address:      nameDetails.Record.Address,
			Blockchain:   "bitcoin",
			ExpireBlock:  nameDetails.Record.ExpireBlock,
			LastTxid:     nameDetails.Record.Txid,
			Status:       status,
			ZonefileHash: nameDetails.Record.ValueHash,
			Zonefile:     map[string]string{"error": "No zone file loaded"},
		}
		w.Write(out.JSON())
		return
	}
	w.Write(jsonKV("error", "slipped request"))
}

// V1GetNameHistoryHandler handles response for /v1/names/{name}/history
func (h *Handlers) V1GetNameHistoryHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	spl := strings.Split(name, ".")
	if len(spl) != 3 && len(spl) != 2 {
		w.Write(jsonKV("error", "invalid name"))
		return
	} else if spl[len(spl)-1] != "id" {
		w.Write(jsonKV("error", "invalid namespace"))
		return
	}
	res, err := h.Client.GetNameBlockchainRecord(name)
	if err != nil {
		er := strings.Split(err.Error(), ": ")[1]
		if er == "invalid name" {
			w.Write(jsonKV("error", er))
			return
		}
	}
	out := V1GetNameHistoryResponse{}
	for k := range res.Record.History {
		// TODO: Maybe check length here. We will see
		tx := res.Record.History[k][0]
		out[k] = []Transaction{Transaction{
			Address:              tx.Address,
			BlockNumber:          tx.BlockNumber,
			ConsensusHash:        tx.ConsensusHash,
			FirstRegistered:      tx.FirstRegistered,
			Importer:             tx.Importer,
			ImporterAddress:      tx.ImporterAddress,
			LastCreationOp:       tx.LastCreationOp,
			LastRenewed:          tx.LastRenewed,
			NamespaceBlockNumber: tx.NamespaceBlockNumber,
			Op:                   tx.Op,
			OpFee:                tx.OpFee,
			Opcode:               tx.Opcode,
			PreorderBlockNumber:  tx.PreorderBlockNumber,
			PreorderHash:         tx.PreorderHash,
			Revoked:              tx.Revoked,
			Sender:               tx.Sender,
			SenderPubkey:         tx.SenderPubkey,
			TransferSendBlockID:  tx.TransferSendBlockID,
			Txid:                 tx.Txid,
			ValueHash:            tx.ValueHash,
			Vtxindex:             tx.Vtxindex,
		}}
	}
	w.Write([]byte(out.JSON()))
}

// V1GetNamesInNamespaceHandler handles response for /v1/namespaces/{namespace}/names?page={page}
func (h *Handlers) V1GetNamesInNamespaceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := r.FormValue("page")
	pg, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		w.Write(jsonKV("error", "invalid integer for page"))
		return
	}
	res, err := h.Client.GetNamesInNamespace(vars["namespace"], (int(pg) * 100), 100)
	if err != nil {
		w.Write([]byte("[]"))
		return
	}
	out := V1GetNamesInNamespaceResponse(res.Names)
	w.Write(out.JSON())
}

// V2GetUserProfileHandler handles response for /v2/users/{name} route
// NOTE: The big one
func (h *Handlers) V2GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {}

// V1GetNameOpsAtHeightHandler handles response for /v1/blockchains/{blockchain}/operations/{blockHeight} route
func (h *Handlers) V1GetNameOpsAtHeightHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockHeight := vars["blockHeight"]
	blockchain := vars["blockchain"]
	bh, err := strconv.ParseInt(blockHeight, 10, 64)
	if err != nil {
		w.Write(jsonKV("error", "invalid integer for blockHeight"))
		return
	} else if blockchain != "bitcoin" {
		w.Write(jsonKV("error", "blockstack runs on the bitcoin blockchain"))
		return
	} else if bh < blockstack.StartBlock {
		w.Write(jsonKV("error", "invalid block height"))
		return
	}
	res, err := h.Client.GetNameOpsAffectedAt(int(bh), 0, 10)
	if err != nil {
		w.Write(jsonKV("error", err.Error()))
		return
	} else if len(res.Nameops) == 0 {
		w.Write([]byte("[]"))
		return
	}
	j, _ := json.Marshal(res.Nameops)
	w.Write(j)
}

// V1GetNamesOwnedByAddressHandler handles response for /v1/addresses/bitcoin/{address} route
func (h *Handlers) V1GetNamesOwnedByAddressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("address", vars["address"])
	w.Write([]byte("ok\n"))
}

// V1GetZonefileHandler handles response for /v1/names/{name}/zonefile route
func (h *Handlers) V1GetZonefileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("name", vars["name"])
	w.Write([]byte("ok\n"))
}

// V1GetNamespaceBlockchainRecordHandler handles response for /v1/namespaces/{namespace} route
// TODO: seeing some fishy differences between responses from core.blockstack.org and this.
// I'm copying over the data from the other transactions for the top level object
// and it looks like core.blockstack.org has data from some other transaction
func (h *Handlers) V1GetNamespaceBlockchainRecordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("NAMESPACE", vars)
	res, err := h.Client.GetNamespaceBlockchainRecord(vars["namespace"])
	if err != nil {
		// TODO: return error to client here and mention that theres a connection error to core node
		log.Fatal(err)
	}

	out := V1GetNamespaceBlockchainRecordResponse{
		History: map[int]NamespaceTransaction{},
	}
	for k := range res.Record.History {
		// TODO: Maybe check length here. We will see
		tx := res.Record.History[k][0]
		if tx.Address != "" {
			out.Address = tx.Address
		}
		if tx.Base != 0 {
			out.Base = tx.Base
		}
		if tx.BlockNumber != 0 {
			out.BlockNumber = tx.BlockNumber
		}
		if tx.Buckets != nil {
			out.Buckets = tx.Buckets
		}
		if tx.Coeff != 0 {
			out.Coeff = tx.Coeff
		}
		if tx.Lifetime != 0 {
			out.Lifetime = tx.Lifetime
		}
		if tx.NamespaceID != "" {
			out.NamespaceID = tx.NamespaceID
		}
		if tx.NoVowelDiscount != 0 {
			out.NoVowelDiscount = tx.NoVowelDiscount
		}
		if tx.NonalphaDiscount != 0 {
			out.NonalphaDiscount = tx.NonalphaDiscount
		}
		if tx.Op != "" {
			out.Op = tx.Op
		}
		if tx.OpFee != 0 {
			out.OpFee = tx.OpFee
		}
		if tx.PreorderHash != "" {
			out.PreorderHash = tx.PreorderHash
		}
		if tx.Recipient != "" {
			out.Recipient = tx.Recipient
		}
		if tx.RecipientAddress != "" {
			out.RecipientAddress = tx.RecipientAddress
		}
		if tx.RevealBlock != 0 {
			out.RevealBlock = tx.RevealBlock
		}
		if tx.Sender != "" {
			out.Sender = tx.Sender
		}
		if tx.SenderPubkey != "" {
			out.SenderPubkey = tx.SenderPubkey
		}
		if tx.Txid != "" {
			out.Txid = tx.Txid
		}
		if tx.Version != 0 {
			out.Version = tx.Version
		}
		if tx.Vtxindex != 0 {
			out.Vtxindex = tx.Vtxindex
		}
		out.History[k] = NamespaceTransaction{
			Address:          tx.Address,
			Base:             tx.Base,
			BlockNumber:      tx.BlockNumber,
			Buckets:          tx.Buckets,
			BurnAddress:      tx.BurnAddress,
			Coeff:            tx.Coeff,
			ConsensusHash:    tx.ConsensusHash,
			HistorySnapshot:  tx.HistorySnapshot,
			Lifetime:         tx.Lifetime,
			NamespaceID:      tx.NamespaceID,
			NonalphaDiscount: tx.NonalphaDiscount,
			NoVowelDiscount:  tx.NoVowelDiscount,
			Op:               tx.Op,
			Opcode:           tx.Opcode,
			OpFee:            tx.OpFee,
			PreorderHash:     tx.PreorderHash,
			Recipient:        tx.Recipient,
			RecipientAddress: tx.RecipientAddress,
			RevealBlock:      tx.RevealBlock,
			Sender:           tx.Sender,
			SenderPubkey:     tx.SenderPubkey,
			Txid:             tx.Txid,
			Version:          tx.Version,
			Vtxindex:         tx.Vtxindex,
		}
	}
	w.Write(out.JSON())
}

// V1GetNamespacesHandler handles response for /v1/namespaces route
func (h *Handlers) V1GetNamespacesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok\n"))
}

// // NumNamesHandler the /resolver/numnames route
// func (h *Handlers) NumNamesHandler(w http.ResponseWriter, r *http.Request) {
// 	h.Indexer.Lock()
// 	w.Write([]byte(fmt.Sprintf("Resolver Stats:\n  CurrentBlock: %v\n  CurrentNames: %v\n  ExpectedNames: %v\n", h.Indexer.CurrentBlock, len(h.Indexer.Names), h.Indexer.ExpectedNames)))
// 	h.Indexer.Unlock()
// }
//
// // NumNamesHandler the /resolver/names route
// func (h *Handlers) GetNamesHandler(w http.ResponseWriter, r *http.Request) {
// 	var out []string
// 	h.Indexer.Lock()
// 	names := h.Indexer.Names
// 	h.Indexer.Unlock()
// 	for _, name := range names {
// 		out = append(out, name.Name)
// 	}
// 	byt, err := json.Marshal(out)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	w.Write(byt)
// }
