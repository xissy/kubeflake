GOPATH:=$(shell go env GOPATH)

.PHONY: format
## format: format files
format:
	@go install golang.org/x/tools/cmd/goimports@latest
	goimports -local github.com/xissy -w .
	gofmt -s -w .
	go mod tidy

.PHONY: test
## test: run tests
test:
	@go install github.com/rakyll/gotest@latest
	gotest -p 1 -race -cover -v ./...

.PHONY: coverage
## coverage: run tests with coverage
coverage:
	@go install github.com/rakyll/gotest@latest
	gotest -p 1 -race -coverprofile=coverage.out -covermode=atomic -v ./...

.PHONY: lint
## lint: check everything's okay
lint:
	@go install github.com/kyoh86/scopelint@latest
	golangci-lint run ./...
	scopelint --set-exit-status ./...
	go mod verify

.PHONY: generate
## generate: generate files such as mocks
generate:
	@go install github.com/golang/mock/mockgen
	go generate ./...

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':'
