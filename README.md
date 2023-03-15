# ApiGear CLI

## Install

ApiGear CLI is a command line application that runs on Windows, Mac and Linux. You can download the latest version from [here](https://github.com/apigear-io/cli-releases/releases/latest).

Note: *The product has not yet a certification from Microsoft, Apple or Linux. So you may need to disable the security check to run the application.*


## Building from source

Install latest Go (see [Go Dev](https://go.dev))

build using

```
go build ./cmd/apigear
```

and use apigear command line:

```
./apigear --help
```

## Development


For development you need some environment variables:

```
export APIGEAR_GIT_PUBLIC_TOKEN=your_token
export APIGEAR_GIT_AUTH_TOKEN=your_token
```

The git public token provides access to public repositories. The git auth token provides access to private template repositories.

You can build and run the project using the following commands:

```bash
go run cmd/apigear/main.go --help
```

To build the project:

```bash
go build ./cmd/apigear
```
```

To run all the tests:

```bash
go test "./..."
```

To run the debugger

```
dlv debug cmd/apigear/main.go -- --help
```


## Program Settings

Settings are stored in the file `~/.apigear/config.json`. The file is created automatically when the first command is executed.

## Log Files

Log files are stored in the file `~/.apigear/*.log`. Outdated log-files are automatically deleted.

## Template Cache

The template cache is stored in the file `~/.apigear/templates`. The file is created automatically when a template is downloaded from a git-repository.

## Template Registry

The template registry is stored in the file `~/.apigear/registry/registry.json`. The file is created automatically when the registry is updated. The registry is currently hosted on github.