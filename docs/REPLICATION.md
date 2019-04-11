# Simple example of replication usage
Recon supports easy multi master replication

## Setup environment
For setup multi master replication - you need run Recon with env: `RECON_REPLICATION_HOSTS` hosts of others masters in replication separated by comma
```env
RECON_REPLICATION_HOSTS=localhost:8081,localhost:8082
```

## Init new master
After start Recon with `RECON_REPLICATION_HOSTS` env new master will be dumped data from first success response host and restore to yourself.
(to disable it set `RECON_REPLICATION_INIT=off` to env)