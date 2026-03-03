package project

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDocumentInfoJSON(t *testing.T) {
	doc := DocumentInfo{
		Name: "demo.module.yaml",
		Path: "/path/to/demo.module.yaml",
		Type: "module",
	}

	// Marshal to JSON
	data, err := json.Marshal(doc)
	require.NoError(t, err)

	// Verify lowercase field names
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	require.NoError(t, err)

	assert.Equal(t, "demo.module.yaml", result["name"])
	assert.Equal(t, "/path/to/demo.module.yaml", result["path"])
	assert.Equal(t, "module", result["type"])

	// Verify no capitalized fields
	assert.Nil(t, result["Name"])
	assert.Nil(t, result["Path"])
	assert.Nil(t, result["Type"])
}

func TestProjectInfoJSON(t *testing.T) {
	project := ProjectInfo{
		Name: "test-project",
		Path: "/path/to/project",
		Documents: []DocumentInfo{
			{
				Name: "demo.module.yaml",
				Path: "/path/to/demo.module.yaml",
				Type: "module",
			},
		},
	}

	// Marshal to JSON
	data, err := json.Marshal(project)
	require.NoError(t, err)

	// Unmarshal back
	var result ProjectInfo
	err = json.Unmarshal(data, &result)
	require.NoError(t, err)

	assert.Equal(t, "test-project", result.Name)
	assert.Equal(t, "/path/to/project", result.Path)
	assert.Len(t, result.Documents, 1)
	assert.Equal(t, "demo.module.yaml", result.Documents[0].Name)
	assert.Equal(t, "module", result.Documents[0].Type)
}
