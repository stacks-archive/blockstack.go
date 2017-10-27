# Blockstack API

The `blockstack-api` provides a performant interface for fetching data about the Blockstack network. Most of the calls make RPC calls against a configured `blockstack-core` node and return that data to the user. It can connect to multiple `blockstack-core` backends at once to enable scaling. It is written in go and utilizes extensive parallelization for speed. It will expose an API as follows:

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
