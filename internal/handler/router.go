package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// RegisterAPIRoutes registers all REST API routes (health, status)
func RegisterAPIRoutes(router chi.Router) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", Health())
		r.Get("/status", Status())
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
