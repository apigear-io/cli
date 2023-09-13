# ApiGear CLI

The cli repo provides a command line application (cmd/apigear/main.go) and a set of packages (pkg) to access the ApiGear API. The go documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/apigear-io/cli).

## Install

ApiGear CLI is a command line application that runs on Windows, Mac and Linux. You can download the latest version from the [release page](https://github.com/apigear-io/cli/releases/latest).

Note: _The product has not yet a certification from Microsoft, Apple or Linux. So you may need to disable the security check to run the application._

## Tasks

### Preparation

A typical development environment is:

- Install [Visual Studio Code](https://code.visualstudio.com)
- Install latest Go from [Go Dev](https://go.dev)
- Install [Taskfile](https://taskfile.dev/#/installation)

### Build

Build uses the go build command to build the command line application.

```bash
task build
```

### Run

Run just uses the go run command to run the command line application.

```bash
task run
```

### Linting

Lint uses golangci-lint (see https://golangci-lint.run/usage/install/#local-installation)

```bash
task lint
```

## Dependencies

All dependencies are defined in `go.mod`. To see why a dependency is used, see the `go mod why <package name>` command.

## Command Line

The command line is a wrapper around the library functions and is implemented using [cobra](https://github.com/spf13/cobra). The command line is defined in `cmd/apigear/main.go`.
The individual commands are defined in `pkg/cmd`.

Command line documentation is available in the [docs](docs/apigear.md) folder.

## Packages

The packages are defined in `pkg`. The packages are used by the command line and can be used by other applications, such as studio.

- `pkg/cfg` - Configuration management using viper (https://github.com/spf13/viper)
- `pkg/cmd` - Command line commands
- `pkg/gen` - Code generation using Go Text Templates (https://golang.org/pkg/text/template/) and rules document (see `pkg/spec/schema/apigear.rules.schema.yaml`)
- `pkg/git` - Git access using go-git (https://github.com/go-git/go-git), mainly for cloning template repositories
- `pkg/helper` - Various helper functions (e.g. fs, emitter, ids, json, strings, ...)
- `pkg/idl` - IDL parser and code generation using Antlr4 (https://www.antlr.org/)
- `pkg/log` - Logging using zerolog (https://github.com/rs/zerolog)
- `pkg/model` - Core API module model. All API module schemas or module IDLs are converted to this model.
- `pkg/mon` - HTTP monitoring and CSV and NDJSON feed ingestion (the server is in `pkg/net`)
- `pkg/net` - HTTP server for monitoring and olink adapter using (https://github.com/apigear-io/objectlink-core-go)
- `pkg/prj` - API project creation and management
- `pkg/repos` - SDK template repository management using git from `pkg/git`
- `pkg/sim` - Simulation engine using actions (`pkg/sim/actions`) or script (`pkg/sim/script`)
- `pkg/sol` - API solution creation and management using schemas from `pkg/spec/schema`
- `pkg/spec` - Specification and schema validation using gojsonschema (https://github.com/xeipuuv/gojsonschema)
- `pkg/tasks` - Task management using to run and watch tasks (e.g. run solution, run simulation, ...)
- `pkg/up` - Update management using self-updater (github.com/creativeprojects/go-selfupdate)
- `pkg/vfs` - Virtual file system for project creation and management, used by `pkg/prj`

## Concepts

### Command Line

The command line is a wrapper around the library functions and is implemented using [cobra](https://github.com/spf13/cobra). The command line entry point is defined in `cmd/apigear/main.go`. The command are defined in `pkg/cmd`. Each sub-command is defined in a separate package folder.

### Configuration

The configuration is managed using viper (https://github.com/spf13/viper). The configuration is stored in `~/.apigear/config.json`. The configuration is loaded in `pkg/config/config.go#init`, the init function is called automatically when the package is first used.

Note: Viper is not thread safe, so the configuration functions are protected using a mutex.

Note: The cfg package is not allowed to depend on other project packages, besides `pkg/helper`.

### Schema Development

The schema development is done using the [JSON Schema](https://json-schema.org/) specification. The schemas are defined in `pkg/spec/schema`. The schemas are validated using gojsonschema (https://github.com/xeipuuv/gojsonschema).

To update a schema edit the YAML file and run `task schema`. This will generate the JSON schema for code validation.

There are several schema files:

- `apigear.module.schema.yaml` - The main schema for the ApiGear API
- `apigear.rules.schema.yaml` - The rules schema for code generation inside sdk templates
- `apigear.solution.schema.yaml` - The solution schema to bind modules with sdk templates
- `apigear.scenario.schema.yaml` - The simulation scenario schema

Note: These schemas are re-used inside the apigear-vscode extension.

The module schema is used to validate API modules. The schema should be in sync with the module model (see `pkg/model`).

### IDL Development

The IDL provides a different API module format and is defined in `pkg/idl/parser/ObjectApi.g4`. The IDL is parsed using Antlr4 (https://www.antlr.org/). The IDL is converted to the core API module model (see `pkg/model`). The transformation is defined in `pkg/idl/listener.go`. The parser is defined in `pkg/idl/parser.go`.

When changing the IDL, the tokenizer and grammar parser needs to be regenerated. To do this, run `task antlr`.

Note: You need to have Antlr4 installed on your system.

### Code Generation

The code generation is done using Go Text Templates (https://golang.org/pkg/text/template/). The templates are defined in SDK templates who lives in external repositories and installed into a local cache inside the `~/.apigear/cache` folder.

The generator works in several steps:

- read or create a solution document and run it (see `pkg/sol/runner.go#runSolution`)
- read the input modules (YAML, JSON, IDL) (see `pkg/model`) and create a system model (see `pkg/model`)
- find the correct SDK templates (see `pkg/repos`), and try to install it when not found
- read the rules document from the SDK template (see `pkg/spec/schema/apigear.rules.schema.yaml`)
- for each document in the rules document, find the correct template and execute it based on the scope (see `pkg/gen`)
- the templates are executed using the scope as context (see `pkg/gen/generator.go#processFeature`)
- the generated files are written to the output folder (see `pkg/gen/out.go#OutputWriter`)

Note: In expert mode an in-memory solution document is created on the fly and passed to the generator. This allows to generate code without a solution document.

### SDK Templates

A SDK template is a rules document with go templates located in a templates folder. The rules document is defined in `pkg/spec/schema/apigear.rules.schema.yaml`. SDK templates are external repositories and installed into a local cache inside the `~/.apigear/cache` folder.
The installation is done using git (see `pkg/repos/cache.go` and `pkg/git`).

A registry is maintained at `https://github.com/apigear-io/template-registry`. The registry is a JSON document which lists git repositories with SDK templates.

A local copy of the registry is maintained in `~/.apigear/registry/` and is used to lookup SDK templates. The code for the registry is in `pkg/repos/registry.go`.

### Monitoring

Monitoring requires a HTTP server to receive the monitoring data. The server is defined in `pkg/net/server.go`. The http handler is defined in `pkg/net/monitor.go`. An incoming monitoring event is then emitted (`pkg/net/event.go`) using an event emitter (see `pkg/helper/emitter.go`).

To display the event you need to register an listener to the emitter and print the event content.

### Simulation

The simulation engine is defined in `pkg/sim`. The simulation engine is defined as an interface in `pkg/sim/core/engine.go`. A multi engine is used as default implementation (see `pkg/sim/core/multi.go`). The multi engine allows to run multiple simulation engines (actions, script) in parallel.

The actions based simulation engine is defined in `pkg/sim/actions/engine.go`. The actions are defined in `pkg/sim/actions/actions.go`. The actions are evaluated and the result is passed back to the caller.

The script based simulation engine is defined in `pkg/sim/script/engine.go`. The script engine is based on a JS VM (https://github.com/dop251/goja).

Note: The script is not well defined currently and needs to be improved.

### Logging

Logging is done using zerolog (https://github.com/rs/zerolog). The logging is configured in `pkg/log/logger.go`. The logging is configured to write to a file in `~/.apigear/apigear.log` and to stdout. The log file is rotated automatically.

## Release

To create a new release, we use github actions (see `.github/workflows/release.yml`). The release is created using goreleaser (https://goreleaser.com/). The release is created when a new tag is pushed to the repository.

The release configuration is defined in `.goreleaser.yaml`.
