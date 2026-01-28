# Modular Architecture Proposal

This document proposes refactoring the monolithic CLI into independent apps that communicate through interfaces.

**Two approaches are explored:**
1. [Go Interfaces Approach](#proposed-architecture) - Apps as Go packages with interfaces
2. [REST API Approach](#alternative-rest-api-architecture) - Apps as web services shared by CLI and Studio

## Current State

```
cmd ─┬─> gen ─┬─> spec ─┬─> model ─┬─> cfg ──> helper
     │        │         │          │
     │        │         ├─> idl ───┤
     │        │         │          │
     │        ├─> sol ──┤          ├─> log ──> cfg, helper
     │        │         │          │
     │        ├─> repos ┴─> git ───┤
     │        │                    │
     ├─> sim ─┴─> net ─> mon ──────┘
     │
     ├─> prj ──> git, vfs
     │
     ├─> mcp (combines gen + spec + repos)
     │
     └─> up, tpl, tasks
```

### Current Dependencies (simplified)

| Package | Direct Dependencies |
|---------|---------------------|
| `helper` | (none) |
| `vfs` | (none) |
| `evt` | (none) |
| `cfg` | helper |
| `log` | cfg, helper |
| `git` | cfg, helper, log |
| `model` | cfg, helper, log |
| `idl` | cfg, helper, log, model |
| `mon` | cfg, helper, log |
| `net` | cfg, helper, log, mon |
| `tasks` | cfg, helper, log |
| `repos` | cfg, git, helper, log |
| `tpl` | cfg, helper, log |
| `up` | cfg, helper, log |
| `prj` | cfg, git, helper, log, vfs |
| `sim` | cfg, helper, log, mon, net |
| `spec` | cfg, git, helper, idl, log, model, mon, net, repos, sim |
| `gen` | cfg, git, helper, idl, log, model, mon, net, repos, sim, spec |
| `sol` | cfg, gen, git, helper, idl, log, model, mon, net, repos, sim, spec, tasks |
| `mcp` | (almost everything) |
| `cmd` | (everything) |

**Problem**: High coupling - most packages depend on cfg, helper, log, and there are cross-domain dependencies.

---

## Proposed Architecture

### Design Principles

1. **Independent Apps**: Each domain becomes a self-contained app
2. **Interface-Based Communication**: Apps interact through Go interfaces
3. **Duplicate Helpers**: Each app has its own internal utilities
4. **Shared Core**: Only interfaces are shared, not implementations
5. **Dependency Injection**: Apps receive dependencies at construction

### App Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                         apigear (CLI)                           │
│  Entry point that orchestrates all apps via interfaces          │
└─────────────────────────────────────────────────────────────────┘
         │              │              │              │
         ▼              ▼              ▼              ▼
┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│  spec-app   │  │  gen-app    │  │  sim-app    │  │  prj-app    │
│             │  │             │  │             │  │             │
│ - model     │  │ - generator │  │ - engine    │  │ - project   │
│ - idl       │  │ - solution  │  │ - monitor   │  │ - git       │
│ - validate  │  │ - template  │  │ - network   │  │             │
│             │  │ - repos     │  │ - events    │  │             │
└─────────────┘  └─────────────┘  └─────────────┘  └─────────────┘
         │              │              │              │
         └──────────────┴──────────────┴──────────────┘
                                │
                    ┌───────────┴───────────┐
                    │     shared/iface      │
                    │   (interfaces only)   │
                    └───────────────────────┘
```

---

## App Definitions

### 1. `spec-app` - API Specification Domain

**Purpose**: Parse, validate, and represent API specifications

**Current packages**: model, idl, spec (partial)

**Exports Interface**:
```go
package iface

// ISpecLoader loads API specifications from files
type ISpecLoader interface {
    LoadFromIDL(files []string) (ISystem, error)
    LoadFromYAML(files []string) (ISystem, error)
    Validate(system ISystem) error
}

// ISystem represents the root of an API specification
type ISystem interface {
    Name() string
    Modules() []IModule
    LookupModule(name string) IModule
    Checksum() string
}

// IModule represents an API module
type IModule interface {
    Name() string
    Version() string
    Interfaces() []IInterface
    Structs() []IStruct
    Enums() []IEnum
    Externs() []IExtern
}

// IInterface represents an API interface
type IInterface interface {
    Name() string
    Properties() []IProperty
    Operations() []IOperation
    Signals() []ISignal
}

// IStruct, IEnum, IProperty, IOperation, ISignal, etc.
```

**Internal structure**:
```
apps/spec/
├── api.go              # Public interface implementation
├── model/              # System, Module, Interface, etc.
├── idl/                # IDL parser (ANTLR)
├── validate/           # Schema validation
└── internal/
    ├── helper/         # File ops, YAML/JSON parsing
    └── rkw/            # Reserved keywords
```

**Dependencies**: None (leaf app)

---

### 2. `gen-app` - Code Generation Domain

**Purpose**: Generate code from API specifications

**Current packages**: gen, sol, tpl, repos

**Exports Interface**:
```go
package iface

// IGenerator generates code from specifications
type IGenerator interface {
    Generate(opts GenerateOptions) (*GenerateResult, error)
}

type GenerateOptions struct {
    System       ISystem  // From spec-app
    OutputDir    string
    TemplateDir  string
    Features     []string
    Force        bool
    DryRun       bool
}

type GenerateResult struct {
    FilesWritten  int
    FilesSkipped  int
    Duration      time.Duration
}

// ISolutionRunner runs solution-based generation
type ISolutionRunner interface {
    Run(ctx context.Context, solutionPath string, force bool) error
    Watch(ctx context.Context, solutionPath string) error
}

// ITemplateRegistry manages templates
type ITemplateRegistry interface {
    List() ([]TemplateInfo, error)
    Install(repoID string) error
    Update() error
    GetPath(repoID string) (string, error)
}
```

**Internal structure**:
```
apps/gen/
├── api.go              # Public interface implementation
├── generator/          # Template-based generator
├── solution/           # Solution runner
├── template/           # Template creation
├── repos/              # Repository cache
├── filters/            # Language filters (cpp, go, py, etc.)
└── internal/
    ├── helper/         # File ops, path utils
    ├── git/            # Git clone/pull (simplified)
    └── tasks/          # Task execution
```

**Dependencies**: `spec-app` (via ISystem interface)

---

### 3. `sim-app` - Simulation Domain

**Purpose**: Simulate API behavior for testing

**Current packages**: sim, mon, net, evt

**Exports Interface**:
```go
package iface

// ISimulator manages simulation scripts
type ISimulator interface {
    LoadScript(path string) error
    Start(ctx context.Context) error
    Stop() error
}

// IMonitor handles event monitoring
type IMonitor interface {
    OnEvent(fn func(IEvent))
    Emit(event IEvent)
    Start() error
    Stop() error
}

// IEvent represents a monitored event
type IEvent interface {
    ID() string
    Type() string      // "call", "signal", "state"
    Symbol() string
    Timestamp() time.Time
    Data() map[string]any
}

// IServer provides HTTP/WebSocket server
type IServer interface {
    Start(addr string) error
    Stop() error
    Address() string
}
```

**Internal structure**:
```
apps/sim/
├── api.go              # Public interface implementation
├── engine/             # JavaScript simulation engine
├── monitor/            # Event monitoring
├── network/            # HTTP/NATS server
├── events/             # Event bus
├── olink/              # ObjectLink protocol
└── internal/
    └── helper/         # HTTP utils, hooks
```

**Dependencies**: `spec-app` (optional, for type info)

---

### 4. `prj-app` - Project Management Domain

**Purpose**: Manage APIGear projects

**Current packages**: prj, git (partial), vfs

**Exports Interface**:
```go
package iface

// IProjectManager manages projects
type IProjectManager interface {
    Open(path string) (IProject, error)
    Init(path string) error
    Import(gitURL, destPath string) error
    Recent() []IProject
}

// IProject represents an APIGear project
type IProject interface {
    Name() string
    Path() string
    Documents() []IDocument
    AddDocument(docType, name string) error
}

// IDocument represents a project document
type IDocument interface {
    Name() string
    Path() string
    Type() string  // "module", "solution", "scenario"
}
```

**Internal structure**:
```
apps/project/
├── api.go              # Public interface implementation
├── manager/            # Project lifecycle
└── internal/
    ├── helper/         # File ops
    ├── git/            # Git clone (simplified)
    └── vfs/            # Embedded demo files
```

**Dependencies**: None (leaf app)

---

### 5. `shared/iface` - Interface Definitions Only

**Purpose**: Define contracts between apps (NO implementations)

```
shared/
└── iface/
    ├── config.go       # IConfig interface
    ├── logger.go       # ILogger interface
    ├── system.go       # ISystem, IModule, etc. (from spec-app)
    ├── generator.go    # IGenerator, ISolutionRunner
    ├── simulator.go    # ISimulator, IMonitor
    └── project.go      # IProjectManager, IProject
```

**Config Interface**:
```go
type IConfig interface {
    Get(key string) any
    GetString(key string) string
    GetInt(key string) int
    GetBool(key string) bool
    Set(key string, value any)
    ConfigDir() string
}
```

**Logger Interface**:
```go
type ILogger interface {
    Debug() ILogEvent
    Info() ILogEvent
    Warn() ILogEvent
    Error() ILogEvent
}

type ILogEvent interface {
    Str(key, val string) ILogEvent
    Err(err error) ILogEvent
    Msg(msg string)
}
```

---

## Directory Structure

```
apigear-cli/
├── cmd/
│   └── apigear/
│       └── main.go             # CLI entry point
│
├── shared/
│   └── iface/                  # Interface definitions ONLY
│       ├── config.go
│       ├── logger.go
│       ├── system.go
│       ├── generator.go
│       ├── simulator.go
│       └── project.go
│
├── apps/
│   ├── spec/                   # spec-app
│   │   ├── api.go
│   │   ├── model/
│   │   ├── idl/
│   │   ├── validate/
│   │   └── internal/
│   │       ├── helper/
│   │       └── rkw/
│   │
│   ├── gen/                    # gen-app
│   │   ├── api.go
│   │   ├── generator/
│   │   ├── solution/
│   │   ├── template/
│   │   ├── repos/
│   │   ├── filters/
│   │   └── internal/
│   │       ├── helper/
│   │       ├── git/
│   │       └── tasks/
│   │
│   ├── sim/                    # sim-app
│   │   ├── api.go
│   │   ├── engine/
│   │   ├── monitor/
│   │   ├── network/
│   │   ├── events/
│   │   ├── olink/
│   │   └── internal/
│   │       └── helper/
│   │
│   └── project/                # prj-app
│       ├── api.go
│       ├── manager/
│       └── internal/
│           ├── helper/
│           ├── git/
│           └── vfs/
│
├── plugins/                    # Optional extensions
│   ├── mcp/                    # MCP server
│   └── update/                 # Self-update
│
└── internal/
    ├── config/                 # IConfig implementation (Viper)
    └── logger/                 # ILogger implementation (zerolog)
```

---

## Dependency Flow

```
                    ┌──────────────────────────────────────┐
                    │          CLI (cmd/apigear)           │
                    │                                      │
                    │  - Creates IConfig implementation    │
                    │  - Creates ILogger implementation    │
                    │  - Wires apps via interfaces         │
                    └──────────────────────────────────────┘
                                     │
              ┌──────────────────────┼──────────────────────┐
              │                      │                      │
              ▼                      ▼                      ▼
       ┌─────────────┐        ┌─────────────┐        ┌─────────────┐
       │  spec-app   │        │  gen-app    │        │  sim-app    │
       │             │◀───────│             │        │             │
       │  ISystem    │        │ needs:      │        │ needs:      │
       │  IModule    │        │  ISystem    │        │  ISystem    │
       │             │        │             │        │  (optional) │
       └─────────────┘        └─────────────┘        └─────────────┘
              │                      │                      │
              └──────────────────────┼──────────────────────┘
                                     │
                                     ▼
                          ┌───────────────────┐
                          │   shared/iface    │
                          │   (interfaces)    │
                          └───────────────────┘
```

---

## Wiring Example

```go
// cmd/apigear/main.go
package main

import (
    "github.com/apigear-io/cli/internal/config"
    "github.com/apigear-io/cli/internal/logger"
    "github.com/apigear-io/cli/apps/spec"
    "github.com/apigear-io/cli/apps/gen"
    "github.com/apigear-io/cli/apps/sim"
    "github.com/apigear-io/cli/apps/project"
)

func main() {
    // Create shared implementations (injected into apps)
    cfg := config.NewViperConfig()
    log := logger.NewZerologLogger(cfg)

    // Create spec-app (no dependencies)
    specApp := spec.New(spec.Options{
        Config: cfg,
        Logger: log,
    })

    // Create gen-app (depends on spec-app for ISystem)
    genApp := gen.New(gen.Options{
        Config:     cfg,
        Logger:     log,
        SpecLoader: specApp,
    })

    // Create sim-app (optionally uses spec-app)
    simApp := sim.New(sim.Options{
        Config:     cfg,
        Logger:     log,
        SpecLoader: specApp, // optional
    })

    // Create prj-app (no dependencies)
    prjApp := project.New(project.Options{
        Config: cfg,
        Logger: log,
    })

    // Build CLI with wired apps
    cli := NewCLI(CLIOptions{
        Config:   cfg,
        Logger:   log,
        Spec:     specApp,
        Gen:      genApp,
        Sim:      simApp,
        Project:  prjApp,
    })

    os.Exit(cli.Run())
}
```

---

## Helper Duplication Strategy

Each app has its own `internal/helper/` with only what it needs:

### spec-app/internal/helper/
```go
// File operations
func ReadFile(path string) ([]byte, error)
func IsFile(path string) bool
func Join(parts ...string) string

// Document parsing
func ParseYAML(data []byte, v any) error
func ParseJSON(data []byte, v any) error
```

### gen-app/internal/helper/
```go
// File operations (same as spec)
func ReadFile(path string) ([]byte, error)
func WriteFile(path string, data []byte) error
func CopyFile(src, dst string) error
func MakeDir(path string) error

// Path utilities
func Join(parts ...string) string
func BaseName(path string) string
func Dir(path string) string
```

### sim-app/internal/helper/
```go
// Event utilities
type Hook[T any] struct { ... }
func (h *Hook[T]) Add(fn func(*T)) func()
func (h *Hook[T]) Fire(event *T)

// HTTP utilities
func GetFreePort() (int, error)
```

**Trade-off**: ~200-500 lines duplicated per app, but complete independence.

---

## Benefits

| Benefit | Description |
|---------|-------------|
| **Independent Development** | Each app can be developed, tested, and versioned separately |
| **Clear Boundaries** | Interfaces define explicit contracts between domains |
| **Reduced Coupling** | Apps only depend on interfaces, not implementations |
| **Testability** | Easy to mock interfaces for unit testing |
| **Parallel Builds** | Apps can be built in parallel |
| **Plugin Architecture** | New features can be added as plugins |
| **Selective Deployment** | Can build CLI with subset of apps |

---

## Trade-offs

| Trade-off | Mitigation |
|-----------|------------|
| **Code Duplication** | Helper code is small (~500 lines per app), well-defined |
| **Interface Maintenance** | Keep interfaces stable, version them |
| **More Boilerplate** | Use code generation for repetitive patterns |
| **Split Debugging** | Good logging helps trace across app boundaries |

---

## Migration Path

### Phase 1: Define Interfaces (Week 1)
- Create `shared/iface/` with all interface definitions
- Ensure current packages could implement these interfaces
- No code changes to existing packages

### Phase 2: Extract spec-app (Week 2)
- Move model, idl to `apps/spec/`
- Extract relevant parts of spec package
- Create `internal/helper/` with needed utilities
- Implement ISystem, IModule, etc.
- Keep old packages as wrappers (temporarily)

### Phase 3: Extract gen-app (Week 3)
- Move gen, sol, tpl, repos to `apps/gen/`
- Create simplified internal git operations
- Depend on spec-app via ISystem
- Implement IGenerator, ISolutionRunner

### Phase 4: Extract sim-app (Week 4)
- Move sim, mon, net, evt to `apps/sim/`
- Create internal helper with Hook pattern
- Implement ISimulator, IMonitor, IServer

### Phase 5: Extract prj-app (Week 5)
- Move prj to `apps/project/`
- Create internal git and vfs
- Implement IProjectManager, IProject

### Phase 6: Refactor CLI (Week 6)
- Update cmd/apigear to use new app structure
- Wire dependencies via interfaces
- Move mcp, up to plugins
- Remove old pkg/ packages

---

## Summary Table

| App | Contains | Depends On | Exports |
|-----|----------|------------|---------|
| `spec-app` | model, idl, validate | (none) | ISpecLoader, ISystem, IModule |
| `gen-app` | generator, solution, template, repos | spec-app | IGenerator, ISolutionRunner, ITemplateRegistry |
| `sim-app` | engine, monitor, network, events | spec-app (optional) | ISimulator, IMonitor, IServer |
| `prj-app` | manager, git, vfs | (none) | IProjectManager, IProject |
| `shared/iface` | interfaces only | (none) | All interfaces |

---

## Alternative: REST API Architecture

Instead of Go interfaces, expose each app as a REST API module within a single server. Both CLI and Studio (React) become clients of the same backend.

### Architecture Overview

```
┌─────────────────────────────────────────────────────────────────────────┐
│                              Clients                                     │
│  ┌─────────────────────┐              ┌─────────────────────┐           │
│  │   CLI (Go client)   │              │   Studio (React)    │           │
│  │   apigear gen ...   │              │   Web UI            │           │
│  └──────────┬──────────┘              └──────────┬──────────┘           │
│             │                                    │                       │
│             └──────────────┬─────────────────────┘                       │
│                            │ HTTP/REST                                   │
└────────────────────────────┼─────────────────────────────────────────────┘
                             │
┌────────────────────────────┼─────────────────────────────────────────────┐
│                            ▼                                             │
│                   APIGear Server (single process)                        │
│                        localhost:8080                                    │
│                                                                          │
│  ┌─────────────────────────────────────────────────────────────────┐    │
│  │                      Chi Router                                  │    │
│  │  r.Route("/api/spec", specModule.Routes)                        │    │
│  │  r.Route("/api/gen", genModule.Routes)                          │    │
│  │  r.Route("/api/sim", simModule.Routes)                          │    │
│  │  r.Route("/api/project", projectModule.Routes)                  │    │
│  └─────────────────────────────────────────────────────────────────┘    │
│                                                                          │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌─────────────┐        │
│  │ spec module │ │ gen module  │ │ sim module  │ │ prj module  │        │
│  │             │ │             │ │             │ │             │        │
│  │ - model     │ │ - generator │ │ - engine    │ │ - project   │        │
│  │ - idl       │ │ - solution  │ │ - monitor   │ │ - git       │        │
│  │ - validate  │ │ - repos     │ │ - events    │ │             │        │
│  └─────────────┘ └─────────────┘ └─────────────┘ └─────────────┘        │
│                                                                          │
└──────────────────────────────────────────────────────────────────────────┘
```

### Key Design: Single Server, Modular Routes

Each "app" is a module that:
1. Defines its own routes via a `Routes(r chi.Router)` function
2. Contains its business logic internally
3. Registers with the main server at startup

```go
// pkg/api/server.go
func NewServer() *Server {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(cors.Handler(cors.Options{...}))

    // Each module registers its routes
    r.Route("/api/spec", specModule.Routes)
    r.Route("/api/gen", genModule.Routes)
    r.Route("/api/sim", simModule.Routes)
    r.Route("/api/project", projectModule.Routes)

    // Health check
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("ok"))
    })

    return &Server{router: r}
}
```

```go
// pkg/api/spec/routes.go
package spec

