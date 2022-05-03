# Development

## External Packages

This is an incomplete list of standard packages and external packages used by the project.

Standard Packages

- http - handling of http server
- Json -handling of json encoding
- Testing - handling of tests

External Packages

- Cobra - handling of cmd line
- Viper - handing of configuration files
- Yaml - handling of yaml documents
- Chi - better handling of http routes
- Uuid - better uuid generator
- Testify - Extended test asserts to ease testing
- GoJsonSchema - json schema support to valid JSON/YAML documents
- Zap - fast structured logger
- Antlr - parser framework to parse IDL

Considered Future Packages

- Go-difflib handle of pretty printing string diffs (e.g. IDL diffs)
- Standard jsonrpc - expose rpc server to be controlled by an external app
- GRPC - expose rpc server

## Releasing

For releasing we use go-releaser

## Desktop UI

For later to place a desktop UI we plan to use [GoWails](https://wails.io). GoWails is a cross platform desktop application that provides a simple way to develop and debug web applications with an go backend.

## Web UI

The WebUI can be directly embedded in the apigear-cli using the http server and html files.
