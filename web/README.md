# ApiGear CLI Web UI

A modern web interface for the ApiGear CLI, built with React 19, Vite 7, Mantine v8, and TanStack Query.

## Tech Stack

- **React 19** - Latest React with improved performance
- **Vite 7** - Fast build tool and dev server
- **React Router v7** - Client-side routing
- **Mantine v8** - Modern React components library
- **TanStack Query v5** - Data fetching and caching
- **TypeScript 5.7** - Type safety
- **Tabler Icons** - Icon library

## Getting Started

### Prerequisites

- Node.js 18+ and pnpm 9+
- Go 1.21+ (for backend server)

If you don't have pnpm installed:
```bash
npm install -g pnpm
# or
brew install pnpm
```

### Installation

```bash
cd web
pnpm install
```

### Development Mode

Development mode uses Vite's dev server with hot module replacement (HMR) and proxies API requests to the Go backend.

1. Start the Go backend server:
```bash
# From repository root
go run ./cmd/apigear serve --port 8080
```

2. In a new terminal, start the Vite dev server:
```bash
cd web
pnpm dev
```

3. Open your browser to `http://localhost:3000`

The Vite dev server will proxy `/api` and `/swagger` requests to `http://localhost:8080`.

### Production Build

Build the static assets for production:

```bash
cd web
pnpm build
```

This creates optimized static files in `web/dist/` with:
- Minified JavaScript and CSS
- Code splitting for better caching
- Source maps for debugging

### Building the Go Binary with Embedded Web UI

The web UI is embedded into the Go binary at compile time. To build a binary with the web UI:

```bash
# 1. Build the web UI first
cd web
pnpm build
cd ..

# 2. Build the Go binary (web/dist is embedded automatically)
go build -o apigear ./cmd/apigear

# 3. Run the binary with embedded UI
./apigear serve
```

The web UI files from `web/dist/` are embedded into the binary using Go's `embed` package, so the resulting binary is completely standalone.

### Running Production Build

The server uses the following priority for serving the UI:

1. **Custom directory** (if `--web-dir` flag is specified)
2. **Embedded UI** (compiled into the binary)
3. **Swagger UI** (fallback if no web UI is available)

```bash
# Use embedded UI (most common)
go run ./cmd/apigear serve

# Use custom directory (for development or custom builds)
go run ./cmd/apigear serve --web-dir ./web/dist
```

The web UI will be available at `http://localhost:8080/`.
Swagger documentation remains accessible at `http://localhost:8080/swagger/`.

#### Auto-open Browser

Use the `--ui` flag to automatically open the UI in your default browser:

```bash
go run ./cmd/apigear serve --ui
```

This is useful for quickly launching the web UI without manually opening your browser.

## Project Structure

```
web/
├── src/
│   ├── main.tsx              # React entry point
│   ├── App.tsx               # Root component with routing
│   ├── theme.ts              # Mantine theme configuration
│   ├── api/
│   │   ├── client.ts         # HTTP client
│   │   ├── types.ts          # TypeScript interfaces
│   │   └── queries.ts        # TanStack Query hooks
│   ├── components/
│   │   └── Layout/
│   │       ├── AppLayout.tsx # Main layout with AppShell
│   │       └── Navigation.tsx # Sidebar navigation
│   └── pages/
│       ├── Dashboard/        # System status dashboard
│       ├── Templates/        # Template browser (coming soon)
│       ├── Projects/         # Project management (coming soon)
│       ├── CodeGen/          # Code generation UI (coming soon)
│       └── Monitor/          # Monitoring dashboard (coming soon)
├── package.json
├── vite.config.ts
├── tsconfig.json
└── index.html
```

## Available Scripts

- `pnpm dev` - Start Vite dev server with HMR
- `pnpm build` - Build for production
- `pnpm preview` - Preview production build locally
- `pnpm lint` - Lint TypeScript files
- `pnpm type-check` - Run TypeScript compiler checks

## Features

### Current (MVP)

- **Dashboard** - System status display with real-time updates
  - Version, commit, build date
  - Go version and uptime
  - Health check status
- **Responsive Layout** - Mobile-friendly sidebar navigation
- **API Integration** - TanStack Query with auto-refresh
- **SPA Routing** - Client-side routing with URL support

### Coming Soon

- **Templates** - Browse and install code generation templates
- **Projects** - Manage ApiGear projects
- **Code Generation** - Generate SDKs with drag-and-drop
- **Monitor** - Real-time API traffic monitoring

## API Integration

The web UI communicates with the Go backend REST API:

- `GET /api/v1/health` - Health check (refreshes every 30s)
- `GET /api/v1/status` - System status (refreshes every 60s)

Future endpoints will be added for templates, projects, and monitoring.

## Environment Variables

- `VITE_API_BASE_URL` - API base URL (default: `/api/v1`)

## Browser Support

- Modern browsers with ES2020+ support
- Chrome 90+
- Firefox 88+
- Safari 15+
- Edge 90+

## Troubleshooting

### API requests fail in development

Make sure the Go backend is running on port 8080:
```bash
go run ./cmd/apigear serve --port 8080
```

### Production build not showing

1. Verify the build completed: check for `web/dist/index.html`
2. Rebuild the Go binary after building the web UI (files are embedded at compile time)
3. Check server logs for "Serving embedded Web UI" message
4. If using `--web-dir`, ensure the directory path is correct

### Changes not reflecting

In development mode, Vite HMR should update automatically. If not:
1. Hard refresh the browser (Cmd+Shift+R / Ctrl+Shift+R)
2. Restart the Vite dev server (`pnpm dev`)
3. Clear browser cache

### pnpm installation issues

If you encounter issues with pnpm:
```bash
# Update pnpm to latest version
pnpm self-update

# Clear pnpm cache
pnpm store prune
```

## License

Same as ApiGear CLI (check root LICENSE file)
