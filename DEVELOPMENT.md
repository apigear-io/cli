# Development Guide

This guide covers setting up and running the development environment for ApiGear CLI.

## Prerequisites

### Required
- **Go 1.21+** - Backend language
- **Node.js 20+** - Frontend runtime
- **pnpm 9+** - Frontend package manager
- **Task** - Task runner (install: `brew install go-task`)

### Recommended (for best DX)
- **air** - Go live reloading (install: `go install github.com/cosmtrek/air@latest`)
- **overmind** - Process manager (install: `brew install overmind` or `brew install tmux && go install github.com/DarthSim/overmind/v2@latest`)

### Alternative Process Managers
If you don't have overmind, you can use:
- **foreman** - Ruby-based (install: `gem install foreman`)
- **hivemind** - Go-based (install: `brew install hivemind`)
- **goreman** - Go-based (install: `go install github.com/mattn/goreman@latest`)

## Quick Start

### Option 1: Automatic Setup (Recommended)

Single command to start everything with live reloading:

```bash
# Install dependencies and start dev environment
task setup:all
task dev
```

This runs:
- Backend with live reloading (air) on http://localhost:8080
- Frontend dev server (vite) on http://localhost:3000
- Auto-restart on file changes

### Option 2: Without Live Reloading

If you don't have air installed:

```bash
task dev:simple
```

### Option 3: Manual Setup

Run in separate terminals:

```bash
# Terminal 1: Backend with live reload
air -c .air.toml

# Terminal 2: Frontend dev server
cd web && pnpm dev
```

Or without air:

```bash
# Terminal 1: Backend
task run -- serve --port 8080

# Terminal 2: Frontend
cd web && pnpm dev
```

## Development Workflow

### Initial Setup

```bash
# Clone repository
git clone https://github.com/apigear-io/cli.git
cd cli

# Install all dependencies
task setup:all

# Verify everything works
task test:all
task build:all
```

### Daily Development

```bash
# Start development environment
task dev

# Access the application:
# - Backend API:  http://localhost:8080/api/v1
# - Web UI:       http://localhost:8080
# - Frontend Dev: http://localhost:3000 (with HMR)
# - Swagger:      http://localhost:8080/swagger/index.html
```

### Making Changes

**Backend Changes (Go):**
- Edit files in `cmd/`, `internal/`, `pkg/`
- air automatically rebuilds and restarts the server
- See errors in the terminal or `build-errors.log`

**Frontend Changes (React/TypeScript):**
- Edit files in `web/src/`
- Vite Hot Module Replacement (HMR) updates instantly
- See errors in browser console or terminal

### Testing

```bash
# Run all tests (backend + frontend)
task test:all

# Backend tests
task test                 # Run backend tests
task test:cover          # With coverage report
task test:ci             # CI mode (with race detector)

# Frontend unit tests (Vitest)
task web:test            # Run unit tests once
task web:test:watch      # Watch mode
task web:test:ui         # Interactive UI mode
task web:test:coverage   # With coverage report

# Frontend E2E tests (Playwright)
task web:test:e2e        # Run E2E tests
task web:test:e2e:ui     # Interactive UI mode (best for debugging)
task web:test:e2e:debug  # Debug mode

# Frontend linting and type checking
task web:type-check      # TypeScript type checking
task web:lint            # ESLint
```

**Testing Resources:**
- Unit test utilities: `web/src/test/utils.tsx`
- E2E test guide: `web/e2e/README.md`
- Query testing: `QUERY_REFACTORING.md`

### Building

```bash
# Build everything
task build:all

# Build backend only
task build

# Build frontend only
task web:build

# The backend binary embeds the frontend automatically
```

## Available Tasks

See all available commands:

```bash
task --list
```

### Most Common Commands

```bash
task dev              # Start dev environment
task dev:manual       # Show manual setup instructions
task build:all        # Build backend + frontend
task test:all         # Test everything
task lint:all         # Lint everything
task ci:all           # Run full CI pipeline
task run -- <args>    # Run CLI commands
task web:dev          # Start frontend only
```

## Project Structure

```
.
├── cmd/              # Go CLI entry points
├── internal/         # Private Go packages
│   └── handler/      # HTTP handlers (REST API)
├── pkg/              # Public Go packages
├── web/              # Frontend React application
│   ├── src/
│   │   ├── api/      # API client & React Query hooks
│   │   │   ├── client.ts      # Fetch wrapper
│   │   │   ├── queries.ts     # React Query hooks (useSuspenseQuery)
│   │   │   ├── queryKeys.ts   # Query key factory
│   │   │   └── types.ts       # TypeScript types
│   │   ├── components/        # Shared components
│   │   │   ├── ErrorBoundary.tsx
│   │   │   ├── LoadingFallback.tsx
│   │   │   └── Layout/
│   │   ├── pages/     # Page components
│   │   ├── test/      # Test utilities
│   │   │   ├── setup.ts       # Vitest setup
│   │   │   └── utils.tsx      # Custom render with providers
│   │   └── main.tsx   # App entry point
│   ├── e2e/          # Playwright E2E tests
│   ├── dist/         # Built frontend (embedded in Go binary)
│   ├── vitest.config.ts       # Vitest configuration
│   ├── playwright.config.ts   # Playwright configuration
│   └── vite.config.ts         # Vite configuration
├── Procfile          # Development process definitions
├── .air.toml         # Air configuration for live reloading
├── Taskfile.yml      # Task definitions
├── CLAUDE.md         # AI assistant context
└── QUERY_REFACTORING.md  # useSuspenseQuery migration guide
```

