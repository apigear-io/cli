# Test Coverage Expansion Plan

## Current State

### Strong Coverage (70%+)
- `pkg/idl` - 93.2% (excellent!)
- `pkg/gen/filters/*` - 74-86% (good filter coverage)
- `pkg/evt` - 69.9%

### Needs Improvement (0-50%)
- 28 packages with 0% coverage
- Several core packages under 50%

## Priority Recommendations

### 1. High-Impact, Easy Wins (Start Here)

These packages have pure functions that are straightforward to test:

#### `pkg/helper` (0% → Target: 80%+)

Pure utility functions are ideal test candidates:
- `strings.go` - Test `Contains()`, `Abbreviate()`, `MapToArray()`, `ArrayToMap()`
- `ids.go` - Test ID generators
- `maps.go`, `iter.go` - Collection utilities

Example test structure:
```go
func TestAbbreviate(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"HelloWorld", "HW"},
        {"API2Gateway", "AG2"},
        {"simple", "S"},
    }
    for _, tt := range tests {
        assert.Equal(t, tt.expected, Abbreviate(tt.input))
    }
}
```

### 2. Core Business Logic (High Priority)

#### `pkg/cfg` (0% → Target: 70%+)

Configuration management is critical. Test:
- Config loading/saving
- Validation logic
- Default values

#### `pkg/prj` (0% → Target: 60%+)

Project operations. Test:
- Project file reading/parsing
- Model validation
- Demo generation

#### `pkg/repos` (12.3% → Target: 60%+)

Template repository management. Expand:
- Repository ID parsing (already has some tests)
- Version handling
- Repository validation

### 3. Integration Components (Medium Priority)

#### `pkg/git` (0% → Target: 40%+)

Git operations need tests with mocking:
- Use interfaces to mock git operations
- Test URL parsing, version extraction
- Mock file system operations

#### `pkg/net` (0% → Target: 50%+)

Network utilities:
- Mock HTTP requests
- Test error handling
- Validate request/response parsing

### 4. Command Layer (Medium-Low Priority)

#### `pkg/cmd/*` packages (mostly 0%)

CLI commands are harder to test but important:
- Test command validation logic
- Mock underlying service calls
- Test flag parsing and validation
- Focus on `pkg/cmd/cfg` (28.6%) as a template

### 5. Expand Existing Coverage

#### `pkg/model` (34.9% → Target: 70%+)
- Add edge case tests
- Test validation methods
- Test model transformations

#### `pkg/spec` (42.9% → Target: 70%+)
- More complex rule scenarios
- Schema validation edge cases
- Error path testing

## Testing Strategy Recommendations

### 1. Add Test Helpers

Create a `testdata/` directory with:
- Sample IDL files
- Mock configurations
- Test templates
- Fixture data

### 2. Table-Driven Tests

You already use this pattern well. Expand it:
```go
func TestFunction(t *testing.T) {
    tests := []struct {
        name     string
        input    InputType
        expected OutputType
        wantErr  bool
    }{
        // test cases
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test logic
        })
    }
}
```

### 3. Mock External Dependencies

For packages like `git`, `net`, `mcp`:
- Define interfaces for external operations
- Create mock implementations
- Test business logic in isolation

### 4. Integration Tests

Expand the `tests/` package (currently 100%):
- End-to-end workflows
- Multi-package interactions
- Real-world scenarios

### 5. Benchmark Tests

For performance-critical code like filters and generation:
```go
func BenchmarkAbbreviate(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Abbreviate("HelloWorldExample")
    }
}
```

## Quick Start: First 5 Tests to Write

1. **`pkg/helper/strings_test.go`** - Test `Abbreviate()` and `Contains()`
2. **`pkg/helper/ids_test.go`** - Test ID generators
3. **`pkg/cfg/config_test.go`** - Test config loading
4. **`pkg/prj/models_test.go`** - Test model validation
5. **`pkg/repos/repoid_test.go`** - Expand existing tests

## Measuring Progress

Update your Taskfile to track coverage over time:
```yaml
test:cover:report:
  desc: Generate coverage report with statistics
  cmds:
    - go test -coverprofile=coverage.txt ./...
    - go tool cover -func=coverage.txt | grep total
```

## Target Milestones

- **Phase 1**: Get all utility packages (`helper`, `cfg`) to 70%+
- **Phase 2**: Core business logic to 60%+
- **Phase 3**: Overall project coverage to 50%+

## Coverage by Package (Baseline)

### 0% Coverage
- `cmd/apigear`
- `pkg/cfg`
- `pkg/cmd` (base)
- `pkg/cmd/gen`
- `pkg/cmd/mon`
- `pkg/cmd/olink`
- `pkg/cmd/prj`
- `pkg/cmd/spec`
- `pkg/cmd/tpl`
- `pkg/cmd/x`
- `pkg/gen/filters` (base)
- `pkg/git`
- `pkg/helper`
- `pkg/idl/parser`
- `pkg/log`
- `pkg/mcp`
- `pkg/mcp/gen`
- `pkg/mcp/spec`
- `pkg/mcp/tpl`
- `pkg/net`
- `pkg/prj`
- `pkg/sol`
- `pkg/tasks`
- `pkg/tools`
- `pkg/tpl`
- `pkg/up`

### Low Coverage (1-50%)
- `pkg/repos` - 12.3%
- `pkg/cmd/cfg` - 28.6%
- `pkg/model` - 34.9%
- `pkg/mon` - 40.9%
- `pkg/spec` - 42.9%
- `pkg/spec/rkw` - 43.9%
- `pkg/gen/filters/common` - 47.8%

### Good Coverage (51-70%)
- `pkg/gen` - 59.1%
- `pkg/gen/filters/filterjava` - 61.7%
- `pkg/evt` - 69.9%

### Excellent Coverage (71%+)
- `pkg/gen/filters/filterue` - 74.4%
- `pkg/gen/filters/filterjs` - 77.0%
- `pkg/gen/filters/filterts` - 77.0%
- `pkg/gen/filters/filtergo` - 77.3%
- `pkg/gen/filters/filterjni` - 80.1%
- `pkg/gen/filters/filterrs` - 80.9%
- `pkg/gen/filters/filtercpp` - 82.4%
- `pkg/gen/filters/filterpy` - 84.1%
- `pkg/gen/filters/filterqt` - 85.7%
- `pkg/idl` - 93.2%
- `tests` - 100.0%
