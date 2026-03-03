package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/foundation"
	"github.com/apigear-io/cli/pkg/foundation/config"
	"github.com/apigear-io/cli/pkg/foundation/tasks"
	"github.com/apigear-io/cli/pkg/objmodel/spec"
	"github.com/apigear-io/cli/pkg/orchestration/project"
	"github.com/apigear-io/cli/pkg/orchestration/solution"
	"github.com/apigear-io/cli/pkg/stream/logging"
)

// ProjectListResponse represents the list of projects
type ProjectListResponse struct {
	Projects []*project.ProjectInfo `json:"projects"`
	Count    int                    `json:"count"`
}

// CreateProjectRequest represents the create project request body
type CreateProjectRequest struct {
	Name string `json:"name"` // Project name (directory name)
	Path string `json:"path"` // Parent directory path
}

// ProjectDirectoriesResponse represents suggested project directories
type ProjectDirectoriesResponse struct {
	HomeDir     string   `json:"homeDir"`
	WorkingDir  string   `json:"workingDir"`
	Suggestions []string `json:"suggestions"`
}

// DirectoryEntry represents a single directory entry
type DirectoryEntry struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	Accessible bool   `json:"accessible"`
}

// DirectoryListResponse represents a list of directories
type DirectoryListResponse struct {
	CurrentPath string           `json:"currentPath"`
	ParentPath  string           `json:"parentPath"`
	Directories []DirectoryEntry `json:"directories"`
	Count       int              `json:"count"`
}

// ReadFileRequest represents a request to read a file
type ReadFileRequest struct {
	Path string `json:"path"`
}

// ReadFileResponse represents file contents
type ReadFileResponse struct {
	Path     string `json:"path"`
	Content  string `json:"content"`
	Encoding string `json:"encoding"` // "utf-8"
}

// WriteFileRequest represents a request to write a file
type WriteFileRequest struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

// OpenExternalRequest represents a request to open file externally
type OpenExternalRequest struct {
	Path string `json:"path"`
}

// GenerateCodeRequest represents a request to generate code
type GenerateCodeRequest struct {
	SolutionPath string `json:"solutionPath"`
	Force        bool   `json:"force"`
}

// GetProjectDirectories godoc
// @Summary Get suggested project directories
// @Description Returns home directory and other common project locations
// @Tags projects
// @Produce json
// @Success 200 {object} ProjectDirectoriesResponse
// @Router /api/v1/projects/directories [get]
func GetProjectDirectories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			homeDir = ""
		}

		// Get current working directory
		workingDir, err := os.Getwd()
		if err != nil {
			workingDir = ""
		}

		// Build suggestions list
		suggestions := []string{}
		if homeDir != "" {
			suggestions = append(suggestions, homeDir)
			// Common project directories
			suggestions = append(suggestions, foundation.Join(homeDir, "Projects"))
			suggestions = append(suggestions, foundation.Join(homeDir, "projects"))
			suggestions = append(suggestions, foundation.Join(homeDir, "workspace"))
			suggestions = append(suggestions, foundation.Join(homeDir, "dev"))
		}
		if workingDir != "" && workingDir != homeDir {
			suggestions = append(suggestions, workingDir)
		}

		// Filter to only directories that exist
		existingSuggestions := []string{}
		for _, dir := range suggestions {
			if _, err := os.Stat(dir); err == nil {
				existingSuggestions = append(existingSuggestions, dir)
			}
		}

		writeJSON(w, http.StatusOK, ProjectDirectoriesResponse{
			HomeDir:     homeDir,
			WorkingDir:  workingDir,
			Suggestions: existingSuggestions,
		})
	}
}

