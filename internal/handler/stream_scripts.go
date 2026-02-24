package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/apigear-io/cli/pkg/stream/scripting"
)

// ListScripts returns all saved scripts
// @Summary List all saved scripts
// @Description Get a list of all saved scripts
// @Tags stream
// @Produce json
// @Success 200 {object} map[string][]string
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/scripts [get]
func ListScripts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		services := getStreamServices()
		names, err := services.ScriptManager.ListScripts()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to list scripts")
			return
		}
		writeJSON(w, http.StatusOK, map[string][]string{"scripts": names})
	}
}

// LoadScript loads a specific script by name
// @Summary Load a script
// @Description Load a specific script with its code and modification time
// @Tags stream
// @Produce json
// @Param name path string true "Script name"
// @Success 200 {object} scripting.ScriptFile
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/scripts/{name} [get]
func LoadScript() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "script name is required")
			return
		}

		services := getStreamServices()
		script, err := services.ScriptManager.LoadScriptWithInfo(name)
		if err != nil {
			writeError(w, http.StatusNotFound, err, "script not found")
			return
		}

		writeJSON(w, http.StatusOK, script)
	}
}

// SaveScriptRequest represents the request to save a script
type SaveScriptRequest struct {
	Name            string `json:"name"`
	Code            string `json:"code"`
	ExpectedModTime int64  `json:"expectedModTime,omitempty"`
}

// SaveScriptResponse represents the response from saving a script
type SaveScriptResponse struct {
	Name    string `json:"name"`
	ModTime int64  `json:"modTime"`
	Message string `json:"message"`
}

// SaveScript creates or updates a script
// @Summary Save a script
// @Description Save a script with conflict detection
// @Tags stream
// @Accept json
// @Produce json
// @Param request body SaveScriptRequest true "Script data"
// @Success 200 {object} SaveScriptResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/scripts [post]
func SaveScript() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SaveScriptRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "invalid request body")
			return
		}

		if req.Name == "" {
			writeError(w, http.StatusBadRequest, nil, "script name is required")
			return
		}

		if req.Code == "" {
			writeError(w, http.StatusBadRequest, nil, "script code is required")
			return
		}

		services := getStreamServices()
		modTime, err := services.ScriptManager.SaveScriptWithCheck(req.Name, req.Code, req.ExpectedModTime)
		if err != nil {
			if errors.Is(err, scripting.ErrConflict) {
				writeError(w, http.StatusConflict, err, "script was modified by another user")
				return
			}
			writeError(w, http.StatusInternalServerError, err, "failed to save script")
			return
		}

		logOperation("SAVE", "script", map[string]interface{}{
			"name": req.Name,
		})
		writeJSON(w, http.StatusOK, SaveScriptResponse{
			Name:    req.Name,
			ModTime: modTime,
			Message: "Script saved successfully",
		})
	}
}

// UpdateScript updates an existing script (same as SaveScript)
// @Summary Update a script
// @Description Update an existing script with conflict detection
// @Tags stream
// @Accept json
// @Produce json
// @Param name path string true "Script name"
// @Param request body SaveScriptRequest true "Script data"
// @Success 200 {object} SaveScriptResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/scripts/{name} [put]
func UpdateScript() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "script name is required")
			return
		}

		var req SaveScriptRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "invalid request body")
			return
		}

		// Use name from URL
		req.Name = name

		if req.Code == "" {
			writeError(w, http.StatusBadRequest, nil, "script code is required")
			return
		}

		services := getStreamServices()
		modTime, err := services.ScriptManager.SaveScriptWithCheck(req.Name, req.Code, req.ExpectedModTime)
		if err != nil {
			if errors.Is(err, scripting.ErrConflict) {
				writeError(w, http.StatusConflict, err, "script was modified by another user")
				return
			}
			writeError(w, http.StatusInternalServerError, err, "failed to update script")
			return
		}

		writeJSON(w, http.StatusOK, SaveScriptResponse{
			Name:    req.Name,
			ModTime: modTime,
			Message: "Script updated successfully",
		})
	}
}

