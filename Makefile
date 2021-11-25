GO111MODULE=on

.PHONY: all test test-short fix-antlr4-bug build

fix-antlr4-bug:
	sed -i.origin.bak "s/1[ ]*<</int64(1) <</g" antlr/parser/grulev3/grulev3_parser.go

build: fix-antlr4-bug
	export GO111MODULE on; \
	go build ./...

lint: build
	${GOBIN}/golint -set_exit_status builder/... engine/... examples/... ast/... pkg/... antlr/. model/...

test-short: build
	go test ./... -v -covermode=count -coverprofile=coverage.out -short

test: build
	go test ./... -covermode=count -coverprofile=coverage.out

test-coverage: test
	go tool cover -html=coverage.out

mocks:
	go generate ./...
