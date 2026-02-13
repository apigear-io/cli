package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apigear-io/cli/pkg/runtime/monitoring"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMonitor_Success(t *testing.T) {
	events := []*monitoring.Event{
		{Type: monitoring.TypeCall, Symbol: "test.method"},
	}
	body, _ := json.Marshal(events)

	req := httptest.NewRequest(http.MethodPost, "/monitor/test-source", bytes.NewReader(body))
	w := httptest.NewRecorder()

	// Setup Chi URL params
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("source", "test-source")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := Monitor()
	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response MonitorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, "ok", response.Status)
	assert.Equal(t, 1, response.EventsProcessed)
	assert.False(t, response.Timestamp.IsZero())
}

func TestMonitor_MissingSource(t *testing.T) {
	events := []*monitoring.Event{
		{Type: monitoring.TypeCall, Symbol: "test.method"},
	}
	body, _ := json.Marshal(events)

	req := httptest.NewRequest(http.MethodPost, "/monitor/", bytes.NewReader(body))
	w := httptest.NewRecorder()

	// Setup Chi URL params with empty source
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("source", "")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := Monitor()
	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.Contains(t, response.Error, "source id is required")
}

func TestMonitor_InvalidJSON(t *testing.T) {
	invalidJSON := []byte(`{invalid json}`)

	req := httptest.NewRequest(http.MethodPost, "/monitor/test-source", bytes.NewReader(invalidJSON))
	w := httptest.NewRecorder()

	// Setup Chi URL params
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("source", "test-source")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := Monitor()
	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.Contains(t, response.Message, "failed to decode event array")
}

func TestMonitor_EventEnrichment(t *testing.T) {
	// Create event without ID and timestamp
	events := []*monitoring.Event{
		{Type: monitoring.TypeSignal, Symbol: "test.signal"},
	}
	body, _ := json.Marshal(events)

	req := httptest.NewRequest(http.MethodPost, "/monitor/test-app", bytes.NewReader(body))
	w := httptest.NewRecorder()

	// Setup Chi URL params
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("source", "test-app")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// Hook to capture the enriched event
	var capturedEvent *monitoring.Event
	removeHook := monitoring.Emitter.AddHook(func(event *monitoring.Event) {
		capturedEvent = event
	})
	defer removeHook()

	handler := Monitor()
	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Verify event was enriched
	require.NotNil(t, capturedEvent)
	assert.Equal(t, "test-app", capturedEvent.Source)
	assert.NotEmpty(t, capturedEvent.Id)
	assert.False(t, capturedEvent.Timestamp.IsZero())
}

func TestMonitor_EmptyArray(t *testing.T) {
	events := []*monitoring.Event{}
	body, _ := json.Marshal(events)

	req := httptest.NewRequest(http.MethodPost, "/monitor/test-source", bytes.NewReader(body))
	w := httptest.NewRecorder()

	// Setup Chi URL params
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("source", "test-source")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := Monitor()
	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response MonitorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, "ok", response.Status)
	assert.Equal(t, 0, response.EventsProcessed)
}
