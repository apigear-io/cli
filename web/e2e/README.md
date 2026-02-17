# E2E Tests

This directory contains end-to-end tests using Playwright.

## Running Tests

```bash
# Run all E2E tests (headless)
pnpm test:e2e

# Run with Playwright UI (interactive)
pnpm test:e2e:ui

# Run in debug mode
pnpm test:e2e:debug

# Or use task commands
task web:test:e2e
task web:test:e2e:ui
```

## Configuration

The E2E tests are configured in `playwright.config.ts` at the project root.

### Key Settings:
- **Dev Server**: Automatically starts on port 5173 (Vite default)
- **Base URL**: http://localhost:5173
- **Browsers**: Chromium, Firefox, WebKit
- **API Mocking**: Tests mock API responses by default

## API Mocking

The E2E tests include API mocking to allow testing without a running backend. Mock responses are defined in each test file using Playwright's `page.route()` method.

### Testing with Real Backend

To test against the real backend API:

1. Start the backend server:
   ```bash
   task run -- serve --port 8080
   ```

2. Remove or comment out the API mocking in your test files

3. Run the tests:
   ```bash
   pnpm test:e2e
   ```

## Writing Tests

Example test structure:

```typescript
import { test, expect } from '@playwright/test';

test.describe('Feature Name', () => {
  test.beforeEach(async ({ page }) => {
    // Setup, navigation, API mocking
    await page.goto('/path');
  });

  test('should do something', async ({ page }) => {
    // Your test assertions
  });
});
```

## Debugging

Use the Playwright UI mode for the best debugging experience:

```bash
pnpm test:e2e:ui
```

This provides:
- Visual test execution
- Time-travel debugging
- Network inspection
- Console logs
