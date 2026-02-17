package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/apigear-io/cli/pkg/codegen/registry"
	"github.com/apigear-io/cli/pkg/foundation/git"
)

// TemplateInfo represents template information for API responses
type TemplateInfo struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Author       string   `json:"author"`
	Git          string   `json:"git"`
	Version      string   `json:"version"`
	Latest       string   `json:"latest"`
	Versions     []string `json:"versions"`
	InCache      bool     `json:"inCache"`
	InRegistry   bool     `json:"inRegistry"`
	Tags         []string `json:"tags,omitempty"`
	UpdateNeeded bool     `json:"updateNeeded"` // True if cached version < latest version
}

// TemplateListResponse represents the list of templates
type TemplateListResponse struct {
	Templates []*TemplateInfo `json:"templates"`
	Count     int             `json:"count"`
}

// InstallRequest represents the install request body
type InstallRequest struct {
	Version string `json:"version,omitempty"`
}

// InstallProgressEvent represents SSE progress events
type InstallProgressEvent struct {
	Type     string `json:"type"`     // "progress", "complete", "error"
	Message  string `json:"message"`  // Human-readable message
	Progress int    `json:"progress"` // 0-100 percentage
	Error    string `json:"error,omitempty"`
}

// isVersionNewer checks if current version is older than target version using semver
// Returns true if update is needed (current < target)
func isVersionNewer(currentVersion, targetVersion string) bool {
	if currentVersion == "" || targetVersion == "" {
		return false
	}

	// Parse versions, handling both with and without 'v' prefix
	current, err := semver.NewVersion(currentVersion)
	if err != nil {
		return false
	}

	target, err := semver.NewVersion(targetVersion)
	if err != nil {
		return false
	}

	// Return true if current version is less than target (update needed)
	return current.LessThan(target)
}

// convertRepoInfo converts git.RepoInfo to TemplateInfo
func convertRepoInfo(info *git.RepoInfo) *TemplateInfo {
	versions := make([]string, 0, len(info.Versions))
	for _, v := range info.Versions {
		versions = append(versions, v.Name)
	}

	return &TemplateInfo{
		Name:        info.Name,
		Description: info.Description,
		Author:      info.Author,
		Git:         info.Git,
		Version:     info.Version.Name,
		Latest:      info.Latest.Name,
		Versions:    versions,
		InCache:     info.InCache,
		InRegistry:  info.InRegistry,
	}
}

// mergeTemplateInfo merges registry and cache information
func mergeTemplateInfo(registryInfos, cacheInfos []*git.RepoInfo) []*TemplateInfo {
	// Create a map for quick lookup of cache info by name
	cacheMap := make(map[string]*git.RepoInfo)
	for _, info := range cacheInfos {
		name := registry.NameFromRepoID(info.Name)
		cacheMap[name] = info
	}

	// Create map for templates
	templateMap := make(map[string]*TemplateInfo)

	// Add all registry templates
	for _, info := range registryInfos {
		name := registry.NameFromRepoID(info.Name)
		templateInfo := convertRepoInfo(info)
		templateInfo.InRegistry = true

		// Check if template is in cache
		if cached, ok := cacheMap[name]; ok {
			templateInfo.InCache = true
			// Use cached version if available, otherwise use latest from cached info
			if cached.Version.Name != "" {
				templateInfo.Version = cached.Version.Name
			} else if cached.Latest.Name != "" {
				templateInfo.Version = cached.Latest.Name
			}

			// Check if update is needed using semantic versioning
			if templateInfo.Version != "" && templateInfo.Latest != "" {
				templateInfo.UpdateNeeded = isVersionNewer(templateInfo.Version, templateInfo.Latest)
			}
		}

		templateMap[name] = templateInfo
	}

	// Add cache-only templates (not in registry)
	for _, info := range cacheInfos {
		name := registry.NameFromRepoID(info.Name)
		if _, exists := templateMap[name]; !exists {
			templateInfo := convertRepoInfo(info)
			templateInfo.InCache = true
			templateInfo.InRegistry = false
			templateMap[name] = templateInfo
		}
	}

	// Convert map to slice
	templates := make([]*TemplateInfo, 0, len(templateMap))
	for _, t := range templateMap {
		templates = append(templates, t)
	}

	// Sort templates by name for consistent ordering
	sort.Slice(templates, func(i, j int) bool {
		return templates[i].Name < templates[j].Name
	})

	return templates
}

