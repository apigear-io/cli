import { test, expect } from '@playwright/test';

test.describe('Stream Dashboard', () => {
  test.beforeEach(async ({ page }) => {
    // Mock API responses
    await page.route('**/api/v1/**', (route) => {
      const url = route.request().url();

      if (url.includes('/stream/dashboard')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({
            proxies: {
              total: 2,
              running: 1,
              stopped: 1,
            },
            clients: {
              total: 1,
              connected: 1,
              disconnected: 0,
            },
            messages: {
              total: 1234,
              rate: 12.5,
            },
          }),
        });
      } else if (url.includes('/health')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ status: 'ok', timestamp: new Date().toISOString() }),
        });
      } else if (url.includes('/status')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({
            version: '0.1.0',
            commit: 'test',
            buildDate: '2025-01-01',
            goVersion: '1.21',
            uptime: '1h',
          }),
        });
      } else {
        route.continue();
      }
    });

    await page.goto('/stream/dashboard');
  });

  test('should display dashboard title', async ({ page }) => {
    await expect(page.getByRole('heading', { name: /stream dashboard/i })).toBeVisible();
  });

  test('should display proxy statistics card', async ({ page }) => {
    const cards = page.locator('.mantine-Card-root');
    const proxyCard = cards.filter({ hasText: 'Proxies' }).first();
    await expect(proxyCard).toBeVisible();
    await expect(proxyCard.getByText('2')).toBeVisible();
    await expect(proxyCard.getByText(/1 running/i)).toBeVisible();
    await expect(proxyCard.getByText(/1 stopped/i)).toBeVisible();
  });

  test('should display client statistics card', async ({ page }) => {
    const cards = page.locator('.mantine-Card-root');
    const clientCard = cards.filter({ hasText: 'Clients' }).first();
    await expect(clientCard).toBeVisible();
    await expect(clientCard.getByText(/1 connected/i)).toBeVisible();
  });

  test('should display message statistics card', async ({ page }) => {
    const cards = page.locator('.mantine-Card-root');
    const messageCard = cards.filter({ hasText: 'Messages' }).first();
    await expect(messageCard).toBeVisible();
    await expect(messageCard.getByText('1,234')).toBeVisible();
    await expect(messageCard.getByText(/12\.50 msg\/s/i)).toBeVisible();
  });

  test('should have quick action buttons', async ({ page }) => {
    await expect(page.getByRole('button', { name: /manage proxies/i })).toBeVisible();
    await expect(page.getByRole('button', { name: /manage clients/i })).toBeVisible();
  });

  test('should navigate to proxies page when clicking proxy card', async ({ page }) => {
    const cards = page.locator('.mantine-Card-root');
    const proxyCard = cards.filter({ hasText: 'Proxies' }).first();
    await proxyCard.click();
    await expect(page).toHaveURL(/\/stream\/proxies/);
  });

  test('should navigate to clients page when clicking client card', async ({ page }) => {
    const cards = page.locator('.mantine-Card-root');
    const clientCard = cards.filter({ hasText: 'Clients' }).first();
    await clientCard.click();
    await expect(page).toHaveURL(/\/stream\/clients/);
  });

  test('should display info cards', async ({ page }) => {
    await expect(page.getByRole('heading', { name: /about websocket streaming/i })).toBeVisible();
    await expect(page.getByRole('heading', { name: /getting started/i })).toBeVisible();
  });
});

