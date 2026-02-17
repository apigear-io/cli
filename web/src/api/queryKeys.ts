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
} as const;