// ListTemplates godoc
// @Summary List all registry templates
// @Description Returns all templates available in the registry with their cache status
// @Tags templates
// @Produce json
// @Success 200 {object} TemplateListResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/templates [get]
func ListTemplates() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get registry templates
		registryInfos, err := registry.Registry.List()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to list registry templates")
			return
		}

		// Get cached templates
		cacheInfos, err := registry.Cache.List()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to list cached templates")
			return
		}

		// Merge information
		templates := mergeTemplateInfo(registryInfos, cacheInfos)

		writeJSON(w, http.StatusOK, TemplateListResponse{
			Templates: templates,
			Count:     len(templates),
		})
	}
}

// GetTemplate godoc
// @Summary Get template details
// @Description Returns detailed information about a specific template
// @Tags templates
// @Produce json
// @Param id query string true "Template ID (e.g., apigear-io/template-ts)"
// @Success 200 {object} TemplateInfo
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/templates/get [get]
func GetTemplate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			writeError(w, http.StatusBadRequest, fmt.Errorf("missing template id"), "Template ID is required")
			return
		}

		// Try to get from registry first
		registryInfo, err := registry.Registry.Get(id)
		if err != nil {
			writeError(w, http.StatusNotFound, err, "Template not found")
			return
		}

		templateInfo := convertRepoInfo(registryInfo)
		templateInfo.InRegistry = true

		// Check if it's in cache
		name := registry.NameFromRepoID(id)
		if registry.Cache.Exists(name) {
			cacheInfo, err := registry.Cache.Info(name)
			if err == nil {
				templateInfo.InCache = true
				templateInfo.Version = cacheInfo.Version.Name
			}
		}

		writeJSON(w, http.StatusOK, templateInfo)
	}
}

// InstallTemplate godoc
// @Summary Install a template
// @Description Installs a template from the registry using Server-Sent Events for progress updates
// @Tags templates
// @Accept json
// @Produce text/event-stream
// @Param id query string true "Template ID (e.g., apigear-io/template-ts)"
// @Param request body InstallRequest false "Install request with optional version"
// @Success 200 {object} InstallProgressEvent
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/templates/install [post]
func InstallTemplate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			writeError(w, http.StatusBadRequest, fmt.Errorf("missing template id"), "Template ID is required")
			return
		}

		// Parse request body
		var req InstallRequest
		if r.Body != nil {
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil && err.Error() != "EOF" {
				writeError(w, http.StatusBadRequest, err, "Invalid request body")
				return
			}
		}

		// Build repo ID with version
		repoID := id
		if req.Version != "" {
			repoID = registry.MakeRepoID(id, req.Version)
		}

		// Set SSE headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		flusher, ok := w.(http.Flusher)
		if !ok {
			writeError(w, http.StatusInternalServerError, fmt.Errorf("streaming not supported"), "Streaming not supported")
			return
		}

		// Helper to send SSE events
		sendSSE := func(event InstallProgressEvent) {
			data, _ := json.Marshal(event)
			fmt.Fprintf(w, "data: %s\n\n", data)
			flusher.Flush()
		}

		// Start installation
		sendSSE(InstallProgressEvent{
			Type:     "progress",
			Message:  "Starting installation...",
			Progress: 10,
		})

		sendSSE(InstallProgressEvent{
			Type:     "progress",
			Message:  "Resolving template from registry...",
			Progress: 25,
		})

		// Install template
		installedID, err := registry.GetOrInstallTemplateFromRepoID(repoID)
		if err != nil {
			sendSSE(InstallProgressEvent{
				Type:    "error",
				Message: "Installation failed",
				Error:   err.Error(),
			})
			return
		}

		sendSSE(InstallProgressEvent{
			Type:     "progress",
			Message:  "Cloning repository...",
			Progress: 50,
		})

		sendSSE(InstallProgressEvent{
			Type:     "progress",
			Message:  "Checking out version...",
			Progress: 75,
		})

		sendSSE(InstallProgressEvent{
			Type:     "complete",
			Message:  fmt.Sprintf("Template %s installed successfully", installedID),
			Progress: 100,
		})
	}
}

