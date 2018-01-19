# Blockstack Core RPC Client

This is an implementation of the [`blockstack-core`](https://github.com/blockstack/blockstack-core) RPC protocol in Golang. To use it in your project import `github.com/blockstack/blockstack.go/blockstack`

### RPC Implementation status

 The following RPC methods are supported:

- `ping`
- `getinfo`
- `get_zonefiles_by_block`
- `get_name_blockchain_record`
- `get_name_history_blocks`
- `get_name_at`
- `get_names_owned_by_address`
- `get_name_cost`
- `get_namespace_cost`
- `get_num_names`
- `get_all_names`
- `get_all_namespaces`
- `get_names_in_namespace`
- `get_num_names_in_namespace`
- `get_consensus_at`
- `get_block_from_consensus`
- `get_atlas_peers`
- `get_zonefile_inventory`
- `get_namespace_blockchain_record`
- `get_nameops_hash_at`
- `get_op_history_rows`
- `get_num_op_history_rows`
- `get_num_nameops_affected_at`
- `get_consensus_hashes`
- `get_nameops_affected_at`

The following methods have not been implemented:

- `get_mutable_data`
- `get_immutable_data`
- `put_mutable_data`
- `get_last_nameops`
- `get_data_servers`
- `get_all_neighbor_info`

The following methods don't appear to be implemented

- `get_zonefiles_by_names`
