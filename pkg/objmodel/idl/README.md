# idl

Interface Definition Language (IDL) parser for API specifications.

## Purpose

The `idl` package provides parsing functionality for the APIGear IDL format. It implements an ANTLR4-based parser that reads IDL documents and builds a `model.System` containing:

- Modules with version information
- Interfaces with properties, operations, and signals
- Structs with typed fields
- Enums with members
- External type references

The parser supports metadata annotations via documentation comments and tags.

## Key Exports

- `Parser` - Main parser struct wrapping a model.System
- `NewParser()` - Creates a new parser instance
- `LoadIdlFromString()` - Parse IDL from string content
- `LoadIdlFromFiles()` - Parse IDL from one or more files
- `ParseFile()`, `ParseString()` - Parser methods
- `ObjectApiListener` - ANTLR4 listener implementation

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Configuration access |
| `helper` | File existence checking |
| `log` | Logging |
| `model` | AST data structures |
| `spec/rkw` | Reserved keyword validation |
