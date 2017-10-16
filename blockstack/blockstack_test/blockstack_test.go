package blockstack_test

import (
	"testing"

	"github.com/jackzampolin/go-blockstack/blockstack"
)

var conf = blockstack.ServerConfig{
	Address: "node.blockstack.org",
	Port:    6263,
	TLS:     true,
}

// TestPing tests the blockstack.Client.Ping method
func TestPing(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.Ping()
	if err != nil {
		t.Fail()
	}
	if res.Status == "" {
		t.Fail()
	}
}

// TestGetInfo tests the blockstack.Client.GetInfo method
func TestGetInfo(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetInfo()
	if err != nil {
		t.Fail()
	}
	if res.Consensus == "" {
		t.Fail()
	}
}

// TestGetZonefilesByBlock tests the blockstack.Client.GetZonefilesByBlock method
func TestGetZonefilesByBlock(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetZonefilesByBlock(480000, 480004, 0, 100)
	if err != nil {
		t.Fail()
	}
	if len(res.ZonefileInfo) != 3 {
		t.Fail()
	}
}

// TestGetNameBlockchainRecord tests the blockstack.Client.GetNameBlockchainRecord method
func TestGetNameBlockchainRecord(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNameBlockchainRecord("muneeb.id")
	if err != nil {
		t.Fail()
	}
	if len(res.Record.History) <= 4 {
		t.Fail()
	}
}

// TestGetNameHistoryBlocks tests the blockstack.Client.GetNameHistoryBlocks method
func TestGetNameHistoryBlocks(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNameHistoryBlocks("muneeb.id")
	if err != nil {
		t.Fail()
	}
	if len(res.HistoryBlocks) <= 4 {
		t.Fail()
	}
}

// TestGetNameAt tests the blockstack.Client.GetNameAt method
func TestGetNameAt(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNameAt("muneeb.id", 480004)
	if err != nil {
		t.Fail()
	}
	if len(res.Records) != 1 {
		t.Fail()
	}
}

// TestGetNamesOwnedByAddress tests the blockstack.Client.GetNamesOwnedByAddress method
func TestGetNamesOwnedByAddress(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNamesOwnedByAddress("17hEAjUUWp5wN9SEGYqxpdtjHKzWVkmHEo")
	if err != nil {
		t.Fail()
	}
	if len(res.Names) <= 0 {
		t.Fail()
	}
}

// TestGetNameCost tests the blockstack.Client.GetNameCost method
func TestGetNameCost(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNameCost("muneeb.id")
	if err != nil {
		t.Fail()
	}
	if res.Satoshis <= 1000 {
		t.Fail()
	}
}

// TestGetNamespaceCost tests the blockstack.Client.GetNamespaceCost method
func TestGetNamespaceCost(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNamespaceCost("foobar")
	if err != nil {
		t.Fail()
	}
	if res.Satoshis <= 10000 {
		t.Fail()
	}
}

// TestGetNumNames tests the blockstack.Client.GetNumNames method
func TestGetNumNames(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNumNames()
	if err != nil {
		t.Fail()
	}
	if res.Count <= 74000 {
		t.Fail()
	}
}

// TestGetAllNames tests the blockstack.Client.GetAllNames method
func TestGetAllNames(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetAllNames(0, 100)
	if err != nil {
		t.Fail()
	}
	if len(res.Names) != 100 {
		t.Fail()
	}
}

// TestGetAllNamespaces tests the blockstack.Client.GetAllNamespaces method
func TestGetAllNamespaces(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetAllNamespaces()
	if err != nil {
		t.Fail()
	}
	if len(res.Namespaces) <= 0 {
		t.Fail()
	}
}

// TestGetNamesInNamespace tests the blockstack.Client.GetNamesInNamespace method
func TestGetNamesInNamespace(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNamesInNamespace("id", 0, 100)
	if err != nil {
		t.Fail()
	}
	if len(res.Names) != 100 {
		t.Fail()
	}
}

