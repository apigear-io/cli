# Claude AI Context - ApiGear CLI

This file provides context for AI assistants (particularly Claude) working on this project. It complements the other documentation files and focuses on recent architectural decisions, patterns, and conventions.

## Project Overview

ApiGear CLI is a command-line tool and web UI for managing API templates, code generation, and development workflows. It consists of:

- **Backend**: Go 1.21+ REST API server
- **Frontend**: React 19 + TypeScript + Vite web application
- **Templates**: Code generation templates for various frameworks

## Recent Major Changes (Feb 2025)

### 1. Testing Infrastructure Setup

**Unit Testing (Vitest)**
- Configured Vitest with jsdom for component testing
- Created test utilities with proper provider wrappers
- Location: `web/src/test/`
- Run: `task web:test` or `pnpm test`

**E2E Testing (Playwright)**
- Configured Playwright for cross-browser testing
- Includes API mocking for tests without backend
- Location: `web/e2e/`
- Run: `task web:test:e2e` or `task web:test:e2e:ui`

**Test Scripts:**
```bash
task web:test              # Unit tests
task web:test:watch        # Unit tests (watch mode)
task web:test:ui           # Unit tests (UI mode)
task web:test:coverage     # Unit tests with coverage
task web:test:e2e          # E2E tests
task web:test:e2e:ui       # E2E tests with UI
task web:test:all          # All frontend tests
```

### 2. React Query Migration to useSuspenseQuery

**Query Key Factory**
- Location: `web/src/api/queryKeys.ts`
- Provides type-safe, hierarchical query keys
- Example: `queryKeys.templates.cache()` or `queryKeys.templates.detail(id)`
- Benefits: Easy invalidation, prevents typos, better refactoring

**useSuspenseQuery Pattern**
- Migrated from `useQuery` to `useSuspenseQuery` (TanStack Query v5)
- Data is guaranteed to exist - no optional chaining needed
- Loading states handled by `<Suspense>` boundaries
- Error states handled by `<ErrorBoundary>` components

**Component Structure:**
```typescript
// Inner component - uses data directly
function PageContent() {
  const { data } = useSuspenseQuery({...});
  // data is guaranteed to exist!
  return <div>{data.items.map(...)}</div>;
}

// Outer component - provides boundaries
export function Page() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingFallback />}>
        <PageContent />
      </Suspense>
    </ErrorBoundary>
  );
}
```

**Current Status:**
- ✅ Templates page migrated
- 🔲 Dashboard, Projects, CodeGen, Monitor pages - still using useQuery

## Architecture & Tech Stack

### Backend (Go)
- **Framework**: net/http with custom router
- **Structure**:
  - `cmd/apigear/` - CLI entry point
  - `internal/handler/` - HTTP handlers (private)
  - `pkg/` - Public packages
- **API**: RESTful API at `/api/v1/*`
- **Testing**: Standard Go testing, `task test`

### Frontend (React)
- **React 19** with TypeScript
- **Vite** - Build tool (dev server on port 5173)
- **Routing**: React Router v7
- **UI Library**: Mantine v8
- **State Management**:
  - TanStack Query v5 for server state (prefer useSuspenseQuery)
  - React hooks for local state
  - URL state for navigation

### Key Dependencies
- **@tanstack/react-query** v5 - Server state management
- **@mantine/core** v8 - UI components
- **@mantine/notifications** - Toast notifications
- **@mantine/modals** - Modal dialogs
- **react-router-dom** v7 - Routing
- **@tabler/icons-react** - Icons

## Project Structure

```
.
├── cmd/apigear/           # CLI entry point
├── internal/              # Private Go packages
│   └── handler/           # API handlers
├── pkg/                   # Public Go packages
├── web/                   # Frontend application
│   ├── src/
│   │   ├── api/           # API client & React Query hooks
│   │   │   ├── client.ts         # Fetch wrapper
│   │   │   ├── queries.ts        # React Query hooks
│   │   │   ├── queryKeys.ts      # Query key factory
│   │   │   └── types.ts          # TypeScript types
│   │   ├── components/
│   │   │   ├── ErrorBoundary.tsx # Error boundary component
│   │   │   ├── LoadingFallback.tsx # Loading component
│   │   │   └── Layout/           # Layout components
│   │   ├── pages/         # Page components
│   │   │   ├── Templates/        # Template management (uses Suspense)
│   │   │   ├── Dashboard/        # Dashboard page
│   │   │   ├── Projects/         # Projects page
│   │   │   ├── CodeGen/          # Code generation
│   │   │   └── Monitor/          # Monitoring
│   │   ├── test/          # Test utilities
│   │   │   ├── setup.ts          # Global test setup
│   │   │   └── utils.tsx         # Custom render with providers
│   │   └── main.tsx       # App entry point
│   ├── e2e/               # Playwright E2E tests
│   ├── vitest.config.ts   # Vitest configuration
│   ├── playwright.config.ts # Playwright configuration
│   └── vite.config.ts     # Vite configuration
├── Taskfile.yml           # Task runner definitions
├── DEVELOPMENT.md         # Development setup guide
├── ARCHITECTURE.md        # Architecture documentation
└── QUERY_REFACTORING.md   # useSuspenseQuery migration guide
```

## Coding Conventions

### Frontend Code Style

**Prefer TypeScript features:**
- Use interfaces for props
- Avoid `any` - use proper types
- Use const assertions for query keys: `as const`

**React Patterns:**
- Function components with hooks
- Extract complex logic to custom hooks
- Use Suspense boundaries for async data
- Use Error Boundaries for error handling
- Prefer composition over prop drilling

