GO111MODULE=on

.PHONY: all test clean build docker

build:
	export GO111MODULE on; \
	go build ./...

lint: build
	go get -u golang.org/x/lint/golint
	golint -set_exit_status builder/... engine/... examples/... ast/... pkg/... antlr/. model/...

test-short: lint
	go install github.com/newm4n/goornogo
	go test ./... -v -covermode=count -coverprofile=coverage.out -short
	goornogo -i coverage.out -c 45.3

test: lint
	go install github.com/newm4n/goornogo
	go test ./... -covermode=count -coverprofile=coverage.out
	goornogo -i coverage.out -c 47.5

test-coverage: test
	go tool cover -html=coverage.out

mocks:
	go generate ./...
