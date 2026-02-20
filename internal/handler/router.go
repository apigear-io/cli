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

// RegisterAPIRoutes registers all REST API routes (health, status, templates, stream)
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

		// Stream endpoints
		r.Route("/stream", func(r chi.Router) {
			// Dashboard
			r.Get("/dashboard", GetStreamDashboard())

			// Proxies
			r.Route("/proxies", func(r chi.Router) {
				r.Get("/", ListStreamProxies())
				r.Post("/", CreateStreamProxy())
				r.Get("/{name}", GetStreamProxy())
				r.Put("/{name}", UpdateStreamProxy())
				r.Delete("/{name}", DeleteStreamProxy())
				r.Post("/{name}/start", StartStreamProxy())
				r.Post("/{name}/stop", StopStreamProxy())
				r.Get("/{name}/stats", GetStreamProxyStats())
				r.Get("/{name}/events", StreamProxyEvents(getStreamServices()))
			})

			// Clients
			r.Route("/clients", func(r chi.Router) {
				r.Get("/", ListStreamClients())
				r.Post("/", CreateStreamClient())
				r.Get("/{name}", GetStreamClient())
				r.Put("/{name}", UpdateStreamClient())
				r.Delete("/{name}", DeleteStreamClient())
				r.Post("/{name}/connect", ConnectStreamClient())
				r.Post("/{name}/disconnect", DisconnectStreamClient())
			})

			// Scripts
			r.Route("/scripts", func(r chi.Router) {
				r.Get("/", ListScripts())
				r.Post("/", SaveScript())
				r.Get("/running", ListRunningScripts())
				r.Post("/run", RunCode())
				r.Get("/output", StreamScriptOutput(getStreamServices()))
				r.Get("/{name}", LoadScript())
				r.Put("/{name}", UpdateScript())
				r.Delete("/{name}", DeleteScript())
				r.Post("/{name}/run", RunScript())
				r.Post("/stop/{id}", StopScript())
			})

			// Traces
			r.Route("/traces", func(r chi.Router) {
				r.Get("/", ListTraceFiles())
				r.Get("/stats", GetTraceStats())
				r.Post("/search", SearchTraces())
				r.Post("/edit", EditTrace())
				r.Post("/merge", MergeTraces())
				r.Post("/export", ExportTrace())
				r.Get("/{name}", GetTraceFile())
				r.Delete("/{name}", DeleteTraceFile())
			})

			// Stream Editor
			r.Route("/editor", func(r chi.Router) {
				r.Post("/load", LoadStreamEditor())
				r.Get("/messages", GetStreamEditorMessages())
				r.Get("/timeline", GetStreamEditorTimeline())
				r.Get("/seek", SeekStreamEditor())
				r.Post("/export", ExportStreamEditor())
				r.Post("/jq", RunStreamEditorJQ())
			})

			// Stream Player
			r.Route("/player", func(r chi.Router) {
				r.Get("/", ListPlayerStreams())
				r.Post("/", CreatePlayerStream())
				r.Get("/{id}", GetPlayerStream())
				r.Post("/{id}/play", PlayPlayerStream())
				r.Post("/{id}/pause", PausePlayerStream())
				r.Post("/{id}/resume", ResumePlayerStream())
				r.Post("/{id}/stop", StopPlayerStream())
				r.Delete("/{id}", DeletePlayerStream())
			})

			// Application Logs
			r.Get("/logs", GetLogs())
			r.Delete("/logs", ClearLogs())
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
