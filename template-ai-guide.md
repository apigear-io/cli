# ApiGear Template AI Coding Guide

A comprehensive reference for AI coding agents working with ApiGear templates using Go text/template language and custom filters.

---

## Go Text/Template Quick Reference

### Basic Syntax

```go
{{ .Variable }}              // Output variable value
{{ .Object.Field }}          // Access nested field
{{ .Method }}                // Call method with no args
{{ .Method arg1 arg2 }}      // Call method with args
```

### Actions

```go
{{/* This is a comment */}}

{{ if .Condition }}...{{ end }}
{{ if .Condition }}...{{ else }}...{{ end }}
{{ if .Condition }}...{{ else if .Other }}...{{ end }}

{{ range .Items }}
  {{ . }}                    // Current item
  {{ $.RootVar }}            // Access root context with $
{{ end }}

{{ range $index, $item := .Items }}
  {{ $index }}: {{ $item }}
{{ end }}

{{ with .Object }}
  {{ .Field }}               // Scoped to .Object
{{ end }}
```

### Variables

```go
{{ $var := .Value }}         // Declare variable
{{ $var }}                   // Use variable
{{ $var = .NewValue }}       // Reassign variable
```

### Whitespace Control

```go
{{- .Var }}                  // Trim left whitespace
{{ .Var -}}                  // Trim right whitespace
{{- .Var -}}                 // Trim both sides
```

### Pipelines

```go
{{ .Name | upper }}          // Pipe to filter
{{ .Name | upper | trim }}   // Chain filters
{{ printf "%s: %d" .Name .Count }}  // printf formatting
```

### Built-in Functions

```go
{{ and .A .B }}              // Logical AND
{{ or .A .B }}               // Logical OR
{{ not .A }}                 // Logical NOT
{{ eq .A .B }}               // Equal
{{ ne .A .B }}               // Not equal
{{ lt .A .B }}               // Less than
{{ le .A .B }}               // Less than or equal
{{ gt .A .B }}               // Greater than
{{ ge .A .B }}               // Greater than or equal
{{ len .Array }}             // Length of array/string/map
{{ index .Array 0 }}         // Index into array
{{ index .Map "key" }}       // Index into map
{{ printf "%s" .Val }}       // Formatted printing
{{ print .Val }}             // Simple printing
{{ println .Val }}           // Print with newline
```

### Template Inclusion

```go
{{ template "name" . }}      // Include template with data
{{ define "name" }}...{{ end }}  // Define named template
{{ block "name" . }}...{{ end }} // Define with default content
```

---

## ApiGear Model Context

Templates receive a context with the following structure:

```go
// Root context variables
.Module          // Current module being processed
.System          // System-wide information
.Imports         // Import declarations
.Externs         // External type definitions

// Module fields
.Module.Name     // Module name (e.g., "org.example")
.Module.Interfaces
.Module.Structs
.Module.Enums
.Module.Externs

// Interface fields
.Interface.Name
.Interface.Properties
.Interface.Operations
.Interface.Signals

// Property/Parameter fields (TypedNode)
.Name            // Variable name
.Schema          // Type information
.Description     // Documentation
```

---

## Common Filters

### Case Conversion

