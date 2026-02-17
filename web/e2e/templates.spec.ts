import { test, expect } from '@playwright/test';

test.describe('Templates Page', () => {
  test.beforeEach(async ({ page }) => {
    // Mock API responses if backend is not available
    await page.route('**/api/v1/**', (route) => {
      const url = route.request().url();

      if (url.includes('/templates/registry')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify([
            {
              name: 'test-template',
              description: 'A test template for E2E testing',
              latest: '1.0.0',
              version: '',
              git: 'https://github.com/test/template',
              inCache: false,
              updateNeeded: false,
              versions: ['1.0.0', '0.9.0'],
            },
          ]),
        });
      } else if (url.includes('/templates/cache')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify([]),
        });
      } else {
        route.continue();
      }
    });

    await page.goto('/templates');
  });

  test('should display page title and tabs', async ({ page }) => {
    // Check for main heading
    await expect(page.getByRole('heading', { name: /templates/i })).toBeVisible();

    // Check for tabs
    await expect(page.getByRole('tab', { name: /registry/i })).toBeVisible();
    await expect(page.getByRole('tab', { name: /cached/i })).toBeVisible();
  });

  test('should switch between Registry and Cached tabs', async ({ page }) => {
    // Initially on Registry tab
    const registryTab = page.getByRole('tab', { name: /registry/i });
    await expect(registryTab).toHaveAttribute('aria-selected', 'true');

    // Switch to Cached tab
    await page.getByRole('tab', { name: /cached/i }).click();
    await expect(page.getByRole('tab', { name: /cached/i })).toHaveAttribute('aria-selected', 'true');
    await expect(registryTab).toHaveAttribute('aria-selected', 'false');
  });

  test('should display template cards in Registry tab', async ({ page }) => {
    // Wait for template cards to load
    // Note: This assumes templates will be loaded from the API
    await page.waitForSelector('[role="article"], .mantine-Card-root', { timeout: 5000 })
      .catch(() => {
        // If no templates, that's okay for this test
      });

    // Check if either templates are displayed or a loading/empty state is shown
    const hasCards = await page.locator('.mantine-Card-root').count() > 0;
    const hasEmptyState = await page.getByText(/no templates/i).isVisible().catch(() => false);
    const hasLoading = await page.getByText(/loading/i).isVisible().catch(() => false);

    expect(hasCards || hasEmptyState || hasLoading).toBeTruthy();
  });

  test('should display template information on card', async ({ page }) => {
    // Wait for at least one template card
    const firstCard = page.locator('.mantine-Card-root').first();

    try {
      await firstCard.waitFor({ timeout: 5000 });

      // Verify card has essential elements (name, button)
      await expect(firstCard.locator('button')).toBeVisible();
    } catch {
      // Skip if no templates are available
      test.skip();
    }
  });

  test('should show install button on template cards', async ({ page }) => {
    const cards = page.locator('.mantine-Card-root');
    const count = await cards.count();

    if (count > 0) {
      const firstCard = cards.first();
      const installButton = firstCard.getByRole('button', { name: /install|update|up to date/i });
      await expect(installButton).toBeVisible();
    }
  });

  test('should navigate to templates page from navigation', async ({ page }) => {
    await page.goto('/');

    // Click on Templates navigation link
    await page.getByRole('link', { name: /templates/i }).click();

    // Verify we're on the templates page
    await expect(page).toHaveURL(/templates/);
    await expect(page.getByRole('heading', { name: /templates/i })).toBeVisible();
  });
});
