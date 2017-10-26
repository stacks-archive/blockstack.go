# Blockstack Resolver/Indexer

This Indexer crawls all the namespaces on Blockstack and resolves all the names. It then (WIP) persists these resolved profiles in a MongoDB database for easy access by an index enabled version of the `blockstack-api`.

It does this by connecting to one or more [`blockstack-core`](https://github.com/blockstack/blockstack-core) instances and running a series of RPC calls to gather all of the Zonefiles in the network. Those Zonefiles are queried for valid URI records that point to the name/domain storage (see [`gaia`](https://github.com/blockstack/gaia)). Those storage URLs are then resolved and the associated profiles are saved, along with the Zonefile info in the persistent storage for the Indexer. The indexer has two modes, by name and by block. The RPC calls and flow for each method are described below.

### By Name

First all namespaces are discovered with the `get_all_namespaces`. That returns an array of namespaces:

```json
[
  "id",
  "helloWorld",
]
```

Each namespace is then queried for all of the names it contains with the `get_all_names_in_namespace` method:

```json
[
  "foo.id",
  "bar.id",
  ...
]
```

These names are returned in batches of 100. Each batch is managed by a separate `goroutine` and details for each name are fetched serially within that `goroutine` using the `get_name_at [currentBlock]` RPC method. Once all of the names in the batch have details all of the `zonefile_hash`es for that batch are fetched using the `get_zonefiles` method. As each `zonefile` is associated with a name it is then sent to be resolved individually. Once resolved the profiles are batched for efficient insert/update on the database layer. Each step of this process is easily parallelize-able. Knobs are provided for managing the concurrency at each step to allow for the application to be fit to the hardware running it.

The current build of this method takes around ~45 minutes to complete on the test setup (a 2CPU machine connected to 8 `blockstack-core` nodes). The bottleneck is RPC calls. Scaling the number of core nodes and optimizing concurrency to effectively use resources on the indexer machine should be able to reduce that number to ~20-30 minutes.

### By Block

Another method is to fetch all of the names/domains and then iterate through blockchain and pull out all the blocks. This method starts like the `By Names` method by pulling all the namespaces (`get_all_namespaces`) and all the names(`get_all_names_in_namespace`). Once those are fetched the blockchain is iterated through starting with the first Blockstack block `373601` using the `get_zonefiles_by_block` method. The zonefiles are then associated with the names `map[string]*Domain`. Newer zonefiles replace older ones until the current block is reached. Once that happens the zonefile hashes are batched up and decoded with `get_zonefiles`. The names are then resolved.

A preliminary build of this setup showed faster performance than the `By Names` method. It was completing in ~30 min. A full build needs to be completed and benchmarked against the `By Names` method before deciding on the proper approach. The application has been designed to accommodate this.

### Metrics

Metrics are exposed on `locahost:3000/metrics` using a Prometheus server. This is to provide visibility into the different parts of pipeline. The metrics collected are subject to change but are currently designed to show the progress of the indexing operation. Information about core call latency and indexing performance will be added when the Indexer is running periodically.

To get just the indexer metrics run `curl -s localhost:3000/metrics | grep "^indexer"`

### 