// ListCachedTemplates godoc
// @Summary List installed templates
// @Description Returns all templates currently installed in the local cache
// @Tags templates
// @Produce json
// @Success 200 {object} TemplateListResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/templates/cache [get]
func ListCachedTemplates() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cacheInfos, err := registry.Cache.List()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to list cached templates")
			return
		}

		templates := make([]*TemplateInfo, 0, len(cacheInfos))
		for _, info := range cacheInfos {
			templateInfo := convertRepoInfo(info)
			templateInfo.InCache = true
			templates = append(templates, templateInfo)
		}

		// Sort templates by name for consistent ordering
		sort.Slice(templates, func(i, j int) bool {
			return templates[i].Name < templates[j].Name
		})

		writeJSON(w, http.StatusOK, TemplateListResponse{
			Templates: templates,
			Count:     len(templates),
		})
	}
}

// RemoveTemplate godoc
// @Summary Remove a template from cache
// @Description Removes an installed template from the local cache
// @Tags templates
// @Produce json
// @Param id query string true "Template ID (e.g., apigear-io/template-ts)"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/templates/cache/remove [delete]
func RemoveTemplate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			writeError(w, http.StatusBadRequest, fmt.Errorf("missing template id"), "Template ID is required")
			return
		}

		err := registry.Cache.Remove(id)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to remove template")
			return
		}

		writeJSON(w, http.StatusOK, map[string]string{
			"message": fmt.Sprintf("Template %s removed successfully", id),
		})
	}
}

// CleanCache godoc
// @Summary Clean template cache
// @Description Removes all templates from the local cache
// @Tags templates
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/templates/cache/clean [post]
func CleanCache() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := registry.Cache.Clean()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to clean cache")
			return
		}

		writeJSON(w, http.StatusOK, map[string]string{
			"message": "Cache cleaned successfully",
		})
	}
}

// UpdateRegistry godoc
// @Summary Update template registry
// @Description Updates the template registry from the remote repository
// @Tags templates
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/templates/registry/update [post]
func UpdateRegistry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := registry.Registry.Update()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to update registry")
			return
		}

		writeJSON(w, http.StatusOK, map[string]string{
			"message": "Registry updated successfully",
		})
	}
}

// SearchTemplates godoc
// @Summary Search templates
// @Description Searches for templates by name or description
// @Tags templates
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} TemplateListResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/templates/search [get]
func SearchTemplates() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query == "" {
			writeError(w, http.StatusBadRequest, fmt.Errorf("missing query parameter"), "Search query is required")
			return
		}

		// Search in registry
		registryInfos, err := registry.Registry.Search(query)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to search registry")
			return
		}

		// Search in cache
		cacheInfos, err := registry.Cache.Search(query)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "Failed to search cache")
			return
		}

		// Merge results
		templates := mergeTemplateInfo(registryInfos, cacheInfos)

		// Filter by query in description as well
		filtered := make([]*TemplateInfo, 0)
		queryLower := strings.ToLower(query)
		for _, t := range templates {
			if strings.Contains(strings.ToLower(t.Name), queryLower) ||
				strings.Contains(strings.ToLower(t.Description), queryLower) {
				filtered = append(filtered, t)
			}
		}

		writeJSON(w, http.StatusOK, TemplateListResponse{
			Templates: filtered,
			Count:     len(filtered),
		})
	}
}
