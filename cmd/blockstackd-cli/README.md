# `blockstackd-cli`

This is a cli tool to interact with the RPC interface on `blockstack-core`. It implements the following RPC methods:

- `get_all_names`
- `get_all_namespaces`
- `get_atlas_peers`
- `get_block_from_consensus`
- `get_consensus_at`
- `get_consensus_hashes`
- `getinfo`
- `get_name_at`
- `get_name_blockchain_record`
- `get_name_cost`
- `get_name_history_blocks`
- `get_nameops_hash_at`
- `get_nameops_affected_at`
- `get_names_in_namespace`
- `get_names_owned_by_address`
- `get_namespace_cost`
- `get_num_nameops_affected_at`
- `get_num_names`
- `get_num_names_in_namespace`
- `get_num_op_history_rows`
- `get_op_history_rows`
- `get_zonefiles`
- `get_zonefiles_by_block`
- `ping`

### Config

This CLI is a [Cobra](https://github.com/spf13/cobra) application. It is configurable via file (default `$HOME/.blockstack.yaml`) or flags. A sample config file is below:

```yaml
# $HOME/.blockstackd-cli.yaml
node: https://node.blockstack.org:6263
```

### Build

To build the binary you must have [Glide](https://github.com/Masterminds/glide) (a dependency manager for go) installed. Then go to the project root and run:

```shell
$ glide i
$ go install cmd/blockstackd-cli/main.go
```
