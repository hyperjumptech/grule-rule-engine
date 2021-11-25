GO111MODULE=on
LINTER="${GOPATH}/bin/golint"

.PHONY: all test test-short fix-antlr4-bug build

fix-antlr4-bug:
	sed -i.origin.bak "s/1[ ]*<</int64(1) <</g" antlr/parser/grulev3/grulev3_parser.go

build: fix-antlr4-bug
	export GO111MODULE on; \
	go build ./...

lint: build
	go get -d golang.org/x/lint/golint
	go install golang.org/x/lint/golint
	${LINTER} -set_exit_status builder/... engine/... examples/... ast/... pkg/... antlr/. model/...

test-short: lint
	go test ./... -v -covermode=count -coverprofile=coverage.out -short

test: lint
	go test ./... -covermode=count -coverprofile=coverage.out

test-coverage: test
	go tool cover -html=coverage.out

mocks:
	go generate ./...
