# ApiGear CLI Architecture Guide

This document provides a comprehensive overview of the ApiGear CLI architecture, covering project structure, package organization, core concepts, and design patterns.

**Last Updated:** 2026-02-09 (after domain-based reorganization)

## Table of Contents

1. [Overview](#overview)
2. [Project Structure](#project-structure)
3. [Domain Architecture](#domain-architecture)
4. [Core Data Model](#core-data-model)
5. [Key Workflows](#key-workflows)
6. [CLI Architecture](#cli-architecture)
7. [Design Patterns](#design-patterns)
8. [Technology Stack](#technology-stack)
9. [Future Architecture](#future-architecture)

---

## Overview

ApiGear CLI is a command-line tool for API specification, code generation, and monitoring. It enables developers to:

- **Define APIs** using IDL (Interface Definition Language) or YAML/JSON specifications
- **Generate code** for multiple target languages using customizable templates
- **Monitor** API calls in real-time
- **Manage projects** with templates, versioning, and sharing capabilities

### High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         CLI Commands                             │
│  (gen, mon, prj, tpl, spec, cfg, x, olink, mcp)                 │
├─────────────────────────────────────────────────────────────────┤
│                      Domain Services                             │
│  ┌────────────┐ ┌────────────┐ ┌────────────┐ ┌────────────┐   │
│  │  ObjModel  │ │  Codegen   │ │Orchestrate │ │  Runtime   │   │
│  │ (API Spec) │ │(Templates) │ │(Solutions) │ │ (Monitor)  │   │
│  └────────────┘ └────────────┘ └────────────┘ └────────────┘   │
├─────────────────────────────────────────────────────────────────┤
│                      Foundation Layer                            │
│  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐   │
│  │  Config │ │ Logging │ │   Git   │ │   VFS   │ │  Tasks  │   │
│  └─────────┘ └─────────┘ └─────────┘ └─────────┘ └─────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

---

## Project Structure

### Directory Layout

```
apigear-io/cli/
├── cmd/                          # Application entry point
│   └── apigear/                  # Main CLI binary
│       └── main.go               # Entry point
│
├── pkg/                          # Core packages (5 domains + CLI + MCP)
│   ├── foundation/               # 🏗️ Foundation - Shared Infrastructure
│   │   ├── *.go                  # Core utilities (fs, http, strings, async)
│   │   ├── config/               # Configuration (Viper wrapper)
│   │   ├── logging/              # Logging (zerolog + rotation)
│   │   ├── git/                  # Git operations
│   │   ├── vfs/                  # Virtual file system
│   │   ├── tasks/                # Task execution framework
│   │   ├── tools/                # Low-level tools
│   │   └── updater/              # Self-update mechanism
│   │
│   ├── objmodel/                 # 📐 ObjectAPI Model - API Specification
│   │   ├── *.go                  # System, Module, Interface, Struct, Enum
│   │   ├── idl/                  # IDL parser (ANTLR4)
│   │   │   ├── parser/           # Generated parser/lexer
│   │   │   └── *.go              # Listener, helper functions
│   │   └── spec/                 # Specification validation
│   │       ├── schema/           # JSON schemas
│   │       └── rkw/              # Reserved keywords
│   │
│   ├── codegen/                  # ⚙️ Code Generation - Templates & Generation
│   │   ├── *.go                  # Generator, rules engine
│   │   ├── filters/              # Language-specific filters
│   │   │   ├── common/           # Shared filter functions
│   │   │   ├── filtercpp/        # C++ filters
│   │   │   ├── filtergo/         # Go filters
│   │   │   ├── filterjs/         # JavaScript filters
│   │   │   ├── filterts/         # TypeScript filters
│   │   │   ├── filterpy/         # Python filters
│   │   │   ├── filterqt/         # Qt filters
│   │   │   ├── filterrs/         # Rust filters
│   │   │   └── filterue/         # Unreal Engine filters
│   │   ├── template/             # Template operations
│   │   └── registry/             # Template registry & cache
│   │
│   ├── orchestration/            # 🎯 Orchestration - High-level Workflows
│   │   ├── solution/             # Solution execution
│   │   └── project/              # Project management
│   │
│   ├── runtime/                  # 🔄 Runtime - Monitoring & Services
│   │   ├── monitoring/           # Event monitoring & recording
│   │   ├── events/               # Event bus (stub after NATS removal)
│   │   ├── network/              # HTTP/WebSocket network layer
│   │   ├── simulation/           # API simulation
│   │   └── streams/              # Event streaming (under development)
│   │
│   ├── cmd/                      # CLI command implementations
│   │   ├── gen/                  # Generate commands
│   │   ├── mon/                  # Monitor commands
│   │   ├── prj/                  # Project commands
│   │   ├── tpl/                  # Template commands
│   │   ├── spec/                 # Spec commands
│   │   ├── cfg/                  # Config commands
│   │   ├── x/                    # Experimental commands
│   │   └── olink/                # ObjectLink REPL
│   │
│   └── mcp/                      # Model Context Protocol server
│       ├── gen/                  # MCP generation tools
│       ├── spec/                 # MCP spec tools
│       └── tpl/                  # MCP template tools
│
├── internal/                     # Private application code
│   └── (reserved for future REST API server implementation)
│
├── data/                         # Static data and samples
│   ├── mon/                      # Monitoring samples
│   ├── project/                  # Project templates
│   ├── spec/                     # Specification schemas
│   └── template/                 # Template samples
│
├── examples/                     # Example projects
│   ├── counter/                  # Counter example
│   └── tpl/                      # Template examples
│
├── tests/                        # Integration tests
├── docs/                         # Documentation
│   ├── ARCHITECTURE.md           # This document
│   └── ARCHITECTURE-REST-WEB.md  # REST API + Web UI plan
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

## Domain Architecture

### Architectural Principles

The codebase follows a **domain-based architecture** with clear separation of concerns:

1. **Foundation Layer** - Shared infrastructure with no business logic
2. **Domain Layers** - Business logic organized by domain boundaries
3. **CLI Layer** - User interface commands
4. **Clean Dependencies** - Unidirectional dependency flow

### Dependency Hierarchy

```
┌────────────────────────────────────────────────────────────┐
│ Layer 1: CLI Commands (pkg/cmd/*)                          │
│ User interface, Cobra command handlers                     │
├────────────────────────────────────────────────────────────┤
│ Layer 2: Orchestration & Runtime                           │
│ orchestration/solution, orchestration/project              │
│ runtime/monitoring, runtime/network, runtime/simulation    │
├────────────────────────────────────────────────────────────┤
│ Layer 3: Code Generation                                   │
│ codegen (generator, filters, template, registry)           │
├────────────────────────────────────────────────────────────┤
│ Layer 4: ObjectAPI Model                                   │
│ objmodel (model, idl, spec)                                │
├────────────────────────────────────────────────────────────┤
│ Layer 5: Foundation                                        │
│ foundation (config, logging, git, vfs, tasks, tools)       │
└────────────────────────────────────────────────────────────┘
```

**Dependency Rules:**
- Higher layers can depend on lower layers
- Lower layers CANNOT depend on higher layers
- No circular dependencies between domains
- Foundation has zero dependencies on other domains

### Domain Descriptions

#### 1. Foundation Domain (`pkg/foundation/`)

**Purpose:** Shared infrastructure used by all other domains.

**Key Packages:**

| Package | Purpose | Key Types |
|---------|---------|-----------|
| `foundation` | Core utilities (fs, http, strings, async, ids) | Helper functions |
| `foundation/config` | Configuration management (Viper wrapper) | Thread-safe config access |
| `foundation/logging` | Logging with zerolog and file rotation | Logger, EventWriter, Rotator |
| `foundation/git` | Git operations (clone, checkout, tags) | Git helper functions |
| `foundation/vfs` | Virtual file system with embedded demos | Demo files |
| `foundation/tasks` | Task execution framework | Manager, Task |
| `foundation/tools` | Low-level tools (colorwriter, hooks) | ColorWriter |
| `foundation/updater` | Self-update mechanism | Updater |

**Dependencies:** None (bottom layer)

#### 2. ObjectAPI Model Domain (`pkg/objmodel/`)

**Purpose:** Define, parse, and validate ObjectAPI specifications.

**Key Packages:**

| Package | Purpose | Key Types |
|---------|---------|-----------|
| `objmodel` | Core API model | `System`, `Module`, `Interface`, `Struct`, `Enum` |
| `objmodel/idl` | ANTLR4-based IDL parser | `Parser`, `Listener`, AST builder |
| `objmodel/spec` | YAML/JSON specification validation | Schema validators, rules |
| `objmodel/spec/rkw` | Reserved keyword checking | Reserved word lists |

**Dependencies:** `foundation`

**Note:** Named `objmodel` (not `apimodel`) to avoid confusion with future REST API models.

#### 3. Code Generation Domain (`pkg/codegen/`)

**Purpose:** Generate source code from ObjectAPI models using templates.

**Key Packages:**

| Package | Purpose | Key Types |
|---------|---------|-----------|
| `codegen` | Template-based code generator | `Generator`, `Options`, `Stats` |
| `codegen/filters/*` | Language-specific template filters | 12 language filters |
| `codegen/filters/common` | Shared filter functions | String/array helpers |
| `codegen/filters/filtercpp` | C++ template filters | Type conversions, namespaces |
| `codegen/filters/filtergo` | Go template filters | Type conversions, packages |
| `codegen/filters/filterjs` | JavaScript template filters | Type conversions |
| `codegen/filters/filterts` | TypeScript template filters | Type conversions |
| `codegen/filters/filterpy` | Python template filters | Type conversions |
| `codegen/filters/filterqt` | Qt/QML template filters | Qt type conversions |
| `codegen/filters/filterrs` | Rust template filters | Type conversions |
| `codegen/filters/filterue` | Unreal Engine filters | UE4/5 type conversions |
| `codegen/template` | Template operations | Create, publish templates |
| `codegen/registry` | Template registry & cache | Registry, cache management |

**Dependencies:** `foundation`, `objmodel`

#### 4. Orchestration Domain (`pkg/orchestration/`)

**Purpose:** Orchestrate high-level workflows for building solutions and managing projects.

**Key Packages:**

| Package | Purpose | Key Types |
|---------|---------|-----------|
| `orchestration/solution` | Solution document execution | Runner, parser |
| `orchestration/project` | Project lifecycle management | ProjectInfo, DocumentInfo |

**Dependencies:** `foundation`, `objmodel`, `codegen`

#### 5. Runtime Domain (`pkg/runtime/`)

**Purpose:** Runtime services for monitoring, networking, simulation, and event streaming.

**Key Packages:**

| Package | Purpose | Key Types |
|---------|---------|-----------|
| `runtime/monitoring` | Event monitoring & recording | Event, EventFactory |
| `runtime/events` | Event bus (stub, NATS removed) | IEventBus interface |
| `runtime/network` | HTTP/WebSocket network layer | NetworkManager, OlinkServer |
| `runtime/simulation` | API simulation engine | Manager |
| `runtime/streams` | Event streaming (under development) | Manager |

**Dependencies:** `foundation`, `objmodel`

**Note:** Some packages are stubs after NATS removal, awaiting redesign.

#### 6. CLI Commands (`pkg/cmd/`)

**Purpose:** User-facing command implementations.

| Package | Purpose |
|---------|---------|
| `cmd/gen` | Code generation commands |
| `cmd/mon` | Monitoring commands |
| `cmd/prj` | Project management commands |
| `cmd/tpl` | Template management commands |
| `cmd/spec` | Specification validation commands |
| `cmd/cfg` | Configuration commands |
| `cmd/x` | Experimental/utility commands |
| `cmd/olink` | ObjectLink REPL commands |

**Dependencies:** All domains (top layer)

#### 7. Model Context Protocol (`pkg/mcp/`)

**Purpose:** MCP server for AI agent integration.

| Package | Purpose |
|---------|---------|
| `mcp/gen` | MCP generation tools |
| `mcp/spec` | MCP spec tools |
| `mcp/tpl` | MCP template tools |

**Dependencies:** All domains

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

**Location:** `pkg/objmodel/visitor.go`

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
├── generate (gen)     # Generate code from APIs
│   ├── expert (x)     # Expert mode with flags
│   └── solution (sol) # Generate from solution document
├── monitor (mon)      # Display/record API calls
├── config (cfg)       # Display/edit configuration
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
**Location:** `pkg/objmodel/visitor.go`

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
**Location:** `pkg/runtime/monitoring/event.go`, `pkg/objmodel/`

Creates events and model nodes with proper initialization.

```go
type EventFactory struct {
    device string
}

func (f *EventFactory) NewCallEvent(symbol string, data Payload) *Event {
    return &Event{
        Id:        foundation.NewID(),
        Device:    f.device,
        Type:      "call",
        Symbol:    symbol,
        Timestamp: time.Now(),
        Data:      data,
    }
}
```

### Manager Pattern
**Location:** `pkg/runtime/network/`, `pkg/runtime/streams/`

Manages lifecycle of complex components with startup/shutdown handling.

```go
type NetworkManager struct {
    server   *http.Server
    // ...
}

func (m *NetworkManager) Start(ctx context.Context) error {
    // Start HTTP server
    return m.server.ListenAndServe()
}

func (m *NetworkManager) Stop() error {
    // Graceful shutdown
    return m.server.Shutdown(context.Background())
}
```

### Strategy Pattern
**Location:** `pkg/codegen/filters/`

Language-specific code generation filters implement common interfaces.

```go
// Each filter package provides language-specific template functions
// pkg/codegen/filters/filtergo/
// pkg/codegen/filters/filtercpp/
// pkg/codegen/filters/filterjs/
// pkg/codegen/filters/filterts/
// pkg/codegen/filters/filterpy/
// etc. (12 language filters total)
```

### Builder Pattern
**Location:** `pkg/objmodel/idl/listener.go`

Builds the model from parsed AST incrementally.

```go
type Listener struct {
    system  *objmodel.System
    module  *objmodel.Module
    current interface{}
}

func (l *Listener) EnterModule(ctx *parser.ModuleContext) {
    l.module = &objmodel.Module{
        Name: ctx.Identifier().GetText(),
    }
    l.system.Modules = append(l.system.Modules, l.module)
}
```

### Adapter Pattern
**Location:** `pkg/runtime/network/`

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

Configuration is accessed through a thread-safe wrapper in `pkg/foundation/config`:

```go
func Get(key string) any
func Set(key string, value any)
func GetString(key string) string
func GetStringSlice(key string) []string
```

---

## Future Architecture

### REST API + Web UI (Planned)

**Status:** Design phase (see `docs/ARCHITECTURE-REST-WEB.md`)

The CLI will be extended with a REST API server and React-based web UI:

```
┌─────────────────────────────────────────────────────────────┐
│                    Web UI (React + Vite)                     │
│  Features: codegen, templates, specs, projects              │
├─────────────────────────────────────────────────────────────┤
│                REST API Server (apigear serve)               │
│  internal/server/ - Chi router + http.HandlerFunc           │
│  internal/restmodel/ - REST DTOs                            │
├─────────────────────────────────────────────────────────────┤
│              Existing Domain Services (pkg/)                 │
│  Reused by both CLI and REST API                            │
└─────────────────────────────────────────────────────────────┘
```

**Key Decisions:**
- Server runs as `apigear serve` subcommand (not separate binary)
- Chi router with stdlib `http.HandlerFunc` pattern
- Swag for OpenAPI generation (annotations in code)
- AI-written TypeScript SDKs (not codegen)
- Separate `restmodel` package for REST DTOs (avoid confusion with `objmodel`)

**Directory Structure:**
```
pkg/cmd/serve/           # Serve subcommand
internal/server/         # HTTP handlers and router
internal/restmodel/      # REST API DTOs
web/                     # Vite + React frontend
  src/
    api/                 # TypeScript SDK (AI-written)
    features/            # Feature modules
    components/          # Shared components
docs/swagger/            # Auto-generated OpenAPI specs
```

**Benefits:**
- Single binary distribution
- Reuses all existing domain logic
- Type-safe APIs (Go + TypeScript)
- Auto-generated documentation
- Parallel frontend/backend development

---

## Package Reorganization History

### February 2026 - Domain-Based Consolidation

The package structure was reorganized from 23 fragmented packages into 5 logical domains:

**Before:**
- 23 small packages: `helper`, `cfg`, `log`, `git`, `vfs`, `tasks`, `tools`, `up`, `model`, `idl`, `spec`, `gen`, `tpl`, `repos`, `sol`, `prj`, `mon`, `evt`, `net`, `sim`, `streams`, etc.
- Imports scattered across many paths
- Unclear boundaries between concerns

**After:**
- 5 domains: `foundation`, `objmodel`, `codegen`, `orchestration`, `runtime`
- Clear dependency hierarchy
- Better code discoverability
- Easier to work on isolated features

**Migration:**
- All import paths updated (1000+ changes)
- Package declarations updated
- No circular dependencies
- All tests passing

---

## Further Reading

- [README.md](README.md) - Quick start guide
- [ARCHITECTURE-REST-WEB.md](docs/ARCHITECTURE-REST-WEB.md) - REST API + Web UI plan
- [examples/](examples/) - Example projects
- [data/spec/](data/spec/) - Specification schemas
- [API Documentation](https://apigear.io/docs) - Online documentation