func Routes(r chi.Router) {
    s := NewService()

    r.Post("/parse", s.HandleParse)
    r.Post("/validate", s.HandleValidate)
    r.Get("/schema/{type}", s.HandleSchema)
}
```

### Service Definitions

#### 1. Spec Service (`/api/spec`)

Parse and validate API specifications.

```yaml
# OpenAPI-style definition
paths:
  /api/spec/parse:
    post:
      summary: Parse IDL or YAML files
      requestBody:
        content:
          multipart/form-data:
            schema:
              type: object
              properties:
                files:
                  type: array
                  items:
                    type: file
      responses:
        200:
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/System'

  /api/spec/validate:
    post:
      summary: Validate a system
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/System'
      responses:
        200:
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ValidationResult'

  /api/spec/schema/{type}:
    get:
      summary: Get JSON schema for document type
      parameters:
        - name: type
          in: path
          enum: [module, solution, scenario, rules]
      responses:
        200:
          content:
            application/json:
              schema:
                type: object
```

**Go Handler Example:**
```go
// pkg/api/spec/handlers.go
func (s *SpecService) HandleParse(w http.ResponseWriter, r *http.Request) {
    files, err := parseMultipartFiles(r)
    if err != nil {
        writeError(w, http.StatusBadRequest, err)
        return
    }

    system, err := s.loader.LoadFromFiles(files)
    if err != nil {
        writeError(w, http.StatusUnprocessableEntity, err)
        return
    }

    writeJSON(w, http.StatusOK, system)
}
```

#### 2. Gen Service (`/api/gen`)

Generate code from specifications.

```yaml
paths:
  /api/gen/generate:
    post:
      summary: Generate code
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
                system:
                  $ref: '#/components/schemas/System'
                template:
                  type: string
                  example: "apigear-io/template-cpp@latest"
                features:
                  type: array
                  items:
                    type: string
                outputDir:
                  type: string
                force:
                  type: boolean
      responses:
        200:
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/GenerateResult'

  /api/gen/solution:
    post:
      summary: Run solution-based generation
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
                solutionPath:
                  type: string
                watch:
                  type: boolean
                force:
                  type: boolean
      responses:
        200:
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/SolutionResult'

  /api/gen/templates:
    get:
      summary: List available templates
      responses:
        200:
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/TemplateInfo'

  /api/gen/templates/{id}:
    post:
      summary: Install template
      parameters:
        - name: id
          in: path
          example: "apigear-io/template-cpp@v1.0.0"
