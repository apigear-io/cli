# spec

Specification types and validation framework for APIGear documents.

## Purpose

The `spec` package defines and validates the core document types used in code generation:

- **Module** documents - API definitions (.idl files)
- **Solution** documents - Generation targets and configuration
- **Scenario** documents - Test/simulation scenarios
- **Rules** documents - Transformation rules for code generation

It provides JSON Schema validation, format conversion, and reserved keyword checking.

## Key Exports

### Document Types
- `DocumentType` - Enum (Module, Solution, Scenario, Rules, Unknown)
- `SolutionDoc` - Solution configuration with targets
- `SolutionTarget` - Individual generation target
- `ScenarioDoc` - Test scenario definitions
- `RulesDoc` - Rules with features and version constraints
- `FeatureRule`, `ScopeRule`, `DocumentRule` - Rule components

### Validation Functions
- `CheckFile()`, `CheckFileAndType()` - Validate specification files
- `CheckJson()` - Validate JSON against schemas
- `CheckCsvFile()`, `CheckIdlFile()`, `CheckJsFile()` - Format-specific validation

### Schema Functions
- `LoadSchema()`, `ShowSchemaFile()` - Schema access
- `GetDocumentType()`, `DocumentTypeFromFileName()` - Type detection
- `YamlToJson()`, `JsonToYaml()` - Format conversion

### Sub-package: rkw (Reserved Keywords)
- `Lang` - Enum for languages (C++, Python, TypeScript, JavaScript, Go, Unreal, Qt)
- Reserved keyword lists for each language

## Dependencies

| Package | Purpose |
|---------|---------|
| `cfg` | Configuration access |
| `git` | Git operations |
| `helper` | File operations, input expansion |
| `idl` | IDL file parsing |
| `log` | Logging |
| `model` | System model for validation |
| `mon` | Monitoring |
| `net` | Network operations |
| `repos` | Template directory access |
| `sim` | JavaScript compilation |