| Filter | Input | Output | Example |
|--------|-------|--------|---------|
| `snake` | `MyVar` | `my_var` | `{{ .Name \| snake }}` |
| `Snake` | `MyVar` | `My_Var` | `{{ .Name \| Snake }}` |
| `SNAKE` | `MyVar` | `MY_VAR` | `{{ .Name \| SNAKE }}` |
| `camel` | `my_var` | `myVar` | `{{ .Name \| camel }}` |
| `Camel` | `my_var` | `MyVar` | `{{ .Name \| Camel }}` |
| `CAMEL` | `my_var` | `MYVAR` | `{{ .Name \| CAMEL }}` |
| `kebap` | `MyVar` | `my-var` | `{{ .Name \| kebap }}` |
| `Kebab` | `MyVar` | `My-Var` | `{{ .Name \| Kebab }}` |
| `KEBAP` | `MyVar` | `MY-VAR` | `{{ .Name \| KEBAP }}` |
| `dot` | `MyVar` | `my.var` | `{{ .Name \| dot }}` |
| `Dot` | `MyVar` | `My.Var` | `{{ .Name \| Dot }}` |
| `DOT` | `MyVar` | `MY.VAR` | `{{ .Name \| DOT }}` |
| `space` | `MyVar` | `my var` | `{{ .Name \| space }}` |
| `Space` | `MyVar` | `My Var` | `{{ .Name \| Space }}` |
| `SPACE` | `MyVar` | `MY VAR` | `{{ .Name \| SPACE }}` |
| `path` | `MyVar` | `my/var` | `{{ .Name \| path }}` |
| `Path` | `MyVar` | `My/Var` | `{{ .Name \| Path }}` |
| `PATH` | `MyVar` | `MY/VAR` | `{{ .Name \| PATH }}` |
| `lower` | `MyVar` | `myvar` | `{{ .Name \| lower }}` |
| `upper` | `MyVar` | `MYVAR` | `{{ .Name \| upper }}` |
| `upper1` | `myVar` | `MyVar` | `{{ .Name \| upper1 }}` |
| `lower1` | `MyVar` | `myVar` | `{{ .Name \| lower1 }}` |
| `first` | `MyVar` | `m` | `{{ .Name \| first }}` |
| `First` | `myVar` | `m` | `{{ .Name \| First }}` |
| `FIRST` | `myVar` | `M` | `{{ .Name \| FIRST }}` |

### String Manipulation

| Filter | Description | Example |
|--------|-------------|---------|
| `join` | Join array with separator | `{{ join ", " .Items }}` |
| `split` | Split string by separator | `{{ split .Name "." }}` |
| `splitFirst` | Get first part before separator | `{{ splitFirst .Name "." }}` |
| `splitLast` | Get last part after separator | `{{ splitLast .Name "." }}` |
| `trim` | Remove leading/trailing whitespace | `{{ .Name \| trim }}` |
| `trimPrefix` | Remove prefix | `{{ trimPrefix .Name "pre_" }}` |
| `trimSuffix` | Remove suffix | `{{ trimSuffix .Name "_suf" }}` |
| `replace` | Replace all occurrences | `{{ replace .Name "old" "new" }}` |
| `contains` | Check if array contains string | `{{ if contains .Tags "api" }}` |
| `indexOf` | Get index of element (-1 if not found) | `{{ indexOf .Items "value" }}` |

### Array Operations

| Filter | Description | Example |
|--------|-------------|---------|
| `appendList` | Append to string list | `{{ $list = appendList $list "item" }}` |
| `getEmptyStringList` | Create empty string slice | `{{ $list := getEmptyStringList }}` |
| `unique` | Get sorted unique elements | `{{ unique .Items }}` |
| `collectFields` | Extract field from struct array | `{{ collectFields .Items "Name" }}` |
| `strSlice` | Create string slice | `{{ strSlice "a" "b" "c" }}` |

### Number to Word

| Filter | Description | Example |
|--------|-------------|---------|
| `int2word` | Number to lowercase word | `{{ int2word 1 "" "" }}` → `one` |
| `Int2Word` | Number to title word | `{{ Int2Word 2 "" "" }}` → `Two` |
| `INT2WORD` | Number to uppercase word | `{{ INT2WORD 3 "" "" }}` → `THREE` |
| `plural` | Pluralize if count > 1 | `{{ plural "item" .Count }}` |

### Utility

| Filter | Description | Example |
|--------|-------------|---------|
| `nl` | Insert newline | `{{ nl }}` |
| `toJson` | Convert to JSON | `{{ toJson .Object }}` |
| `abbreviate` | Abbreviate string | `{{ abbreviate .Name }}` |

