# Blockstack Indexer

The `blockstack-indexer` connects to a `blockstack-core` node and pulls the data out for faster access. It is written in go and utilizes extensive parallelization for speed. It will expose an API as follows:

```bash
# implemented
/v1/names/:domainName
# implemented
/v1/names/:domainName/history
# implemented
/v1/namespaces/:namespaceId/names?page=:pageNum
/v2/users/:domainName
/v1/blockchains/bitcoin/operations/:blockHeight

/v1/addresses/bitcoin/:address
/v1/names/:domainName/zonefile
/v1/namespaces/:id
/v1/namespaces
/v1/blockchains/bitcoin/name_count
```
