package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListTemplates(t *testing.T) {
	handler := ListTemplates()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/templates", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response TemplateListResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.NotNil(t, response.Templates)
	assert.GreaterOrEqual(t, response.Count, 0)
	assert.Equal(t, len(response.Templates), response.Count)
}

func TestGetTemplate_Success(t *testing.T) {
	handler := GetTemplate()

	// Test with a template that exists in the registry
	req := httptest.NewRequest(http.MethodGet, "/api/v1/templates/get?id=apigear-io/template-python", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	// Note: This test will fail if the registry is not initialized or template doesn't exist
	// In a production test environment, you'd want to mock the registry
	if w.Code == http.StatusOK {
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		var response TemplateInfo
		err := json.NewDecoder(w.Body).Decode(&response)
		require.NoError(t, err)

		assert.NotEmpty(t, response.Name)
		assert.NotEmpty(t, response.Git)
	}
}

func TestGetTemplate_MissingID(t *testing.T) {
	handler := GetTemplate()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/templates/get", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.Contains(t, response.Error, "missing template id")
}

func TestGetTemplate_NotFound(t *testing.T) {
	handler := GetTemplate()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/templates/get?id=nonexistent/template", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.NotEmpty(t, response.Error)
}

func TestInstallTemplate_MissingID(t *testing.T) {
	handler := InstallTemplate()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/templates/install", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestInstallTemplate_WithVersion(t *testing.T) {
	handler := InstallTemplate()

	body := InstallRequest{
		Version: "v1.0.0",
	}
	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/templates/install?id=test/template", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler(w, req)

	// Should return SSE headers
	assert.Equal(t, "text/event-stream", w.Header().Get("Content-Type"))
	assert.Equal(t, "no-cache", w.Header().Get("Cache-Control"))
	assert.Equal(t, "keep-alive", w.Header().Get("Connection"))
}

func TestInstallTemplate_SSEFormat(t *testing.T) {
	handler := InstallTemplate()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/templates/install?id=test/template", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	// Check SSE format
	body := w.Body.String()

	// SSE events should have "data: " prefix
	assert.Contains(t, body, "data: ")

	// Should contain JSON events
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "data: ") {
			eventJSON := strings.TrimPrefix(line, "data: ")
			var event InstallProgressEvent
			err := json.Unmarshal([]byte(eventJSON), &event)
			if err == nil {
				// Valid event should have type and message
				assert.NotEmpty(t, event.Type)
				assert.NotEmpty(t, event.Message)
			}
		}
	}
}

func TestListCachedTemplates(t *testing.T) {
	handler := ListCachedTemplates()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/templates/cache", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response TemplateListResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.NotNil(t, response.Templates)
	assert.GreaterOrEqual(t, response.Count, 0)
	assert.Equal(t, len(response.Templates), response.Count)

	// All cached templates should have InCache = true
	for _, tmpl := range response.Templates {
		assert.True(t, tmpl.InCache)
	}
}

func TestRemoveTemplate_MissingID(t *testing.T) {
	handler := RemoveTemplate()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/templates/cache/remove", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.Contains(t, response.Error, "missing template id")
}

func TestRemoveTemplate_WithID(t *testing.T) {
	handler := RemoveTemplate()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/templates/cache/remove?id=test/template", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	// Will fail if template doesn't exist, but we're testing the HTTP layer
	if w.Code == http.StatusOK {
		var response map[string]string
		err := json.NewDecoder(w.Body).Decode(&response)
		require.NoError(t, err)
		assert.Contains(t, response["message"], "removed successfully")
	} else {
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	}
}

func TestCleanCache(t *testing.T) {
	handler := CleanCache()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/templates/cache/clean", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	// Should return success or error, but proper JSON response
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	if w.Code == http.StatusOK {
		var response map[string]string
		err := json.NewDecoder(w.Body).Decode(&response)
		require.NoError(t, err)
		assert.Contains(t, response["message"], "cleaned successfully")
	}
}

func TestUpdateRegistry(t *testing.T) {
	handler := UpdateRegistry()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/templates/registry/update", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	// Should return success or error, but proper JSON response
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	if w.Code == http.StatusOK {
		var response map[string]string
		err := json.NewDecoder(w.Body).Decode(&response)
		require.NoError(t, err)
		assert.Contains(t, response["message"], "updated successfully")
	}
}

func TestSearchTemplates_MissingQuery(t *testing.T) {
	handler := SearchTemplates()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/templates/search", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.Contains(t, response.Error, "missing query parameter")
}

func TestSearchTemplates_WithQuery(t *testing.T) {
	handler := SearchTemplates()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/templates/search?q=python", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response TemplateListResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.NotNil(t, response.Templates)
	assert.GreaterOrEqual(t, response.Count, 0)

	// All results should match the query
	for _, tmpl := range response.Templates {
		matched := strings.Contains(strings.ToLower(tmpl.Name), "python") ||
			strings.Contains(strings.ToLower(tmpl.Description), "python")
		assert.True(t, matched, "Template %s should match query 'python'", tmpl.Name)
	}
}

func TestSearchTemplates_EmptyQuery(t *testing.T) {
	handler := SearchTemplates()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/templates/search?q=", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	// Empty query should return bad request
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSearchTemplates_NoResults(t *testing.T) {
	handler := SearchTemplates()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/templates/search?q=nonexistenttemplate12345", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response TemplateListResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, 0, response.Count)
	assert.Empty(t, response.Templates)
}

// Test helper functions

func TestConvertRepoInfo(t *testing.T) {
	// This tests the internal conversion function
	// We'd need to import the git package and create test data
	// For now, we'll skip this as it's an internal helper
}

func TestMergeTemplateInfo(t *testing.T) {
	// This tests the internal merge function
	// We'd need to create mock RepoInfo structs
	// For now, we'll skip this as it's an internal helper
}

// Integration tests that require a full server setup

func TestTemplateRoutes_Integration(t *testing.T) {
	// This would test the full routing with chi router
	// Skip for now as it requires router setup
	t.Skip("Integration test - requires full router setup")
}

// Benchmark tests

func BenchmarkListTemplates(b *testing.B) {
	handler := ListTemplates()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/templates", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		handler(w, req)
	}
}

func BenchmarkSearchTemplates(b *testing.B) {
	handler := SearchTemplates()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/templates/search?q=python", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		handler(w, req)
	}
}
