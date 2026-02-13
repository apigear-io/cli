package handler

import (
	"net/http"
	"time"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

// Health godoc
// @Summary Health check endpoint
// @Description Returns the health status of the API
// @Tags system
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, HealthResponse{
			Status:    "ok",
			Timestamp: time.Now(),
		})
	}
}
