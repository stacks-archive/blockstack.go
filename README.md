# Blockstack Indexer

The `blockstack-indexer` connects to a `blockstack-core` node and pulls the data out for faster access. It is written in go and utilizes extensive parallelization for speed. It will expose an API as follows:

```
/v1/names/:domainName
/v1/names/:domainName/history
/v1/namespaces/:namespaceId/names?page=:pageNum
/v2/users/:domainName
/v1/blockchains/bitcoin/operations/:blockHeight
/v1/addresses/bitcoin/:address
/v1/names/:domainName/zonefile
/v1/namespaces/:id
/v1/namespaces
/v1/blockchains/bitcoin/name_count
```

Test core node connectivity with the following curl call:

```
curl -L -XPOST -H "Content-Type: application/xml" {addresss}:{port}/RPC2 -d '<?xml version="1.0"?><methodCall><methodName>getinfo</methodName><params></params></methodCall>'
```