// DeleteScript deletes a script by name
// @Summary Delete a script
// @Description Delete a saved script
// @Tags stream
// @Produce json
// @Param name path string true "Script name"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/scripts/{name} [delete]
func DeleteScript() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "script name is required")
			return
		}

		services := getStreamServices()
		err := services.ScriptManager.DeleteScript(name)
		if err != nil {
			writeError(w, http.StatusNotFound, err, "script not found")
			return
		}

		logOperation("DELETE", "script", map[string]interface{}{
			"name": name,
		})
		writeJSON(w, http.StatusOK, map[string]string{
			"name":    name,
			"message": "Script deleted successfully",
		})
	}
}

// ListRunningScripts returns all currently running scripts
// @Summary List running scripts
// @Description Get a list of all currently running scripts
// @Tags stream
// @Produce json
// @Success 200 {object} map[string][]scripting.ScriptInfo
// @Router /api/v1/stream/scripts/running [get]
func ListRunningScripts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		services := getStreamServices()
		scripts := services.ScriptManager.GetRunningScripts()
		writeJSON(w, http.StatusOK, map[string][]scripting.ScriptInfo{"scripts": scripts})
	}
}

// RunScriptResponse represents the response from running a script
type RunScriptResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

// RunScript runs a saved script by name
// @Summary Run a saved script
// @Description Run a saved script by name
// @Tags stream
// @Produce json
// @Param name path string true "Script name"
// @Success 200 {object} RunScriptResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/scripts/{name}/run [post]
func RunScript() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "script name is required")
			return
		}

		services := getStreamServices()

		// Load the script
		code, err := services.ScriptManager.LoadScript(name)
		if err != nil {
			writeError(w, http.StatusNotFound, err, "script not found")
			return
		}

		// Run the script
		id, err := services.ScriptManager.RunScript(name, code)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to run script")
			return
		}

		logOperation("RUN", "script", map[string]interface{}{
			"name": name,
			"id":   id,
		})
		writeJSON(w, http.StatusOK, RunScriptResponse{
			ID:      id,
			Name:    name,
			Message: "Script started successfully",
		})
	}
}

// RunCodeRequest represents a request to run ad-hoc code
type RunCodeRequest struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code"`
}

// RunCode runs ad-hoc JavaScript code without saving
// @Summary Run ad-hoc code
// @Description Run JavaScript code without saving it
// @Tags stream
// @Accept json
// @Produce json
// @Param request body RunCodeRequest true "Code to run"
// @Success 200 {object} RunScriptResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/scripts/run [post]
func RunCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RunCodeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "invalid request body")
			return
		}

		if req.Code == "" {
			writeError(w, http.StatusBadRequest, nil, "code is required")
			return
		}

		name := req.Name
		if name == "" {
			name = "ad-hoc"
		}

		services := getStreamServices()
		id, err := services.ScriptManager.RunScript(name, req.Code)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to run code")
			return
		}

		writeJSON(w, http.StatusOK, RunScriptResponse{
			ID:      id,
			Name:    name,
			Message: "Code started successfully",
		})
	}
}

// StopScript stops a running script by ID
// @Summary Stop a running script
// @Description Stop a running script by its ID
// @Tags stream
// @Produce json
// @Param id path string true "Script ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/stream/scripts/stop/{id} [post]
func StopScript() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			writeError(w, http.StatusBadRequest, nil, "script ID is required")
			return
		}

		services := getStreamServices()
		err := services.ScriptManager.StopScript(id)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to stop script")
			return
		}

		logOperation("STOP", "script", map[string]interface{}{
			"id": id,
		})
		writeJSON(w, http.StatusOK, map[string]string{
			"id":      id,
			"message": "Script stopped successfully",
		})
	}
}
