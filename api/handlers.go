package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackzampolin/blockstack-indexer/blockstack"
)

type Handlers struct {
	Client *blockstack.Client
}

func NewHandlers(conf blockstack.ServerConfig) *Handlers {
	return &Handlers{
		Client: blockstack.NewClient(conf),
	}
}

// V1GetNameHandler handles the /v1/names/{name} route
// NOTE: This needs to be derived from a DB call
func (h *Handlers) V1GetNameHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	var out V1GetNameResponse
	// Do stuff...
	w.Write([]byte(out.JSON()))
}

// V1GetNameHistoryHandler handles response for /v1/names/{name}/history
func (h *Handlers) V1GetNameHistoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res, err := h.Client.GetNameBlockchainRecord(vars["name"])
	if err != nil {
		// TODO: return error to client
		log.Fatal(err)
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
		// TODO: return error to client here and mention that its an ill formed page value
		log.Fatal(err)
	}
	res, err := h.Client.GetNamesInNamespace(vars["namespace"], (int(pg) * 100), 100)
	if err != nil {
		// TODO: return error to client here and mention that theres a connection error to core node
		log.Fatal(err)
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
	log.Println("blockchain", vars["blockchain"], "blockHeight", vars["blockHeight"])
	w.Write([]byte("ok\n"))
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
