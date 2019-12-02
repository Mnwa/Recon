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