```

#### 3. Sim Service (`/api/sim`)

Simulation and monitoring.

```yaml
paths:
  /api/sim/start:
    post:
      summary: Start simulation
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
                scriptPath:
                  type: string
      responses:
        200:
          content:
            application/json:
              schema:
                type: object
                properties:
                  sessionId:
                    type: string

  /api/sim/stop:
    post:
      summary: Stop simulation
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
                sessionId:
                  type: string

  /api/sim/events:
    get:
      summary: Stream events (SSE)
      responses:
        200:
          content:
            text/event-stream:
              schema:
                $ref: '#/components/schemas/Event'

  /api/sim/events:
    post:
      summary: Emit event
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Event'
```

#### 4. Project Service (`/api/project`)

Project management.

```yaml
paths:
  /api/project:
    get:
      summary: List recent projects
      responses:
        200:
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Project'

  /api/project:
    post:
      summary: Create or open project
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
                path:
                  type: string
                action:
                  type: string
                  enum: [create, open, import]
                gitUrl:
                  type: string

  /api/project/{id}/documents:
    get:
      summary: List project documents
    post:
      summary: Add document to project
```

### Directory Structure

```
apigear-cli/
├── cmd/
│   ├── apigear/              # CLI (can run standalone or connect to server)
│   │   └── main.go
│   └── apigear-server/       # Standalone API server (optional)
│       └── main.go
│
├── pkg/
│   ├── api/                  # REST API layer (thin wrappers)
│   │   ├── server.go         # Server setup, route registration
│   │   ├── middleware.go     # Auth, CORS, logging
│   │   ├── response.go       # JSON response helpers
│   │   │
│   │   ├── spec/             # /api/spec module
│   │   │   ├── routes.go     # Route registration
│   │   │   ├── handlers.go   # HTTP handlers
│   │   │   └── types.go      # Request/response types
│   │   │
│   │   ├── gen/              # /api/gen module
│   │   │   ├── routes.go
│   │   │   ├── handlers.go
│   │   │   └── types.go
│   │   │
│   │   ├── sim/              # /api/sim module
│   │   │   ├── routes.go
│   │   │   ├── handlers.go
│   │   │   └── types.go
│   │   │
│   │   └── project/          # /api/project module
│   │       ├── routes.go
│   │       ├── handlers.go
│   │       └── types.go
│   │
│   ├── client/               # Go HTTP client (for CLI remote mode)
│   │   ├── client.go         # Base client with auth, retries
│   │   ├── spec.go           # Spec API methods
│   │   ├── gen.go            # Gen API methods
│   │   ├── sim.go            # Sim API methods
│   │   └── project.go        # Project API methods
│   │
│   │   # Existing packages (business logic - unchanged)
│   ├── model/
│   ├── idl/
│   ├── gen/
│   ├── sim/
│   ├── spec/
│   ├── prj/
│   ├── repos/
│   └── ...
│
└── studio/                   # React frontend (separate repo or subdir)
    └── src/
        ├── api/              # Auto-generated TypeScript client
        │   └── index.ts      # Generated from OpenAPI spec
        └── ...
