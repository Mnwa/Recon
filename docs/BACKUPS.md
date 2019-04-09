# Simple example of environment usage

## Get backup data
```http
GET http://localhost:8080/backup
Accept: */*
Cache-Control: no-cache

###

{
  "projectName/default/var_six": "3",
  "projectName/projectType/var_four": "4",
  "projectName/projectType/var_one": "3",
  "projectName/projectType/var_three": "3",
  "projectName/projectType/var_two": "1"
}
```
## Get key of project environment
```http
POST http://localhost:8080/backup
Accept: */*
Cache-Control: no-cache

{
  "projectName/default/var_six": "3",
  "projectName/projectType/var_four": "4",
  "projectName/projectType/var_one": "3",
  "projectName/projectType/var_three": "3",
  "projectName/projectType/var_two": "1"
}

###

Success restored
```