// BrowseDirectories godoc
// @Summary Browse server filesystem directories
// @Description Lists subdirectories in the specified path for directory picker
// @Tags projects
// @Produce json
// @Param path query string false "Directory path to browse (defaults to home directory)"
// @Success 200 {object} DirectoryListResponse
// @Failure 400 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Router /api/v1/projects/browse [get]
func BrowseDirectories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get path from query parameter
		dirPath := r.URL.Query().Get("path")

		// If no path provided, use home directory
		if dirPath == "" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				writeError(w, http.StatusInternalServerError, err, "Failed to get home directory")
				return
			}
			dirPath = homeDir
		}

		// Clean and validate the path
		dirPath = filepath.Clean(dirPath)

		// Check if directory exists
		info, err := os.Stat(dirPath)
		if err != nil {
			if os.IsNotExist(err) {
				writeError(w, http.StatusNotFound, err, "Directory not found")
			} else if os.IsPermission(err) {
				writeError(w, http.StatusForbidden, err, "Permission denied")
			} else {
				writeError(w, http.StatusBadRequest, err, "Invalid directory path")
			}
			return
		}

		// Ensure it's a directory
		if !info.IsDir() {
			writeError(w, http.StatusBadRequest, fmt.Errorf("not a directory"), "Path is not a directory")
			return
		}

		// Read directory contents
		entries, err := os.ReadDir(dirPath)
		if err != nil {
			if os.IsPermission(err) {
				writeError(w, http.StatusForbidden, err, "Permission denied")
			} else {
				writeError(w, http.StatusInternalServerError, err, "Failed to read directory")
			}
			return
		}

		// Filter to only directories and build response
		directories := []DirectoryEntry{}
		for _, entry := range entries {
			if entry.IsDir() {
				// Skip hidden directories (starting with .)
				if len(entry.Name()) > 0 && entry.Name()[0] == '.' {
					continue
				}

				fullPath := foundation.Join(dirPath, entry.Name())

				// Check if we can read this directory (to show if it's accessible)
				accessible := true
				if _, err := os.ReadDir(fullPath); err != nil {
					accessible = false
				}

				directories = append(directories, DirectoryEntry{
					Name:       entry.Name(),
					Path:       fullPath,
					Accessible: accessible,
				})
			}
		}

		// Get parent directory
		parentDir := filepath.Dir(dirPath)
		if parentDir == dirPath {
			parentDir = "" // We're at root
		}

		writeJSON(w, http.StatusOK, DirectoryListResponse{
			CurrentPath: dirPath,
			ParentPath:  parentDir,
			Directories: directories,
			Count:       len(directories),
		})
	}
}

// ListRecentProjects godoc
// @Summary List recent projects
// @Description Returns recently opened projects from config
// @Tags projects
// @Produce json
// @Success 200 {object} ProjectListResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/projects/recent [get]
func ListRecentProjects() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get recent projects from config
		projects := project.RecentProjectInfos()

		writeJSON(w, http.StatusOK, ProjectListResponse{
			Projects: projects,
			Count:    len(projects),
		})
	}
}

// CreateProject godoc
// @Summary Create new project
// @Description Creates a new project with default demo files
// @Tags projects
// @Accept json
// @Produce json
// @Param request body CreateProjectRequest true "Project creation request"
// @Success 201 {object} project.ProjectInfo
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/projects [post]
func CreateProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateProjectRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "Invalid request body")
			return
		}

		// Validate input
		if req.Name == "" {
			writeError(w, http.StatusBadRequest, fmt.Errorf("name is required"), "Project name cannot be empty")
			return
		}
		if req.Path == "" {
			writeError(w, http.StatusBadRequest, fmt.Errorf("path is required"), "Parent path cannot be empty")
			return
		}

		// Check if parent directory exists
		if _, err := os.Stat(req.Path); os.IsNotExist(err) {
			writeError(w, http.StatusBadRequest, err, "Parent directory does not exist")
			return
		}

		// Build full project path
		fullPath := foundation.Join(req.Path, req.Name)

		// Check if project already exists
		if _, err := os.Stat(fullPath); err == nil {
			writeError(w, http.StatusConflict, fmt.Errorf("project already exists"), "Project directory already exists")
			return
		}

		// Initialize project
		info, err := project.InitProject(fullPath)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to initialize project")
			return
		}

		// Add to recent entries
		if err := config.AppendRecentEntry(fullPath); err != nil {
			// Log warning but don't fail the request
			logging.Warn("Failed to add project to recent entries", map[string]interface{}{
				"error": err.Error(),
				"path":  fullPath,
			})
		}

		writeJSON(w, http.StatusCreated, info)
	}
}