test.describe('Proxies Page', () => {
  test.beforeEach(async ({ page }) => {
    await page.route('**/api/v1/**', (route) => {
      const url = route.request().url();
      const method = route.request().method();

      if (url.includes('/stream/proxies') && method === 'GET') {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify([
            {
              name: 'test-proxy',
              listen: 'ws://localhost:5550/ws',
              backend: 'ws://localhost:5560/ws',
              mode: 'proxy',
              status: 'running',
              messagesReceived: 100,
              messagesSent: 95,
              activeConnections: 2,
              bytesReceived: 10240,
              bytesSent: 9728,
              uptime: 3600,
            },
            {
              name: 'echo-server',
              listen: 'ws://localhost:5551/ws',
              backend: '',
              mode: 'echo',
              status: 'stopped',
              messagesReceived: 0,
              messagesSent: 0,
              activeConnections: 0,
              bytesReceived: 0,
              bytesSent: 0,
              uptime: 0,
            },
          ]),
        });
      } else if (url.includes('/stream/proxies') && method === 'POST') {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({
            name: 'new-proxy',
            listen: 'ws://localhost:5552/ws',
            backend: 'ws://localhost:5562/ws',
            mode: 'proxy',
            status: 'stopped',
            messagesReceived: 0,
            messagesSent: 0,
            activeConnections: 0,
            bytesReceived: 0,
            bytesSent: 0,
            uptime: 0,
          }),
        });
      } else if (url.includes('/start') && method === 'POST') {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ status: 'running' }),
        });
      } else if (url.includes('/stop') && method === 'POST') {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ status: 'stopped' }),
        });
      } else if (url.includes('/health')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ status: 'ok', timestamp: new Date().toISOString() }),
        });
      } else if (url.includes('/status')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({
            version: '0.1.0',
            commit: 'test',
            buildDate: '2025-01-01',
            goVersion: '1.21',
            uptime: '1h',
          }),
        });
      } else {
        route.continue();
      }
    });

    await page.goto('/stream/proxies');
  });

  test('should display page title and proxy count', async ({ page }) => {
    await expect(page.getByRole('heading', { name: /^proxies$/i })).toBeVisible();
    await expect(page.getByText('2 total')).toBeVisible();
  });

  test('should display create proxy button', async ({ page }) => {
    await expect(page.getByRole('button', { name: /create proxy/i })).toBeVisible();
  });

  test('should display proxy cards', async ({ page }) => {
    await expect(page.getByText('test-proxy')).toBeVisible();
    await expect(page.getByText('echo-server')).toBeVisible();
  });

  test('should display proxy status badges', async ({ page }) => {
    const cards = page.locator('.mantine-Card-root');
    await expect(cards.first().getByText('running')).toBeVisible();
    await expect(cards.nth(1).getByText('stopped')).toBeVisible();
  });

  test('should display proxy mode badges', async ({ page }) => {
    const cards = page.locator('.mantine-Card-root');
    // Check for badges with mode text
    await expect(cards.first().locator('.mantine-Badge-root').filter({ hasText: 'proxy' })).toBeVisible();
    await expect(cards.nth(1).locator('.mantine-Badge-root').filter({ hasText: 'echo' })).toBeVisible();
  });

  test('should display proxy statistics', async ({ page }) => {
    const firstCard = page.locator('.mantine-Card-root').first();
    await expect(firstCard.getByText(/↓100 ↑95/)).toBeVisible();
    await expect(firstCard.getByText('2')).toBeVisible(); // active connections
  });

  test('should show start button for stopped proxy', async ({ page }) => {
    const echoCard = page.locator('.mantine-Card-root').nth(1);
    await expect(echoCard.getByRole('button', { name: /start/i })).toBeVisible();
  });

  test('should show stop button for running proxy', async ({ page }) => {
    const testProxyCard = page.locator('.mantine-Card-root').first();
    await expect(testProxyCard.getByRole('button', { name: /stop/i })).toBeVisible();
  });

  test('should open create proxy modal', async ({ page }) => {
    await page.getByRole('button', { name: /create proxy/i }).first().click();
    const modal = page.locator('.mantine-Modal-root');
    await expect(modal.getByRole('heading', { name: /create proxy/i })).toBeVisible();
    await expect(modal.getByText('Name')).toBeVisible();
    await expect(modal.getByText('Mode')).toBeVisible();
    await expect(modal.getByText('Listen Address')).toBeVisible();
  });

  test('should show backend field for proxy mode by default', async ({ page }) => {
    await page.getByRole('button', { name: /create proxy/i }).first().click();
    const modal = page.locator('.mantine-Modal-root');
    await expect(modal.getByRole('heading', { name: /create proxy/i })).toBeVisible();

    // Default mode is proxy, backend field should be visible
    await expect(modal.getByText('Backend Address')).toBeVisible();
  });

  test('should display proxy listen and backend addresses', async ({ page }) => {
    const firstCard = page.locator('.mantine-Card-root').first();
    await expect(firstCard.getByText('ws://localhost:5550/ws')).toBeVisible();
    await expect(firstCard.getByText('ws://localhost:5560/ws')).toBeVisible();
  });
});