```

### Module Structure Pattern

Each API module follows the same pattern:

```
pkg/api/spec/
├── routes.go      # func Routes(r chi.Router) - registers all routes
├── handlers.go    # HTTP handlers that call business logic
├── types.go       # Request/Response DTOs (separate from domain models)
└── service.go     # Optional: module-specific service layer
```

```go
// pkg/api/spec/types.go
package spec

// Request/Response types - decoupled from internal models
type ParseRequest struct {
    Files []string `json:"files"`
}

type ParseResponse struct {
    System  *SystemDTO `json:"system"`
    Errors  []string   `json:"errors,omitempty"`
}

type SystemDTO struct {
    Name     string      `json:"name"`
    Modules  []ModuleDTO `json:"modules"`
    Checksum string      `json:"checksum"`
}

// Convert from internal model
func SystemToDTO(s *model.System) *SystemDTO {
    return &SystemDTO{
        Name:     s.Name,
        Modules:  modulesToDTO(s.Modules),
        Checksum: s.Checksum(),
    }
}
```

```go
// pkg/api/spec/handlers.go
package spec

import (
    "github.com/apigear-io/cli/pkg/model"
    "github.com/apigear-io/cli/pkg/idl"
)

type Service struct {
    // Can inject dependencies here
}

func NewService() *Service {
    return &Service{}
}

func (s *Service) HandleParse(w http.ResponseWriter, r *http.Request) {
    var req ParseRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        writeError(w, http.StatusBadRequest, err)
        return
    }

    // Call existing business logic
    system := model.NewSystem("api")
    parser := idl.NewParser(system)

    for _, file := range req.Files {
        if err := parser.ParseFile(file); err != nil {
            writeError(w, http.StatusUnprocessableEntity, err)
            return
        }
    }

    if err := system.Validate(); err != nil {
        writeError(w, http.StatusUnprocessableEntity, err)
        return
    }

    // Convert to DTO and return
    writeJSON(w, http.StatusOK, ParseResponse{
        System: SystemToDTO(system),
    })
}
```

### CLI as HTTP Client

```go
// cmd/apigear/main.go
func main() {
    // CLI connects to local or remote server
    serverURL := os.Getenv("APIGEAR_SERVER")
    if serverURL == "" {
        serverURL = "http://localhost:8080"
    }

    client := client.New(serverURL)

    // Commands use HTTP client
    app := &cli.App{
        Commands: []*cli.Command{
            {
                Name: "gen",
                Subcommands: []*cli.Command{
                    {
                        Name: "solution",
                        Action: func(c *cli.Context) error {
                            return client.Gen.RunSolution(c.Context, c.String("file"))
                        },
                    },
                },
            },
        },
    }
}
```

```go
// pkg/client/gen.go
type GenClient struct {
    baseURL string
    http    *http.Client
}

func (c *GenClient) RunSolution(ctx context.Context, path string) error {
    req := GenerateSolutionRequest{
        SolutionPath: path,
        Force:        false,
    }

    resp, err := c.post(ctx, "/api/gen/solution", req)
    if err != nil {
        return err
    }

    var result SolutionResult
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return err
    }

    fmt.Printf("Generated %d files\n", result.FilesWritten)
    return nil
}
```

### React Studio Client

```typescript
// studio/src/api/client.ts
const API_BASE = process.env.REACT_APP_API_URL || 'http://localhost:8080';

export const specApi = {
  parse: async (files: File[]): Promise<System> => {
    const formData = new FormData();
    files.forEach(f => formData.append('files', f));

    const resp = await fetch(`${API_BASE}/api/spec/parse`, {
      method: 'POST',
      body: formData,
    });
    return resp.json();
  },

  validate: async (system: System): Promise<ValidationResult> => {
    const resp = await fetch(`${API_BASE}/api/spec/validate`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(system),
    });
    return resp.json();
  },
};

export const genApi = {
  templates: async (): Promise<TemplateInfo[]> => {
    const resp = await fetch(`${API_BASE}/api/gen/templates`);
    return resp.json();
  },

  generate: async (opts: GenerateOptions): Promise<GenerateResult> => {
    const resp = await fetch(`${API_BASE}/api/gen/generate`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(opts),
    });
    return resp.json();
  },
};
```

### Deployment Modes

#### Mode 1: Local Development (Embedded Server)

CLI starts server automatically:

```go
// CLI starts embedded server if not running
func ensureServer() (*client.Client, error) {
    c := client.New("http://localhost:8080")

    if err := c.Health(); err != nil {
        // Start embedded server
        go server.Start(":8080")
        time.Sleep(100 * time.Millisecond)
    }

    return c, nil
}
```

#### Mode 2: Standalone Server

Server runs separately (Docker, systemd):

```bash
# Start server
apigear-server --port 8080

# CLI connects to it
export APIGEAR_SERVER=http://localhost:8080
apigear gen solution my.solution.yaml
```

#### Mode 3: Remote/Cloud

Server runs in cloud, multiple clients connect:

```bash
# CLI connects to remote
export APIGEAR_SERVER=https://api.apigear.io
apigear gen solution my.solution.yaml

# Studio also connects to same server
# (configured in environment)
```

### Comparison: Go Interfaces vs REST API

| Aspect | Go Interfaces | REST API (Single Server) |
|--------|---------------|--------------------------|
| **Latency** | Nanoseconds (in-process) | Milliseconds (HTTP) |
| **Complexity** | Lower | Medium (HTTP, DTOs) |
| **CLI standalone** | Yes (single binary) | Yes (embedded server) |
| **Studio sharing** | No (separate Go/React) | Yes (same API) |
| **Testing** | Unit tests | Unit + API tests |
| **Deployment** | Single binary | Single binary (server included) |
| **Language agnostic** | No (Go only) | Yes (any HTTP client) |
| **Offline mode** | Always works | Works (embedded server) |
| **Multi-user** | No | Yes (shared server mode) |
| **Real-time updates** | Via channels | Via SSE/WebSocket |
| **OpenAPI docs** | Manual | Auto-generated |
| **Existing code changes** | Significant | Minimal (add API layer) |

### Effort Estimate for REST API Approach

| Phase | Work | Estimate |
|-------|------|----------|
| **1. Create API scaffolding** | server.go, middleware, response helpers | 2-3 days |
| **2. Define OpenAPI spec** | Document all endpoints | 3-5 days |
| **3. Implement spec module** | /api/spec handlers | 3-5 days |
| **4. Implement gen module** | /api/gen handlers | 1 week |
| **5. Implement sim module** | /api/sim handlers + SSE | 1 week |
| **6. Implement project module** | /api/project handlers | 2-3 days |
| **7. Create Go client** | HTTP client for CLI | 3-5 days |
| **8. Generate TypeScript client** | From OpenAPI spec | 1-2 days |
| **9. Embedded server mode** | CLI auto-starts server | 2-3 days |
| **10. Testing** | API integration tests | 1 week |

**Total: 5-7 weeks**

### Incremental Migration Path for REST API

The REST API approach can be done incrementally without breaking existing CLI:

**Week 1-2: Foundation**
```
1. Create pkg/api/server.go with basic Chi setup
2. Add /health endpoint
3. Create pkg/api/middleware.go (logging, CORS)
4. Create pkg/api/response.go (JSON helpers)
5. Wire into existing `apigear serve` command
```

**Week 3: First Module (spec)**
```
1. Create pkg/api/spec/routes.go
2. Create pkg/api/spec/types.go (DTOs)
3. Implement POST /api/spec/parse
4. Implement POST /api/spec/validate
5. Test with curl/Postman
```

**Week 4: Gen Module**
```
1. Create pkg/api/gen/routes.go
2. Implement GET /api/gen/templates
3. Implement POST /api/gen/generate
4. Implement POST /api/gen/solution
```

**Week 5: Sim Module**
```
1. Create pkg/api/sim/routes.go
2. Implement POST /api/sim/start, /stop
3. Implement GET /api/sim/events (SSE)
4. Implement POST /api/sim/events
```

**Week 6: Project Module + Client**
```
1. Create pkg/api/project/routes.go
2. Implement CRUD endpoints
3. Create pkg/client/ for Go HTTP client
4. Add --server flag to CLI commands
```

**Week 7: Polish**
```
1. Generate OpenAPI spec from code (swag)
2. Generate TypeScript client (openapi-generator)
3. Add authentication middleware (optional)
4. Write API tests
```

### CLI Server Lifecycle Management

The CLI automatically manages the server:

1. **Check** if server is running on standard port (e.g., `:8080`)
2. **Start** embedded server if not found
3. **Execute** command via HTTP API
4. **Stop** embedded server when CLI exits

```go
// pkg/client/lifecycle.go
package client

