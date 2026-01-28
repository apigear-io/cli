# ApiGear CLI Architecture Guide

This document provides a comprehensive overview of the ApiGear CLI architecture, covering project structure, package organization, core concepts, and design patterns.

## Table of Contents

1. [Overview](#overview)
2. [Project Structure](#project-structure)
3. [Package Architecture](#package-architecture)
4. [Core Data Model](#core-data-model)
5. [Key Workflows](#key-workflows)
6. [CLI Architecture](#cli-architecture)
7. [Design Patterns](#design-patterns)
8. [Technology Stack](#technology-stack)

---

## Overview

ApiGear CLI is a command-line tool for API specification, code generation, monitoring, and simulation. It enables developers to:

- **Define APIs** using IDL (Interface Definition Language) or YAML/JSON specifications
- **Generate code** for multiple target languages using customizable templates
- **Monitor** API calls in real-time
- **Simulate** API behavior using JavaScript-based simulation scripts
- **Manage projects** with templates, versioning, and sharing capabilities

### High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         CLI Commands                             │
│  (gen, mon, sim, prj, tpl, spec, cfg, x, serve, olink, mcp)     │
├─────────────────────────────────────────────────────────────────┤
│                       Domain Services                            │
│  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐   │
│  │   Gen   │ │   Sim   │ │   Mon   │ │   Prj   │ │   Tpl   │   │
│  └─────────┘ └─────────┘ └─────────┘ └─────────┘ └─────────┘   │
├─────────────────────────────────────────────────────────────────┤
│                        Core Model                                │
│  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐               │
│  │  Model  │ │   IDL   │ │  Spec   │ │   Evt   │               │
│  └─────────┘ └─────────┘ └─────────┘ └─────────┘               │
├─────────────────────────────────────────────────────────────────┤
│                    Infrastructure                                │
│  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐   │
│  │   Net   │ │ Streams │ │  Server │ │   Cfg   │ │  Helper │   │
│  └─────────┘ └─────────┘ └─────────┘ └─────────┘ └─────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

---

## Project Structure

### Directory Layout

```
apigear-io/cli/
├── cmd/                          # Application entry points
│   ├── apigear/                  # Main CLI binary
│   │   └── main.go               # Entry point
│   └── apigear-streams/          # Streams CLI binary
│       └── main.go
├── pkg/                          # Core packages (27+ packages)
│   ├── cfg/                      # Configuration management
│   ├── cmd/                      # CLI command implementations
│   ├── gen/                      # Code generation engine
│   ├── model/                    # Core API model
│   ├── idl/                      # IDL parser (ANTLR4)
│   ├── spec/                     # Specification validation
│   ├── sim/                      # Simulation engine
│   ├── mon/                      # Monitoring
│   ├── net/                      # Network management
│   ├── streams/                  # Event streaming (NATS)
│   ├── server/                   # Server orchestration
│   ├── prj/                      # Project management
│   ├── tpl/                      # Template management
│   ├── repos/                    # Template repository cache
│   ├── git/                      # Git operations
│   ├── vfs/                      # Virtual file system
│   ├── evt/                      # Event system
│   ├── helper/                   # Utility functions
│   ├── log/                      # Logging (zerolog)
│   ├── sol/                      # Solution documents
│   ├── olnk/                     # ObjectLink protocol
│   ├── mcp/                      # Model Context Protocol
│   ├── app/                      # Application utilities
│   ├── tools/                    # Miscellaneous tools
│   ├── tasks/                    # Task execution
│   └── up/                       # Self-update mechanism
├── data/                         # Static data and samples
│   ├── mon/                      # Monitoring samples
│   ├── project/                  # Project templates
│   ├── simu/                     # Simulation demos
│   ├── spec/                     # Specification schemas
│   └── template/                 # Template samples
├── examples/                     # Example projects
│   ├── counter/                  # Counter example
│   ├── sim/                      # Simulation examples
│   ├── stim/                     # Stimulus examples
│   └── tpl/                      # Template examples
├── tests/                        # Integration tests
├── docs/                         # Generated documentation
├── .github/                      # GitHub workflows
├── go.mod                        # Go module definition
├── go.sum                        # Dependency checksums
├── Taskfile.yml                  # Task automation
├── .goreleaser.yaml              # Release configuration
└── README.md                     # Project documentation
```

### Entry Points

**Primary Entry Point:** `cmd/apigear/main.go`
```go
func main() {
    info := build.NewInfo(version, commit, date)
    code := cmd.Run(info)
    os.Exit(code)
}
```

**Root Command:** `pkg/cmd/root.go`
- Initializes Cobra command hierarchy
- Registers all subcommands
- Sets up persistent flags

### Build System

**Taskfile.yml** provides common development tasks:

| Task | Description |
|------|-------------|
| `setup` | Run `go mod tidy` |
| `build` | Compile binary to `./bin/apigear` |
| `install` | Install globally |
| `lint` | Run golangci-lint |
| `test` | Run all tests |
| `test:ci` | Run tests with race detection |
| `cover` | Generate coverage report |
| `ci` | Full CI pipeline |
| `antlr` | Regenerate ANTLR parser |
| `docs` | Generate CLI documentation |

**GoReleaser** handles cross-platform releases:
- Linux (x86_64, arm64)
- macOS (x86_64, arm64)
- Windows (x86_64, arm64)

---

## Package Architecture

### Layer Overview

```
┌────────────────────────────────────────────────────────────┐
│ Layer 1: CLI Commands (pkg/cmd/*)                          │
│ Cobra command handlers, user interaction                   │
├────────────────────────────────────────────────────────────┤
│ Layer 2: Domain Services                                   │
│ gen, sim, mon, prj, tpl, spec, sol                        │
├────────────────────────────────────────────────────────────┤
│ Layer 3: Core Model                                        │
│ model, idl, evt                                            │
├────────────────────────────────────────────────────────────┤
│ Layer 4: Infrastructure                                    │
│ net, streams, server, cfg, helper, log, git, vfs          │
└────────────────────────────────────────────────────────────┘
```

### Package Descriptions

#### Core Infrastructure

| Package | Purpose | Key Types |
|---------|---------|-----------|
| `cfg` | Configuration management using Viper | Thread-safe config wrapper |
| `log` | Logging with zerolog and file rotation | Logger configuration |
| `helper` | Utilities (fs, http, strings, async) | Various helper functions |
| `git` | Git operations for project management | Clone, checkout functions |

#### Data Model

| Package | Purpose | Key Types |
|---------|---------|-----------|
| `model` | Core API module representation | `System`, `Module`, `Interface`, `Struct`, `Enum` |
| `idl` | ANTLR4-based IDL parser | `Listener`, parser/lexer |
| `spec` | Schema validation (YAML/JSON) | Document validators |
| `evt` | Event system | `Event` struct |

#### Code Generation

| Package | Purpose | Key Types |
|---------|---------|-----------|
| `gen` | Template-based code generator | `Generator`, `Options`, `Stats` |
| `gen/filters/*` | Language-specific template filters | `filtercpp`, `filtergo`, `filterjs`, etc. |
| `tpl` | Template repository management | Cache, registry operations |
| `repos` | SDK template cache | Template storage |

#### Simulation & Monitoring

| Package | Purpose | Key Types |
|---------|---------|-----------|
| `sim` | JavaScript simulation engine (Goja) | `Engine`, `World`, `ObjectService` |
| `mon` | HTTP monitoring and recording | `Event`, `EventFactory` |

#### Network & Communication

| Package | Purpose | Key Types |
|---------|---------|-----------|
| `net` | Network management | `NetworkManager`, `OlinkServer` |
| `streams` | NATS JetStream integration | `Manager`, `Controller` |
| `server` | Server orchestration | `Server` lifecycle |

#### Project Management

| Package | Purpose | Key Types |
|---------|---------|-----------|
| `prj` | Project handling | `ProjectInfo`, `DocumentInfo` |
| `sol` | Solution documents | Solution parsing |
| `vfs` | Virtual file system | Embedded demo files |

#### CLI Commands

| Package | Purpose |
|---------|---------|
| `cmd/gen` | Code generation commands |
| `cmd/mon` | Monitoring commands |
| `cmd/sim` | Simulation commands |
| `cmd/prj` | Project management commands |
| `cmd/tpl` | Template management commands |
| `cmd/spec` | Specification validation commands |
| `cmd/cfg` | Configuration commands |
| `cmd/x` | Experimental/utility commands |
| `cmd/stim` | Stimulus commands |
| `cmd/olink` | ObjectLink REPL commands |

---

## Core Data Model

### Model Hierarchy

```
System
└── Module[]
    ├── name, version, description
    ├── imports[]
    ├── externs[]
    ├── interfaces[]
    │   ├── name, description
    │   ├── properties[]
    │   │   └── name, type (Schema)
    │   ├── operations[]
    │   │   ├── name, params[], return type
    │   │   └── Schema for each param/return
    │   └── signals[]
    │       └── name, params[]
    ├── structs[]
    │   └── fields[]
    │       └── name, type (Schema)
    └── enums[]
        └── members[]
            └── name, value
```

### Base Types

**NamedNode** - Base for all named entities:
```go
type NamedNode struct {
    Name        string
    Kind        string
    Description string
    Meta        map[string]any
}
```

**TypedNode** - Extends NamedNode with type information:
```go
type TypedNode struct {
    NamedNode
    Schema Schema
}
```

### Type System

**Primitive Types:**
- `void`, `bool`, `int`, `int32`, `int64`
- `float`, `float32`, `float64`
- `string`, `bytes`, `any`

**Symbol Types:**
- `enum` - Enumeration reference
- `struct` - Structure reference
- `interface` - Interface reference

**Schema Properties:**
```go
type Schema struct {
    Type      string      // Primitive or symbol type name
    Module    string      // Module containing the type
    IsArray   bool        // Array type flag
    IsPrimitive bool      // Primitive type flag
    IsSymbol  bool        // Symbol type flag
    KindType  string      // Kind of symbol (enum/struct/interface)
}
```

### Model Visitor Pattern

The `ModelVisitor` interface enables traversal of the model hierarchy:

```go
type ModelVisitor interface {
    VisitSystem(s *System) error
    VisitModule(m *Module) error
    VisitExtern(e *Extern) error
    VisitInterface(i *Interface) error
    VisitOperation(o *Operation) error
    VisitSignal(g *Signal) error
    VisitProperty(p *Property) error
    VisitStruct(s *Struct) error
    VisitStructField(f *TypedNode) error
    VisitEnum(e *Enum) error
    VisitEnumMember(m *EnumMember) error
}
```

Used for:
- Type validation and resolution
- Reserved word checking
- Code generation traversal

---

## Key Workflows

### Code Generation Pipeline

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│ Read Source │───▶│ Parse/Load  │───▶│  Validate   │
│ (IDL/YAML)  │    │ (idl/spec)  │    │  (spec)     │
└─────────────┘    └─────────────┘    └─────────────┘
                                             │
                                             ▼
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│ Write Files │◀───│  Execute    │◀───│ Load Rules  │
│  (output)   │    │ Templates   │    │ & Templates │
└─────────────┘    └─────────────┘    └─────────────┘
```

1. **Read Source** - Load IDL or YAML/JSON API specifications
2. **Parse/Load** - Convert to internal model using `idl` or `spec` packages
3. **Validate** - Validate against JSON schemas
4. **Load Rules** - Read generation rules document
5. **Execute Templates** - Apply Go templates with language-specific filters
6. **Write Files** - Output generated code to target directory

**Generator Options:**
```go
type Options struct {
    OutputDir    string
    TemplatesDir string
    System       *model.System
    Features     []string
    Force        bool
    DryRun       bool
    Meta         map[string]any
}
```

### Simulation Engine Flow

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│ Load Script │───▶│ Create Goja │───▶│ Register    │
│    (.js)    │    │   Runtime   │    │  World API  │
└─────────────┘    └─────────────┘    └─────────────┘
                                             │
                                             ▼
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Events    │◀───│   Execute   │◀───│   Create    │
│  via OLink  │    │   Script    │    │  Services   │
└─────────────┘    └─────────────┘    └─────────────┘
```

1. **Load Script** - Read JavaScript simulation file
2. **Create Runtime** - Initialize Goja JavaScript engine
3. **Register World API** - Expose `$createService`, `$createChannel`, etc.
4. **Create Services** - Script creates simulated API services
5. **Execute Script** - Run simulation logic
6. **Events via OLink** - Communicate with clients over ObjectLink protocol

**World API:**
```javascript
// Available in simulation scripts
$createService(name)    // Create a service proxy
$createClient(name)     // Create a client proxy
$createChannel(name)    // Create a communication channel
```

### Monitoring & Event Streaming

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│ API Events  │───▶│   HTTP/WS   │───▶│   NATS      │
│  (calls)    │    │   Server    │    │  JetStream  │
└─────────────┘    └─────────────┘    └─────────────┘
                                             │
                                             ▼
                   ┌─────────────┐    ┌─────────────┐
                   │   Export    │◀───│   Record    │
                   │ (CSV/NDJSON)│    │  Sessions   │
                   └─────────────┘    └─────────────┘
```

**Server Ports:**
- HTTP Server: `:5555` (REST API, monitoring)
- WebSocket: `:5555/ws` (ObjectLink protocol)
- NATS Server: `:4222` (message bus with JetStream)

**Event Structure:**
```go
type Event struct {
    Id        string
    Device    string
    Type      string    // "call", "signal", "state"
    Symbol    string
    Timestamp time.Time
    Data      Payload
}
```

---

## CLI Architecture

### Command Framework

The CLI uses **Cobra** for command structure and **Viper** for configuration.

### Command Hierarchy

```
apigear
├── serve              # Start server for monitoring/simulation
├── generate (gen)     # Generate code from APIs
│   ├── expert (x)     # Expert mode with flags
│   └── solution (sol) # Generate from solution document
├── monitor (mon)      # Display/record API calls
├── config (cfg)       # Display/edit configuration
├── simulate (sim)     # Simulate API behavior
├── stimulate (stim)   # Stimulate API services
├── spec (s)           # Load and validate specs
├── project (prj)      # Manage projects
│   ├── create         # Create new project
│   ├── add            # Add document to project
│   ├── edit           # Edit project
│   ├── info           # Display project info
│   ├── import         # Import project
│   ├── open           # Open project
│   ├── pack           # Pack project
│   ├── recent         # Show recent projects
│   └── share          # Share project
├── template (tpl)     # Manage templates
│   ├── list (ls)      # List templates
│   ├── install (i)    # Install template
│   ├── update         # Update template
│   ├── info           # Template information
│   ├── cache          # List cached templates
│   ├── remove         # Remove from cache
│   ├── clean          # Clean cache
│   ├── import         # Import template
│   ├── create         # Create template
│   ├── lint           # Lint template
│   └── publish        # Publish template
├── x                  # Experimental commands
│   ├── yaml2json      # Convert YAML to JSON
│   ├── idl2yaml       # Convert IDL to YAML
│   └── wscat          # WebSocket client
├── update             # Update the program
├── version            # Display version
├── olink (ol)         # ObjectLink REPL
├── mcp                # Start MCP server
└── stream             # Manage message streams
```

### Command Implementation Pattern

```go
// Standard command structure
func NewExampleCommand() *cobra.Command {
    var options struct {
        input  string
        output string
        force  bool
    }

    cmd := &cobra.Command{
        Use:     "example",
        Short:   "Short description",
        Long:    "Long description with details",
        Args:    cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            // Implementation
            return nil
        },
    }

    // Define flags
    cmd.Flags().StringVarP(&options.input, "input", "i", "", "Input file")
    cmd.Flags().StringVarP(&options.output, "output", "o", ".", "Output directory")
    cmd.Flags().BoolVarP(&options.force, "force", "f", false, "Force overwrite")

    // Mark required flags
    cmd.MarkFlagRequired("input")

    return cmd
}
```

### Flag Patterns

| Type | Example |
|------|---------|
| String | `--input, -i` |
| Bool | `--force, -f` |
| Int | `--port, -p` |
| StringSlice | `--features, -f` |
| Duration | `--timeout` |
| Persistent | Applies to subcommands |

### Context and Signal Handling

Commands support cancellation and signal handling:

```go
func withSignalContext(ctx context.Context, fn func(context.Context) error) error {
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    // Handle interrupt signals
    go func() {
        sigCh := make(chan os.Signal, 1)
        signal.Notify(sigCh, os.Interrupt)
        <-sigCh
        cancel()
    }()

    return fn(ctx)
}
```

---

## Design Patterns

### Visitor Pattern
**Location:** `pkg/model/visitor.go`

Used for traversing the model hierarchy for validation, code generation, and analysis.

```go
type ModelVisitor interface {
    VisitSystem(s *System) error
    VisitModule(m *Module) error
    // ... other visit methods
}

func WalkModule(m *Module, v ModelVisitor) error {
    if err := v.VisitModule(m); err != nil {
        return err
    }
    for _, iface := range m.Interfaces {
        if err := WalkInterface(iface, v); err != nil {
            return err
        }
    }
    // ... walk other elements
    return nil
}
```

### Factory Pattern
**Location:** `pkg/mon/event.go`, `pkg/model/`

Creates events and model nodes with proper initialization.

```go
type EventFactory struct {
    device string
}

func (f *EventFactory) NewCallEvent(symbol string, data Payload) *Event {
    return &Event{
        Id:        helper.NewID(),
        Device:    f.device,
        Type:      "call",
        Symbol:    symbol,
        Timestamp: time.Now(),
        Data:      data,
    }
}
```

### Manager Pattern
**Location:** `pkg/server/`, `pkg/net/`, `pkg/streams/`

Manages lifecycle of complex components with startup/shutdown handling.

```go
type Server struct {
    network  *net.NetworkManager
    streams  *streams.Manager
    sim      *sim.Manager
}

func (s *Server) Start(ctx context.Context) error {
    if err := s.network.Start(ctx); err != nil {
        return err
    }
    if err := s.streams.Start(ctx); err != nil {
        return err
    }
    return nil
}

func (s *Server) Stop() error {
    s.streams.Stop()
    s.network.Stop()
    return nil
}
```

### Strategy Pattern
**Location:** `pkg/gen/filters/`

Language-specific code generation filters implement common interfaces.

```go
// Each filter package provides language-specific template functions
// pkg/gen/filters/filtergo/
// pkg/gen/filters/filtercpp/
// pkg/gen/filters/filterjs/
// etc.
```

### Builder Pattern
**Location:** `pkg/idl/listener.go`

Builds the model from parsed AST incrementally.

```go
type Listener struct {
    system  *model.System
    module  *model.Module
    current interface{}
}

func (l *Listener) EnterModule(ctx *parser.ModuleContext) {
    l.module = &model.Module{
        Name: ctx.Identifier().GetText(),
    }
    l.system.Modules = append(l.system.Modules, l.module)
}
```

### Proxy Pattern
**Location:** `pkg/sim/`

Service proxies for JavaScript integration.

### Adapter Pattern
**Location:** `pkg/net/`

Protocol adapters (OLink, WebSocket) adapt between different communication protocols.

---

## Technology Stack

| Category | Technology | Version/Notes |
|----------|------------|---------------|
| Language | Go | 1.25.0 |
| CLI Framework | Cobra | v1.10.1 |
| Configuration | Viper | v1.21.0 |
| Parsing | ANTLR4 | IDL grammar |
| Schema Validation | gojsonschema | JSON Schema |
| JavaScript VM | Goja | Simulation scripts |
| Message Bus | NATS | JetStream enabled |
| Logging | zerolog | With lumberjack rotation |
| WebSocket | gorilla/websocket | Protocol communication |
| HTTP Router | go-chi | REST API |
| Git | go-git | v5 |
| Testing | testify | Assertions |

### External Dependencies

Key dependencies from `go.mod`:

```
github.com/spf13/cobra          # CLI framework
github.com/spf13/viper          # Configuration
github.com/apigear-io/objectlink-core-go  # ObjectLink protocol
github.com/dop251/goja          # JavaScript engine
github.com/go-git/go-git/v5     # Git operations
github.com/nats-io/nats-server/v2  # Message bus
github.com/gorilla/websocket    # WebSocket
github.com/rs/zerolog           # Logging
github.com/mark3labs/mcp-go     # MCP protocol
github.com/antlr4-go/antlr/v4   # Parser generator
github.com/xeipuuv/gojsonschema # JSON Schema validation
```

---

## Configuration

### Configuration Storage

Location: `~/.apigear/config.json`

### Configuration Keys

| Key | Description |
|-----|-------------|
| `recent` | Recent project paths |
| `server_port` | Default server port |
| `editor_command` | Editor for opening files |
| `update_channel` | Update channel (stable/beta) |
| `templates_dir` | Template cache directory |
| `registry_dir` | Registry directory |
| `registry_url` | Template registry URL |
| `version` | Current version |

### Thread-Safe Access

Configuration is accessed through a thread-safe wrapper in `pkg/cfg`:

```go
func Get(key string) any
func Set(key string, value any)
func GetString(key string) string
func GetStringSlice(key string) []string
```

---

## Further Reading

- [README.md](README.md) - Quick start guide
- [examples/](examples/) - Example projects
- [data/spec/](data/spec/) - Specification schemas
- [API Documentation](https://apigear.io/docs) - Online documentation
