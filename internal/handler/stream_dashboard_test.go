package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/apigear-io/cli/pkg/stream/config"
)

func TestGetStreamDashboard_Empty(t *testing.T) {
	setupTestStreamServices()
	handler := GetStreamDashboard()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stream/dashboard", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var stats StreamDashboardStats
	err := json.NewDecoder(w.Body).Decode(&stats)
	require.NoError(t, err)

	// Empty dashboard should have zeros
	assert.Equal(t, 0, stats.Proxies.Total)
	assert.Equal(t, 0, stats.Proxies.Running)
	assert.Equal(t, 0, stats.Proxies.Stopped)
	assert.Equal(t, 0, stats.Clients.Total)
	assert.Equal(t, 0, stats.Clients.Connected)
	assert.Equal(t, 0, stats.Clients.Disconnected)
	assert.Equal(t, int64(0), stats.Messages.Total)
	assert.Equal(t, float64(0), stats.Messages.Rate)
}

func TestGetStreamDashboard_WithProxies(t *testing.T) {
	t.Skip("Skipping due to race condition in proxy.Start() - TODO: fix in proxy package")

	services := setupTestStreamServices()

	// Create proxies
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

	// Start one proxy
	err = services.ProxyManager.StartProxy("proxy1")
	require.NoError(t, err)
	defer services.ProxyManager.StopProxy("proxy1")

	handler := GetStreamDashboard()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stream/dashboard", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var stats StreamDashboardStats
	err = json.NewDecoder(w.Body).Decode(&stats)
	require.NoError(t, err)

	assert.Equal(t, 2, stats.Proxies.Total)
	assert.Equal(t, 1, stats.Proxies.Running)
	assert.Equal(t, 1, stats.Proxies.Stopped)
}

func TestGetStreamDashboard_WithClients(t *testing.T) {
	services := setupTestStreamServices()

	// Create clients
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

	handler := GetStreamDashboard()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stream/dashboard", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var stats StreamDashboardStats
	err = json.NewDecoder(w.Body).Decode(&stats)
	require.NoError(t, err)

	assert.Equal(t, 2, stats.Clients.Total)
	// Clients start disconnected
	assert.Equal(t, 0, stats.Clients.Connected)
	assert.Equal(t, 2, stats.Clients.Disconnected)
}

func TestGetStreamDashboard_MixedState(t *testing.T) {
	t.Skip("Skipping due to race condition in proxy.Start() - TODO: fix in proxy package")

	services := setupTestStreamServices()

	// Create proxies
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

	err = services.ProxyManager.AddProxy("proxy3", config.ProxyConfig{
		Listen: "ws://localhost:15552/ws",
		Mode:   "echo",
	})
	require.NoError(t, err)

	// Start two proxies
	err = services.ProxyManager.StartProxy("proxy1")
	require.NoError(t, err)
	defer services.ProxyManager.StopProxy("proxy1")

	err = services.ProxyManager.StartProxy("proxy2")
	require.NoError(t, err)
	defer services.ProxyManager.StopProxy("proxy2")

	// Create clients
	err = services.ClientManager.AddClient("client1", config.ClientConfig{
		URL:     "ws://localhost:15560/ws",
		Enabled: true,
	})
	require.NoError(t, err)

	err = services.ClientManager.AddClient("client2", config.ClientConfig{
		URL:     "ws://localhost:15561/ws",
		Enabled: true,
	})
	require.NoError(t, err)

	handler := GetStreamDashboard()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stream/dashboard", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var stats StreamDashboardStats
	err = json.NewDecoder(w.Body).Decode(&stats)
	require.NoError(t, err)

	// Verify proxy stats
	assert.Equal(t, 3, stats.Proxies.Total)
	assert.Equal(t, 2, stats.Proxies.Running)
	assert.Equal(t, 1, stats.Proxies.Stopped)

	// Verify client stats
	assert.Equal(t, 2, stats.Clients.Total)
	assert.Equal(t, 0, stats.Clients.Connected)
	assert.Equal(t, 2, stats.Clients.Disconnected)
}