test.describe('Clients Page', () => {
  test.beforeEach(async ({ page }) => {
    await page.route('**/api/v1/**', (route) => {
      const url = route.request().url();
      const method = route.request().method();

      if (url.includes('/stream/clients') && method === 'GET') {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify([
            {
              name: 'test-client',
              url: 'ws://localhost:5560/ws',
              interfaces: ['demo.Counter', 'demo.Calculator'],
              status: 'connected',
              autoReconnect: true,
              enabled: true,
            },
            {
              name: 'offline-client',
              url: 'ws://localhost:5561/ws',
              interfaces: ['demo.Timer'],
              status: 'disconnected',
              autoReconnect: false,
              enabled: true,
              lastError: 'Connection refused',
            },
          ]),
        });
      } else if (url.includes('/stream/clients') && method === 'POST') {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({
            name: 'new-client',
            url: 'ws://localhost:5562/ws',
            interfaces: [],
            status: 'disconnected',
            autoReconnect: true,
            enabled: true,
          }),
        });
      } else if (url.includes('/connect') && method === 'POST') {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ status: 'connected' }),
        });
      } else if (url.includes('/disconnect') && method === 'POST') {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ status: 'disconnected' }),
        });
      } else if (url.includes('/health')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ status: 'ok', timestamp: new Date().toISOString() }),
        });
      } else if (url.includes('/status')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({
            version: '0.1.0',
            commit: 'test',
            buildDate: '2025-01-01',
            goVersion: '1.21',
            uptime: '1h',
          }),
        });
      } else {
        route.continue();
      }
    });

    await page.goto('/stream/clients');
  });

  test('should display page title and client count', async ({ page }) => {
    await expect(page.getByRole('heading', { name: /^clients$/i })).toBeVisible();
    await expect(page.getByText('2 total')).toBeVisible();
  });

  test('should display create client button', async ({ page }) => {
    await expect(page.getByRole('button', { name: /create client/i })).toBeVisible();
  });

  test('should display client cards', async ({ page }) => {
    await expect(page.getByText('test-client')).toBeVisible();
    await expect(page.getByText('offline-client')).toBeVisible();
  });

  test('should display client status badges', async ({ page }) => {
    const cards = page.locator('.mantine-Card-root');
    await expect(cards.first().getByText('connected')).toBeVisible();
    await expect(cards.nth(1).getByText('disconnected')).toBeVisible();
  });

  test('should display client interfaces', async ({ page }) => {
    const firstCard = page.locator('.mantine-Card-root').first();
    await expect(firstCard.getByText('demo.Counter')).toBeVisible();
    await expect(firstCard.getByText('demo.Calculator')).toBeVisible();
  });

  test('should display client URL', async ({ page }) => {
    const firstCard = page.locator('.mantine-Card-root').first();
    await expect(firstCard.getByText('ws://localhost:5560/ws')).toBeVisible();
  });

  test('should display auto-reconnect badge', async ({ page }) => {
    const firstCard = page.locator('.mantine-Card-root').first();
    await expect(firstCard.getByText('auto-reconnect')).toBeVisible();
  });

  test('should show disconnect button for connected client', async ({ page }) => {
    const testClientCard = page.locator('.mantine-Card-root').first();
    await expect(testClientCard.getByRole('button', { name: /disconnect/i })).toBeVisible();
  });

  test('should show connect button for disconnected client', async ({ page }) => {
    const offlineCard = page.locator('.mantine-Card-root').nth(1);
    await expect(offlineCard.getByRole('button', { name: /connect/i })).toBeVisible();
  });

  test('should display error message if present', async ({ page }) => {
    const offlineCard = page.locator('.mantine-Card-root').nth(1);
    await expect(offlineCard.getByText('Connection refused')).toBeVisible();
  });

  test('should open create client modal', async ({ page }) => {
    await page.getByRole('button', { name: /create client/i }).first().click();
    const modal = page.locator('.mantine-Modal-root');
    await expect(modal.getByRole('heading', { name: /create client/i })).toBeVisible();
    await expect(modal.getByText('Name')).toBeVisible();
    await expect(modal.getByText('WebSocket URL')).toBeVisible();
    await expect(modal.getByText('ObjectLink Interfaces')).toBeVisible();
    await expect(modal.getByText('Enabled')).toBeVisible();
    await expect(modal.getByText('Auto-reconnect')).toBeVisible();
  });

  test('should have enabled and auto-reconnect switches checked by default', async ({ page }) => {
    await page.getByRole('button', { name: /create client/i }).first().click();
    const modal = page.locator('.mantine-Modal-root');
    await expect(modal.getByRole('heading', { name: /create client/i })).toBeVisible();
    // Check that switches are rendered
    await expect(modal.getByText('Enabled')).toBeVisible();
    await expect(modal.getByText('Auto-reconnect')).toBeVisible();
  });
});

