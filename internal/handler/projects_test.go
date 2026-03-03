package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/apigear-io/cli/pkg/foundation"
	"github.com/apigear-io/cli/pkg/foundation/config"
	"github.com/apigear-io/cli/pkg/orchestration/project"
)

// setupTestProject creates a test project in a temporary directory
func setupTestProject(t *testing.T, name string) (string, string) {
	tempDir, err := os.MkdirTemp("", "project-test-*")
	require.NoError(t, err)

	projectPath := foundation.Join(tempDir, name)
	info, err := project.InitProject(projectPath)
	require.NoError(t, err)
	require.NotNil(t, info)

	// Add to recent entries for testing
	err = config.AppendRecentEntry(projectPath)
	require.NoError(t, err)

	return tempDir, projectPath
}

// cleanupTestProject removes the test directory
func cleanupTestProject(t *testing.T, tempDir string) {
	err := os.RemoveAll(tempDir)
	if err != nil {
		t.Logf("Failed to cleanup test directory: %v", err)
	}
}

func TestGetProjectDirectories_Success(t *testing.T) {
	handler := GetProjectDirectories()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects/directories", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response ProjectDirectoriesResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	// Should have at least homeDir or workingDir
	assert.True(t, response.HomeDir != "" || response.WorkingDir != "")

	// Suggestions should be valid
	assert.NotNil(t, response.Suggestions)
}

func TestBrowseDirectories_DefaultHome(t *testing.T) {
	handler := BrowseDirectories()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects/browse", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response DirectoryListResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	// Should have current path
	assert.NotEmpty(t, response.CurrentPath)

	// Should have directories list (even if empty)
	assert.NotNil(t, response.Directories)
}

func TestBrowseDirectories_WithPath(t *testing.T) {
	handler := BrowseDirectories()

	// Use /tmp which should exist on most systems
	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects/browse?path=/tmp", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response DirectoryListResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, "/tmp", response.CurrentPath)
	assert.NotNil(t, response.Directories)
}

func TestBrowseDirectories_NotFound(t *testing.T) {
	handler := BrowseDirectories()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects/browse?path=/nonexistent/path/12345", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Message, "not found")
}

func TestListRecentProjects_ReturnsSuccessfully(t *testing.T) {
	// This test verifies the API returns successfully
	// Count may vary depending on existing entries from other tests
	handler := ListRecentProjects()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects/recent", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response ProjectListResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, response.Count, 0)
	assert.Equal(t, response.Count, len(response.Projects))
}

func TestListRecentProjects_WithProjects(t *testing.T) {
	// Setup test projects
	tempDir1, projectPath1 := setupTestProject(t, "project1")
	defer cleanupTestProject(t, tempDir1)
	defer config.RemoveRecentEntry(projectPath1)

	tempDir2, projectPath2 := setupTestProject(t, "project2")
	defer cleanupTestProject(t, tempDir2)
	defer config.RemoveRecentEntry(projectPath2)

	handler := ListRecentProjects()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects/recent", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response ProjectListResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, response.Count, 2)
	assert.GreaterOrEqual(t, len(response.Projects), 2)

	// Check that at least one of our test projects exists in the list
	projectPaths := make(map[string]bool)
	for _, p := range response.Projects {
		projectPaths[p.Path] = true
	}
	// At least one should be present (may not be both if one failed)
	assert.True(t, projectPaths[projectPath1] || projectPaths[projectPath2],
		"Expected at least one test project in recent projects")
}

func TestCreateProject_Success(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "project-create-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	handler := CreateProject()

	reqBody := CreateProjectRequest{
		Name: "myproject",
		Path: tempDir,
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var info project.ProjectInfo
	err = json.NewDecoder(w.Body).Decode(&info)
	require.NoError(t, err)
	assert.Equal(t, "myproject", info.Name)
	assert.Equal(t, foundation.Join(tempDir, "myproject"), info.Path)
	assert.NotEmpty(t, info.Documents)

	// Verify project was created on disk
	projectPath := foundation.Join(tempDir, "myproject")
	_, err = os.Stat(projectPath)
	assert.NoError(t, err)

	// Verify apigear directory exists
	apigearPath := foundation.Join(projectPath, "apigear")
	_, err = os.Stat(apigearPath)
	assert.NoError(t, err)

	// Verify demo files exist
	demoModule := foundation.Join(apigearPath, "demo.module.yaml")
	_, err = os.Stat(demoModule)
	assert.NoError(t, err)
}

func TestCreateProject_EmptyName(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "project-create-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	handler := CreateProject()

	reqBody := CreateProjectRequest{
		Name: "",
		Path: tempDir,
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Message, "name")
}

func TestCreateProject_EmptyPath(t *testing.T) {
	handler := CreateProject()

	reqBody := CreateProjectRequest{
		Name: "myproject",
		Path: "",
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Message, "path")
}

func TestCreateProject_ParentNotExists(t *testing.T) {
	handler := CreateProject()

	reqBody := CreateProjectRequest{
		Name: "myproject",
		Path: "/nonexistent/path/that/does/not/exist",
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Message, "not exist")
}

func TestCreateProject_AlreadyExists(t *testing.T) {
	tempDir, projectPath := setupTestProject(t, "existing")
	defer cleanupTestProject(t, tempDir)

	handler := CreateProject()

	reqBody := CreateProjectRequest{
		Name: "existing",
		Path: tempDir,
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)

	var response ErrorResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Message, "already exists")

	// Cleanup
	config.RemoveRecentEntry(projectPath)
}

