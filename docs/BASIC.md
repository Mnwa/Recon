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