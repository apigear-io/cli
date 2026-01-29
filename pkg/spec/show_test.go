package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShowSchemaFile(t *testing.T) {
	t.Run("returns module schema in JSON format", func(t *testing.T) {
		result, err := ShowSchemaFile(DocumentTypeModule, SchemaFormatJson)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.NotEmpty(t, *result)
		// Should contain JSON schema content
		assert.Contains(t, *result, "apigear.module")
	})

	t.Run("returns module schema in YAML format", func(t *testing.T) {
		result, err := ShowSchemaFile(DocumentTypeModule, SchemaFormatYaml)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.NotEmpty(t, *result)
		// Should contain YAML schema content
		assert.Contains(t, *result, "apigear.module")
	})

	t.Run("returns solution schema in JSON format", func(t *testing.T) {
		result, err := ShowSchemaFile(DocumentTypeSolution, SchemaFormatJson)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.NotEmpty(t, *result)
		// Should contain JSON schema content
		assert.Contains(t, *result, "apigear.solution")
	})

	t.Run("returns solution schema in YAML format", func(t *testing.T) {
		result, err := ShowSchemaFile(DocumentTypeSolution, SchemaFormatYaml)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.NotEmpty(t, *result)
		// Should contain YAML schema content
		assert.Contains(t, *result, "apigear.solution")
	})

	t.Run("returns rules schema in JSON format", func(t *testing.T) {
		result, err := ShowSchemaFile(DocumentTypeRules, SchemaFormatJson)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.NotEmpty(t, *result)
		// Should contain JSON schema content
		assert.Contains(t, *result, "apigear.rules")
	})

	t.Run("returns rules schema in YAML format", func(t *testing.T) {
		result, err := ShowSchemaFile(DocumentTypeRules, SchemaFormatYaml)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.NotEmpty(t, *result)
		// Should contain YAML schema content
		assert.Contains(t, *result, "apigear.rules")
	})

	t.Run("returns error for unsupported document type", func(t *testing.T) {
		result, err := ShowSchemaFile(DocumentTypeUnknown, SchemaFormatJson)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unsupported document type")
	})

	t.Run("returns error for unsupported schema format - module", func(t *testing.T) {
		result, err := ShowSchemaFile(DocumentTypeModule, SchemaFormat("invalid"))
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unsupported schema format")
	})

	t.Run("returns error for unsupported schema format - solution", func(t *testing.T) {
		result, err := ShowSchemaFile(DocumentTypeSolution, SchemaFormat("invalid"))
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unsupported schema format")
	})

	t.Run("returns error for unsupported schema format - rules", func(t *testing.T) {
		result, err := ShowSchemaFile(DocumentTypeRules, SchemaFormat("invalid"))
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unsupported schema format")
	})
}
