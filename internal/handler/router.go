package handler

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// RegisterAPIRoutes registers all REST API routes (health, status, templates)
func RegisterAPIRoutes(router chi.Router) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", Health())
		r.Get("/status", Status())

		// Template endpoints
		r.Route("/templates", func(r chi.Router) {
			r.Get("/", ListTemplates())
			r.Get("/get", GetTemplate())        // Use query param: ?id=apigear-io/template-ts
			r.Post("/install", InstallTemplate()) // Use query param: ?id=apigear-io/template-ts
			r.Get("/search", SearchTemplates())

			r.Route("/cache", func(r chi.Router) {
				r.Get("/", ListCachedTemplates())
				r.Delete("/remove", RemoveTemplate()) // Use query param: ?id=apigear-io/template-ts
				r.Post("/clean", CleanCache())
			})

			r.Route("/registry", func(r chi.Router) {
				r.Post("/update", UpdateRegistry())
			})
		})
	})
}

// RegisterSwaggerRoutes registers Swagger documentation routes
func RegisterSwaggerRoutes(router chi.Router) {
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	// Root redirect to Swagger UI
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})
}

// RegisterMonitorRoutes registers monitoring endpoint routes
func RegisterMonitorRoutes(router chi.Router) {
	router.Post("/monitor/{source}", Monitor())
}

// RegisterWebUIRoutes registers Web UI static file serving with SPA fallback
func RegisterWebUIRoutes(router chi.Router, staticDir string) {
	// Serve static files
	fileServer := http.FileServer(http.Dir(staticDir))

	// Handler that implements SPA fallback
	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		// Skip API and Swagger routes - they're handled separately
		if strings.HasPrefix(r.URL.Path, "/api/") || strings.HasPrefix(r.URL.Path, "/swagger/") {
			http.NotFound(w, r)
			return
		}

		// Try to serve the requested file
		path := filepath.Join(staticDir, r.URL.Path)

		// Check if file exists
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// File doesn't exist, serve index.html for SPA routing
			http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
			return
		}

		// File exists, serve it
		fileServer.ServeHTTP(w, r)
	})

	// Serve Swagger docs at /swagger/
	router.Get("/swagger/*", httpSwagger.WrapHandler)
}

// RegisterEmbeddedWebUIRoutes registers embedded Web UI static file serving with SPA fallback
func RegisterEmbeddedWebUIRoutes(router chi.Router, webFS fs.FS) {
	// Create file server from embedded filesystem
	fileServer := http.FileServer(http.FS(webFS))

	// Handler that implements SPA fallback for embedded files
	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		// Skip API and Swagger routes - they're handled separately
		if strings.HasPrefix(r.URL.Path, "/api/") || strings.HasPrefix(r.URL.Path, "/swagger/") {
			http.NotFound(w, r)
			return
		}

		// Try to open the requested file from embedded FS
		path := strings.TrimPrefix(r.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}

		_, err := webFS.Open(path)
		if err != nil {
			// File doesn't exist, serve index.html for SPA routing
			r.URL.Path = "/"
			fileServer.ServeHTTP(w, r)
			return
		}

		// File exists, serve it
		fileServer.ServeHTTP(w, r)
	})

	// Serve Swagger docs at /swagger/
	router.Get("/swagger/*", httpSwagger.WrapHandler)
}
