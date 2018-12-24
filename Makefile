# vi: ft=make

GOPATH:=$(shell go env GOPATH)

.PHONY: test
test:
	@go get github.com/rakyll/gotest
	gotest -p 1 -race -cover -v ./...

.PHONY: bench
bench:
	@go get github.com/rakyll/gotest
	gotest -p 1 -bench=.

.PHONY: lint
lint:
	golangci-lint run ./...
