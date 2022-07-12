# ApiGear CLI

## Install

Download and install from the release page (see [Releases](./releases))


## Manual

* Install Go 1.18.1 or later (see [Go Dev](https://go.dev))

build using

```
go build -o apigear .
```

and use apigear command line:

```
./apigear
```

## Development

You can build and run the project using the following commands:

```bash
go run main.go --help
```

To build the project:

```bash
go build .
```

To run all the tests:

```bash
go test "./..."
```

To run the debugger

```
dlv debug . -- --help
```
