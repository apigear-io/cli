export CGO_ENABLED=0

.PHONY: test
test:
	go test ./...


.PHONY: build
build:
	go build -o ./bin/apigear ./cmd/apigear

.PHONY: run
run:
	go run github.com/apigear-io/cli/cmd/apigear

.PHONY: watch
watch:
	air

.PHONY: deb
debug:
	dlv debug cmd/apigear/main.go

.PHONY: antlr
antlr:
	antlr -Dlanguage=Go pkg/idl/parser/ObjectApi.g4

.PHONY: schema
schema:
	go run main.go yaml2json "pkg/spec/schema/*.yaml"

.PHONY: deps
deps:
	@echo "https://go.dev/doc/install"
	@echo "https://goreleaser.com/install/"
	@echo "https://github.com/cosmtrek/air"
	@echo "https://golangci-lint.run/usage/install/"
	@echo "https://go.dev/blog/vuln"

.PHONY: lint
lint:
	golangci-lint run

.PHONY: vuln
vuln:
	govulncheck ./...

.PHONY: cover
cover:
	go test -covermode=count -coverprofile=coverage.out -coverpkg=apigear/... ./...
	go tool cover -func coverage.out
