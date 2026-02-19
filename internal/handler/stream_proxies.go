package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/apigear-io/cli/pkg/stream/config"
	_ "github.com/apigear-io/cli/pkg/stream/proxy" // Used in Swagger docs
)

// ListStreamProxies returns all stream proxies
// @Summary List all stream proxies
// @Description Get a list of all configured stream proxies with their status
// @Tags stream
// @Produce json
// @Success 200 {array} proxy.Info
// @Router /api/v1/stream/proxies [get]
func ListStreamProxies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		services := getStreamServices()
		proxies := services.ProxyManager.ListProxies()
		writeJSON(w, http.StatusOK, proxies)
	}
}

// GetStreamProxy returns a specific proxy by name
// @Summary Get a stream proxy
// @Description Get details for a specific stream proxy
// @Tags stream
// @Produce json
// @Param name path string true "Proxy name"
// @Success 200 {object} proxy.Info
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/proxies/{name} [get]
func GetStreamProxy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "proxy name is required")
			return
		}

		services := getStreamServices()
		p, err := services.ProxyManager.GetProxy(name)
		if err != nil {
			writeError(w, http.StatusNotFound, err, "proxy not found")
			return
		}

		info := p.Info()
		writeJSON(w, http.StatusOK, info)
	}
}

// CreateStreamProxyRequest represents the request to create a proxy
type CreateStreamProxyRequest struct {
	Name    string              `json:"name"`
	Config  config.ProxyConfig  `json:"config"`
}

// CreateStreamProxy creates a new stream proxy
// @Summary Create a stream proxy
// @Description Create a new stream proxy with the specified configuration
// @Tags stream
// @Accept json
// @Produce json
// @Param request body CreateStreamProxyRequest true "Proxy configuration"
// @Success 201 {object} proxy.Info
// @Failure 400 {object} ErrorResponse
// @Router /api/v1/stream/proxies [post]
func CreateStreamProxy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateStreamProxyRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err, "invalid request body")
			return
		}

		if req.Name == "" {
			writeError(w, http.StatusBadRequest, nil, "proxy name is required")
			return
		}

		// Validate config
		if req.Config.Listen == "" {
			writeError(w, http.StatusBadRequest, nil, "listen address is required")
			return
		}

		if req.Config.Mode == "" {
			req.Config.Mode = "proxy"
		}

		services := getStreamServices()
		if err := services.ProxyManager.AddProxy(req.Name, req.Config); err != nil {
			writeError(w, http.StatusBadRequest, err, "failed to create proxy")
			return
		}

		p, err := services.ProxyManager.GetProxy(req.Name)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to get created proxy")
			return
		}

		info := p.Info()
		writeJSON(w, http.StatusCreated, info)
	}
}

// UpdateStreamProxy updates an existing stream proxy
// @Summary Update a stream proxy
// @Description Update an existing stream proxy configuration
// @Tags stream
// @Accept json
// @Produce json
// @Param name path string true "Proxy name"
// @Param request body config.ProxyConfig true "Proxy configuration"
// @Success 200 {object} proxy.Info
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/proxies/{name} [put]
func UpdateStreamProxy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "proxy name is required")
			return
		}

		var cfg config.ProxyConfig
		if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
			writeError(w, http.StatusBadRequest, err, "invalid request body")
			return
		}

		services := getStreamServices()

		// Check if proxy exists
		_, err := services.ProxyManager.GetProxy(name)
		if err != nil {
			writeError(w, http.StatusNotFound, err, "proxy not found")
			return
		}

		// Remove and re-add with new config
		if err := services.ProxyManager.RemoveProxy(name); err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to remove proxy")
			return
		}

		if err := services.ProxyManager.AddProxy(name, cfg); err != nil {
			writeError(w, http.StatusBadRequest, err, "failed to update proxy")
			return
		}

		p, err := services.ProxyManager.GetProxy(name)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to get updated proxy")
			return
		}

		info := p.Info()
		writeJSON(w, http.StatusOK, info)
	}
}

// DeleteStreamProxy deletes a stream proxy
// @Summary Delete a stream proxy
// @Description Delete an existing stream proxy
// @Tags stream
// @Param name path string true "Proxy name"
// @Success 204
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/proxies/{name} [delete]
func DeleteStreamProxy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "proxy name is required")
			return
		}

		services := getStreamServices()
		if err := services.ProxyManager.RemoveProxy(name); err != nil {
			writeError(w, http.StatusNotFound, err, "proxy not found")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// StartStreamProxy starts a stream proxy
// @Summary Start a stream proxy
// @Description Start an existing stream proxy
// @Tags stream
// @Param name path string true "Proxy name"
// @Success 200 {object} proxy.Info
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/proxies/{name}/start [post]
func StartStreamProxy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "proxy name is required")
			return
		}

		services := getStreamServices()
		if err := services.ProxyManager.StartProxy(name); err != nil {
			writeError(w, http.StatusBadRequest, err, "failed to start proxy")
			return
		}

		p, err := services.ProxyManager.GetProxy(name)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to get proxy")
			return
		}

		info := p.Info()
		writeJSON(w, http.StatusOK, info)
	}
}

// StopStreamProxy stops a stream proxy
// @Summary Stop a stream proxy
// @Description Stop a running stream proxy
// @Tags stream
// @Param name path string true "Proxy name"
// @Success 200 {object} proxy.Info
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/proxies/{name}/stop [post]
func StopStreamProxy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "proxy name is required")
			return
		}

		services := getStreamServices()
		if err := services.ProxyManager.StopProxy(name); err != nil {
			writeError(w, http.StatusBadRequest, err, "failed to stop proxy")
			return
		}

		p, err := services.ProxyManager.GetProxy(name)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err, "failed to get proxy")
			return
		}

		info := p.Info()
		writeJSON(w, http.StatusOK, info)
	}
}

// GetStreamProxyStats returns statistics for a proxy
// @Summary Get proxy statistics
// @Description Get detailed statistics for a specific proxy
// @Tags stream
// @Produce json
// @Param name path string true "Proxy name"
// @Success 200 {object} proxy.Info
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/stream/proxies/{name}/stats [get]
func GetStreamProxyStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			writeError(w, http.StatusBadRequest, nil, "proxy name is required")
			return
		}

		services := getStreamServices()
		p, err := services.ProxyManager.GetProxy(name)
		if err != nil {
			writeError(w, http.StatusNotFound, err, "proxy not found")
			return
		}

		info := p.Info()
		writeJSON(w, http.StatusOK, info)
	}
}
