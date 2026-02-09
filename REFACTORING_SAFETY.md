# Refactoring Safety - CLI Regression Testing

## Overview

The CLI now has comprehensive end-to-end regression tests that lock down the user-facing API. These tests ensure that refactoring internal code doesn't break the command-line interface that users depend on.

## What's Protected

The regression test suite verifies:

1. **Command Structure** - All commands and subcommands exist
2. **Aliases** - Short forms like `gen`, `cfg`, `mon` still work
3. **Flags** - Required and optional flags are present
4. **Help Output** - Usage information is accessible
5. **Error Handling** - Invalid flags are properly rejected
6. **Exit Codes** - Commands fail appropriately

## Test Infrastructure

- **Location**: `tests/cli_regression_test.go`
- **Test Scripts**: `tests/testscripts/*.txtar`
- **Technology**: [testscript](https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript)
- **Coverage**: 9 test scenarios covering all major command groups

## Commands Covered

| Test File | Coverage |
|-----------|----------|
| `help_commands.txtar` | Root help, subcommand help, all aliases |
| `version.txtar` | Version output and flag validation |
| `config_commands.txtar` | Config info, get, env + aliases |
| `spec_check.txtar` | Spec validation and schema commands |
| `generate_expert.txtar` | Code generation structure |
| `template_list.txtar` | Template management |
| `monitor_commands.txtar` | Monitor/debug commands |
| `project_commands.txtar` | Project management |
| `experimental_commands.txtar` | Format conversion commands |

## Running Tests

### Before Starting Refactoring
```bash
# Establish green baseline
go test -v -run TestCLIRegression ./tests
```

### During Refactoring
```bash
# Quick check
go test -run TestCLIRegression ./tests

# Verbose output for debugging
go test -v -run TestCLIRegression ./tests

# Run specific test
go test -v -run TestCLIRegression/help_commands ./tests
```

### After Completing Changes
```bash
# Full test suite
go test ./tests
```

## Interpreting Failures

If a test fails during refactoring:

1. **Unintentional Breaking Change**
   - Review the failure output
   - Fix your code to maintain backward compatibility
   - Re-run tests to verify the fix

2. **Intentional CLI Change**
   - Discuss with team - is this a breaking change?
   - Update the test to reflect new behavior
   - Document the change in release notes
   - Consider deprecation warnings for removed features

## Example Failure

```
FAIL: testscripts/generate_expert.txtar:10: no match for `--template` found in stdout
```

This means:
- The test expected to find `--template` flag in help output
- The flag might have been renamed or removed
- Action: Either restore the flag or update the test (with approval)

## Testing Strategy

We now have two complementary test layers:

### 1. Unit Tests (Existing)
- **Location**: `tests/*_test.go`
- **Purpose**: Fast development feedback
- **Method**: In-process command execution
- **Use When**: Developing new features, testing logic

### 2. E2E Regression Tests (New)
- **Location**: `tests/testscripts/*.txtar`
- **Purpose**: Prevent breaking changes
- **Method**: Actual binary execution
- **Use When**: Before/during refactoring, before releases

## Recommended Workflow

1. **Start Refactoring**
   ```bash
   go test -run TestCLIRegression ./tests  # Green baseline
   ```

2. **Make Changes**
   - Refactor internal code
   - Run unit tests frequently for quick feedback
   ```bash
   go test ./pkg/...
   ```

3. **Check CLI Stability**
   - After significant changes, verify CLI integrity
   ```bash
   go test -run TestCLIRegression ./tests
   ```

4. **Before Commit**
   - Run full test suite
   ```bash
   go test ./...
   ```

## Extending Coverage

When adding new CLI features:

1. Add unit tests first (TDD approach)
2. Implement the feature
3. Add testscript regression test:
   ```txtar
   # Test new command
   exec apigear newcmd --help
   stdout 'Usage:'
   stdout 'expected-flag'

   # Test alias
   exec apigear nc --help
   stdout 'newcmd'
   ```

See `tests/testscripts/README.md` for detailed examples.

## Benefits for Refactoring

This safety net allows you to:

- **Refactor Confidently** - Internal changes won't break user workflows
- **Catch Regressions Early** - Failing tests show exactly what broke
- **Document Behavior** - Tests serve as executable specifications
- **Speed Up Reviews** - Reviewers can trust that CLI behavior is preserved
- **Automate Verification** - CI can catch breaking changes before merge

## CI Integration

Add to your CI pipeline:

```yaml
- name: Run CLI Regression Tests
  run: go test -v -run TestCLIRegression ./tests
```

This ensures no breaking changes reach production.

## Next Steps

1. Add these tests to your CI/CD pipeline
2. Run baseline test before starting refactoring:
   ```bash
   go test -v -run TestCLIRegression ./tests > baseline.txt
   ```
3. Begin refactoring with confidence
4. Extend coverage as needed for critical workflows

## Questions?

See `tests/testscripts/README.md` for detailed documentation on the test framework and how to add new tests.
