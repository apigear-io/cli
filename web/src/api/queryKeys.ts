/**
 * Query key factory for consistent and type-safe query keys.
 * Follow the pattern: [resource, operation, ...params]
 *
 * Benefits:
 * - Type-safe query keys
 * - Easier to invalidate related queries
 * - Prevents typos and inconsistencies
 * - Self-documenting query structure
 */

export const queryKeys = {
  // Health & Status
  health: () => ['health'] as const,
  status: () => ['status'] as const,

  // Templates
  templates: {
    all: () => ['templates'] as const,

    // Registry templates
    registry: () => [...queryKeys.templates.all(), 'registry'] as const,

    // Cached/installed templates
    cache: () => [...queryKeys.templates.all(), 'cache'] as const,

    // Single template detail
    detail: (id: string) => [...queryKeys.templates.all(), 'detail', id] as const,

    // Search
    search: (query: string) => [...queryKeys.templates.all(), 'search', query] as const,
  },

  // Stream
  stream: {
    all: () => ['stream'] as const,

    // Dashboard
    dashboard: () => [...queryKeys.stream.all(), 'dashboard'] as const,

    // Proxies
    proxies: {
      all: () => [...queryKeys.stream.all(), 'proxies'] as const,
      list: () => [...queryKeys.stream.proxies.all(), 'list'] as const,
      detail: (name: string) => [...queryKeys.stream.proxies.all(), 'detail', name] as const,
      stats: (name: string) => [...queryKeys.stream.proxies.all(), 'stats', name] as const,
    },

    // Clients
    clients: {
      all: () => [...queryKeys.stream.all(), 'clients'] as const,
      list: () => [...queryKeys.stream.clients.all(), 'list'] as const,
      detail: (name: string) => [...queryKeys.stream.clients.all(), 'detail', name] as const,
    },

    // Scripts
    scripts: {
      all: () => [...queryKeys.stream.all(), 'scripts'] as const,
      list: () => [...queryKeys.stream.scripts.all(), 'list'] as const,
      detail: (name: string) => [...queryKeys.stream.scripts.all(), 'detail', name] as const,
      running: () => [...queryKeys.stream.scripts.all(), 'running'] as const,
    },

    // Traces
    traces: {
      all: () => [...queryKeys.stream.all(), 'traces'] as const,
      list: () => [...queryKeys.stream.traces.all(), 'list'] as const,
      detail: (name: string) => [...queryKeys.stream.traces.all(), 'detail', name] as const,
      stats: () => [...queryKeys.stream.traces.all(), 'stats'] as const,
    },

    // Player
    player: {
      all: () => [...queryKeys.stream.all(), 'player'] as const,
      list: () => [...queryKeys.stream.player.all(), 'list'] as const,
      detail: (id: string) => [...queryKeys.stream.player.all(), 'detail', id] as const,
    },

    // Logs
    logs: {
      all: () => [...queryKeys.stream.all(), 'logs'] as const,
      list: (level?: string, search?: string) =>
        [...queryKeys.stream.logs.all(), 'list', level || '', search || ''] as const,
    },
  },
} as const;
