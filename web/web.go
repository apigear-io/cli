package web

import (
	"embed"
	"io/fs"
	"net/http"
)

// Embed the entire dist directory at compile time.
//
// IMPORTANT: The web/dist directory must exist and contain built web UI files
// before compiling the Go binary. If the directory is empty or doesn't exist,
// the embed will succeed but the UI will not be available.
//
// To build the web UI:
//   cd web && pnpm install && pnpm build
//
//go:embed dist
var distFS embed.FS

// FS returns the embedded filesystem containing the web UI static files.
// This is the dist subdirectory of the embedded filesystem.
func FS() (fs.FS, error) {
	return fs.Sub(distFS, "dist")
}

// Handler returns an http.Handler that serves the embedded web UI with SPA fallback.
// It serves static files and falls back to index.html for client-side routing.
func Handler() (http.Handler, error) {
	webFS, err := FS()
	if err != nil {
		return nil, err
	}

	return http.FileServer(http.FS(webFS)), nil
}

// Available returns true if the embedded web UI files are available.
// This checks if the dist directory was embedded at build time.
func Available() bool {
	entries, err := distFS.ReadDir("dist")
	if err != nil {
		return false
	}
	return len(entries) > 0
}