---

## Language-Specific Filters

Each language has a consistent set of filters with the prefix pattern:

- `<lang>Return` / `<lang>Type` - Convert to language type
- `<lang>Default` - Get default/zero value
- `<lang>Param` - Format single parameter
- `<lang>Params` - Format parameter list
- `<lang>Var` - Get variable name
- `<lang>Vars` - Get comma-separated variable names

### C++ Filters (prefix: `cpp`)

```go
{{ cppReturn "" .Property }}     // string → std::string, int → int32_t
{{ cppType "" .Property }}       // Alias for cppReturn
{{ cppTypeRef "" .Property }}    // const std::string& (reference type)
{{ cppDefault "" .Property }}    // "", 0, false, nullptr
{{ cppParam "" .Property }}      // "const std::string& name"
{{ cppParams "" .Properties }}   // "const std::string& a, int32_t b"
{{ cppVar .Property }}           // "name"
{{ cppVars .Properties }}        // "a, b, c"
{{ cppNs .Module }}              // "org::example" (namespace)
{{ cppNsOpen .Module }}          // "namespace org { namespace example {"
{{ cppNsClose .Module }}         // "} // namespace example } // namespace org"
{{ cppGpl .Module }}             // GPL license header
{{ cppExtern .Extern }}          // Parse extern metadata
{{ cppTestValue "" .Property }}  // Test/example value
```

### Go Filters (prefix: `go`)

```go
{{ goReturn "" .Property }}      // string, int32, []string
{{ goType "" .Property }}        // Alias for goReturn
{{ goDefault "" .Property }}     // "", int32(0), []string{}, nil
{{ goParam "" .Property }}       // "name string"
{{ goParams "" .Properties }}    // "a string, b int32"
{{ goVar .Property }}            // "name"
{{ goPublicVar .Property }}      // "Name" (PascalCase)
{{ goVars .Properties }}         // "a, b, c"
{{ goPublicVars .Properties }}   // "A, B, C"
{{ goDoc .Interface }}           // "// Documentation comment"
{{ goExtern .Extern }}           // Parse extern metadata
```

### TypeScript Filters (prefix: `ts`)

```go
{{ tsReturn "" .Property }}      // string, number, boolean
{{ tsType "" .Property }}        // Alias for tsReturn
{{ tsDefault "" .Property }}     // "", 0, false, null
{{ tsParam "" .Property }}       // "name: string"
{{ tsParams "" .Properties }}    // "a: string, b: number"
{{ tsVar .Property }}            // "name"
{{ tsVars .Properties }}         // "a, b, c"
```

### Python Filters (prefix: `py`)

```go
{{ pyReturn "" .Property }}      // str, int, float, bool, list[Type]
{{ pyType "" .Property }}        // Alias for pyReturn
{{ pyDefault "" .Property }}     // "", 0, 0.0, False, [], None
{{ pyParam "" .Property }}       // "name: str" (snake_case)
{{ pyParams "" .Properties }}    // "self, a: str, b: int" (includes self)
{{ pyFuncParams "" .Properties }} // "a: str, b: int" (no self)
{{ pyVar .Property }}            // "name" (snake_case)
{{ pyVars .Properties }}         // "a, b, c"
{{ pyExtern .Extern }}           // Parse extern metadata
{{ pyTestValue "" .Property }}   // Test/example value
```

### Java Filters (prefix: `java`)

```go
{{ javaReturn "" .Property }}    // String, Integer, Long, Double, Boolean
{{ javaType "" .Property }}      // Alias for javaReturn
{{ javaDefault "" .Property }}   // null, 0, false
{{ javaParam "" .Property }}     // "String name" or "String[] names"
{{ javaParams "" .Properties }}  // "String a, Integer b"
{{ javaVar .Property }}          // "name"
{{ javaVars .Properties }}       // "a, b, c"
{{ javaAsyncReturn "" .Property }} // CompletableFuture return type
{{ javaElementType .Property }}  // Element type for arrays
{{ javaExtern .Extern }}         // Parse extern metadata
{{ javaTestValue "" .Property }} // Test/example value
```