**Component Organization:**
```typescript
// 1. Imports
import { useState } from 'react';
import { Stack, Title } from '@mantine/core';
import { useSomeQuery } from '@/api/queries';

// 2. Types/Interfaces
interface MyComponentProps {
  id: string;
}

// 3. Component
export function MyComponent({ id }: MyComponentProps) {
  // Hooks first
  const { data } = useSomeQuery();
  const [state, setState] = useState();

  // Event handlers
  const handleClick = () => {...};

  // Render
  return <div>...</div>;
}
```

**Query Hooks (TanStack Query):**
- Use `useSuspenseQuery` for new code
- Use query key factory: `queryKeys.resource.operation(params)`
- Invalidate at the parent level: `queryKeys.templates.all()`
- Mutations should invalidate related queries

**Testing:**
- Test file next to component: `Component.test.tsx`
- Use `render` from `@/test/utils` (includes providers)
- Mock API calls in tests
- Focus on user behavior, not implementation

### Backend Code Style

**Go Conventions:**
- Follow standard Go style (gofmt, golangci-lint)
- Use meaningful package names
- Keep handlers thin - business logic in services
- Write tests alongside code: `*_test.go`

**API Design:**
- RESTful endpoints under `/api/v1/`
- JSON request/response
- Proper HTTP status codes
- Swagger documentation in handler comments

## Common Tasks

### Adding a New API Endpoint

1. Create handler in `internal/handler/`
2. Add route in router
3. Write tests in `*_test.go`
4. Add Swagger comments
5. Add TypeScript types in `web/src/api/types.ts`
6. Add query key in `web/src/api/queryKeys.ts`
7. Create React Query hook in `web/src/api/queries.ts`
8. Use in component with Suspense

### Adding a New Frontend Page

1. Create page component in `web/src/pages/NewPage/NewPage.tsx`
2. Create inner content component that uses queries
3. Wrap in `<ErrorBoundary>` + `<Suspense>`
4. Add route in `web/src/App.tsx`
5. Add navigation link in layout
6. Write tests in `NewPage.test.tsx`
7. Write E2E test in `e2e/new-page.spec.ts`

### Migrating a Page to useSuspenseQuery

See `QUERY_REFACTORING.md` for detailed guide. Quick steps:

1. Import `useSuspenseQuery` instead of `useQuery`
2. Update query keys to use factory: `queryKeys.resource.operation()`
3. Remove optional chaining: `data.field` instead of `data?.field`
4. Remove manual loading/error handling
5. Split into content component + wrapper with Suspense
6. Update tests if needed

## Important Notes

### Port Configuration
- **Backend**: 8080
- **Frontend Dev**: 5173 (Vite default)
- **Frontend Prod**: Served by backend at 8080

### Dev Server Proxy
The frontend dev server proxies `/api` and `/swagger` requests to `http://localhost:8080`.

### Query Key Invalidation
When mutating data, invalidate at the parent level:
```typescript
// Good - invalidates all template queries
queryClient.invalidateQueries({ queryKey: queryKeys.templates.all() });

// Bad - only invalidates registry
queryClient.invalidateQueries({ queryKey: queryKeys.templates.registry() });
```

### Testing Best Practices
- Unit tests should be fast and isolated
- E2E tests include API mocking by default
- Use `task web:test:ui` for debugging tests
- Mock external dependencies

### Error Handling
- Frontend errors caught by ErrorBoundary
- API errors shown via notifications
- Suspense handles loading states
- Retry logic in Error Boundaries

## Task Runner Commands

Most common commands:

```bash
# Development
task dev                   # Start dev environment
task web:dev              # Frontend only

# Testing
task test:all             # All tests (backend + frontend)
task web:test             # Frontend unit tests
task web:test:e2e         # Frontend E2E tests
task test                 # Backend tests

# Building
task build:all            # Build everything
task web:build            # Frontend only

# Linting
task lint:all             # Lint everything
task web:lint             # Frontend only
task web:type-check       # TypeScript

# CI
task ci:all               # Full CI pipeline
```

## Resources

### Documentation
- [DEVELOPMENT.md](./DEVELOPMENT.md) - Setup and daily workflows
- [ARCHITECTURE.md](./ARCHITECTURE.md) - System architecture
- [QUERY_REFACTORING.md](./web/QUERY_REFACTORING.md) - useSuspenseQuery guide
- [E2E Testing Guide](./web/e2e/README.md) - Playwright setup

### External Resources
- [TanStack Query v5 Docs](https://tanstack.com/query/latest)
- [Mantine UI Components](https://mantine.dev/)
- [React Router v7](https://reactrouter.com/)
- [Vitest](https://vitest.dev/)
- [Playwright](https://playwright.dev/)

## Future Improvements

### Potential Migrations
- [ ] Migrate remaining pages to useSuspenseQuery
- [ ] Add more E2E test coverage
- [ ] Implement global error tracking
- [ ] Add performance monitoring
- [ ] Consider React Server Components (when stable)

### Testing Enhancements
- [ ] Visual regression testing
- [ ] API contract testing
- [ ] Performance testing
- [ ] Accessibility testing

## Tips for AI Assistants

1. **Always check existing patterns** before creating new ones
2. **Use the query key factory** for all new queries
3. **Prefer useSuspenseQuery** for new components
4. **Write tests** for new features
5. **Follow the established file structure**
6. **Update this file** when making architectural changes
7. **Check DEVELOPMENT.md** for setup commands
8. **Run `task test:all`** before committing

## Questions?

Check the documentation files:
- Setup issues → DEVELOPMENT.md
- Architecture questions → ARCHITECTURE.md
- Query patterns → QUERY_REFACTORING.md
- Testing → web/e2e/README.md or vitest.config.ts
