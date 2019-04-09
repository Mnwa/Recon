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

## Introduction 
Recon have are simple standard based on LDAP. 
It's used default http requests methods (PUT, POST, GET, DELETE) and path like are:
`/:projectName/:projectType/:dataType`  where:
* `:projectName String` name of you project
* `:projectType String` project type 
* `:dataType String` type of you data input/output (Supports [env](./docs/ENV.md), [config](./docs/CONFIG.md))

## Project type configuration
Every project have `default` type and it type will be merged with other types of you project.
When you try to GET `/:projectName/:projectType/:dataType` it's will be return `/:projectName/default/:dataType` + `/:projectName/:projectType/:dataType`
with `:projectType` merge priority

## Usages
* [Default kv usage](./docs/CONFIG.md)
* [Environment usage](./docs/ENV.md)
* [Make backups](./docs/BACKUPS.md)
* [Prometheus metrics](./docs/PROMETHEUS.md)

# Thanks for community! 🎉 🎉 🎉 🎉
Recon required that's packages:
* [fasthttp](https://github.com/valyala/fasthttp)
* [fasthttprouter](https://github.com/buaazp/fasthttprouter)
* [fasthttp-prometheus](https://github.com/flf2ko/fasthttp-prometheus)
* [bitcask](https://github.com/prologic/bitcask)