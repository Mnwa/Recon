# Recon
[![Github all releases](https://img.shields.io/github/release/Mnwa/Recon.svg)](https://github.com/Mnwa/Recon/releases)
[![Build Status](https://cloud.drone.io/api/badges/Mnwa/Recon/status.svg)](https://cloud.drone.io/Mnwa/Recon)
[![Go Report Card](https://goreportcard.com/badge/Mnwa/Recon)](https://goreportcard.com/report/Mnwa/Recon)
[![GitHub license](https://img.shields.io/github/license/Mnwa/Recon.svg)](https://github.com/Mnwa/Recon)
[![Repository Size](https://img.shields.io/github/repo-size/Mnwa/Recon.svg)](https://github.com/Mnwa/Recon)

LDAP based key/value store with env support

Now it have only env data support

# Run
Recon it's simple golang app and it's requirements only [go](https://golang.org/) and [dep](https://github.com/golang/dep)

## Unix
Install dependencies:
```bash
dep ensure
```

Run tests:
```bash
go test
```
Compile sources:
```bash
go build -o ./build/service
```
Run binary:
```bash
RECON_DB_DIR=/tmp/recon ./build/service
```

## Docker
### Docker Hub
```bash
docker pull mnwamnowich/recon
```
Run:
```bash
docker run -p 8080:8080 --name recon mnwamnowich/recon
```
### Build from sources
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
* `RECON_REPLICATION_INIT` will be restore backup from other masters if not set `off`. Default is `nil`
* `RECON_REPLICATION_HOSTS` hosts of others masters in replication separated by comma (`localhost:8081,localhost:443`) default is `nil`

# Recon usage

## Introduction 
Recon have are simple standard based on LDAP. 
It's used default http requests methods (PUT, POST, GET, DELETE) and path like are:
`/projects/:projectName/:projectType/:dataType`  where:
* `:projectName String` name of you project
* `:projectType String` project type 
* `:dataType String` type of you data input/output (Supports [env](./docs/ENV.md), [config](./docs/CONFIG.md))

## Project type configuration
Every project have `default` type and it type will be merged with other types of you project.
When you try to GET `/projects/:projectName/:projectType/:dataType` it's will be return `/projects/:projectName/default/:dataType` + `/:projectName/:projectType/:dataType`
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