### JNI Filters (prefix: `jni`)

```go
{{ jniToReturnType .Property }}  // jstring, jint, jlong, jobject
{{ jniJavaParam "" .Property }}  // Java param for JNI
{{ jniJavaParams "" .Properties }} // JNI Java params
{{ jniSignatureType .Property }} // JNI signature format
{{ jniJavaSignatureParam "" .Property }}  // JNI signature param
{{ jniJavaSignatureParams "" .Properties }} // JNI signature params
{{ jniToEnvNameType .Property }} // Env name and type
{{ jniEmptyReturn .Property }}   // Check if void return
```

### Rust Filters (prefix: `rs`)

```go
{{ rsReturn "" .Property }}      // &str, i32, i64, f32, f64, bool
{{ rsType "" .Property }}        // Alias for rsReturn
{{ rsTypeRef "" .Property }}     // Type with reference qualifier
{{ rsDefault "" .Property }}     // Default/zero value
{{ rsParam "" "" .Property }}    // Parameter with reference handling
{{ rsParams "" "" .Properties }} // Comma-separated params
{{ rsVar .Property }}            // Variable name
{{ rsVars .Properties }}         // Comma-separated names
{{ rsNs .Module }}               // Rust module namespace
{{ rsNsOpen .Module }}           // Module opening
{{ rsNsClose .Module }}          // Module closing
{{ rsExtern .Extern }}           // Parse extern metadata
```

### JavaScript Filters (prefix: `js`)

```go
{{ jsReturn "" .Property }}      // Type info (no explicit types)
{{ jsType "" .Property }}        // Alias for jsReturn
{{ jsDefault "" .Property }}     // Default value
{{ jsParam "" .Property }}       // Parameter name (no type hints)
{{ jsParams "" .Properties }}    // Comma-separated param names
{{ jsVar .Property }}            // Variable name
{{ jsVars .Properties }}         // Comma-separated names
```

### Qt (C++ Qt) Filters (prefix: `qt`)

```go
{{ qtReturn "" .Property }}      // QString, QList, qint32, qreal
{{ qtType "" .Property }}        // Alias for qtReturn
{{ qtDefault "" .Property }}     // Default value
{{ qtParam "" .Property }}       // "const QString& name"
{{ qtParams "" .Properties }}    // Comma-separated params
{{ qtVar .Property }}            // Variable name
{{ qtVars .Properties }}         // Comma-separated names
{{ qtNamespace .Module.Name }}   // Qt namespace format
{{ qtExtern .Extern }}           // Parse extern (namespace, include)
{{ qtExterns .Externs }}         // Array of QtExtern structs
{{ qtTestValue "" .Property }}   // Test/example value
```

### Unreal Engine Filters (prefix: `ue`)

```go
{{ ueReturn "" .Property }}      // FString, TArray, int32, float, bool
{{ ueType "" .Property }}        // Alias for ueReturn
{{ ueConstType "" .Property }}   // Const type representation
{{ ueDefault "" .Property }}     // Default value
{{ ueParam "" .Property }}       // "const FString& Name"
{{ ueParams "" .Properties }}    // Comma-separated params
{{ ueVar .Property }}            // "Name" (PascalCase)
{{ ueVars .Properties }}         // Comma-separated PascalCase names
{{ ueIsStdSimpleType .Property }} // true for int, float, bool, enum
{{ ueExtern .Extern }}           // Parse extern metadata
{{ ueTestValue "" .Property }}   // Test/example value
```

---

## Common Template Patterns

### Iterating Over Interfaces

