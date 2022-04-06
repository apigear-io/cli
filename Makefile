.PHONY: antlr deb test build

antlr:
	antlr -Dlanguage=Go pkg/idl/parser/ObjectApi.g4

deb:
	dlv debug

test:
	go test ./...

build:
	go build .

schema:
	go run main.go yaml2json "pkg/spec/schema/*.yaml"
