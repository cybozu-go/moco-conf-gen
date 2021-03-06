# Build the moco-conf-gen binary
FROM quay.io/cybozu/golang:1.13-bionic as builder

WORKDIR /workspace

# Copy the go source
COPY go.mod go.mod
COPY cmd/ cmd/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -mod=vendor -a -o moco-conf-gen ./cmd/moco-conf-gen/main.go

FROM quay.io/cybozu/ubuntu:18.04
WORKDIR /
COPY --from=builder /workspace/moco-conf-gen ./

RUN mkdir -p /etc/mysql \
  && chown -R 10000:10000 /etc/mysql
VOLUME /etc/mysql
USER 10000:10000

ENTRYPOINT ["/moco-conf-gen"]
