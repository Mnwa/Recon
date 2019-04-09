# Simple example of default usage

## Create project environment
```http
PUT http://localhost:8080/projectName/projectType/config
Accept: */*
Cache-Control: no-cache
Content-Type: text/plain

var_one=1
var_two=2

###

var_one=1
var_two=2
```
## Add key to project environment
```http
PUT http://localhost:8080/projectName/projectType/config/var_three
Accept: */*
Cache-Control: no-cache
Content-Type: text/plain

3

###

3
```

## Update project environment
```http
POST http://localhost:8080/projectName/projectType/config
Accept: */*
Cache-Control: no-cache
Content-Type: text/plain

var_four=4
var_five=5

###

var_one=1
var_two=2
var_three=3
var_four=4
var_five=5
```
## Update key of project environment
```http
POST http://localhost:8080/projectName/projectType/config/var_five
Accept: */*
Cache-Control: no-cache
Content-Type: text/plain

10

###

10
```

## Get project environment
```http
GET http://localhost:8080/projectName/projectType/config
Accept: */*
Cache-Control: no-cache

###

var_one=1
var_two=2
var_three=3
var_four=4
var_five=10
```
## Get key of project environment
```http
GET http://localhost:8080/projectName/projectType/config/var_five
Accept: */*
Cache-Control: no-cache

###

10
```

## DELETE project environment
```http
DELETE http://localhost:8080/projectName/projectType/config
Accept: */*
Cache-Control: no-cache

###

Deleted
```
## DELETE key of project environment
```http
DELETE http://localhost:8080/projectName/projectType/config/var_five
Accept: */*
Cache-Control: no-cache

### 

Deleted
```