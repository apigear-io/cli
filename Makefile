.PHONY: antlr deb test build check


test:
	go test ./...

build:
	go build .

watch:
	air
run:
	go run main.go

debug:
	dlv debug

antlr:
	antlr -Dlanguage=Go pkg/idl/parser/ObjectApi.g4

schema:
	go run main.go yaml2json "pkg/spec/schema/*.yaml"

deps:
	@echo "https://go.dev/doc/install"
	@echo "https://goreleaser.com/install/"
	@echo "https://github.com/cosmtrek/air"
	@echo "https://golangci-lint.run/usage/install/"

check:
	golangci-lint run