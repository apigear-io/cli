package handler

import (
	"net/http"

	"github.com/apigear-io/cli/pkg/stream/logging"
)

// GetLogs returns application log entries
// @Summary Get application logs
// @Description Get application log entries with optional filtering
// @Tags stream
// @Produce json
// @Param level query string false "Filter by log level (DEBUG, INFO, WARN, ERROR)"
// @Param search query string false "Search term for message or fields"
// @Success 200 {object} map[string][]logging.LogEntry
// @Router /api/v1/stream/logs [get]
func GetLogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		level := logging.LogLevel(r.URL.Query().Get("level"))
		search := r.URL.Query().Get("search")

		logger := logging.GetGlobalLogger()
		entries := logger.GetEntriesFiltered(level, search)

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"entries": entries,
			"count":   len(entries),
		})
	}
}

// ClearLogs clears all log entries
// @Summary Clear application logs
// @Description Clear all application log entries
// @Tags stream
// @Success 204
// @Router /api/v1/stream/logs [delete]
func ClearLogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logging.GetGlobalLogger()
		logger.Clear()
		w.WriteHeader(http.StatusNoContent)
	}
}
