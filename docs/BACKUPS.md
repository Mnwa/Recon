# Simple example of environment usage
Recon backups worked on protobuf standards and must be download/restored as binary file

## Get backup data
```http
GET http://localhost:8080/backup
Accept: */*
Cache-Control: no-cache

###

***BACKUP DATA***
```
## Get key of project environment
```http
POST http://localhost:8080/backup
Accept: */*
Cache-Control: no-cache
Content-Type: application/protobuf

***BACKUP DATA***

###

Success restored
```