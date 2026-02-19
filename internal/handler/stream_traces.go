package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/apigear-io/cli/pkg/stream/tracing"
)

// ListTraceFiles returns all trace files
// @Summary List trace files
// @Description Get a list of all trace files with metadata
// @Tags stream
// @Produce json
// @Success 200 {object} map[string][]tracing.TraceFileInfo
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/traces [get]
func ListTraceFiles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files, err := tracing.ListTraceFiles()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to list trace files")
			return
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{"files": files})
	}
}

// GetTraceStats returns trace storage statistics
// @Summary Get trace statistics
// @Description Get storage statistics for trace files
// @Tags stream
// @Produce json
// @Success 200 {object} tracing.TraceStats
// @Router /api/v1/stream/traces/stats [get]
func GetTraceStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats := tracing.GetTraceStats()
		writeJSON(w, http.StatusOK, stats)
	}
}

// GetTraceFileRequest represents query parameters for reading a trace file
type GetTraceFileRequest struct {
	Direction string `json:"direction"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	Limit     int    `json:"limit"`
}

// GetTraceFile reads and returns trace entries from a specific file
// @Summary Get trace file entries
// @Description Read entries from a trace file with optional filtering
// @Tags stream
// @Produce json
// @Param name path string true "Trace filename"
// @Param direction query string false "Filter by direction (SEND/RECV)"
// @Param startTime query integer false "Filter by start time (Unix ms)"
// @Param endTime query integer false "Filter by end time (Unix ms)"
// @Param limit query integer false "Maximum number of entries (default: 1000)"
// @Success 200 {object} map[string][]tracing.TraceEntry
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/traces/{name} [get]
func GetTraceFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "trace filename is required")
			return
		}

		// Parse query parameters
		direction := r.URL.Query().Get("direction")
		var startTime, endTime int64
		var limit int = 1000 // Default limit

		if st := r.URL.Query().Get("startTime"); st != "" {
			if val, err := strconv.ParseInt(st, 10, 64); err == nil {
				startTime = val
			}
		}
		if et := r.URL.Query().Get("endTime"); et != "" {
			if val, err := strconv.ParseInt(et, 10, 64); err == nil {
				endTime = val
			}
		}
		if l := r.URL.Query().Get("limit"); l != "" {
			if val, err := strconv.Atoi(l); err == nil {
				limit = val
			}
		}

		// Build filter
		filter := tracing.FilterOptions{
			Direction: direction,
			StartTime: startTime,
			EndTime:   endTime,
			Limit:     limit,
		}

		// Read entries
		entries, err := tracing.ReadTraceFileFiltered(name, filter)
		if err != nil {
			writeError(w, http.StatusNotFound, err, "failed to read trace file")
			return
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"filename": name,
			"entries":  entries,
			"count":    len(entries),
		})
	}
}

// DeleteTraceFile deletes a trace file
// @Summary Delete trace file
// @Description Delete a specific trace file
// @Tags stream
// @Produce json
// @Param name path string true "Trace filename"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/traces/{name} [delete]
func DeleteTraceFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "trace filename is required")
			return
		}

		err := tracing.DeleteTraceFile(name)
		if err != nil {
			writeError(w, http.StatusNotFound, err, "failed to delete trace file")
			return
		}

		writeJSON(w, http.StatusOK, map[string]string{
			"filename": name,
			"message":  "Trace file deleted successfully",
		})
	}
}

// SearchTracesRequest represents a request to search traces
type SearchTracesRequest struct {
	ProxyName  string `json:"proxyName"`
	Direction  string `json:"direction"`
	StartTime  int64  `json:"startTime"`
	EndTime    int64  `json:"endTime"`
	MaxFiles   int    `json:"maxFiles"`
	MaxEntries int    `json:"maxEntries"`
}

// SearchTraces searches across trace files
// @Summary Search traces
// @Description Search across multiple trace files with filters
// @Tags stream
// @Accept json
// @Produce json
// @Param request body SearchTracesRequest true "Search criteria"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/traces/search [post]
func SearchTraces() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SearchTracesRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "invalid request body")
			return
		}

		// For now, return a simple filtered read from a single file
		// Full search implementation would use tracing.Browser
		filter := tracing.FilterOptions{
			Direction: req.Direction,
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
			Limit:     req.MaxEntries,
		}

		// Get all files
		files, err := tracing.ListTraceFiles()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to list trace files")
			return
		}

		// Filter by proxy name if specified
		var filteredFiles []tracing.TraceFileInfo
		if req.ProxyName != "" {
			for _, f := range files {
				if f.ProxyName == req.ProxyName {
					filteredFiles = append(filteredFiles, f)
				}
			}
		} else {
			filteredFiles = files
		}

		// Limit files
		if req.MaxFiles > 0 && len(filteredFiles) > req.MaxFiles {
			filteredFiles = filteredFiles[:req.MaxFiles]
		}

		// Collect entries from all files
		var allEntries []tracing.TraceEntry
		for _, file := range filteredFiles {
			entries, err := tracing.ReadTraceFileFiltered(file.Name, filter)
			if err != nil {
				continue // Skip files that can't be read
			}
			allEntries = append(allEntries, entries...)

			// Check limit
			if req.MaxEntries > 0 && len(allEntries) >= req.MaxEntries {
				allEntries = allEntries[:req.MaxEntries]
				break
			}
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"entries": allEntries,
			"count":   len(allEntries),
			"files":   len(filteredFiles),
		})
	}
}