import (
    "context"
    "net/http"
    "time"

    "github.com/apigear-io/cli/pkg/api"
)

const (
    DefaultPort    = "8080"
    DefaultAddress = "http://localhost:" + DefaultPort
    HealthEndpoint = "/health"
    StartupTimeout = 2 * time.Second
)

type ManagedClient struct {
    *Client
    server   *api.Server
    embedded bool
}

// GetOrCreateClient returns a client, starting embedded server if needed
func GetOrCreateClient(ctx context.Context) (*ManagedClient, error) {
    client := New(DefaultAddress)

    // Check if server is already running
    if err := client.Health(ctx); err == nil {
        // Server already running (maybe Studio started it)
        return &ManagedClient{Client: client, embedded: false}, nil
    }

    // Start embedded server
    server := api.NewServer()
    go func() {
        if err := server.Start(":" + DefaultPort); err != nil {
            log.Error().Err(err).Msg("embedded server failed")
        }
    }()

    // Wait for server to be ready
    deadline := time.Now().Add(StartupTimeout)
    for time.Now().Before(deadline) {
        if err := client.Health(ctx); err == nil {
            return &ManagedClient{
                Client:   client,
                server:   server,
                embedded: true,
            }, nil
        }
        time.Sleep(50 * time.Millisecond)
    }

    return nil, fmt.Errorf("timeout waiting for embedded server")
}

// Close shuts down the embedded server if we started it
func (c *ManagedClient) Close() error {
    if c.embedded && c.server != nil {
        return c.server.Stop()
    }
    return nil
}
```

```go
// pkg/cmd/gen/solution.go
func runSolution(cmd *cobra.Command, args []string) error {
    ctx := cmd.Context()

    // Get or create client (auto-starts server if needed)
    client, err := client.GetOrCreateClient(ctx)
    if err != nil {
        return fmt.Errorf("failed to connect to server: %w", err)
    }
    defer client.Close()  // Auto-stops embedded server

    // Execute via API
    result, err := client.Gen.RunSolution(ctx, args[0])
    if err != nil {
        return err
    }

    fmt.Printf("Generated %d files in %s\n", result.FilesWritten, result.Duration)
    return nil
}
```

### Server Discovery Flow

```
┌─────────────────────────────────────────────────────────────────┐
│                    CLI Command Execution                         │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
                 ┌────────────────────────┐
                 │  Check localhost:8080  │
                 │    GET /health         │
                 └────────────────────────┘
                              │
              ┌───────────────┴───────────────┐
              │                               │
              ▼                               ▼
     ┌─────────────────┐             ┌─────────────────┐
     │ Server Running  │             │ Server Not Found│
     │ (Studio or other)│            │                 │
     └────────┬────────┘             └────────┬────────┘
              │                               │
              │                               ▼
              │                  ┌────────────────────────┐
              │                  │ Start Embedded Server  │
              │                  │ (in background)        │
              │                  └────────────┬───────────┘
              │                               │
              └───────────────┬───────────────┘
                              │
                              ▼
                 ┌────────────────────────┐
                 │  Execute API Request   │
                 │  POST /api/gen/...     │
                 └────────────────────────┘
                              │
                              ▼
                 ┌────────────────────────┐
                 │  Command Complete      │
                 └────────────────────────┘
                              │
              ┌───────────────┴───────────────┐
              │                               │
              ▼                               ▼
     ┌─────────────────┐             ┌─────────────────┐
     │ External Server │             │ Embedded Server │
     │ (leave running) │             │ (shut down)     │
     └─────────────────┘             └─────────────────┘
```

### Usage Scenarios

**Scenario 1: CLI only (typical developer)**
```bash
$ apigear gen sol my.solution.yaml
# Server auto-starts on :8080
# Generates code
# Server auto-stops

$ apigear gen sol another.solution.yaml
# Server auto-starts again
# Generates code
# Server auto-stops
```

**Scenario 2: Studio running (GUI user)**
```bash
# Studio is running, server already on :8080

$ apigear gen sol my.solution.yaml
# Detects existing server
# Uses it (no embedded server started)
# Server keeps running (Studio manages it)
```

**Scenario 3: Long-running server (power user)**
```bash
# Terminal 1: Start server explicitly
$ apigear serve
Server running on :8080

# Terminal 2: CLI commands use existing server
$ apigear gen sol my.solution.yaml
# Uses existing server
# Server keeps running
```

**Scenario 4: Watch mode (keeps server alive)**
```bash
$ apigear gen sol --watch my.solution.yaml
# Server starts
# Watches for changes
# Re-generates on change
# Server stays alive until Ctrl+C
# Server stops on exit
```

### Configuration

```yaml
# ~/.apigear/config.yaml
server:
  port: 8080                    # Default port
  auto_start: true              # Auto-start if not running
  auto_stop: true               # Auto-stop embedded server on exit
  startup_timeout: 2s           # Wait time for server startup
  external_url: ""              # Override: use remote server instead
```

```go
// Environment variables also work
// APIGEAR_SERVER_PORT=8080
// APIGEAR_SERVER_URL=https://api.apigear.io  (use remote)
```

### Edge Cases

| Scenario | Behavior |
|----------|----------|
| Port in use (not apigear) | Error: "port 8080 in use by another process" |
| Server crashes mid-request | Retry once, then error |
| Multiple CLI instances | All share same server (first starts, last may stop) |
| Ctrl+C during command | Graceful shutdown, server stops if embedded |
| `--no-server` flag | Direct mode (bypass API, like current behavior) |

### Reference Counting (Optional Enhancement)

For multiple concurrent CLI processes:

```go
// Track how many CLI processes are using the embedded server
type ServerManager struct {
    refCount int32
    server   *api.Server
    mu       sync.Mutex
}

func (m *ServerManager) Acquire() (*Client, error) {
    m.mu.Lock()
    defer m.mu.Unlock()

    if m.refCount == 0 {
        // Start server
        m.server = api.NewServer()
        go m.server.Start(":8080")
    }
    atomic.AddInt32(&m.refCount, 1)
    return NewClient(DefaultAddress), nil
}

func (m *ServerManager) Release() {
    if atomic.AddInt32(&m.refCount, -1) == 0 {
        // Last user, stop server
        m.server.Stop()
    }
}
```

This could use a lock file or Unix socket for cross-process coordination.

---

### Multi-User / Shared Server Scenarios

The REST API architecture naturally enables multiple users to share the same server:

```
┌─────────────────────────────────────────────────────────────────────────┐
│                         Shared APIGear Server                           │
│                     (Team Server / Cloud Instance)                      │
│                                                                         │
│  ┌─────────────────────────────────────────────────────────────────┐   │
│  │                      localhost:8080 or                           │   │
│  │                   https://apigear.company.com                    │   │
│  └─────────────────────────────────────────────────────────────────┘   │
│                                    │                                    │
└────────────────────────────────────┼────────────────────────────────────┘
                                     │
        ┌────────────────────────────┼────────────────────────────────┐
        │                            │                                │
        ▼                            ▼                                ▼
┌───────────────┐          ┌───────────────┐              ┌───────────────┐
│  Developer A  │          │  Developer B  │              │  CI/CD        │
│               │          │               │              │               │
│  CLI + Studio │          │  CLI only     │              │  CLI          │
│  (macOS)      │          │  (Linux)      │              │  (Docker)     │
└───────────────┘          └───────────────┘              └───────────────┘
```

#### Deployment Scenarios

**1. Local Development (Single User)**
```bash
# Default: each developer runs their own embedded server
$ apigear gen sol my.solution.yaml
# Server auto-starts, runs locally, auto-stops
```

**2. Team Development Server**
```bash
# Ops: Deploy shared server
$ docker run -p 8080:8080 apigear/server

# Developers: Point to shared server
$ export APIGEAR_SERVER=http://dev-server.local:8080
$ apigear gen sol my.solution.yaml

# Or in config file
$ cat ~/.apigear/config.yaml
server:
  url: http://dev-server.local:8080
```

**3. CI/CD Pipeline**
```yaml
# .github/workflows/generate.yml
jobs:
  generate:
    runs-on: ubuntu-latest
    services:
      apigear:
        image: apigear/server
        ports:
          - 8080:8080
    steps:
      - uses: actions/checkout@v4
      - name: Generate SDK
        run: |
          export APIGEAR_SERVER=http://localhost:8080
          apigear gen sol solution.yaml
```

**4. Cloud/SaaS Deployment**
```bash
# Central company server
$ export APIGEAR_SERVER=https://apigear.company.com

