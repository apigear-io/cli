# Linting Guide

## ESLint Configuration

This project uses **ESLint for linting only** - no formatting rules.

### What ESLint Checks

✅ **Code Quality & Correctness:**
- Potential bugs and errors
- Best practices
- TypeScript type safety
- React Hooks rules of hooks
- Unused variables and imports
- React Fast Refresh compatibility

❌ **NOT Checked by ESLint:**
- Code formatting (indentation, spacing, etc.)
- Semicolons vs no semicolons
- Quote styles
- Line length
- Trailing commas

### Running Linting

```bash
# Lint all files
pnpm lint

# Via task
task web:lint

# Fix auto-fixable issues
pnpm lint --fix
```

### Why Linting Only?

We separate concerns:
- **ESLint** = Code quality and correctness
- **TypeScript** = Type safety
- **Formatter** = Code style (if needed, use Prettier separately)

This keeps ESLint fast and focused on catching real issues, not arguing about style preferences.

### Configured Rules

Our ESLint config uses:

1. **JavaScript (`js.configs.recommended`)**
   - Basic JavaScript best practices
   - Potential error detection

2. **TypeScript (`tseslint.configs.recommended`)**
   - TypeScript-specific linting
   - Type-aware rules
   - No stylistic rules

3. **React Hooks (`react-hooks/recommended`)**
   - Rules of Hooks enforcement
   - Dependencies array validation

4. **React Refresh**
   - Fast refresh compatibility warnings
   - Disabled for test files

5. **Custom Rules**
   - Unused variables as errors (with `_` prefix exception)

### Disabling Rules

If you need to disable a rule for a specific line:

```typescript
// eslint-disable-next-line rule-name
const something = dangerous();
```

For entire files (use sparingly):
```typescript
/* eslint-disable rule-name */
```

### Configuration Files

- `eslint.config.js` - Main ESLint config (flat config format)
- `package.json` - Contains lint scripts
- `.eslintignore` - Not needed (ignores in config file)

### Common Issues

**"no-unused-vars" errors:**
Prefix unused parameters with underscore:
```typescript
function handler(_req, res) { // _req not used
  res.send('ok');
}
```

**"react-hooks/exhaustive-deps" warnings:**
Add missing dependencies or use `// eslint-disable-next-line` if intentional.

**"react-refresh/only-export-components" warnings:**
Only export components from component files. Disabled in test files.

## IDE Integration

### VS Code
Install the ESLint extension:
```bash
code --install-extension dbaeumer.vscode-eslint
```

### IntelliJ/WebStorm
ESLint support is built-in. Enable it in:
Settings → Languages & Frameworks → JavaScript → Code Quality Tools → ESLint

## CI/CD

Linting runs automatically in CI via:
```bash
task ci:all
```

The build fails if linting errors are found.