test.describe('Stream Navigation', () => {
  test.beforeEach(async ({ page }) => {
    await page.route('**/api/v1/**', (route) => {
      const url = route.request().url();

      if (url.includes('/health')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ status: 'ok', timestamp: new Date().toISOString() }),
        });
      } else if (url.includes('/status')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({
            version: '0.1.0',
            commit: 'test',
            buildDate: '2025-01-01',
            goVersion: '1.21',
            uptime: '1h',
          }),
        });
      } else if (url.includes('/stream/dashboard')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({
            proxies: { total: 0, running: 0, stopped: 0 },
            clients: { total: 0, connected: 0, disconnected: 0 },
            messages: { total: 0, rate: 0 },
          }),
        });
      } else if (url.includes('/stream/proxies')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify([]),
        });
      } else if (url.includes('/stream/clients')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify([]),
        });
      } else {
        route.continue();
      }
    });
  });

  test('should have stream section in navigation', async ({ page }) => {
    await page.goto('/');
    // Check that Stream section exists in navigation
    await expect(page.locator('.mantine-NavLink-root').filter({ hasText: 'Stream' })).toBeVisible();
  });

  test('should show stream sub-items when on stream pages', async ({ page }) => {
    await page.goto('/stream/dashboard');
    // When on a stream page, sub-items should be visible
    await expect(page.locator('.mantine-NavLink-root').filter({ hasText: /^Dashboard$/ }).first()).toBeVisible();
    await expect(page.locator('.mantine-NavLink-root').filter({ hasText: /^Proxies$/ })).toBeVisible();
    await expect(page.locator('.mantine-NavLink-root').filter({ hasText: /^Clients$/ })).toBeVisible();
  });

  test('should navigate between stream pages', async ({ page }) => {
    // Start at dashboard
    await page.goto('/stream/dashboard');
    await expect(page.getByRole('heading', { name: /stream dashboard/i })).toBeVisible();

    // Navigate to proxies page (using quick action button)
    await page.getByRole('button', { name: /manage proxies/i }).click();
    await expect(page).toHaveURL(/\/stream\/proxies/);
    await expect(page.getByRole('heading', { name: /^proxies$/i })).toBeVisible();

    // Navigate to clients page
    await page.goto('/stream/clients');
    await expect(page).toHaveURL(/\/stream\/clients/);
    await expect(page.getByRole('heading', { name: /^clients$/i })).toBeVisible();
  });
});

test.describe('Empty States', () => {
  test.beforeEach(async ({ page }) => {
    await page.route('**/api/v1/**', (route) => {
      const url = route.request().url();

      if (url.includes('/stream/proxies')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify([]),
        });
      } else if (url.includes('/stream/clients')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify([]),
        });
      } else if (url.includes('/health')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ status: 'ok', timestamp: new Date().toISOString() }),
        });
      } else if (url.includes('/status')) {
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({
            version: '0.1.0',
            commit: 'test',
            buildDate: '2025-01-01',
            goVersion: '1.21',
            uptime: '1h',
          }),
        });
      } else {
        route.continue();
      }
    });
  });

  test('should display empty state for proxies', async ({ page }) => {
    await page.goto('/stream/proxies');
    // Wait for page to load
    await page.waitForLoadState('networkidle');
    const emptyCard = page.locator('.mantine-Card-root').filter({ hasText: 'No proxies configured' });
    await expect(emptyCard).toBeVisible();
    await expect(emptyCard.getByRole('button', { name: /create proxy/i })).toBeVisible();
  });

  test('should display empty state for clients', async ({ page }) => {
    await page.goto('/stream/clients');
    // Wait for page to load
    await page.waitForLoadState('networkidle');
    const emptyCard = page.locator('.mantine-Card-root').filter({ hasText: 'No clients configured' });
    await expect(emptyCard).toBeVisible();
    await expect(emptyCard.getByRole('button', { name: /create client/i })).toBeVisible();
  });
});
