package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/apigear-io/cli/pkg/stream/config"
)

func TestListStreamClients_Empty(t *testing.T) {
	setupTestStreamServices()
	handler := ListStreamClients()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stream/clients", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var clients []interface{}
	err := json.NewDecoder(w.Body).Decode(&clients)
	require.NoError(t, err)
	assert.Empty(t, clients)
}

func TestCreateStreamClient_Success(t *testing.T) {
	setupTestStreamServices()
	handler := CreateStreamClient()

	reqBody := CreateStreamClientRequest{
		Name: "test-client",
		Config: config.ClientConfig{
			URL:           "ws://localhost:15560/ws",
			Interfaces:    []string{"demo.Counter", "demo.Calculator"},
			Enabled:       true,
			AutoReconnect: true,
		},
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/clients", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, "test-client", response["name"])
	// Note: Status can be "disconnected" or "error" depending on whether connection attempt was made
	status := response["status"].(string)
	assert.Contains(t, []string{"disconnected", "error"}, status)
}

func TestCreateStreamClient_MissingName(t *testing.T) {
	setupTestStreamServices()
	handler := CreateStreamClient()

	reqBody := CreateStreamClientRequest{
		Config: config.ClientConfig{
			URL: "ws://localhost:15560/ws",
		},
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/clients", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Error, "client name is required")
}

func TestCreateStreamClient_MissingURL(t *testing.T) {
	setupTestStreamServices()
	handler := CreateStreamClient()

	reqBody := CreateStreamClientRequest{
		Name: "test-client",
		Config: config.ClientConfig{
			Interfaces: []string{"demo.Counter"},
		},
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/clients", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Error, "URL is required")
}

func TestCreateStreamClient_InvalidJSON(t *testing.T) {
	setupTestStreamServices()
	handler := CreateStreamClient()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/clients", bytes.NewReader([]byte("invalid json")))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Error, "invalid")
}

func TestGetStreamClient_Success(t *testing.T) {
	services := setupTestStreamServices()

	// Create a client first
	err := services.ClientManager.AddClient("test-client", config.ClientConfig{
		URL:        "ws://localhost:15560/ws",
		Interfaces: []string{"demo.Counter"},
		Enabled:    true,
	})
	require.NoError(t, err)

	handler := GetStreamClient()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stream/clients/test-client", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "test-client")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, "test-client", response["name"])
	assert.Equal(t, "ws://localhost:15560/ws", response["url"])
}

func TestGetStreamClient_NotFound(t *testing.T) {
	setupTestStreamServices()
	handler := GetStreamClient()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stream/clients/nonexistent", nil)
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

func TestUpdateStreamClient_Success(t *testing.T) {
	services := setupTestStreamServices()

	// Create a client first
	err := services.ClientManager.AddClient("test-client", config.ClientConfig{
		URL:        "ws://localhost:15560/ws",
		Interfaces: []string{"demo.Counter"},
		Enabled:    true,
	})
	require.NoError(t, err)

	handler := UpdateStreamClient()

	updatedConfig := config.ClientConfig{
		URL:        "ws://localhost:15561/ws",
		Interfaces: []string{"demo.Counter", "demo.Calculator"},
		Enabled:    true,
	}
	body, err := json.Marshal(updatedConfig)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/stream/clients/test-client", bytes.NewReader(body))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "test-client")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, "test-client", response["name"])
	assert.Equal(t, "ws://localhost:15561/ws", response["url"])
}

func TestUpdateStreamClient_NotFound(t *testing.T) {
	setupTestStreamServices()
	handler := UpdateStreamClient()

	updatedConfig := config.ClientConfig{
		URL:     "ws://localhost:15561/ws",
		Enabled: true,
	}
	body, err := json.Marshal(updatedConfig)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/stream/clients/nonexistent", bytes.NewReader(body))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "nonexistent")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteStreamClient_Success(t *testing.T) {
	services := setupTestStreamServices()

	// Create a client first
	err := services.ClientManager.AddClient("test-client", config.ClientConfig{
		URL:     "ws://localhost:15560/ws",
		Enabled: true,
	})
	require.NoError(t, err)

	handler := DeleteStreamClient()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/stream/clients/test-client", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "test-client")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// Verify client was deleted
	_, err = services.ClientManager.GetClient("test-client")
	assert.Error(t, err)
}

func TestDeleteStreamClient_NotFound(t *testing.T) {
	setupTestStreamServices()
	handler := DeleteStreamClient()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/stream/clients/nonexistent", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "nonexistent")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestConnectStreamClient_NotFound(t *testing.T) {
	setupTestStreamServices()
	handler := ConnectStreamClient()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/clients/nonexistent/connect", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "nonexistent")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDisconnectStreamClient_NotFound(t *testing.T) {
	setupTestStreamServices()
	handler := DisconnectStreamClient()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/clients/nonexistent/disconnect", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "nonexistent")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListStreamClients_Multiple(t *testing.T) {
	services := setupTestStreamServices()

	// Create multiple clients
	err := services.ClientManager.AddClient("client1", config.ClientConfig{
		URL:     "ws://localhost:15560/ws",
		Enabled: true,
	})
	require.NoError(t, err)

	err = services.ClientManager.AddClient("client2", config.ClientConfig{
		URL:     "ws://localhost:15561/ws",
		Enabled: true,
	})
	require.NoError(t, err)

	handler := ListStreamClients()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stream/clients", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var clients []map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&clients)
	require.NoError(t, err)
	assert.Len(t, clients, 2)

	// Verify client names
	names := []string{}
	for _, c := range clients {
		names = append(names, c["name"].(string))
	}
	assert.Contains(t, names, "client1")
	assert.Contains(t, names, "client2")
}

func TestCreateStreamClient_DuplicateName(t *testing.T) {
	services := setupTestStreamServices()

	// Create first client
	err := services.ClientManager.AddClient("test-client", config.ClientConfig{
		URL:     "ws://localhost:15560/ws",
		Enabled: true,
	})
	require.NoError(t, err)

	// Try to create another with same name
	handler := CreateStreamClient()

	reqBody := CreateStreamClientRequest{
		Name: "test-client",
		Config: config.ClientConfig{
			URL:     "ws://localhost:15561/ws",
			Enabled: true,
		},
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/clients", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Error, "already exists")
}

func TestCreateStreamClient_WithInterfaces(t *testing.T) {
	setupTestStreamServices()
	handler := CreateStreamClient()

	reqBody := CreateStreamClientRequest{
		Name: "test-client",
		Config: config.ClientConfig{
			URL: "ws://localhost:15560/ws",
			Interfaces: []string{
				"demo.Counter",
				"demo.Calculator",
				"demo.Timer",
			},
			Enabled:       true,
			AutoReconnect: true,
		},
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/stream/clients", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	interfaces := response["interfaces"].([]interface{})
	assert.Len(t, interfaces, 3)
	assert.Contains(t, interfaces, "demo.Counter")
	assert.Contains(t, interfaces, "demo.Calculator")
	assert.Contains(t, interfaces, "demo.Timer")
}
