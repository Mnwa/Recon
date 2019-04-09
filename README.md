# Recon
LDAP based key/value store with env support

Now it have only env data support

# Run
Recon it's simple golang app and it's don't need some requirements, only go

## Unix
Compile sources:
```bash
go build -o ./build/service
```
Run binary:
```bash
cd ./build
RECON_DB_DIR=/tmp/recon ./service
```

## Docker
Build the image:
```bash
docker build -t recon_build . 
```
Run:
```bash
docker run -p 8080:8080 --name recon recon_build
```

For saving data you need create volume from ```/var/lib/recon``` to you system
# Options

Recon can be configured with environment:
* `RECON_DB_DIR` default is `/var/lib/recon`
* `RECON_ADDR` default is `:8080`

# Recon usage

* [Default kv usage](./docs/CONFIG.md)
* [Environment usage](./docs/ENV.md)