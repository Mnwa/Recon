# Recon
[![Github all releases](https://img.shields.io/github/release/Mnwa/Recon.svg)](https://github.com/Mnwa/Recon/releases)
[![Build Status](https://cloud.drone.io/api/badges/Mnwa/Recon/status.svg)](https://cloud.drone.io/Mnwa/Recon)
[![Go Report Card](https://goreportcard.com/badge/Mnwa/Recon)](https://goreportcard.com/report/Mnwa/Recon)
[![GitHub license](https://img.shields.io/github/license/Mnwa/Recon.svg)](https://github.com/Mnwa/Recon)
[![Repository Size](https://img.shields.io/github/repo-size/Mnwa/Recon.svg)](https://github.com/Mnwa/Recon)

Recon is the simple solution for storing configs of you application.
There is no specified instruments, no specified data protocols. For the full power using of Recon you needed only `curl`.

## Examples

Put the data to project `myLittleProduction` with `default` project type
```bash
curl -X PUT -d '<<<EOF
TEST_ONE=1
TEST_TWO=2
TEST_THREE=3
' http://localhost:8080/projects/myLittleProduction/default/env

###

TEST_ONE=1
TEST_TWO=2
TEST_THREE=3
```

Get the data from `myLittleProduction` with `default` project type
```bash
curl http://localhost:8080/projects/myLittleProduction/default/env

###

TEST_ONE=1
TEST_TWO=2
TEST_THREE=3
```

## Data storing types
Recon DB support two data storing types:
* `default` snake case key value store (data_key=data)
* `env` environment like data storing (DATA_KEY=data)

## Structure of projects

Every project on Recon DB have are `default` type. Also you can add anything types to every projects for differentiation of configs and when you get custom project type, it will be merged with the default project type.

As example lets imagine the simple application, who stored on two data centres `1` and `2`. 
It's application have a `AWS secret key`, what's same for all data centres and it's have `Database url`, what's different on `1` and `2` data centres.
In order not to duplicate configs of `AWS secret key` we can add it to `default` project type

```bash
curl -X POST -d 'AWS_KEY=123' http://localhost:8080/projects/myApp/default/env
```

And after add `DATABASE URL` to the different project types

```bash
curl -X POST -d 'DATABASE_URL=localhost1' http://localhost:8080/projects/myApp/usa/env
curl -X POST -d 'DATABASE_URL=localhost2' http://localhost:8080/projects/myApp/europe/env
```

Now you can get fully config for every place

```bash
curl http://localhost:8080/projects/myApp/usa/env

###

AWS_KEY=123
DATABASE_URL=localhost1
```

```bash
curl http://localhost:8080/projects/myApp/europe/env

###

AWS_KEY=123
DATABASE_URL=localhost2
```

## Documentation
* [Basic](./docs/BASIC.md)
* [Build and run](docs/BUILD.md)
* [Default kv usage](./docs/CONFIG.md)
* [Environment usage](./docs/ENV.md)
* [Make backups](./docs/BACKUPS.md)
* [Prometheus metrics](./docs/PROMETHEUS.md)
* [Replication](./docs/REPLICATION.md)

## Recon requirements
* [Recon Engine](https://github.com/Mnwa/ReconEngine)
* [fasthttp](https://github.com/valyala/fasthttp)
* [fasthttp router](https://github.com/fasthttp/router)
* [fasthttp prometheus](https://github.com/Mnwa/fasthttprouter-prometheus)

### Thanks for community! 🎉 🎉 🎉 🎉