# Recon
LDAP based key/value store with env support

Now it have only env data support

# Examples

## Create project environment
```http
PUT http://localhost:8080/projectName/projectType/env
Accept: */*
Cache-Control: no-cache
Content-Type: text/plain

VAR_ONE=1
VAR_TWO=2

###

VAR_ONE=1
VAR_TWO=2
```
## Add key to project environment
```http
PUT http://localhost:8080/projectName/projectType/env/VAR_THREE
Accept: */*
Cache-Control: no-cache
Content-Type: text/plain

3

###

3
```

## Update project environment
```http
POST http://localhost:8080/projectName/projectType/env
Accept: */*
Cache-Control: no-cache
Content-Type: text/plain

VAR_FOUR=4
VAR_FIVE=5

###

VAR_ONE=1
VAR_TWO=2
VAR_THREE=3
VAR_FOUR=4
VAR_FIVE=5
```
## Update key of project environment
```http
POST http://localhost:8080/projectName/projectType/env/VAR_FIVE
Accept: */*
Cache-Control: no-cache
Content-Type: text/plain

10

###

10
```

## Get project environment
```http
GET http://localhost:8080/projectName/projectType/env
Accept: */*
Cache-Control: no-cache

###

VAR_ONE=1
VAR_TWO=2
VAR_THREE=3
VAR_FOUR=4
VAR_FIVE=10
```
## Get key of project environment
```http
GET http://localhost:8080/projectName/projectType/env/VAR_FIVE
Accept: */*
Cache-Control: no-cache

###

10
```

## DELETE project environment
```http
DELETE http://localhost:8080/projectName/projectType/env
Accept: */*
Cache-Control: no-cache

###

Deleted
```
## DELETE key of project environment
```http
DELETE http://localhost:8080/projectName/projectType/env/VAR_FIVE
Accept: */*
Cache-Control: no-cache

### 

Deleted
```
