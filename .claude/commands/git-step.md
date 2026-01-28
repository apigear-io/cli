# Git Step - Commit Changes with Conventional Commits

Commit current changes using conventional commit format.

## Instructions

1. Run `git status` to check for changes
2. Run `git diff` to see the changes
3. Analyze the changes and determine the appropriate conventional commit type
4. Draft a conventional commit message following the format:
   ```
   <type>(<scope>): <description>

   [optional body]

   [optional footer]
   ```
5. Present the commit message to the user for approval
6. Stage the relevant files using `git add`
7. Create the commit with the approved message
8. Confirm the commit was successful with `git log -1`

## Conventional Commit Types

- `feat` - A new feature
- `fix` - A bug fix
- `docs` - Documentation only changes
- `style` - Changes that don't affect code meaning (formatting, etc.)
- `refactor` - Code change that neither fixes a bug nor adds a feature
- `perf` - Performance improvement
- `test` - Adding or correcting tests
- `build` - Changes to build system or dependencies
- `ci` - Changes to CI configuration
- `chore` - Other changes that don't modify src or test files

## Scope Examples

- Package names: `cfg`, `gen`, `mcp`, `idl`, etc.
- Component names: `filters`, `parser`, `commands`
- Feature areas: `auth`, `templates`, `monitoring`

## Message Guidelines

- Use imperative mood in description ("add" not "added" or "adds")
- Don't capitalize first letter of description
- No period at the end of description
- Keep description under 72 characters
- Use body to explain what and why (not how)
- Reference issues in footer: `Fixes #123` or `Closes #456`

## Examples

```
feat(gen): add support for external types in JNI filter
fix(mcp): correct tool annotations for registry operations
docs: add comprehensive package documentation
test(helper): add unit tests for string utilities
refactor(cmd): simplify command flag parsing
```

## Breaking Changes

For breaking changes, add `!` after type/scope and include `BREAKING CHANGE:` in footer:
```
feat(api)!: change configuration file format

BREAKING CHANGE: Configuration files now use YAML instead of JSON.
Migration guide available in docs/migration.md
```
