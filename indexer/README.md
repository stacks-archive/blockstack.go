# Blockstack Indexer

The resolver runs through all the names in the Blockstack network, pulls their Zonefiles and resolves their profiles by fetching the data there. It persists this data in a Mongodb instance to survive restarts and re-syncs the data if the process dies.

The blockstack api runs an instance of the indexer to help manage responses. The resolver also has the database connection.


# Implement Retries for the following methods:



# Sample Subdomain record:

```
created_equal.self_evident_truth.id.	3600	IN	TXT	"owner=1AYddAnfHbw6bPNvnsQFFrEuUdhMhf2XG9" "seqn=0" "parts=1" "zf0=JE9SSUdJTiBjcmVhdGVkX2VxdWFsCiRUVEwgMzYwMApfaHR0cHMuX3RjcCBVUkkgMTAgMSAiaHR0cHM6Ly93d3cuY3MucHJpbmNldG9uLmVkdS9+YWJsYW5rc3QvY3JlYXRlZF9lcXVhbC5qc29uIgpfZmlsZSBVUkkgMTAgMSAiZmlsZTovLy90bXAvY3JlYXRlZF9lcXVhbC5qc29uIgo="
```