func TestGetProject_Success(t *testing.T) {
	tempDir, projectPath := setupTestProject(t, "gettest")
	defer cleanupTestProject(t, tempDir)
	defer config.RemoveRecentEntry(projectPath)

	handler := GetProject()

	encodedPath := url.QueryEscape(projectPath)
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/projects/get?path=%s", encodedPath), nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var info project.ProjectInfo
	err := json.NewDecoder(w.Body).Decode(&info)
	require.NoError(t, err)
	assert.Equal(t, "gettest", info.Name)
	assert.Equal(t, projectPath, info.Path)
	assert.NotEmpty(t, info.Documents)
}

func TestGetProject_MissingPath(t *testing.T) {
	handler := GetProject()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects/get", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Message, "required")
}

func TestGetProject_NotFound(t *testing.T) {
	handler := GetProject()

	nonExistentPath := "/tmp/nonexistent-project-12345"
	encodedPath := url.QueryEscape(nonExistentPath)
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/projects/get?path=%s", encodedPath), nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Message, "not found")
}

func TestDeleteProject_Success(t *testing.T) {
	tempDir, projectPath := setupTestProject(t, "deletetest")
	defer cleanupTestProject(t, tempDir)

	// Verify project exists before delete
	_, err := os.Stat(projectPath)
	require.NoError(t, err)

	handler := DeleteProject()

	encodedPath := url.QueryEscape(projectPath)
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/projects?path=%s", encodedPath), nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// Verify project was deleted
	_, err = os.Stat(projectPath)
	assert.True(t, os.IsNotExist(err))
}

func TestDeleteProject_MissingPath(t *testing.T) {
	handler := DeleteProject()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/projects", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Message, "required")
}

func TestDeleteProject_NotFound(t *testing.T) {
	handler := DeleteProject()

	nonExistentPath := "/tmp/nonexistent-project-12345"
	encodedPath := url.QueryEscape(nonExistentPath)
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/projects?path=%s", encodedPath), nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Message, "not found")
}

func TestDeleteProject_NotValidProject(t *testing.T) {
	// Create a directory without apigear subdirectory
	tempDir, err := os.MkdirTemp("", "not-a-project-*")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	handler := DeleteProject()

	encodedPath := url.QueryEscape(tempDir)
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/projects?path=%s", encodedPath), nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Message, "apigear")
}

func TestCreateProject_InvalidJSON(t *testing.T) {
	handler := CreateProject()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", bytes.NewReader([]byte("invalid json")))
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)
	assert.Contains(t, response.Message, "Invalid")
}

func TestGetProject_DocumentsHaveLowercaseFields(t *testing.T) {
	// Create a test project
	tempDir, projectPath := setupTestProject(t, "json-test")
	defer cleanupTestProject(t, tempDir)
	defer config.RemoveRecentEntry(projectPath)

	handler := GetProject()

	encodedPath := url.QueryEscape(projectPath)
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/projects/get?path=%s", encodedPath), nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Decode as raw JSON to check field names
	var rawResult map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &rawResult)
	require.NoError(t, err)

	// Check that documents array exists
	documents, ok := rawResult["documents"].([]interface{})
	require.True(t, ok, "documents field should be an array")
	require.NotEmpty(t, documents, "documents should not be empty")

	// Check first document has lowercase fields
	doc := documents[0].(map[string]interface{})
	assert.NotNil(t, doc["name"], "should have lowercase 'name' field")
	assert.NotNil(t, doc["path"], "should have lowercase 'path' field")
	assert.NotNil(t, doc["type"], "should have lowercase 'type' field")

	// Verify no capitalized fields
	assert.Nil(t, doc["Name"], "should NOT have capitalized 'Name' field")
	assert.Nil(t, doc["Path"], "should NOT have capitalized 'Path' field")
	assert.Nil(t, doc["Type"], "should NOT have capitalized 'Type' field")
}