// GetProject godoc
// @Summary Get project details
// @Description Returns project information including documents list
// @Tags projects
// @Produce json
// @Param path query string true "Project path (URL encoded)"
// @Success 200 {object} project.ProjectInfo
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/projects/get [get]
func GetProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get path from query parameter
		encodedPath := r.URL.Query().Get("path")
		if encodedPath == "" {
			writeError(w, http.StatusBadRequest, fmt.Errorf("path is required"), "Path parameter is required")
			return
		}

		// URL decode the path
		projectPath, err := url.QueryUnescape(encodedPath)
		if err != nil {
			writeError(w, http.StatusBadRequest, err, "Invalid path encoding")
			return
		}

		// Check if project directory exists
		if _, err := os.Stat(projectPath); os.IsNotExist(err) {
			writeError(w, http.StatusNotFound, err, "Project not found")
			return
		}

		// Read project info
		info, err := project.ReadProject(projectPath)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to read project")
			return
		}

		writeJSON(w, http.StatusOK, info)
	}
}

// DeleteProject godoc
// @Summary Delete project
// @Description Removes project from disk and recent entries
// @Tags projects
// @Param path query string true "Project path (URL encoded)"
// @Success 204 "Project deleted successfully"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/projects [delete]
func DeleteProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get path from query parameter
		encodedPath := r.URL.Query().Get("path")
		if encodedPath == "" {
			writeError(w, http.StatusBadRequest, fmt.Errorf("path is required"), "Path parameter is required")
			return
		}

		// URL decode the path
		projectPath, err := url.QueryUnescape(encodedPath)
		if err != nil {
			writeError(w, http.StatusBadRequest, err, "Invalid path encoding")
			return
		}

		// Check if project exists
		if _, err := os.Stat(projectPath); os.IsNotExist(err) {
			writeError(w, http.StatusNotFound, err, "Project not found")
			return
		}

		// Validate path to prevent deletion of system directories
		// Ensure it's an absolute path and contains "apigear" directory
		absPath, err := filepath.Abs(projectPath)
		if err != nil {
			writeError(w, http.StatusBadRequest, err, "Invalid path")
			return
		}

		// Security check: ensure it has an apigear subdirectory
		apigearPath := foundation.Join(absPath, "apigear")
		if _, err := os.Stat(apigearPath); os.IsNotExist(err) {
			writeError(w, http.StatusBadRequest, fmt.Errorf("not a valid project"), "Directory does not contain apigear folder")
			return
		}

		// Remove from filesystem
		if err := os.RemoveAll(absPath); err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to delete project")
			return
		}

		// Remove from recent entries
		if err := config.RemoveRecentEntry(absPath); err != nil {
			// Log warning but don't fail the request
			logging.Warn("Failed to remove project from recent entries", map[string]interface{}{
				"error": err.Error(),
				"path":  absPath,
			})
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// ReadFile godoc
// @Summary Read file contents
// @Description Reads the contents of a file for editing
// @Tags projects
// @Produce json
// @Param path query string true "File path"
// @Success 200 {object} ReadFileResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/projects/files/read [get]
func ReadFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Query().Get("path")
		if filePath == "" {
			writeError(w, http.StatusBadRequest, fmt.Errorf("path is required"), "File path is required")
			return
		}

		// Security: ensure file exists and is readable
		info, err := os.Stat(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				writeError(w, http.StatusNotFound, err, "File not found")
			} else if os.IsPermission(err) {
				writeError(w, http.StatusForbidden, err, "Permission denied")
			} else {
				writeError(w, http.StatusBadRequest, err, "Invalid file path")
			}
			return
		}

		// Ensure it's a file, not a directory
		if info.IsDir() {
			writeError(w, http.StatusBadRequest, fmt.Errorf("path is a directory"), "Path must be a file")
			return
		}

		// Read file contents
		content, err := os.ReadFile(filePath)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to read file")
			return
		}

		writeJSON(w, http.StatusOK, ReadFileResponse{
			Path:     filePath,
			Content:  string(content),
			Encoding: "utf-8",
		})
	}
}

