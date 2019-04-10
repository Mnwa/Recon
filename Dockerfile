FROM golang:alpine as builder

COPY . $GOPATH/src/Recon/
WORKDIR $GOPATH/src/Recon/

RUN apk add git gcc && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep ensure
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /home/go/Recon/service
FROM scratch
ENV RECON_DB_DIR='/var/lib/recon'
ENV RECON_ADDR=':8080'
# Copy our static executable
COPY --from=builder /home/go/Recon/service /app/service
ENTRYPOINT ["/app/service"]
