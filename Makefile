

# For Go
GO111MODULE = on
GOFLAGS     = -mod=vendor
export GO111MODULE GOFLAGS

all: test

test:
	cd /tmp; GO111MODULE=on GOFLAGS= go install github.com/cybozu/neco-containers/golang/analyzer/cmd/custom-checker
	test -z "$$(gofmt -s -l . | grep -v '^vendor' | tee /dev/stderr)"
	test -z "$$(golint $$(go list -mod=vendor ./... | grep -v /vendor/) | tee /dev/stderr)"
	test -z "$$(nilerr ./... 2>&1 | tee /dev/stderr)"
	test -z "$$(custom-checker -restrictpkg.packages=html/template,log $$(go list -tags='$(GOTAGS)' ./... | grep -v /vendor/ ) 2>&1 | tee /dev/stderr)"
	go build -mod=vendor ./...
	go test -race -v ./...
	go vet ./...
	test -z "$$(go vet ./... | grep -v '^vendor' | tee /dev/stderr)"
	ineffassign .

mod:
	go mod tidy
	go mod vendor
	git add -f vendor
	git add go.mod

.PHONY:	all test mod
