# CLI Regression Tests

This directory contains end-to-end regression tests for the apigear CLI using [testscript](https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript).

## Purpose

These tests lock down the user-facing CLI API (commands, arguments, flags, aliases) to ensure that refactoring doesn't break the command-line interface. They test the actual compiled binary, not just the Go code.

## Test Coverage

The test suite covers:

- **help_commands.txtar** - Root command help, subcommand help, and all aliases
- **version.txtar** - Version command output and flag validation
- **config_commands.txtar** - Config commands (info, get, env) and aliases
- **spec_check.txtar** - Spec validation commands and flags
- **generate_expert.txtar** - Code generation command structure and flags
- **template_list.txtar** - Template management commands
- **monitor_commands.txtar** - Monitor/debugging commands
- **project_commands.txtar** - Project management commands
- **experimental_commands.txtar** - Experimental (x) format conversion commands

## Running the Tests

Run all CLI regression tests:
```bash
go test -v -run TestCLIRegression ./tests
```

Run a specific test:
```bash
go test -v -run TestCLIRegression/help_commands ./tests
```

## Test Format

Tests use the `.txtar` format:
- Scripts contain commands and assertions
- Files can be embedded using `-- filename --` separator
- Commands are executed line by line
- `!` prefix means command should fail
- `stdout` and `stderr` assertions use regex patterns

### Example Test

```txtar
# Test a command
exec apigear generate --help
stdout 'Usage:'
stdout 'expert'
stdout 'solution'

# Test that invalid flags are rejected
! exec apigear generate --invalid-flag
stderr 'unknown flag'

-- embedded-file.yaml --
schema: apigear.module/1.0
name: test
```

## Key Assertions

- **Command exists**: `exec apigear command --help`
- **Alias works**: `exec apigear alias --help`
- **Flag exists**: `stdout '\-\-flag-name'`
- **Output format**: `stdout 'expected pattern'` or `stderr 'expected pattern'`
- **Flag rejection**: `! exec apigear cmd --invalid-flag` + `stderr 'unknown flag'`

## Adding New Tests

When adding new CLI commands or flags:

1. Create a new `.txtar` file in `tests/testscripts/` or add to existing file
2. Test the command help output
3. Test all aliases
4. Test required and optional flags
5. Test error cases (invalid flags, missing arguments)
6. Run `go test -v -run TestCLIRegression ./tests` to verify

## Before Major Refactoring

1. Run the full test suite to capture current behavior:
   ```bash
   go test -v -run TestCLIRegression ./tests
   ```

2. Ensure all tests pass (green baseline)

3. During refactoring, run tests frequently to catch breaking changes early

4. If tests fail, either:
   - Fix the code to maintain backward compatibility
   - Update tests if the CLI change is intentional (with team approval)

## Testing Strategy

These e2e tests complement the existing unit tests:

- **Unit tests** (`tests/*_test.go`) - Fast, in-process command testing
- **E2e tests** (`tests/testscripts/*.txtar`) - Actual binary execution testing

Both are important:
- Use unit tests for rapid development and detailed logic testing
- Use e2e tests as regression safeguards before releases

## Limitations

- Tests require network access for some commands (template registry)
- Some tests focus on command structure rather than full behavior
- Output format changes may require test updates

## Resources

- [testscript documentation](https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript)
- [testscript tutorial](https://bitfieldconsulting.com/golang/test-scripts)
- [txtar format](https://pkg.go.dev/golang.org/x/tools/txtar)
