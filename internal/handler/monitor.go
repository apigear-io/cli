package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/apigear-io/cli/pkg/runtime/monitoring"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// MonitorResponse represents the confirmation response for processed events
type MonitorResponse struct {
	Status          string    `json:"status"`
	EventsProcessed int       `json:"eventsProcessed"`
	Timestamp       time.Time `json:"timestamp"`
}

// Monitor godoc
// @Summary Monitor events endpoint
// @Description Receives monitoring events from client applications
// @Tags monitoring
// @Accept json
// @Produce json
// @Param source path string true "Event source identifier"
// @Param events body []monitoring.Event true "Array of monitoring events"
// @Success 200 {object} MonitorResponse
// @Failure 400 {object} ErrorResponse
// @Router /monitor/{source} [post]
func Monitor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		source := chi.URLParam(r, "source")
		if source == "" {
			writeError(w, http.StatusBadRequest,
				fmt.Errorf("source id is required"),
				"source parameter must not be empty")
			return
		}

		var events []*monitoring.Event
		err := json.NewDecoder(r.Body).Decode(&events)
		if err != nil {
			writeError(w, http.StatusBadRequest, err,
				"failed to decode event array")
			return
		}

		for _, event := range events {
			event.Source = source
			if event.Id == "" {
				event.Id = uuid.New().String()
			}
			if event.Timestamp.IsZero() {
				event.Timestamp = time.Now()
			}
			monitoring.Emitter.FireHook(event)
		}

		writeJSON(w, http.StatusOK, MonitorResponse{
			Status:          "ok",
			EventsProcessed: len(events),
			Timestamp:       time.Now(),
		})
	}
}
