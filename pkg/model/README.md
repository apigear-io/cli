# model

Domain model and metadata representation for API specifications.

## Purpose

The `model` package defines the core data structures representing an API specification system. It provides:

- **Hierarchical Model**: System -> Modules -> Interfaces/Structs/Enums -> Members
- **Type System**: Primitives, symbols (custom types), arrays, type resolution
- **Schema Validation**: Type checking and cross-module reference resolution
- **Visitor Pattern**: Tree traversal for code generation
- **Serialization**: JSON/YAML parsing and unmarshaling
- **Reserved Word Checking**: Identifier validation across languages

## Key Exports

### Core Types
- `System` - Root container for all modules
- `Module` - Collection of interfaces, structs, enums
- `Interface` - Properties, operations, signals
- `Struct` - Named composite type with fields
- `Enum` - Enumeration with members
- `Extern` - External/opaque types

### Type System
- `Schema` - Type information with lazy resolution
- `TypedNode` - Node with type schema
- `KindType` - Type classifiers (void, bool, int, string, etc.)

### Scopes (for code generation)
- `SystemScope`, `ModuleScope`, `InterfaceScope`, `StructScope`, `EnumScope`, `ExternScope`

### Utilities
- `ModelVisitor` - Interface for tree traversal
- `DataParser` - JSON/YAML parser for API definitions

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Configuration access |
| `helper` | Utility functions |
| `log` | Logging |
| `spec/rkw` | Reserved keyword validation |
