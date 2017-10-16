package blockstack

import (
	"encoding/json"
	"fmt"
)

// Ping calls the ping RPC method for blockstack server
func (bsk *Client) Ping() (PingResult, error) {
	rpcCall := "ping"
	var callResult string

	err := bsk.node.Call(rpcCall, nil, &callResult)
	if err != nil {
		return PingResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return PingResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return PingResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out PingResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return PingResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetInfo calls the getinfo RPC method for blockstack server
func (bsk *Client) GetInfo() (GetInfoResult, error) {
	rpcCall := "getinfo"
	var callResult string

	err := bsk.node.Call(rpcCall, nil, &callResult)
	if err != nil {
		return GetInfoResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetInfoResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetInfoResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetInfoResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetInfoResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetZonefilesByBlock calls the get_zonefiles_by_block RPC method for blockstack server
func (bsk *Client) GetZonefilesByBlock(startBlock, endBlock, offset, count int) (GetZonefilesByBlockResult, error) {
	rpcCall := "get_zonefiles_by_block"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{startBlock, endBlock, offset, count}, &callResult)
	if err != nil {
		return GetZonefilesByBlockResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetZonefilesByBlockResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetZonefilesByBlockResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetZonefilesByBlockResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetZonefilesByBlockResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNameBlockchainRecord calls the get_name_blockchain_record RPC method for blockstack server
func (bsk *Client) GetNameBlockchainRecord(name string) (GetNameBlockchainRecordResult, error) {
	rpcCall := "get_name_blockchain_record"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{name}, &callResult)
	if err != nil {
		return GetNameBlockchainRecordResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNameBlockchainRecordResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetNameBlockchainRecordResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetNameBlockchainRecordResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNameBlockchainRecordResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNameHistoryBlocks calls the get_name_history_blocks RPC method for blockstack server
func (bsk *Client) GetNameHistoryBlocks(name string) (GetNameHistoryBlocksResult, error) {
	rpcCall := "get_name_history_blocks"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{name}, &callResult)
	if err != nil {
		return GetNameHistoryBlocksResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNameHistoryBlocksResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetNameHistoryBlocksResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetNameHistoryBlocksResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNameHistoryBlocksResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNameAt calls the get_name_at RPC method for blockstack server
func (bsk *Client) GetNameAt(name string, blockHeight int) (GetNameAtResult, error) {
	rpcCall := "get_name_at"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{name, blockHeight}, &callResult)
	if err != nil {
		return GetNameAtResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNameAtResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetNameAtResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetNameAtResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNameAtResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNamesOwnedByAddress calls the get_names_owned_by_address RPC method for blockstack server
func (bsk *Client) GetNamesOwnedByAddress(address string) (GetNamesOwnedByAddressResult, error) {
	rpcCall := "get_names_owned_by_address"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{address}, &callResult)
	if err != nil {
		return GetNamesOwnedByAddressResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNamesOwnedByAddressResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetNamesOwnedByAddressResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetNamesOwnedByAddressResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNamesOwnedByAddressResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNameCost calls the get_name_cost RPC method for blockstack server
func (bsk *Client) GetNameCost(name string) (GetNameCostResult, error) {
	rpcCall := "get_name_cost"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{name}, &callResult)
	if err != nil {
		return GetNameCostResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNameCostResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetNameCostResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetNameCostResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNameCostResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNamespaceCost calls the get_namespace_cost RPC method for blockstack server
func (bsk *Client) GetNamespaceCost(namespace string) (GetNamespaceCostResult, error) {
	rpcCall := "get_namespace_cost"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{namespace}, &callResult)
	if err != nil {
		return GetNamespaceCostResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNamespaceCostResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetNamespaceCostResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetNamespaceCostResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNamespaceCostResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNumNames calls the get_num_names RPC method for blockstack server
func (bsk *Client) GetNumNames() (CountResult, error) {
	rpcCall := "get_num_names"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{}, &callResult)
	if err != nil {
		return CountResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return CountResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return CountResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out CountResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return CountResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetAllNames calls the get_all_names RPC method for blockstack server
func (bsk *Client) GetAllNames(offset, count int) (GetAllNamesResult, error) {
	rpcCall := "get_all_names"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{offset, count}, &callResult)
	if err != nil {
		return GetAllNamesResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetAllNamesResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetAllNamesResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetAllNamesResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetAllNamesResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetAllNamespaces calls the get_all_namespaces RPC method for blockstack server
func (bsk *Client) GetAllNamespaces() (GetAllNamespacesResult, error) {
	rpcCall := "get_all_namespaces"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{}, &callResult)
	if err != nil {
		return GetAllNamespacesResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetAllNamespacesResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetAllNamespacesResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetAllNamespacesResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetAllNamespacesResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNamesInNamespace calls the get_names_in_namespace RPC method for blockstack server
func (bsk *Client) GetNamesInNamespace(ns string, offset int, count int) (GetNamesInNamespaceResult, error) {
	rpcCall := "get_names_in_namespace"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{ns, offset, count}, &callResult)
	if err != nil {
		return GetNamesInNamespaceResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNamesInNamespaceResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetNamesInNamespaceResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetNamesInNamespaceResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNamesInNamespaceResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNumNamesInNamespace calls the get_num_names_in_namespace RPC method for blockstack server
func (bsk *Client) GetNumNamesInNamespace(namespace string) (CountResult, error) {
	rpcCall := "get_num_names_in_namespace"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{namespace}, &callResult)
	if err != nil {
		return CountResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return CountResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return CountResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out CountResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return CountResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetConsensusAt calls the get_consensus_at RPC method for blockstack server
func (bsk *Client) GetConsensusAt(blockHeight int) (GetConsensusAtResult, error) {
	rpcCall := "get_consensus_at"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{blockHeight}, &callResult)
	if err != nil {
		return GetConsensusAtResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetConsensusAtResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetConsensusAtResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetConsensusAtResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetConsensusAtResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetBlockFromConsensus calls the get_block_from_consensus RPC method for blockstack server
func (bsk *Client) GetBlockFromConsensus(consensusHash string) (GetBlockFromConsensusResult, error) {
	rpcCall := "get_block_from_consensus"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{consensusHash}, &callResult)
	if err != nil {
		return GetBlockFromConsensusResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetBlockFromConsensusResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetBlockFromConsensusResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetBlockFromConsensusResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetBlockFromConsensusResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetAtlasPeers calls the get_atlas_peers RPC method for blockstack server
func (bsk *Client) GetAtlasPeers() (GetAtlasPeersResult, error) {
	rpcCall := "get_atlas_peers"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{}, &callResult)
	if err != nil {
		return GetAtlasPeersResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetAtlasPeersResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetAtlasPeersResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetAtlasPeersResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetAtlasPeersResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetZonefileInventory calls the get_zonefile_inventory RPC method for blockstack server
func (bsk *Client) GetZonefileInventory(offset, length int) (GetZonefileInventoryResult, error) {
	rpcCall := "get_zonefile_inventory"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{offset, length}, &callResult)
	if err != nil {
		return GetZonefileInventoryResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetZonefileInventoryResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetZonefileInventoryResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetZonefileInventoryResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetZonefileInventoryResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNameOpsHashAt calls the get_nameops_hash_at RPC method for blockstack server
func (bsk *Client) GetNameOpsHashAt(blockHeight int) (GetNameOpsHashAtResult, error) {
	rpcCall := "get_nameops_hash_at"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{blockHeight}, &callResult)
	if err != nil {
		return GetNameOpsHashAtResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNameOpsHashAtResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetNameOpsHashAtResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetNameOpsHashAtResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNameOpsHashAtResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNamespaceBlockchainRecord calls the get_namespace_blockchain_record RPC method for blockstack server
func (bsk *Client) GetNamespaceBlockchainRecord(namespace string) (GetNamespaceBlockchainRecordResult, error) {
	rpcCall := "get_namespace_blockchain_record"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{namespace}, &callResult)
	if err != nil {
		return GetNamespaceBlockchainRecordResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNamespaceBlockchainRecordResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetNamespaceBlockchainRecordResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetNamespaceBlockchainRecordResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNamespaceBlockchainRecordResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetZonefiles calls the get_zonefiles RPC method for blockstack server
func (bsk *Client) GetZonefiles(zonefiles []string) (GetZonefilesResult, error) {
	rpcCall := "get_zonefiles"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{zonefiles}, &callResult)
	if err != nil {
		return GetZonefilesResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetZonefilesResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetZonefilesResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetZonefilesResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetZonefilesResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetOpHistoryRows calls the get_op_history_rows RPC method for blockstack server
func (bsk *Client) GetOpHistoryRows(historyID string, offset int, count int) (GetOpHistoryRowsResult, error) {
	rpcCall := "get_op_history_rows"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{historyID, offset, count}, &callResult)
	if err != nil {
		return GetOpHistoryRowsResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetOpHistoryRowsResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetOpHistoryRowsResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetOpHistoryRowsResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetOpHistoryRowsResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNameOpsAffectedAt calls the get_nameops_affected_at RPC method for blockstack server
func (bsk *Client) GetNameOpsAffectedAt(blockID, offset, count int) (GetNameOpsAffectedAtResult, error) {
	rpcCall := "get_nameops_affected_at"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{blockID, offset, count}, &callResult)
	if err != nil {
		return GetNameOpsAffectedAtResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNameOpsAffectedAtResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetNameOpsAffectedAtResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetNameOpsAffectedAtResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNameOpsAffectedAtResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetConsensusHashes calls the get_consensus_hashes RPC method for blockstack server
func (bsk *Client) GetConsensusHashes(blocks []int) (GetConsensusHashesResult, error) {
	rpcCall := "get_consensus_hashes"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{blocks}, &callResult)
	if err != nil {
		return GetConsensusHashesResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetConsensusHashesResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return GetConsensusHashesResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out GetConsensusHashesResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetConsensusHashesResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNumOpHistoryRows calls the get_num_op_history_rows RPC method for blockstack server
func (bsk *Client) GetNumOpHistoryRows(historyID string) (CountResult, error) {
	rpcCall := "get_num_op_history_rows"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{historyID}, &callResult)
	if err != nil {
		return CountResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return CountResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return CountResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out CountResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return CountResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}

// GetNumNameOpsAffectedAt calls the get_num_nameops_affected_at RPC method for blockstack server
func (bsk *Client) GetNumNameOpsAffectedAt(blockID int) (CountResult, error) {
	rpcCall := "get_num_nameops_affected_at"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{blockID}, &callResult)
	if err != nil {
		return CountResult{}, fmt.Errorf("RPC call %v failed: %v", rpcCall, err)
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return CountResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	if rpcError.Error() != "" {
		return CountResult{}, fmt.Errorf("RPC Error on call %v: %v", rpcCall, rpcError)
	}

	var out CountResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return CountResult{}, fmt.Errorf("Failed to Unmarshall %v RPC result: %v", rpcCall, err)
	}

	return out, nil
}
