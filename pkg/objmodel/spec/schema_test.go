package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDocumentType(t *testing.T) {
	tests := []struct {
		name string
		want DocumentType
	}{
		{
			name: "demo.idl",
			want: DocumentTypeModule,
		},
		{
			name: "demo.module.yaml",
			want: DocumentTypeModule,
		},
		{
			name: "demo.module.json",
			want: DocumentTypeModule,
		},
		{
			name: "demo.solution.yaml",
			want: DocumentTypeSolution,
		},
		{
			name: "demo.solution.json",
			want: DocumentTypeSolution,
		},
		{
			name: "rules.yaml",
			want: DocumentTypeRules,
		},
		{
			name: "rules.json",
			want: DocumentTypeRules,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDocumentType(tt.name)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDocumentTypeFromFileName(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
		wantErr  bool
	}{
		{
			name:     "module yaml",
			filename: "demo.module.yaml",
			want:     "module",
			wantErr:  false,
		},
		{
			name:     "solution json",
			filename: "demo.solution.json",
			want:     "solution",
			wantErr:  false,
		},
		{
			name:     "rules yaml",
			filename: "rules.yaml",
			want:     "rules",
			wantErr:  false,
		},
		{
			name:     "idl file",
			filename: "demo.idl",
			want:     "module",
			wantErr:  false,
		},
		{
			name:     "invalid filename - no extension",
			filename: "demo",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "simple filename with extension",
			filename: "demo.yaml",
			want:     "demo",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DocumentTypeFromFileName(tt.filename)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestYamlToJson(t *testing.T) {
	t.Run("converts valid yaml to json", func(t *testing.T) {
		yamlData := []byte(`
name: test
version: "1.0"
count: 42
enabled: true
`)
		jsonData, err := YamlToJson(yamlData)
		assert.NoError(t, err)
		assert.NotEmpty(t, jsonData)
		assert.Contains(t, string(jsonData), `"name"`)
		assert.Contains(t, string(jsonData), `"test"`)
		assert.Contains(t, string(jsonData), `"version"`)
		assert.Contains(t, string(jsonData), `"1.0"`)
		assert.Contains(t, string(jsonData), `"count"`)
		assert.Contains(t, string(jsonData), `42`)
	})

	t.Run("handles empty yaml", func(t *testing.T) {
		yamlData := []byte(`{}`)
		jsonData, err := YamlToJson(yamlData)
		assert.NoError(t, err)
		assert.NotEmpty(t, jsonData)
	})

	t.Run("handles nested structures", func(t *testing.T) {
		yamlData := []byte(`
parent:
  child:
    name: nested
    value: 123
`)
		jsonData, err := YamlToJson(yamlData)
		assert.NoError(t, err)
		assert.Contains(t, string(jsonData), `"parent"`)
		assert.Contains(t, string(jsonData), `"child"`)
		assert.Contains(t, string(jsonData), `"nested"`)
	})

	t.Run("handles arrays", func(t *testing.T) {
		yamlData := []byte(`
items:
  - name: first
  - name: second
  - name: third
`)
		jsonData, err := YamlToJson(yamlData)
		assert.NoError(t, err)
		assert.Contains(t, string(jsonData), `"items"`)
		assert.Contains(t, string(jsonData), `"first"`)
		assert.Contains(t, string(jsonData), `"second"`)
	})

	t.Run("returns error for invalid yaml", func(t *testing.T) {
		yamlData := []byte(`
invalid yaml:
  - unclosed bracket: [
  - unmatched quote: "test
`)
		_, err := YamlToJson(yamlData)
		assert.Error(t, err)
	})
}

func TestJsonToYaml(t *testing.T) {
	t.Run("converts valid json to yaml", func(t *testing.T) {
		jsonData := []byte(`{
  "name": "test",
  "version": "1.0",
  "count": 42,
  "enabled": true
}`)
		yamlData, err := JsonToYaml(jsonData)
		assert.NoError(t, err)
		assert.NotEmpty(t, yamlData)
		assert.Contains(t, string(yamlData), "name:")
		assert.Contains(t, string(yamlData), "test")
		assert.Contains(t, string(yamlData), "version:")
		assert.Contains(t, string(yamlData), "count:")
	})

	t.Run("handles empty json object", func(t *testing.T) {
		jsonData := []byte(`{}`)
		yamlData, err := JsonToYaml(jsonData)
		assert.NoError(t, err)
		assert.NotEmpty(t, yamlData)
	})

	t.Run("handles nested structures", func(t *testing.T) {
		jsonData := []byte(`{
  "parent": {
    "child": {
      "name": "nested",
      "value": 123
    }
  }
}`)
		yamlData, err := JsonToYaml(jsonData)
		assert.NoError(t, err)
		assert.Contains(t, string(yamlData), "parent:")
		assert.Contains(t, string(yamlData), "child:")
		assert.Contains(t, string(yamlData), "nested")
	})

	t.Run("handles arrays", func(t *testing.T) {
		jsonData := []byte(`{
  "items": [
    {"name": "first"},
    {"name": "second"},
    {"name": "third"}
  ]
}`)
		yamlData, err := JsonToYaml(jsonData)
		assert.NoError(t, err)
		assert.Contains(t, string(yamlData), "items:")
		assert.Contains(t, string(yamlData), "first")
		assert.Contains(t, string(yamlData), "second")
	})

	t.Run("returns error for invalid json", func(t *testing.T) {
		jsonData := []byte(`{
  "invalid": "json",
  "missing": "closing brace"
`)
		_, err := JsonToYaml(jsonData)
		assert.Error(t, err)
	})
}

func TestLoadSchema(t *testing.T) {
	t.Run("loads module schema", func(t *testing.T) {
		schema, err := LoadSchema(DocumentTypeModule)
		assert.NoError(t, err)
		assert.NotNil(t, schema)
	})

	t.Run("loads solution schema", func(t *testing.T) {
		schema, err := LoadSchema(DocumentTypeSolution)
		assert.NoError(t, err)
		assert.NotNil(t, schema)
	})

	t.Run("loads rules schema", func(t *testing.T) {
		schema, err := LoadSchema(DocumentTypeRules)
		assert.NoError(t, err)
		assert.NotNil(t, schema)
	})

	t.Run("panics for unknown document type", func(t *testing.T) {
		assert.Panics(t, func() {
			_, _ = LoadSchema(DocumentTypeUnknown)
		})
	})

	t.Run("panics for invalid document type", func(t *testing.T) {
		assert.Panics(t, func() {
			_, _ = LoadSchema(DocumentType("invalid"))
		})
	})
}

func TestCheckJson(t *testing.T) {
	t.Run("validates valid module json", func(t *testing.T) {
		// Minimal valid module JSON
		jsonDoc := []byte(`{
  "schema": "apigear.module/1.0",
  "name": "test.module",
  "version": "1.0.0"
}`)
		result, err := CheckJson(DocumentTypeModule, jsonDoc)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		// Result should be valid (no errors)
		assert.True(t, result.Valid())
	})

	t.Run("detects invalid module json", func(t *testing.T) {
		// Invalid module JSON - missing required fields
		jsonDoc := []byte(`{
  "schema": "apigear.module/1.0"
}`)
		result, err := CheckJson(DocumentTypeModule, jsonDoc)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		// Result should be invalid (has errors)
		assert.False(t, result.Valid())
		assert.NotEmpty(t, result.Errors)
	})

	t.Run("returns error for malformed json", func(t *testing.T) {
		// Malformed JSON
		jsonDoc := []byte(`{
  "schema": "apigear.module/1.0",
  "name": "test.module"
  "missing comma": true
}`)
		_, err := CheckJson(DocumentTypeModule, jsonDoc)
		assert.Error(t, err)
	})

	t.Run("validates valid solution json", func(t *testing.T) {
		// Minimal valid solution JSON
		jsonDoc := []byte(`{
  "schema": "apigear.solution/1.0",
  "name": "test.solution",
  "version": "1.0.0",
  "targets": []
}`)
		result, err := CheckJson(DocumentTypeSolution, jsonDoc)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.True(t, result.Valid())
	})
}