## Configuration Files

- **Procfile** - Process definitions for overmind/foreman
- **Procfile.dev** - Alternative without air (no live reload)
- **.air.toml** - Air configuration for Go live reloading
- **Taskfile.yml** - Task runner definitions
- **web/vite.config.ts** - Frontend build configuration
- **web/tsconfig.json** - TypeScript configuration

## Troubleshooting

### Port Already in Use

If port 8080 or 3000 is in use:

```bash
# Kill processes on port 8080
lsof -ti:8080 | xargs kill

# Or use different ports
task run -- serve --port 8081
cd web && PORT=3001 pnpm dev
```

### Air Not Found

```bash
go install github.com/cosmtrek/air@latest
```

### Overmind Not Found

```bash
# macOS
brew install overmind

# Or use alternative
brew install hivemind
hivemind Procfile
```

### Frontend Build Errors

```bash
cd web
rm -rf node_modules pnpm-lock.yaml
pnpm install
pnpm build
```

### Backend Build Errors

```bash
go clean -cache
go mod tidy
task build
```

### Live Reload Not Working

Check that you have correct file permissions and your editor isn't causing issues:

```bash
# Some editors need this for file watching
# Add to air config or use polling mode
echo 'fs.inotify.max_user_watches=524288' | sudo tee -a /etc/sysctl.conf
```

## API Development

### Adding New Endpoints

1. Create handler in `internal/handler/`
2. Add tests in `internal/handler/*_test.go`
3. Register route in `internal/handler/router.go`
4. Update Swagger docs (comments in handler)
5. Add TypeScript types in `web/src/api/types.ts`
6. Add TanStack Query hooks in `web/src/api/queries.ts`

### Testing API Endpoints

```bash
# Health check
curl http://localhost:8080/api/v1/health

# List templates
curl http://localhost:8080/api/v1/templates

# With jq for pretty output
curl -s http://localhost:8080/api/v1/templates | jq
```

## Frontend Development

### Adding New Pages

1. Create page component in `web/src/pages/NewPage/`
2. Add route in `web/src/App.tsx`
3. Add navigation link in `web/src/components/Layout/AppLayout.tsx`
4. Use TanStack Query for data fetching (prefer `useSuspenseQuery`)
5. Use Mantine UI components for consistency
6. Write tests: `NewPage.test.tsx` and `e2e/new-page.spec.ts`

### State Management

- **TanStack Query v5** - Server state (API data, prefer `useSuspenseQuery`)
- **React Hooks** - Local component state
- **URL State** - Route parameters and query strings

**Query Best Practices:**
- Use query key factory: `queryKeys.resource.operation()`
- Prefer `useSuspenseQuery` for simpler code
- Wrap components in `<Suspense>` + `<ErrorBoundary>`
- See `QUERY_REFACTORING.md` for migration guide

## CI/CD

The CI pipeline runs these checks:

```bash
task ci:all
```

Which includes:
- Backend linting (golangci-lint)
- Backend tests (with race detector)
- Frontend TypeScript type checking
- Frontend linting (ESLint)
- Frontend unit tests (Vitest)
- Full build (backend + frontend)

## Performance

### Backend Optimization

- Use `task build` (production) instead of `go run` (dev)
- Profile with `pprof`: `go tool pprof http://localhost:8080/debug/pprof/profile`

### Frontend Optimization

- Production builds are optimized automatically
- Check bundle size: `cd web && pnpm build --report`
- Analyze with: `cd web && pnpm build && npx vite-bundle-visualizer`

## Additional Documentation

- **[CLAUDE.md](./CLAUDE.md)** - Context for AI assistants
- **[ARCHITECTURE.md](./ARCHITECTURE.md)** - System architecture
- **[QUERY_REFACTORING.md](./web/QUERY_REFACTORING.md)** - useSuspenseQuery migration guide
- **[E2E Testing Guide](./web/e2e/README.md)** - Playwright E2E test setup

## Resources

- [Task Documentation](https://taskfile.dev/)
- [Air Documentation](https://github.com/cosmtrek/air)
- [Overmind Documentation](https://github.com/DarthSim/overmind)
- [Vite Documentation](https://vitejs.dev/)
- [TanStack Query v5](https://tanstack.com/query/latest)
- [Mantine UI v8](https://mantine.dev/)
- [Vitest](https://vitest.dev/)
- [Playwright](https://playwright.dev/)