// TestGetNumNamesInNamespace tests the blockstack.Client.GetNumNamesInNamespace method
func TestGetNumNamesInNamespace(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNumNamesInNamespace("id")
	if err != nil {
		t.Fail()
	}
	if res.Count <= 74000 {
		t.Fail()
	}
}

// TestGetConsensusAt tests the blockstack.Client.GetConsensusAt method
func TestGetConsensusAt(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetConsensusAt(480004)
	if err != nil {
		t.Fail()
	}
	if len(res.Consensus) != 32 {
		t.Fail()
	}
}

// TestGetBlockFromConsensus tests the blockstack.Client.GetBlockFromConsensus method
func TestGetBlockFromConsensus(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetBlockFromConsensus("28aaae61809c6292b187b5cd9a92aa25")
	if err != nil {
		t.Fail()
	}
	if res.BlockID == 480004 {
		t.Fail()
	}
}

// TestGetAtlasPeers tests the blockstack.Client.GetAtlasPeers method
func TestGetAtlasPeers(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetAtlasPeers()
	if err != nil {
		t.Fail()
	}
	if len(res.Peers) <= 5 {
		t.Fail()
	}
}

// TestGetZonefileInventory tests the blockstack.Client.GetZonefileInventory method
func TestGetZonefileInventory(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetZonefileInventory(0, 524288)
	if err != nil {
		t.Fail()
	}
	if res.Inv == "" {
		t.Fail()
	}
}

// TestGetNameOpsHashAt tests the blockstack.Client.GetNameOpsHashAt method
func TestGetNameOpsHashAt(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNameOpsHashAt(480003)
	if err != nil {
		t.Fail()
	}
	if len(res.OpsHash) != 64 {
		t.Fail()
	}
}

// TestGetNamespaceBlockchainRecord tests the blockstack.Client.GetNamespaceBlockchainRecord method
func TestGetNamespaceBlockchainRecord(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNamespaceBlockchainRecord("id")
	if err != nil {
		t.Fail()
	}
	if len(res.Record.History) < 2 {
		t.Fail()
	}
}

// TestGetZonefiles tests the blockstack.Client.TestGetZonefiles method
// func TestGetZonefiles(t *testing.T) {
// 	t.Parallel()
// 	bsk := blockstack.NewClient(conf)
// 	res, err := bsk.GetZonefiles()
// 	if err != nil {
// 		t.Fail()
// 	}
// 	if res.Lastblock == 12313 {
// 		t.Fail()
// 	}
// }

// TestGetOpHistoryRows tests the blockstack.Client.GetOpHistoryRows method
func TestGetOpHistoryRows(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetOpHistoryRows("id", 0, 10)
	if err != nil {
		t.Fail()
	}
	if len(res.HistoryRows) != 2 {
		t.Fail()
	}
}

// TestGetNameOpsAffectedAt tests the blockstack.Client.GetNameOpsAffectedAt method
func TestGetNameOpsAffectedAt(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNameOpsAffectedAt(480003, 0, 10)
	if err != nil {
		t.Fail()
	}
	if len(res.Nameops) != 3 {
		t.Fail()
	}
}

// TestGetConsensusHashes tests the blockstack.Client.GetConsensusHashes method
func TestGetConsensusHashes(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetConsensusHashes([]int{480003, 480005})
	if err != nil {
		t.Fail()
	}
	if res.ConsensusHashes[480003] != "1b8536675c9248dfddf4669985d5359f" {
		t.Fail()
	}
}

// TestGetNumOpHistoryRows tests the blockstack.Client.GetNumOpHistoryRows method
func TestGetNumOpHistoryRows(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNumOpHistoryRows("id")
	if err != nil {
		t.Fail()
	}
	if res.Count != 2 {
		t.Fail()
	}
}

// TestGetNumNameOpsAffectedAt tests the blockstack.Client.GetNumNameOpsAffectedAt method
func TestGetNumNameOpsAffectedAt(t *testing.T) {
	t.Parallel()
	bsk := blockstack.NewClient(conf)
	res, err := bsk.GetNumNameOpsAffectedAt(480003)
	if err != nil {
		t.Fail()
	}
	if res.Count != 3 {
		t.Fail()
	}
}
