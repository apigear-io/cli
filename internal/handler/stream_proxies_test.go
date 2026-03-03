package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/apigear-io/cli/pkg/stream"
	"github.com/apigear-io/cli/pkg/stream/config"
)

// setupTestStreamServices creates a fresh stream services instance for testing
func setupTestStreamServices() *stream.Services {
	// Create a temp directory for test config files
	tempDir, err := os.MkdirTemp("", "stream-test-*")
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(tempDir, "stream.yaml")
	SetStreamConfigPath(configPath)

	services := stream.NewServices()
	setStreamServices(services)
	return services
}

func TestListStreamProxies_Empty(t *testing.T) {
	setupTestStreamServices()
	handler := ListStreamProxies()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stream/proxies", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var proxies []interface{}
	err := json.NewDecoder(w.Body).Decode(&proxies)
	require.NoError(t, err)
	assert.Empty(t, proxies)
}

func TestCreateStreamProxy_Success(t *testing.T) {
	setupTestStreamServices()
	handler := CreateStreamProxy()

	reqBody := CreateStreamProxyRequest{
		Name: "test-proxy",
		Config: config.ProxyConfig{
			Listen:  "ws://localhost:15550/ws",
			Backend: "ws://localhost:15560/ws",
			Mode:    "proxy",
		},
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/proxies", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, "test-proxy", response["name"])
	assert.Equal(t, "stopped", response["status"])
}

func TestCreateStreamProxy_MissingName(t *testing.T) {
	setupTestStreamServices()
	handler := CreateStreamProxy()

	reqBody := CreateStreamProxyRequest{
		Config: config.ProxyConfig{
			Listen: "ws://localhost:15550/ws",
			Mode:   "proxy",
		},
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/proxies", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Error, "proxy name is required")
}

func TestCreateStreamProxy_MissingListenAddress(t *testing.T) {
	setupTestStreamServices()
	handler := CreateStreamProxy()

	reqBody := CreateStreamProxyRequest{
		Name: "test-proxy",
		Config: config.ProxyConfig{
			Mode: "proxy",
		},
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/proxies", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Error, "listen address is required")
}

func TestCreateStreamProxy_InvalidJSON(t *testing.T) {
	setupTestStreamServices()
	handler := CreateStreamProxy()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/proxies", bytes.NewReader([]byte("invalid json")))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Error, "invalid")
}

func TestCreateStreamProxy_DefaultMode(t *testing.T) {
	setupTestStreamServices()
	handler := CreateStreamProxy()

	reqBody := CreateStreamProxyRequest{
		Name: "test-proxy",
		Config: config.ProxyConfig{
			Listen: "ws://localhost:15550/ws",
			// Mode is empty, should default to "proxy"
		},
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/proxies", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, "proxy", response["mode"])
}

func TestGetStreamProxy_Success(t *testing.T) {
	services := setupTestStreamServices()

	// Create a proxy first
	err := services.ProxyManager.AddProxy("test-proxy", config.ProxyConfig{
		Listen:  "ws://localhost:15550/ws",
		Backend: "ws://localhost:15560/ws",
		Mode:    "proxy",
	})
	require.NoError(t, err)

	handler := GetStreamProxy()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stream/proxies/test-proxy", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "test-proxy")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, "test-proxy", response["name"])
}

func TestGetStreamProxy_NotFound(t *testing.T) {
	setupTestStreamServices()
	handler := GetStreamProxy()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stream/proxies/nonexistent", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "nonexistent")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Error, "not found")
}

func TestUpdateStreamProxy_Success(t *testing.T) {
	services := setupTestStreamServices()

	// Create a proxy first
	err := services.ProxyManager.AddProxy("test-proxy", config.ProxyConfig{
		Listen:  "ws://localhost:15550/ws",
		Backend: "ws://localhost:15560/ws",
		Mode:    "proxy",
	})
	require.NoError(t, err)

	handler := UpdateStreamProxy()

	updatedConfig := config.ProxyConfig{
		Listen:  "ws://localhost:15551/ws",
		Backend: "ws://localhost:15561/ws",
		Mode:    "proxy",
	}
	body, err := json.Marshal(updatedConfig)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/stream/proxies/test-proxy", bytes.NewReader(body))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "test-proxy")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, "test-proxy", response["name"])
	assert.Equal(t, "ws://localhost:15551/ws", response["listen"])
}

func TestUpdateStreamProxy_NotFound(t *testing.T) {
	setupTestStreamServices()
	handler := UpdateStreamProxy()

	updatedConfig := config.ProxyConfig{
		Listen: "ws://localhost:15551/ws",
		Mode:   "proxy",
	}
	body, err := json.Marshal(updatedConfig)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/stream/proxies/nonexistent", bytes.NewReader(body))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "nonexistent")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteStreamProxy_Success(t *testing.T) {
	services := setupTestStreamServices()

	// Create a proxy first
	err := services.ProxyManager.AddProxy("test-proxy", config.ProxyConfig{
		Listen: "ws://localhost:15550/ws",
		Mode:   "echo",
	})
	require.NoError(t, err)

	handler := DeleteStreamProxy()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/stream/proxies/test-proxy", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "test-proxy")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// Verify proxy was deleted
	_, err = services.ProxyManager.GetProxy("test-proxy")
	assert.Error(t, err)
}

func TestDeleteStreamProxy_NotFound(t *testing.T) {
	setupTestStreamServices()
	handler := DeleteStreamProxy()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/stream/proxies/nonexistent", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "nonexistent")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestStartStreamProxy_Success(t *testing.T) {
	t.Skip("Skipping due to race condition in proxy.Start() - TODO: fix in proxy package")

	services := setupTestStreamServices()

	// Create a proxy first with echo mode (easiest to start without backend)
	err := services.ProxyManager.AddProxy("test-proxy", config.ProxyConfig{
		Listen: "ws://localhost:15550/ws",
		Mode:   "echo",
	})
	require.NoError(t, err)

	handler := StartStreamProxy()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/proxies/test-proxy/start", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "test-proxy")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, "running", response["status"])

	// Clean up - stop the proxy
	_ = services.ProxyManager.StopProxy("test-proxy")
}

func TestStartStreamProxy_NotFound(t *testing.T) {
	setupTestStreamServices()
	handler := StartStreamProxy()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/proxies/nonexistent/start", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "nonexistent")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestStopStreamProxy_Success(t *testing.T) {
	t.Skip("Skipping due to race condition in proxy.Start() - TODO: fix in proxy package")

	services := setupTestStreamServices()

	// Create and start a proxy first
	err := services.ProxyManager.AddProxy("test-proxy", config.ProxyConfig{
		Listen: "ws://localhost:15552/ws", // Use different port to avoid conflicts
		Mode:   "echo",
	})
	require.NoError(t, err)

	err = services.ProxyManager.StartProxy("test-proxy")
	require.NoError(t, err)

	handler := StopStreamProxy()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/proxies/test-proxy/stop", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "test-proxy")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, "stopped", response["status"])
}

func TestGetStreamProxyStats_Success(t *testing.T) {
	services := setupTestStreamServices()

	// Create a proxy first
	err := services.ProxyManager.AddProxy("test-proxy", config.ProxyConfig{
		Listen: "ws://localhost:15550/ws",
		Mode:   "echo",
	})
	require.NoError(t, err)

	handler := GetStreamProxyStats()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stream/proxies/test-proxy/stats", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "test-proxy")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, "test-proxy", response["name"])
	assert.NotNil(t, response["messagesReceived"])
	assert.NotNil(t, response["messagesSent"])
}

func TestListStreamProxies_Multiple(t *testing.T) {
	services := setupTestStreamServices()

	// Create multiple proxies
	err := services.ProxyManager.AddProxy("proxy1", config.ProxyConfig{
		Listen: "ws://localhost:15550/ws",
		Mode:   "echo",
	})
	require.NoError(t, err)

	err = services.ProxyManager.AddProxy("proxy2", config.ProxyConfig{
		Listen: "ws://localhost:15551/ws",
		Mode:   "echo",
	})
	require.NoError(t, err)

	handler := ListStreamProxies()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stream/proxies", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var proxies []map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&proxies)
	require.NoError(t, err)
	assert.Len(t, proxies, 2)

	// Verify proxy names
	names := []string{}
	for _, p := range proxies {
		names = append(names, p["name"].(string))
	}
	assert.Contains(t, names, "proxy1")
	assert.Contains(t, names, "proxy2")
}

func TestCreateStreamProxy_DuplicateName(t *testing.T) {
	services := setupTestStreamServices()

	// Create first proxy
	err := services.ProxyManager.AddProxy("test-proxy", config.ProxyConfig{
		Listen: "ws://localhost:15550/ws",
		Mode:   "echo",
	})
	require.NoError(t, err)

	// Try to create another with same name
	handler := CreateStreamProxy()

	reqBody := CreateStreamProxyRequest{
		Name: "test-proxy",
		Config: config.ProxyConfig{
			Listen: "ws://localhost:15551/ws",
			Mode:   "echo",
		},
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/proxies", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Error, "already exists")
}