```go
{{- range .Module.Interfaces }}
class {{ .Name | Camel }} {
{{- range .Properties }}
    {{ cppType "" . }} {{ .Name | camel }};
{{- end }}
};
{{- end }}
```

### Generating Method Signatures

```go
{{- range .Interface.Operations }}
{{ cppReturn "" . }} {{ .Name | camel }}({{ cppParams "" .Params }});
{{- end }}
```

### Conditional Type Handling

```go
{{- if eq .Schema.Type "array" }}
std::vector<{{ cppReturn "" .Schema.Items }}>
{{- else }}
{{ cppReturn "" . }}
{{- end }}
```

### Namespace Wrapping

```go
{{ cppNsOpen .Module }}

// Your code here

{{ cppNsClose .Module }}
```

### Building Include Lists

```go
{{- $includes := getEmptyStringList }}
{{- range .Module.Externs }}
{{- $extern := cppExtern . }}
{{- $includes = appendList $includes $extern.Include }}
{{- end }}
{{- range unique $includes }}
#include "{{ . }}"
{{- end }}
```

### Parameter Lists with Commas

```go
void method({{ range $i, $p := .Params }}{{ if $i }}, {{ end }}{{ cppParam "" $p }}{{ end }})
```

### Enum Generation

```go
{{- range .Module.Enums }}
enum class {{ .Name | Camel }} {
{{- range $i, $m := .Members }}
    {{ $m.Name | Camel }} = {{ $m.Value }}{{ if lt $i (sub (len $.Members) 1) }},{{ end }}
{{- end }}
};
{{- end }}
```

### Struct Generation

```go
{{- range .Module.Structs }}
struct {{ .Name | Camel }} {
{{- range .Fields }}
    {{ cppType "" . }} {{ .Name | camel }};
{{- end }}
};
{{- end }}
```

---

## Type Mappings Reference

### Schema Types to Language Types

| Schema Type | C++ | Go | TypeScript | Python | Java | Rust | UE |
|-------------|-----|-----|------------|--------|------|------|-----|
| `string` | `std::string` | `string` | `string` | `str` | `String` | `&str` | `FString` |
| `int` | `int32_t` | `int32` | `number` | `int` | `Integer` | `i32` | `int32` |
| `int32` | `int32_t` | `int32` | `number` | `int` | `Integer` | `i32` | `int32` |
| `int64` | `int64_t` | `int64` | `number` | `int` | `Long` | `i64` | `int64` |
| `float` | `float` | `float32` | `number` | `float` | `Float` | `f32` | `float` |
| `float32` | `float` | `float32` | `number` | `float` | `Float` | `f32` | `float` |
| `float64` | `double` | `float64` | `number` | `float` | `Double` | `f64` | `double` |
| `bool` | `bool` | `bool` | `boolean` | `bool` | `Boolean` | `bool` | `bool` |
| `array` | `std::list<T>` | `[]T` | `T[]` | `list[T]` | `T[]` | `Vec<T>` | `TArray<T>` |

---

## Tips for AI Agents

1. **Always use the prefix**: Language filters require a prefix parameter (usually `""` for default).

2. **TypedNode vs Schema**: Most filters expect `*model.TypedNode` which contains both name and type info.

3. **Whitespace matters**: Use `{{-` and `-}}` to control whitespace in generated code.

4. **Use pipelines**: Chain filters for complex transformations: `{{ .Name | snake | upper }}`

5. **Access root context**: Use `$.` to access root context inside `range` or `with` blocks.

6. **Error handling**: Most filters return `(string, error)` - errors will stop template execution.

7. **Case conventions vary**: Each language has its own naming conventions (PascalCase for UE, snake_case for Python, etc.).

8. **Externs are special**: Use language-specific `*Extern` filters to parse external type metadata.

9. **Test values**: Use `*TestValue` filters when generating test fixtures or mock data.

10. **Parameter order**: For `*Param` filters, prefix comes first, then the node: `{{ cppParam "" .Property }}`
