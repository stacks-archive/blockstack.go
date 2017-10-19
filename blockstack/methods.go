package blockstack

import (
	"encoding/json"
)

// Ping calls the ping RPC method for blockstack server
func (bsk *Client) Ping() (PingResult, Error) {
	rpcCall := "ping"
	var callResult string

	err := bsk.node.Call(rpcCall, nil, &callResult)
	if err != nil {
		return PingResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return PingResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return PingResult{}, rpcError
	}

	var out PingResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return PingResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetInfo calls the getinfo RPC method for blockstack server
func (bsk *Client) GetInfo() (GetInfoResult, Error) {
	rpcCall := "getinfo"
	var callResult string

	err := bsk.node.Call(rpcCall, nil, &callResult)
	if err != nil {
		return GetInfoResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetInfoResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetInfoResult{}, rpcError
	}

	var out GetInfoResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetInfoResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetZonefilesByBlock calls the get_zonefiles_by_block RPC method for blockstack server
func (bsk *Client) GetZonefilesByBlock(startBlock, endBlock, offset, count int) (GetZonefilesByBlockResult, Error) {
	rpcCall := "get_zonefiles_by_block"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{startBlock, endBlock, offset, count}, &callResult)
	if err != nil {
		return GetZonefilesByBlockResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetZonefilesByBlockResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetZonefilesByBlockResult{}, rpcError
	}

	var out GetZonefilesByBlockResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetZonefilesByBlockResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNameBlockchainRecord calls the get_name_blockchain_record RPC method for blockstack server
func (bsk *Client) GetNameBlockchainRecord(name string) (GetNameBlockchainRecordResult, Error) {
	rpcCall := "get_name_blockchain_record"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{name}, &callResult)
	if err != nil {
		return GetNameBlockchainRecordResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNameBlockchainRecordResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetNameBlockchainRecordResult{}, rpcError
	}

	var out GetNameBlockchainRecordResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNameBlockchainRecordResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNameHistoryBlocks calls the get_name_history_blocks RPC method for blockstack server
func (bsk *Client) GetNameHistoryBlocks(name string) (GetNameHistoryBlocksResult, Error) {
	rpcCall := "get_name_history_blocks"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{name}, &callResult)
	if err != nil {
		return GetNameHistoryBlocksResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNameHistoryBlocksResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetNameHistoryBlocksResult{}, rpcError
	}

	var out GetNameHistoryBlocksResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNameHistoryBlocksResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNameAt calls the get_name_at RPC method for blockstack server
func (bsk *Client) GetNameAt(name string, blockHeight int) (GetNameAtResult, Error) {
	rpcCall := "get_name_at"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{name, blockHeight}, &callResult)
	if err != nil {
		return GetNameAtResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNameAtResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetNameAtResult{}, rpcError
	}

	var out GetNameAtResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNameAtResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNamesOwnedByAddress calls the get_names_owned_by_address RPC method for blockstack server
func (bsk *Client) GetNamesOwnedByAddress(address string) (GetNamesOwnedByAddressResult, Error) {
	rpcCall := "get_names_owned_by_address"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{address}, &callResult)
	if err != nil {
		return GetNamesOwnedByAddressResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNamesOwnedByAddressResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetNamesOwnedByAddressResult{}, rpcError
	}

	var out GetNamesOwnedByAddressResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNamesOwnedByAddressResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNameCost calls the get_name_cost RPC method for blockstack server
func (bsk *Client) GetNameCost(name string) (GetNameCostResult, Error) {
	rpcCall := "get_name_cost"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{name}, &callResult)
	if err != nil {
		return GetNameCostResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNameCostResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetNameCostResult{}, rpcError
	}

	var out GetNameCostResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNameCostResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNamespaceCost calls the get_namespace_cost RPC method for blockstack server
func (bsk *Client) GetNamespaceCost(namespace string) (GetNamespaceCostResult, Error) {
	rpcCall := "get_namespace_cost"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{namespace}, &callResult)
	if err != nil {
		return GetNamespaceCostResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNamespaceCostResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetNamespaceCostResult{}, rpcError
	}

	var out GetNamespaceCostResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNamespaceCostResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNumNames calls the get_num_names RPC method for blockstack server
func (bsk *Client) GetNumNames() (CountResult, Error) {
	rpcCall := "get_num_names"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{}, &callResult)
	if err != nil {
		return CountResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return CountResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return CountResult{}, rpcError
	}

	var out CountResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return CountResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetAllNames calls the get_all_names RPC method for blockstack server
func (bsk *Client) GetAllNames(offset, count int) (GetAllNamesResult, Error) {
	rpcCall := "get_all_names"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{offset, count}, &callResult)
	if err != nil {
		return GetAllNamesResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetAllNamesResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetAllNamesResult{}, rpcError
	}

	var out GetAllNamesResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetAllNamesResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetAllNamespaces calls the get_all_namespaces RPC method for blockstack server
func (bsk *Client) GetAllNamespaces() (GetAllNamespacesResult, Error) {
	rpcCall := "get_all_namespaces"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{}, &callResult)
	if err != nil {
		return GetAllNamespacesResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetAllNamespacesResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetAllNamespacesResult{}, rpcError
	}

	var out GetAllNamespacesResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetAllNamespacesResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNamesInNamespace calls the get_names_in_namespace RPC method for blockstack server
func (bsk *Client) GetNamesInNamespace(ns string, offset int, count int) (GetNamesInNamespaceResult, Error) {
	rpcCall := "get_names_in_namespace"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{ns, offset, count}, &callResult)
	if err != nil {
		return GetNamesInNamespaceResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNamesInNamespaceResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetNamesInNamespaceResult{}, rpcError
	}

	var out GetNamesInNamespaceResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNamesInNamespaceResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNumNamesInNamespace calls the get_num_names_in_namespace RPC method for blockstack server
func (bsk *Client) GetNumNamesInNamespace(namespace string) (CountResult, Error) {
	rpcCall := "get_num_names_in_namespace"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{namespace}, &callResult)
	if err != nil {
		return CountResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return CountResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return CountResult{}, rpcError
	}

	var out CountResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return CountResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetConsensusAt calls the get_consensus_at RPC method for blockstack server
func (bsk *Client) GetConsensusAt(blockHeight int) (GetConsensusAtResult, Error) {
	rpcCall := "get_consensus_at"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{blockHeight}, &callResult)
	if err != nil {
		return GetConsensusAtResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetConsensusAtResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetConsensusAtResult{}, rpcError
	}

	var out GetConsensusAtResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetConsensusAtResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetBlockFromConsensus calls the get_block_from_consensus RPC method for blockstack server
func (bsk *Client) GetBlockFromConsensus(consensusHash string) (GetBlockFromConsensusResult, Error) {
	rpcCall := "get_block_from_consensus"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{consensusHash}, &callResult)
	if err != nil {
		return GetBlockFromConsensusResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetBlockFromConsensusResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetBlockFromConsensusResult{}, rpcError
	}

	var out GetBlockFromConsensusResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetBlockFromConsensusResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetAtlasPeers calls the get_atlas_peers RPC method for blockstack server
func (bsk *Client) GetAtlasPeers() (GetAtlasPeersResult, Error) {
	rpcCall := "get_atlas_peers"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{}, &callResult)
	if err != nil {
		return GetAtlasPeersResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetAtlasPeersResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetAtlasPeersResult{}, rpcError
	}

	var out GetAtlasPeersResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetAtlasPeersResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetZonefileInventory calls the get_zonefile_inventory RPC method for blockstack server
func (bsk *Client) GetZonefileInventory(offset, length int) (GetZonefileInventoryResult, Error) {
	rpcCall := "get_zonefile_inventory"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{offset, length}, &callResult)
	if err != nil {
		return GetZonefileInventoryResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetZonefileInventoryResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetZonefileInventoryResult{}, rpcError
	}

	var out GetZonefileInventoryResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetZonefileInventoryResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNameOpsHashAt calls the get_nameops_hash_at RPC method for blockstack server
func (bsk *Client) GetNameOpsHashAt(blockHeight int) (GetNameOpsHashAtResult, Error) {
	rpcCall := "get_nameops_hash_at"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{blockHeight}, &callResult)
	if err != nil {
		return GetNameOpsHashAtResult{}, CallError{Err: err, RPC: rpcCall}
	}

	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNameOpsHashAtResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetNameOpsHashAtResult{}, rpcError
	}

	var out GetNameOpsHashAtResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNameOpsHashAtResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNamespaceBlockchainRecord calls the get_namespace_blockchain_record RPC method for blockstack server
func (bsk *Client) GetNamespaceBlockchainRecord(namespace string) (GetNamespaceBlockchainRecordResult, Error) {
	rpcCall := "get_namespace_blockchain_record"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{namespace}, &callResult)
	if err != nil {
		return GetNamespaceBlockchainRecordResult{}, CallError{Err: err, RPC: rpcCall}
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNamespaceBlockchainRecordResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetNamespaceBlockchainRecordResult{}, rpcError
	}

	var out GetNamespaceBlockchainRecordResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNamespaceBlockchainRecordResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetZonefiles calls the get_zonefiles RPC method for blockstack server
func (bsk *Client) GetZonefiles(zonefiles []string) (GetZonefilesResult, Error) {
	rpcCall := "get_zonefiles"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{zonefiles}, &callResult)
	if err != nil {
		return GetZonefilesResult{}, CallError{Err: err, RPC: rpcCall}
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetZonefilesResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetZonefilesResult{}, rpcError
	}

	var out GetZonefilesResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetZonefilesResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetOpHistoryRows calls the get_op_history_rows RPC method for blockstack server
func (bsk *Client) GetOpHistoryRows(historyID string, offset int, count int) (GetOpHistoryRowsResult, Error) {
	rpcCall := "get_op_history_rows"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{historyID, offset, count}, &callResult)
	if err != nil {
		return GetOpHistoryRowsResult{}, CallError{Err: err, RPC: rpcCall}
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetOpHistoryRowsResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetOpHistoryRowsResult{}, rpcError
	}

	var out GetOpHistoryRowsResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetOpHistoryRowsResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNameOpsAffectedAt calls the get_nameops_affected_at RPC method for blockstack server
func (bsk *Client) GetNameOpsAffectedAt(blockID, offset, count int) (GetNameOpsAffectedAtResult, Error) {
	rpcCall := "get_nameops_affected_at"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{blockID, offset, count}, &callResult)
	if err != nil {
		return GetNameOpsAffectedAtResult{}, CallError{Err: err, RPC: rpcCall}
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetNameOpsAffectedAtResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetNameOpsAffectedAtResult{}, rpcError
	}

	var out GetNameOpsAffectedAtResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetNameOpsAffectedAtResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetConsensusHashes calls the get_consensus_hashes RPC method for blockstack server
func (bsk *Client) GetConsensusHashes(blocks []int) (GetConsensusHashesResult, Error) {
	rpcCall := "get_consensus_hashes"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{blocks}, &callResult)
	if err != nil {
		return GetConsensusHashesResult{}, CallError{Err: err, RPC: rpcCall}
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return GetConsensusHashesResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return GetConsensusHashesResult{}, rpcError
	}

	var out GetConsensusHashesResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return GetConsensusHashesResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNumOpHistoryRows calls the get_num_op_history_rows RPC method for blockstack server
func (bsk *Client) GetNumOpHistoryRows(historyID string) (CountResult, Error) {
	rpcCall := "get_num_op_history_rows"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{historyID}, &callResult)
	if err != nil {
		return CountResult{}, CallError{Err: err, RPC: rpcCall}
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return CountResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return CountResult{}, rpcError
	}

	var out CountResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return CountResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}

// GetNumNameOpsAffectedAt calls the get_num_nameops_affected_at RPC method for blockstack server
func (bsk *Client) GetNumNameOpsAffectedAt(blockID int) (CountResult, Error) {
	rpcCall := "get_num_nameops_affected_at"
	var callResult string

	err := bsk.node.Call(rpcCall, []interface{}{blockID}, &callResult)
	if err != nil {
		return CountResult{}, CallError{Err: err, RPC: rpcCall}
	}
	var rpcError RPCError
	err = json.Unmarshal([]byte(callResult), &rpcError)
	if err != nil {
		return CountResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	if rpcError.Error() != "" {
		rpcError.RPC = rpcCall
		return CountResult{}, rpcError
	}

	var out CountResult
	err = json.Unmarshal([]byte(callResult), &out)
	if err != nil {
		return CountResult{}, JSONUnmarshalError{RPC: rpcCall, Err: err}
	}

	return out, nil
}
