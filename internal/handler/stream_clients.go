package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	_ "github.com/apigear-io/cli/pkg/stream/client" // Used in Swagger docs
	"github.com/apigear-io/cli/pkg/stream/config"
)

// ListStreamClients returns all stream clients
// @Summary List all stream clients
// @Description Get a list of all configured stream clients with their status
// @Tags stream
// @Produce json
// @Success 200 {array} client.Info
// @Router /api/v1/stream/clients [get]
func ListStreamClients() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		services := getStreamServices()
		clients := services.ClientManager.ListClients()
		writeJSON(w, http.StatusOK, clients)
	}
}

// GetStreamClient returns a specific client by name
// @Summary Get a stream client
// @Description Get details for a specific stream client
// @Tags stream
// @Produce json
// @Param name path string true "Client name"
// @Success 200 {object} client.Info
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/clients/{name} [get]
func GetStreamClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "client name is required")
			return
		}

		services := getStreamServices()
		c, err := services.ClientManager.GetClient(name)
		if err != nil {
			writeError(w, http.StatusNotFound, err, "client not found")
			return
		}

		info := c.Info()
		writeJSON(w, http.StatusOK, info)
	}
}

// CreateStreamClientRequest represents the request to create a client
type CreateStreamClientRequest struct {
	Name   string              `json:"name"`
	Config config.ClientConfig `json:"config"`
}

// CreateStreamClient creates a new stream client
// @Summary Create a stream client
// @Description Create a new stream client with the specified configuration
// @Tags stream
// @Accept json
// @Produce json
// @Param request body CreateStreamClientRequest true "Client configuration"
// @Success 201 {object} client.Info
// @Failure 400 {object} ErrorResponse
// @Router /api/v1/stream/clients [post]
func CreateStreamClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateStreamClientRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "invalid request body")
			return
		}

		if req.Name == "" {
			writeError(w, http.StatusBadRequest, nil, "client name is required")
			return
		}

		// Validate config
		if req.Config.URL == "" {
			writeError(w, http.StatusBadRequest, nil, "URL is required")
			return
		}

		// Persist to config file first
		configPath := getStreamConfigPath()
		persistence := config.NewConfigPersistence(configPath)
		if err := persistence.AddClient(req.Name, req.Config); err != nil {
			writeError(w, http.StatusBadRequest, err, "failed to save client to config")
			return
		}

		// Add to in-memory manager
		services := getStreamServices()
		if err := services.ClientManager.AddClient(req.Name, req.Config); err != nil {
			// Try to rollback config file change
			_ = persistence.DeleteClient(req.Name)
			writeError(w, http.StatusBadRequest, err, "failed to create client")
			return
		}

		c, err := services.ClientManager.GetClient(req.Name)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to get created client")
			return
		}

		info := c.Info()
		writeJSON(w, http.StatusCreated, info)
	}
}

// UpdateStreamClient updates an existing stream client
// @Summary Update a stream client
// @Description Update an existing stream client configuration
// @Tags stream
// @Accept json
// @Produce json
// @Param name path string true "Client name"
// @Param request body config.ClientConfig true "Client configuration"
// @Success 200 {object} client.Info
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/clients/{name} [put]
func UpdateStreamClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "client name is required")
			return
		}

		var cfg config.ClientConfig
		if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
			writeError(w, http.StatusBadRequest, err, "invalid request body")
			return
		}

		services := getStreamServices()

		// Check if client exists
		_, err := services.ClientManager.GetClient(name)
		if err != nil {
			writeError(w, http.StatusNotFound, err, "client not found")
			return
		}

		// Persist to config file first (upsert: update if exists, add if not)
		configPath := getStreamConfigPath()
		persistence := config.NewConfigPersistence(configPath)
		err = persistence.UpdateClient(name, cfg)
		if err != nil {
			// If update fails because client doesn't exist in config, add it instead
			if err := persistence.AddClient(name, cfg); err != nil {
				writeError(w, http.StatusBadRequest, err, "failed to save client to config")
				return
			}
		}

		// Remove and re-add with new config in memory
		if err := services.ClientManager.RemoveClient(name); err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to remove client")
			return
		}

		if err := services.ClientManager.AddClient(name, cfg); err != nil {
			writeError(w, http.StatusBadRequest, err, "failed to update client")
			return
		}

		c, err := services.ClientManager.GetClient(name)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to get updated client")
			return
		}

		info := c.Info()
		writeJSON(w, http.StatusOK, info)
	}
}

// DeleteStreamClient deletes a stream client
// @Summary Delete a stream client
// @Description Delete an existing stream client
// @Tags stream
// @Param name path string true "Client name"
// @Success 204
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/clients/{name} [delete]
func DeleteStreamClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "client name is required")
			return
		}

		services := getStreamServices()

		// Remove from in-memory manager first
		if err := services.ClientManager.RemoveClient(name); err != nil {
			writeError(w, http.StatusNotFound, err, "client not found")
			return
		}

		// Persist to config file (best effort - don't fail if it doesn't exist)
		configPath := getStreamConfigPath()
		persistence := config.NewConfigPersistence(configPath)
		_ = persistence.DeleteClient(name)
		// Note: We ignore errors here since the client is already removed from memory.
		// This handles the case where client was in memory but not persisted to config.

		w.WriteHeader(http.StatusNoContent)
	}
}

// ConnectStreamClient connects a stream client
// @Summary Connect a stream client
// @Description Connect an existing stream client to its server
// @Tags stream
// @Param name path string true "Client name"
// @Success 200 {object} client.Info
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/clients/{name}/connect [post]
func ConnectStreamClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "client name is required")
			return
		}

		services := getStreamServices()
		if err := services.ClientManager.ConnectClient(name); err != nil {
			writeError(w, http.StatusBadRequest, err, "failed to connect client")
			return
		}

		c, err := services.ClientManager.GetClient(name)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to get client")
			return
		}

		info := c.Info()
		writeJSON(w, http.StatusOK, info)
	}
}

// DisconnectStreamClient disconnects a stream client
// @Summary Disconnect a stream client
// @Description Disconnect a connected stream client
// @Tags stream
// @Param name path string true "Client name"
// @Success 200 {object} client.Info
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/clients/{name}/disconnect [post]
func DisconnectStreamClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "client name is required")
			return
		}

		services := getStreamServices()
		if err := services.ClientManager.DisconnectClient(name); err != nil {
			writeError(w, http.StatusBadRequest, err, "failed to disconnect client")
			return
		}

		c, err := services.ClientManager.GetClient(name)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to get client")
			return
		}

		info := c.Info()
		writeJSON(w, http.StatusOK, info)
	}
}
