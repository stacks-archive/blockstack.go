# `bsk-load`

This tool is designed to load test `blockstackd` installations. To build the tool, have a working golang installation and `glide` (go dependency management) installed and run the following: 

```bash
# install dependencies
$ glide i

# build binary
$ go build -o bsk-load ./cmd/bsk-load/main.go

# Create configuration file (Format below)
$ touch ~/.bsk-load.yaml 
```

### Configuration

Default location for the configuration file is `~/.bsk-load.yaml`. Users can also pass a configuration file to the process with the `--config` flag. Format is below:

```yaml
testServer:
  address: "node.blockstack.org"
  scheme: "https"
  port: 6263
```

> TODO: Need to add things like concurrency, optional name/zonefile fetching, and reporting frequency in the configuration as well.

### Output

The test emits a summary of all requests performed every 10 seconds as a JSON blob. Operators can pipe `stdout` to a file and process the records with `jq` easily. Example output:

```json
{
  "get_all_namespaces": {
    "max": "336.409532ms",
    "mean": "336.409532ms",
    "min": "336.409532ms",
    "numCalls": "1",
    "p50": "336.409532ms",
    "p90": "336.409532ms",
    "p95": "336.409532ms"
  },
  "get_blockchain_name_record": {
    "max": "9.156900319s",
    "mean": "2.55826621s",
    "min": "2.118160567s",
    "numCalls": "8620",
    "p50": "2.431953575s",
    "p90": "2.803634547s",
    "p95": "3.020158923s"
  },
  "get_names_in_namespace": {
    "max": "5.799410502s",
    "mean": "3.236053718s",
    "min": "457.36694ms",
    "numCalls": "90",
    "p50": "3.056377935s",
    "p90": "4.334941868s",
    "p95": "4.478795734s"
  },
  "get_num_names_in_namespace": {
    "max": "4.02093507s",
    "mean": "4.02093507s",
    "min": "4.02093507s",
    "numCalls": "1",
    "p50": "4.02093507s",
    "p90": "4.02093507s",
    "p95": "4.02093507s"
  },
  "totals": {
    "calls": "8712",
    "callsPerSec": "3.889262614710134",
    "time": "37m20.013304077s"
  }
}
```

### Scenarios

There is currently one scenario built in: `nameScan`. This scenario calls the following RPC methods:

- `get_all_namespaces` - fetch the list of namespaces
- `get_num_names_in_namespace` - find the number of names in each namespace
- `get_names_in_namespace` - fetch a list of all names in all namespaces

These could be behind a feature flag, but currently aren't:

- `get_name_blockchain_record` - get details for each name, including zonefile hash
- `get_zonefiles` - fetches all the zonefiles for all the names

The program currently runs 10 routines fetching pages and then fetching details for each name. I plan to expose that as a configurable variable. 