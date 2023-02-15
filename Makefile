.PHONY: antlr deb test build check cover install


export CGO_ENABLED=0

export CGO_ENABLED=0

test:
	go test ./...


build:
	go build -o ./bin/apigear ./cmd/apigear

run:
	go run github.com/apigear-io/cli/cmd/apigear

watch:
	air

debug:
	dlv debug cmd/apigear/main.go

antlr:
	antlr -Dlanguage=Go pkg/idl/parser/ObjectApi.g4

schema:
	go run main.go yaml2json "pkg/spec/schema/*.yaml"

deps:
	@echo "https://go.dev/doc/install"
	@echo "https://goreleaser.com/install/"
	@echo "https://github.com/cosmtrek/air"
	@echo "https://golangci-lint.run/usage/install/"
	@echo "https://github.com/gotestyourself/gotestsum"

check:
	golangci-lint run

cover:
	go test -covermode=count -coverprofile=coverage.out -coverpkg=apigear/... ./...
	go tool cover -func coverage.out

install:
	go install github.com/apigear-io/cli/cmd/apigear
