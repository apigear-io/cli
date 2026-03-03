package handler

import (
	"encoding/json"
	"net/http"

	"github.com/apigear-io/cli/pkg/stream/logging"
)

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// writeJSON writes a JSON response with the given status code
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// writeError writes a JSON error response with the given status code
// It also logs the error to the application log
func writeError(w http.ResponseWriter, status int, err error, message string) {
	errMsg := message
	if err != nil {
		errMsg = err.Error()
	}

	// Log the error
	fields := map[string]interface{}{
		"status":  status,
		"error":   errMsg,
		"message": message,
	}

	if status >= 500 {
		logging.Error("API Error", fields)
	} else if status >= 400 {
		logging.Warn("API Error", fields)
	}

	response := ErrorResponse{
		Error:   errMsg,
		Message: message,
	}
	writeJSON(w, status, response)
}

// logOperation logs a successful operation (create, update, delete, etc.)
func logOperation(operation string, resource string, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["operation"] = operation
	fields["resource"] = resource
	logging.Info("Operation", fields)
}
