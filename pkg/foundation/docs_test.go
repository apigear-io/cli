package foundation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDocumentType(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{"IDL file", "demo.idl", "module"},
		{"module YAML", "demo.module.yaml", "module"},
		{"solution YAML", "demo.solution.yaml", "solution"},
		{"JS simulation", "demo.sim.js", "simulation"},
		{"JS file", "script.js", "simulation"},
		{"unknown type", "readme.txt", "unknown"},
		{"no extension", "file", "unknown"},
		{"full path IDL", "/path/to/demo.idl", "module"},
		{"full path module", "/path/to/demo.module.yaml", "module"},
		{"full path solution", "/path/to/demo.solution.yaml", "solution"},
		{"full path JS", "/path/to/demo.js", "simulation"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetDocumentType(tt.path)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseJson(t *testing.T) {
	t.Run("parse valid JSON", func(t *testing.T) {
		type TestStruct struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		}

		data := []byte(`{"name": "test", "value": 42}`)
		var result TestStruct

		err := ParseJson(data, &result)
		require.NoError(t, err)
		assert.Equal(t, "test", result.Name)
		assert.Equal(t, 42, result.Value)
	})

	t.Run("parse invalid JSON", func(t *testing.T) {
		data := []byte(`{invalid json}`)
		var result map[string]interface{}

		err := ParseJson(data, &result)
		assert.Error(t, err)
	})

	t.Run("parse empty JSON", func(t *testing.T) {
		data := []byte(`{}`)
		var result map[string]interface{}

		err := ParseJson(data, &result)
		require.NoError(t, err)
		assert.Empty(t, result)
	})

	t.Run("parse JSON array", func(t *testing.T) {
		data := []byte(`[1, 2, 3]`)
		var result []int

		err := ParseJson(data, &result)
		require.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, result)
	})
}

func TestParseYaml(t *testing.T) {
	t.Run("parse valid YAML", func(t *testing.T) {
		type TestStruct struct {
			Name  string `yaml:"name"`
			Value int    `yaml:"value"`
		}

		data := []byte(`name: test
value: 42`)
		var result TestStruct

		err := ParseYaml(data, &result)
		require.NoError(t, err)
		assert.Equal(t, "test", result.Name)
		assert.Equal(t, 42, result.Value)
	})

	t.Run("parse invalid YAML", func(t *testing.T) {
		data := []byte(`invalid: yaml: syntax:`)
		var result map[string]interface{}

		err := ParseYaml(data, &result)
		assert.Error(t, err)
	})

	t.Run("parse empty YAML", func(t *testing.T) {
		data := []byte(``)
		var result map[string]interface{}

		err := ParseYaml(data, &result)
		require.NoError(t, err)
		// Empty YAML returns nil map
	})

	t.Run("parse YAML with nested structure", func(t *testing.T) {
		type Config struct {
			Server struct {
				Host string `yaml:"host"`
				Port int    `yaml:"port"`
			} `yaml:"server"`
		}

		data := []byte(`server:
  host: localhost
  port: 8080`)
		var result Config

		err := ParseYaml(data, &result)
		require.NoError(t, err)
		assert.Equal(t, "localhost", result.Server.Host)
		assert.Equal(t, 8080, result.Server.Port)
	})

	t.Run("parse YAML array", func(t *testing.T) {
		data := []byte(`- item1
- item2
- item3`)
		var result []string

		err := ParseYaml(data, &result)
		require.NoError(t, err)
		assert.Equal(t, []string{"item1", "item2", "item3"}, result)
	})
}