# All teams use same server
$ apigear gen sol my.solution.yaml
# Templates cached centrally
# Consistent versions across teams
```

#### Benefits of Shared Server

| Benefit | Description |
|---------|-------------|
| **Template caching** | Download once, use everywhere |
| **Consistent versions** | All users get same template versions |
| **Centralized config** | Company-wide settings in one place |
| **Audit logging** | Track who generated what, when |
| **Resource sharing** | One server vs. many embedded instances |
| **Studio + CLI parity** | Same backend for both interfaces |

#### Multi-User Features

**Workspaces / Projects**
```
/api/workspaces
├── GET    /                    # List user's workspaces
├── POST   /                    # Create workspace
├── GET    /{id}                # Get workspace
├── DELETE /{id}                # Delete workspace
└── GET    /{id}/projects       # List projects in workspace
```

**User Context**
```go
// Middleware adds user context from auth token
func UserContextMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        user, err := validateToken(token)
        if err != nil {
            writeError(w, http.StatusUnauthorized, err)
            return
        }
        ctx := context.WithValue(r.Context(), "user", user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Handlers can access user
func (s *Service) HandleGenerate(w http.ResponseWriter, r *http.Request) {
    user := r.Context().Value("user").(*User)
    log.Info().Str("user", user.ID).Msg("generating code")
    // ...
}
```

**Shared Template Registry**
```go
// Server maintains central template cache
type TemplateRegistry struct {
    cache   map[string]*Template  // Shared across all users
    mu      sync.RWMutex
}

// Install once, available to all
func (r *TemplateRegistry) Install(repoID string) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    if _, exists := r.cache[repoID]; exists {
        return nil  // Already installed
    }

    // Download and cache
    tpl, err := downloadTemplate(repoID)
    if err != nil {
        return err
    }
    r.cache[repoID] = tpl
    return nil
}
```

#### Authentication Options

| Mode | Use Case | Implementation |
|------|----------|----------------|
| **None** | Local dev, trusted network | No auth middleware |
| **API Key** | CI/CD, scripts | `X-API-Key` header |
| **JWT** | Multi-user, Studio | `Authorization: Bearer <token>` |
| **OAuth2** | Enterprise SSO | OIDC with company IdP |

```go
// pkg/api/middleware/auth.go
func AuthMiddleware(mode string) func(http.Handler) http.Handler {
    switch mode {
    case "none":
        return func(next http.Handler) http.Handler { return next }
    case "apikey":
        return APIKeyAuth(os.Getenv("APIGEAR_API_KEYS"))
    case "jwt":
        return JWTAuth(os.Getenv("APIGEAR_JWT_SECRET"))
    case "oauth2":
        return OAuth2Auth(oauth2Config)
    default:
        return func(next http.Handler) http.Handler { return next }
    }
}
```

#### Rate Limiting & Quotas

For shared servers, prevent abuse:

```go
// Per-user rate limiting
func RateLimitMiddleware(rps int) func(http.Handler) http.Handler {
    limiters := make(map[string]*rate.Limiter)
    var mu sync.Mutex

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            user := getUserID(r)

            mu.Lock()
            limiter, exists := limiters[user]
            if !exists {
                limiter = rate.NewLimiter(rate.Limit(rps), rps*2)
                limiters[user] = limiter
            }
            mu.Unlock()

            if !limiter.Allow() {
                writeError(w, http.StatusTooManyRequests, "rate limit exceeded")
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
```

#### Server Deployment Options

**Docker Compose (Team Server)**
```yaml
# docker-compose.yml
version: '3.8'
services:
  apigear:
    image: apigear/server:latest
    ports:
      - "8080:8080"
    volumes:
      - apigear-templates:/app/templates
      - apigear-data:/app/data
    environment:
      - APIGEAR_AUTH_MODE=apikey
      - APIGEAR_API_KEYS=key1,key2,key3
    restart: unless-stopped

volumes:
  apigear-templates:
  apigear-data:
```

**Kubernetes (Enterprise)**
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: apigear-server
spec:
  replicas: 3
  selector:
    matchLabels:
      app: apigear
  template:
    spec:
      containers:
        - name: apigear
          image: apigear/server:latest
          ports:
            - containerPort: 8080
          env:
            - name: APIGEAR_AUTH_MODE
              value: oauth2
          volumeMounts:
            - name: templates
              mountPath: /app/templates
      volumes:
        - name: templates
          persistentVolumeClaim:
            claimName: apigear-templates
---
apiVersion: v1
kind: Service
metadata:
  name: apigear
spec:
  selector:
    app: apigear
  ports:
    - port: 80
      targetPort: 8080
  type: LoadBalancer
```

#### Summary: Deployment Modes

| Mode | Server | Users | Auth | Use Case |
|------|--------|-------|------|----------|
| **Embedded** | Auto-start/stop | 1 | None | Local dev |
| **Standalone** | `apigear serve` | 1+ | Optional | Power user |
| **Docker** | Container | Team | API Key | Team dev |
| **Kubernetes** | Cluster | Many | OAuth2 | Enterprise |
| **Cloud** | Managed | Many | OAuth2 | SaaS |

### Hybrid Approach (Recommended)

Combine both approaches for flexibility:

```
┌─────────────────────────────────────────────────────────────┐
│                         CLI                                  │
│  ┌─────────────────┐    ┌─────────────────┐                 │
│  │ Direct Mode     │ OR │ Client Mode     │                 │
│  │ (Go interfaces) │    │ (HTTP client)   │                 │
│  └────────┬────────┘    └────────┬────────┘                 │
│           │                      │                           │
│           ▼                      ▼                           │
│  ┌─────────────────────────────────────────┐                │
│  │           Core Business Logic            │                │
│  │    (model, idl, gen, sim, etc.)         │                │
│  └─────────────────────────────────────────┘                │
│                      │                                       │
│                      ▼                                       │
│  ┌─────────────────────────────────────────┐                │
│  │           REST API Layer                 │                │
│  │    (thin wrapper over core logic)       │                │
│  └─────────────────────────────────────────┘                │
│                      │                                       │
└──────────────────────┼───────────────────────────────────────┘
                       │
                       ▼
              ┌─────────────────┐
              │  Studio (React) │
              │  External Tools │
              └─────────────────┘
```

**Benefits of Hybrid:**
- CLI works offline (direct mode)
- CLI can connect to server (client mode)
- Studio uses same API
- Core logic is shared
- Incremental migration possible

---

## Effort and Complexity Analysis

### Codebase Metrics

| Metric | Value |
|--------|-------|
| **Total source files** | 318 |
| **Total lines of code** | ~24,000 |
| **Test files** | 114 |

### Size by Proposed App

| App | Current Packages | Lines | Complexity |
|-----|------------------|-------|------------|
| **spec-app** | model, idl, spec (partial) | ~9,200 | High (ANTLR parser) |
| **gen-app** | gen, sol, tpl, repos | ~6,900 | High (templates, 11 language filters) |
| **sim-app** | sim, mon, net, evt | ~2,800 | Medium (JS runtime, ObjectLink) |
| **prj-app** | prj, git, vfs | ~750 | Low |
| **cli** | cmd, mcp | ~2,600 | Medium |
| **shared** | helper, cfg, log, tasks | ~1,700 | Low (to duplicate) |

### Lines of Code per Package

```
cfg             335 lines
cmd           2,278 lines
evt             234 lines
gen           6,043 lines  (includes filters)
git             373 lines
helper          869 lines
idl           6,116 lines  (includes ANTLR parser)
log             136 lines
mcp             365 lines
model         1,776 lines
mon             326 lines
net             595 lines
prj             358 lines
repos           508 lines
sim           1,674 lines
sol             280 lines
spec          1,323 lines
tasks           373 lines
tools           143 lines
tpl             115 lines
up               85 lines
vfs              18 lines
```

---

## Effort Estimate

### Full Refactoring Timeline

| Phase | Work | Estimate | Risk |
|-------|------|----------|------|
| **1. Define interfaces** | Create `shared/iface/` | 2-3 days | Low |
| **2. Extract spec-app** | model + idl (9k lines, ANTLR) | 1-2 weeks | High |
| **3. Extract gen-app** | gen + filters + repos (7k lines) | 1-2 weeks | High |
| **4. Extract sim-app** | sim + mon + net (3k lines) | 1 week | Medium |
| **5. Extract prj-app** | prj + git (750 lines) | 2-3 days | Low |
| **6. Rewire CLI** | cmd + mcp + wiring | 1 week | Medium |
| **7. Testing & fixes** | Integration, edge cases | 1-2 weeks | High |

**Total: 6-10 weeks** for one experienced developer

### High-Risk Areas

1. **IDL parser (6k lines)**
   - ANTLR-generated code tightly coupled to model
   - Complex listener pattern with state management

2. **Generator filters (3k lines across 11 languages)**
   - Shared patterns between filters
   - Template function registration

3. **Simulation engine**
   - JavaScript runtime (Goja) integration
   - ObjectLink protocol implementation

4. **Circular interface design**
   - Getting the interfaces right requires iteration
   - Changes ripple across all apps

### Hidden Work

| Hidden Cost | Impact |
|-------------|--------|
| **Test rewrites** | 114 test files need updating |
| **Integration tests** | Cross-app workflows need new tests |
| **Build system** | Taskfile, goreleaser updates |
| **Documentation** | README, examples need updating |
| **Edge cases** | Things that work by accident today |
| **CI/CD pipeline** | May need restructuring |

---

## Alternative Approaches

### Option A: Incremental Refactoring (Lower Risk)

Instead of big-bang, evolve gradually:

| Step | Effort | Outcome |
|------|--------|---------|
| 1. Add interfaces alongside existing code | 1-2 weeks | Contracts defined |
| 2. Make packages implement interfaces | 2-3 weeks | Testable boundaries |
| 3. Gradually add dependency injection | Ongoing | Reduced coupling |
| 4. Extract apps one at a time | Months | Full separation |

**Total: 4-6 weeks** for initial improvement, then ongoing

### Option B: Boundaries Only (Minimal Effort)

Keep current structure, improve boundaries:

| Step | Effort | Outcome |
|------|--------|---------|
| 1. Add `api.go` to each package | 3-5 days | Clean public interface |
| 2. Move internals to `internal/` | 1 week | Hidden implementation |
| 3. Reduce exports | 3-5 days | Smaller surface area |
| 4. Document interfaces | 2-3 days | Clear contracts |

**Total: 2-3 weeks** for meaningful improvement

---

## Recommendation

### Pragmatic Path (Recommended)

| Step | Effort | Value |
|------|--------|-------|
| 1. Add interface files to existing packages | 1 week | Define contracts |
| 2. Create `internal/` in each package | 1 week | Hide implementation |
| 3. Extract `helper` duplicates where needed | 1 week | Reduce coupling |
| 4. Extract one app (prj-app is easiest) | 1 week | Prove the pattern |
| 5. Evaluate if full migration is worth it | - | Informed decision |

**Total: 4 weeks** to validate the approach

This gives **80% of the benefits** (clear boundaries, documented interfaces, reduced coupling) with **20% of the effort** and risk.

### Decision Framework

**Choose Full Refactoring if:**
- Multiple developers will work on different domains
- You need to version/release apps independently
- The codebase will grow significantly
- You're willing to invest 2-3 months

**Choose Incremental/Boundaries if:**
- Single developer or small team
- Current structure works reasonably well
- Need to ship features in parallel
- Want lower risk and faster payoff

---

## Risk Mitigation

### Before Starting

1. **Increase test coverage** - Ensure critical paths are tested
2. **Document current behavior** - Capture implicit contracts
3. **Set up feature flags** - Enable gradual rollout
4. **Create rollback plan** - Keep old code path available

### During Migration

1. **One app at a time** - Complete each before starting next
2. **Maintain compatibility** - Old and new code coexist
3. **Continuous integration** - Run full test suite on each change
4. **Regular checkpoints** - Deployable state at each phase end

### Success Metrics

| Metric | Target |
|--------|--------|
| Test pass rate | 100% after each phase |
| Build time | No significant increase |
| Binary size | < 10% increase |
| No regressions | Zero user-facing bugs |

---

## Phase 0: Increase Test Coverage

Before any refactoring, establish a safety net with comprehensive tests.

### Current Test Coverage

| Package | Coverage | Test Files | Priority |
|---------|----------|------------|----------|
| `idl` | 93.2% | 10 | Low (good) |
| `filterqt` | 85.7% | yes | Low (good) |
| `filterpy` | 84.1% | yes | Low (good) |
| `filtercpp` | 82.4% | yes | Low (good) |
| `filterrs` | 80.9% | yes | Low (good) |
| `filterjni` | 80.1% | yes | Low (good) |
| `filtergo` | 77.3% | yes | Low (good) |
| `filterjs` | 77.0% | yes | Low (good) |
| `filterts` | 77.0% | yes | Low (good) |
| `filterue` | 74.4% | yes | Low (good) |
| `evt` | 69.9% | 1 | Low (good) |
| `filterjava` | 61.7% | yes | Medium |
| `gen` | 59.1% | 2 | Medium |
| `common` | 47.8% | yes | Medium |
| `spec/rkw` | 43.9% | yes | Medium |
| `spec` | 42.9% | 4 | **High** |
| `mon` | 40.9% | 3 | Medium |
| `sim` | 38.1% | 6 | **High** |
| `model` | 34.9% | 6 | **High** |
| `cmd/cfg` | 28.6% | yes | Medium |
| `repos` | 12.3% | 1 | **High** |
| `cfg` | 0% | **none** | **Critical** |
| `cmd` | 0% | **none** | Medium |
| `git` | 0% | **none** | **High** |
| `helper` | 0% | **none** | **Critical** |
| `log` | 0% | **none** | Medium |
| `mcp` | 0% | **none** | Low |
| `net` | 0% | **none** | **High** |
| `prj` | 0% | **none** | **High** |
| `sol` | 0% | **none** | **High** |
| `tasks` | 0% | **none** | Medium |
| `tpl` | 0% | **none** | Low |
| `up` | 0% | **none** | Low |
| `vfs` | 0% | **none** | Low |

### Test Coverage Goals

| Phase | Target | Focus |
|-------|--------|-------|
| **Immediate** | 50%+ on critical packages | helper, cfg, model, git |
| **Before refactoring** | 70%+ on packages to extract | model, spec, gen, sim |
| **After refactoring** | 80%+ on new apps | Validate new structure |

### Priority 1: Critical Packages (No Tests)

These packages are used everywhere and have zero tests:

#### `helper` - Foundation utilities
```go
// pkg/helper/helper_test.go
func TestIsDir(t *testing.T) {
    // Test with existing directory
    // Test with file (should return false)
    // Test with non-existent path
}

func TestIsFile(t *testing.T) { ... }
func TestJoin(t *testing.T) { ... }
func TestReadDocument(t *testing.T) { ... }
func TestWriteDocument(t *testing.T) { ... }
func TestCopyFile(t *testing.T) { ... }
func TestParseYAML(t *testing.T) { ... }
func TestParseJSON(t *testing.T) { ... }
```

#### `cfg` - Configuration
```go
// pkg/cfg/cfg_test.go
func TestGetSetString(t *testing.T) { ... }
func TestGetSetBool(t *testing.T) { ... }
func TestConfigDir(t *testing.T) { ... }
func TestRecentEntries(t *testing.T) { ... }
```

#### `git` - Git operations
```go
// pkg/git/git_test.go
func TestIsValidGitUrl(t *testing.T) {
    tests := []struct{
        url string
        valid bool
    }{
        {"https://github.com/org/repo.git", true},
        {"git@github.com:org/repo.git", true},
        {"not-a-url", false},
    }
    // ...
}

func TestParseAsUrl(t *testing.T) { ... }
func TestClone(t *testing.T) { ... }  // May need mocking
```

### Priority 2: Low Coverage Packages

These have tests but need more:

#### `model` (34.9%) - Core data structures
```go
// Focus areas:
// - System.Validate()
// - Module.LookupInterface()
// - Schema type resolution
// - Visitor pattern traversal
```

#### `repos` (12.3%) - Template repository
```go
// Focus areas:
// - Registry.List()
// - Cache.Install()
// - RepoID parsing (EnsureRepoID, SplitRepoID)
```

#### `spec` (42.9%) - Specification validation
```go
// Focus areas:
// - CheckFile() with various file types
// - Schema validation
// - Feature computation
```

### Priority 3: Packages to Extract

Before extracting to apps, ensure high coverage:

| Future App | Packages | Target Coverage |
|------------|----------|-----------------|
| spec-app | model, idl, spec | 80% |
| gen-app | gen, sol, repos | 70% |
| sim-app | sim, mon, net | 70% |
| prj-app | prj, git | 70% |

### Test Writing Strategy

#### 1. Start with Pure Functions
Test functions with no side effects first:

```go
// Easy to test - no I/O, no state
func TestAbbreviate(t *testing.T) {
    assert.Equal(t, "ABC", helper.Abbreviate("ApiBaseClient"))
}

func TestSplitRepoID(t *testing.T) {
    name, version := repos.SplitRepoID("apigear/template@v1.0.0")
    assert.Equal(t, "apigear/template", name)
    assert.Equal(t, "v1.0.0", version)
}
```

#### 2. Use Table-Driven Tests
```go
func TestIsValidGitUrl(t *testing.T) {
    tests := []struct {
        name  string
        url   string
        want  bool
    }{
        {"https url", "https://github.com/org/repo.git", true},
        {"ssh url", "git@github.com:org/repo.git", true},
        {"invalid", "not-a-url", false},
        {"empty", "", false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := git.IsValidGitUrl(tt.url)
            assert.Equal(t, tt.want, got)
        })
    }
}
```

#### 3. Use Test Fixtures
Create `testdata/` directories for file-based tests:

```
pkg/model/
├── testdata/
│   ├── valid_module.yaml
│   ├── invalid_module.yaml
│   └── complex_system.yaml
└── model_test.go
```

#### 4. Mock External Dependencies
For packages that use I/O, create interfaces:

```go
// pkg/git/git.go
type GitClient interface {
    Clone(src, dst string) error
    Pull(dst string) error
}

// In tests, use mock implementation
type mockGitClient struct {
    cloneErr error
}
func (m *mockGitClient) Clone(src, dst string) error {
    return m.cloneErr
}
```

### Test Coverage Checklist

**Week 1-2: Foundation**
- [ ] Add tests for `helper` (target: 80%)
- [ ] Add tests for `cfg` (target: 70%)
- [ ] Add tests for `git` URL parsing (target: 50%)

**Week 3-4: Core Model**
- [ ] Increase `model` coverage (target: 70%)
- [ ] Increase `spec` coverage (target: 70%)
- [ ] Add tests for `repos` (target: 50%)

**Week 5-6: Domain Packages**
- [ ] Add tests for `prj` (target: 70%)
- [ ] Add tests for `sol` (target: 70%)
- [ ] Add tests for `net` (target: 50%)

**Ongoing: Maintain Coverage**
- [ ] Add coverage check to CI (fail if < 50%)
- [ ] Require tests for new code
- [ ] Track coverage trends

### Running Coverage Locally

```bash
# Overall coverage
go test -cover ./pkg/...

# Detailed coverage report
go test -coverprofile=coverage.out ./pkg/...
go tool cover -html=coverage.out -o coverage.html

# Coverage for specific package
go test -cover -coverprofile=pkg.out ./pkg/model/...
go tool cover -func=pkg.out

# Identify uncovered lines
go tool cover -func=coverage.out | grep -v "100.0%"
```

### CI Integration

Add to your CI pipeline:

```yaml
# .github/workflows/test.yml
- name: Run tests with coverage
  run: go test -coverprofile=coverage.out -covermode=atomic ./pkg/...

- name: Check coverage threshold
  run: |
    COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
    if (( $(echo "$COVERAGE < 50" | bc -l) )); then
      echo "Coverage $COVERAGE% is below 50% threshold"
      exit 1
    fi
```

---

## Preparation Steps

Small, low-risk changes that make future refactoring easier. Each can be done independently.

### 1. Add `api.go` to Each Package (1-2 hours per package)

Create a single file that documents the public interface:

```go
// pkg/model/api.go
package model

// Public API for model package
// All other exports are considered internal and may change

// NewSystem creates a new API system
func NewSystem(name string) *System { ... }

// System is the root container for API modules
type System struct { ... }

// Module represents an API module
type Module struct { ... }
```

**Why it helps**: Forces you to think about what's public, documents intent.

### 2. Create `internal/` Subdirectories (30 min per package)

Move implementation details to `internal/`:

```
pkg/model/
├── api.go           # Public interface
├── system.go        # System implementation
├── module.go        # Module implementation
└── internal/
    ├── validate.go  # Validation logic
    └── checksum.go  # Checksum calculation
```

**Why it helps**: Go enforces that `internal/` can't be imported from outside.

### 3. Replace Direct Config Access (1 day)

Currently packages import `cfg` directly. Add config interfaces:

```go
// pkg/model/api.go
type Config interface {
    GetString(key string) string
    GetBool(key string) bool
}

// Accept config as parameter instead of importing cfg
func NewSystemWithConfig(name string, cfg Config) *System { ... }
```

**Why it helps**: Removes global state, enables testing, prepares for DI.

### 4. Replace Direct Log Access (1 day)

Same pattern for logging:

```go
// pkg/model/api.go
type Logger interface {
    Debug() LogEvent
    Info() LogEvent
    Warn() LogEvent
    Error() LogEvent
}

type LogEvent interface {
    Str(key, val string) LogEvent
    Msg(msg string)
}
```

**Why it helps**: Decouples from zerolog, enables testing with mock loggers.

### 5. Reduce Helper Imports (1 day)

Many packages import `helper` for 1-2 functions. Copy those locally:

```go
// Before: pkg/git/clone.go
import "github.com/apigear-io/cli/pkg/helper"

func Clone(src, dst string) error {
    if helper.IsDir(dst) { ... }
}

// After: pkg/git/clone.go (no helper import)
func Clone(src, dst string) error {
    if isDir(dst) { ... }
}

func isDir(path string) bool {
    info, err := os.Stat(path)
    return err == nil && info.IsDir()
}
```

**Why it helps**: Reduces coupling, makes package self-contained.

### 6. Add Interface Files (2-3 days)

Create interface definitions without changing implementations:

```go
// pkg/model/iface.go
package model

// ISystem defines the public contract for System
type ISystem interface {
    Name() string
    Modules() []*Module
    LookupModule(name string) *Module
    Validate() error
}

// Ensure System implements ISystem
var _ ISystem = (*System)(nil)
```

**Why it helps**: Documents contracts, enables mocking, prepares for extraction.

### 7. Add Constructor Functions (1 day)

Replace direct struct creation with constructors:

```go
// Before
system := &model.System{Name: "test"}

// After
system := model.NewSystem("test")
```

**Why it helps**: Hides struct fields, allows internal changes, enables validation.

### 8. Group Related Tests (1 day)

Ensure tests are co-located with code they test:

```
pkg/model/
├── system.go
├── system_test.go     # Tests for system.go
├── module.go
├── module_test.go     # Tests for module.go
└── integration_test.go # Cross-cutting tests
```

**Why it helps**: Tests move with code during extraction.

### 9. Document Cross-Package Contracts (2-3 days)

Add comments documenting expected behavior:

```go
// pkg/gen/generator.go

// Generate processes a System and produces output files.
//
// Contract:
// - system must be validated (system.Validate() called)
// - outputDir must exist and be writable
// - templates must contain valid Go templates
//
// Returns GeneratorStats with counts of files written/skipped.
func (g *Generator) Generate(system *model.System) (*GeneratorStats, error)
```

**Why it helps**: Makes implicit contracts explicit before refactoring.

### 10. Add Package-Level README (Done!)

You've already done this step. Each package now has documentation.

---

## Preparation Checklist

| Step | Effort | Impact | Priority |
|------|--------|--------|----------|
| Add `api.go` files | 1-2 days | High | 1 |
| Create `internal/` dirs | 1 day | Medium | 2 |
| Add interface files | 2-3 days | High | 3 |
| Replace direct cfg access | 1 day | High | 4 |
| Replace direct log access | 1 day | Medium | 5 |
| Reduce helper imports | 1 day | Medium | 6 |
| Add constructor functions | 1 day | Low | 7 |
| Group related tests | 1 day | Low | 8 |
| Document contracts | 2-3 days | Medium | 9 |

**Total preparation: ~2 weeks** of incremental work

### Quick Wins (This Week)

1. **Add `api.go` to `model` and `gen`** - The two most complex packages
2. **Create interface for ISystem** - Most packages depend on this
3. **Copy `IsDir`/`IsFile` locally** - Most common helper functions

### Order of Package Preparation

Prepare leaf packages first (fewer dependencies to manage):

1. `helper` → `vfs` → `evt` → `tools` (no internal deps)
2. `cfg` → `log` (only depend on helper)
3. `git` → `tasks` → `tpl` → `up` (simple deps)
4. `model` → `mon` → `repos` → `prj` (medium complexity)
5. `idl` → `net` → `sim` (higher complexity)
6. `spec` → `gen` → `sol` (highest complexity, most deps)
7. `cmd` → `mcp` (orchestration layer)
