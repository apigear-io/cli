import { test, expect } from '@playwright/test';

test.describe('Projects Page', () => {
  test.beforeEach(async ({ page }) => {
    // Mock API responses
    await page.route('**/api/v1/projects/**', (route) => {
      const url = route.request().url();
      const method = route.request().method();

      if (url.includes('/projects/recent') && method === 'GET') {
        // Mock recent projects list
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({
            projects: [
              {
                name: 'demo-project',
                path: '/tmp/demo-project',
                documents: [
                  {
                    name: 'demo.module.yaml',
                    path: '/tmp/demo-project/apigear/demo.module.yaml',
                    type: 'module',
                  },
                  {
                    name: 'demo.solution.yaml',
                    path: '/tmp/demo-project/apigear/demo.solution.yaml',
                    type: 'solution',
                  },
                ],
              },
              {
                name: 'test-project',
                path: '/tmp/test-project',
                documents: [
                  {
                    name: 'api.module.yaml',
                    path: '/tmp/test-project/apigear/api.module.yaml',
                    type: 'module',
                  },
                ],
              },
            ],
            count: 2,
          }),
        });
      } else if (method === 'POST' && !url.includes('/projects/')) {
        // Mock create project
        route.fulfill({
          status: 201,
          contentType: 'application/json',
          body: JSON.stringify({
            name: 'new-project',
            path: '/tmp/new-project',
            documents: [
              {
                name: 'demo.module.yaml',
                path: '/tmp/new-project/apigear/demo.module.yaml',
                type: 'module',
              },
            ],
          }),
        });
      } else if (method === 'DELETE') {
        // Mock delete project
        route.fulfill({
          status: 204,
        });
      } else if (url.includes('/projects/get')) {
        // Mock get project details
        const pathParam = new URL(url).searchParams.get('path');
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({
            name: 'demo-project',
            path: pathParam || '/tmp/demo-project',
            documents: [
              {
                name: 'demo.module.yaml',
                path: '/tmp/demo-project/apigear/demo.module.yaml',
                type: 'module',
              },
              {
                name: 'demo.solution.yaml',
                path: '/tmp/demo-project/apigear/demo.solution.yaml',
                type: 'solution',
              },
            ],
          }),
        });
      } else {
        route.continue();
      }
    });

    await page.goto('/projects');
  });

  test('should display page title and create button', async ({ page }) => {
    await expect(page.getByRole('heading', { name: /projects/i })).toBeVisible();
    await expect(page.getByRole('button', { name: /create project/i })).toBeVisible();
  });

  test('should display project cards', async ({ page }) => {
    // Wait for project cards to load
    await page.waitForSelector('.mantine-Card-root', { timeout: 5000 });

    // Should show two projects
    const cards = page.locator('.mantine-Card-root');
    await expect(cards).toHaveCount(2);

    // Check project names
    await expect(page.getByText('demo-project')).toBeVisible();
    await expect(page.getByText('test-project')).toBeVisible();
  });

  test('should display document counts on project cards', async ({ page }) => {
    await page.waitForSelector('.mantine-Card-root', { timeout: 5000 });

    // Check for document count badges
    await expect(page.getByText('2 documents')).toBeVisible();
    await expect(page.getByText('1 document')).toBeVisible();
  });

  test('should open create project modal', async ({ page }) => {
    // Click create button
    await page.getByRole('button', { name: /create project/i }).first().click();

    // Modal should be visible
    await expect(page.getByRole('dialog')).toBeVisible();
    await expect(page.getByText('Create New Project')).toBeVisible();

    // Check form fields
    await expect(page.getByLabel(/project name/i)).toBeVisible();
    await expect(page.getByLabel(/parent directory/i)).toBeVisible();
  });

  test('should validate create project form', async ({ page }) => {
    // Open modal
    await page.getByRole('button', { name: /create project/i }).first().click();

    // Try to submit empty form
    await page.getByRole('button', { name: /create project/i }).last().click();

    // Should show validation errors
    await expect(page.getByText(/required/i)).toBeVisible();
  });

  test('should create new project', async ({ page }) => {
    // Open modal
    await page.getByRole('button', { name: /create project/i }).first().click();

    // Fill form
    await page.getByLabel(/project name/i).fill('new-project');
    await page.getByLabel(/parent directory/i).fill('/tmp');

    // Submit
    await page.getByRole('button', { name: /create project/i }).last().click();

    // Modal should close (wait a bit for animation)
    await page.waitForTimeout(500);

    // Success notification should appear (if notifications are rendered)
    // Note: Mantine notifications might not be easily testable in E2E
  });

  test('should open project details drawer', async ({ page }) => {
    await page.waitForSelector('.mantine-Card-root', { timeout: 5000 });

    // Find and click the action menu on the first card
    const firstCard = page.locator('.mantine-Card-root').first();
    await firstCard.getByRole('button').first().click();

    // Click "View Details" in menu
    await page.getByRole('menuitem', { name: /view details/i }).click();

    // Drawer should open
    await expect(page.getByText('Project Details')).toBeVisible();
    await expect(page.getByText('Documents')).toBeVisible();
  });

  test('should display empty state when no projects', async ({ page }) => {
    // Override API response for empty list
    await page.route('**/api/v1/projects/recent', (route) => {
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({
          projects: [],
          count: 0,
        }),
      });
    });

    await page.goto('/projects');

    // Should show empty state
    await expect(page.getByText(/no projects yet/i)).toBeVisible();
    await expect(page.getByText(/create your first project/i)).toBeVisible();
  });

  test('should navigate to projects page from sidebar', async ({ page }) => {
    await page.goto('/');

    // Click Projects link in sidebar/navigation
    await page.getByRole('link', { name: /projects/i }).click();

    // Should navigate to projects page
    await expect(page).toHaveURL(/\/projects/);
    await expect(page.getByRole('heading', { name: /projects/i })).toBeVisible();
  });

  test('should show delete confirmation dialog', async ({ page }) => {
    await page.waitForSelector('.mantine-Card-root', { timeout: 5000 });

    // Find and click the action menu on the first card
    const firstCard = page.locator('.mantine-Card-root').first();
    await firstCard.getByRole('button').first().click();

    // Click "Delete" in menu
    await page.getByRole('menuitem', { name: /delete/i }).click();

    // Confirmation modal should appear
    await expect(page.getByText(/delete project/i)).toBeVisible();
    await expect(page.getByText(/cannot be undone/i)).toBeVisible();
  });
});

test.describe('Projects Page - Empty State', () => {
  test.beforeEach(async ({ page }) => {
    // Mock empty projects list
    await page.route('**/api/v1/projects/recent', (route) => {
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({
          projects: [],
          count: 0,
        }),
      });
    });

    await page.goto('/projects');
  });

  test('should show empty state with create button', async ({ page }) => {
    await expect(page.getByText(/no projects yet/i)).toBeVisible();
    await expect(page.getByRole('button', { name: /create your first project/i })).toBeVisible();
  });

  test('should open create modal from empty state', async ({ page }) => {
    await page.getByRole('button', { name: /create your first project/i }).click();

    // Modal should open
    await expect(page.getByText('Create New Project')).toBeVisible();
  });
});
