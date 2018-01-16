# Core Deployment

This folder has the following files:

- `ops`: A shell script to run docker images that initialize a `blockstack-core` node's `~/.blockstack-server` directory with the `fast_sync` directive
- `docker-compose.yaml`: Runs 4 `blockstack-core` nodes, each on a separate host port.

This allows for running tightly packed `blockstack-core` nodes on multi-core servers.

```bash
# First run the fast_sync
$ ./ops init-core 1
$ ./ops init-core 2
$ ./ops init-core 3
$ ./ops init-core 4

# You should now have the following file structure
$ tree -L 2
.
├── core-1
│   ├── api
│   └── server
├── core-2
│   ├── api
│   └── server
├── core-3
│   ├── api
│   └── server
├── core-4
│   ├── api
│   └── server
├── docker-compose.yaml
└── ops

# Now you can start the containers
$ docker-compose up -d
```