// WriteFile godoc
// @Summary Write file contents
// @Description Writes content to a file
// @Tags projects
// @Accept json
// @Produce json
// @Param request body WriteFileRequest true "Write file request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/projects/files/write [post]
func WriteFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req WriteFileRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "Invalid request body")
			return
		}

		if req.Path == "" {
			writeError(w, http.StatusBadRequest, fmt.Errorf("path is required"), "File path is required")
			return
		}

		// Security: ensure parent directory exists
		parentDir := filepath.Dir(req.Path)
		if _, err := os.Stat(parentDir); os.IsNotExist(err) {
			writeError(w, http.StatusBadRequest, err, "Parent directory does not exist")
			return
		}

		// Write file
		err := os.WriteFile(req.Path, []byte(req.Content), 0644)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to write file")
			return
		}

		writeJSON(w, http.StatusOK, map[string]string{
			"message": "File saved successfully",
			"path":    req.Path,
		})
	}
}

// OpenFileExternal godoc
// @Summary Open file in external editor
// @Description Opens a file in the configured external editor
// @Tags projects
// @Accept json
// @Produce json
// @Param request body OpenExternalRequest true "Open external request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/projects/files/open-external [post]
func OpenFileExternal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req OpenExternalRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "Invalid request body")
			return
		}

		if req.Path == "" {
			writeError(w, http.StatusBadRequest, fmt.Errorf("path is required"), "File path is required")
			return
		}

		// Check if file exists
		if _, err := os.Stat(req.Path); os.IsNotExist(err) {
			writeError(w, http.StatusNotFound, err, "File not found")
			return
		}

		// Open in external editor using project.OpenEditor
		err := project.OpenEditor(req.Path)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to open file in external editor")
			return
		}

		writeJSON(w, http.StatusOK, map[string]string{
			"message": "File opened in external editor",
			"path":    req.Path,
		})
	}
}

// GenerateCode godoc
// @Summary Generate code from solution file
// @Description Streams code generation events via Server-Sent Events
// @Tags projects
// @Produce text/event-stream
// @Param path query string true "Solution file path"
// @Param force query boolean false "Force overwrite existing files"
// @Success 200 {object} tasks.TaskEvent
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/projects/generate [get]
func GenerateCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		solutionPath := r.URL.Query().Get("path")
		if solutionPath == "" {
			writeError(w, http.StatusBadRequest, fmt.Errorf("path is required"), "Solution file path is required")
			return
		}

		// Parse force parameter
		force := r.URL.Query().Get("force") == "true"

		// Validate solution file
		result, err := spec.CheckFileAndType(solutionPath, spec.DocumentTypeSolution)
		if err != nil {
			writeError(w, http.StatusBadRequest, err, "Invalid solution file")
			return
		}

		if !result.Valid() {
			// Return validation errors
			errors := make([]string, 0, len(result.Errors))
			for _, err := range result.Errors {
				errors = append(errors, fmt.Sprintf("%s: %s", err.Field, err.Description))
			}
			writeError(w, http.StatusBadRequest,
				fmt.Errorf("validation failed"),
				fmt.Sprintf("Solution validation failed: %v", errors))
			return
		}

		// Set SSE headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		flusher, ok := w.(http.Flusher)
		if !ok {
			writeError(w, http.StatusInternalServerError,
				fmt.Errorf("streaming not supported"),
				"Streaming not supported")
			return
		}

		// Create context that cancels when client disconnects
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		// Send initial connection message
		sendSSEEvent(w, flusher, "connected", map[string]interface{}{
			"message": "Code generation started",
			"path":    solutionPath,
			"force":   force,
		})

		// Create solution runner with task event callback
		runner := solution.NewRunner()
		runner.OnTask(func(evt *tasks.TaskEvent) {
			// Send task event via SSE
			sendSSEEvent(w, flusher, "task", map[string]interface{}{
				"name":  evt.Name,
				"state": evt.State.String(),
				"meta":  evt.Meta,
			})
		})

		// Run code generation
		err = runner.RunSource(ctx, solutionPath, force)
		if err != nil {
			sendSSEEvent(w, flusher, "error", map[string]interface{}{
				"message": err.Error(),
			})
			return
		}

		// Send completion message
		sendSSEEvent(w, flusher, "completed", map[string]interface{}{
			"message": "Code generation completed successfully",
		})
	}
}
