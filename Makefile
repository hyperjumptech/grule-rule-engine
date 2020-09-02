GO111MODULE=on

.PHONY: all test clean build docker

build:
	export GO111MODULE on; \
	go build ./...

lint: build
	golint -set_exit_status builder/... engine/... examples/... ast/... pkg/... antlr/. model/...

test-short: lint
	go test ./... -v -covermode=count -coverprofile=coverage.out -short
	goornogo -i coverage.out -c 45.3

test: lint
	go test ./... -covermode=count -coverprofile=coverage.out
	goornogo -i coverage.out -c 47.5

test-coverage: test
	go tool cover -html=coverage.out

mocks:
	go generate ./...
