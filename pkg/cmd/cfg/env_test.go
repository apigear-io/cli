package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonIdent(t *testing.T) {
	t.Run("marshals string value", func(t *testing.T) {
		result := jsonIdent("test")
		assert.Equal(t, `"test"`, result)
	})

	t.Run("marshals number value", func(t *testing.T) {
		result := jsonIdent(42)
		assert.Equal(t, "42", result)
	})

	t.Run("marshals boolean value", func(t *testing.T) {
		result := jsonIdent(true)
		assert.Equal(t, "true", result)
	})

	t.Run("marshals map value", func(t *testing.T) {
		result := jsonIdent(map[string]string{"key": "value"})
		assert.Contains(t, result, `"key"`)
		assert.Contains(t, result, `"value"`)
	})

	t.Run("marshals slice value", func(t *testing.T) {
		result := jsonIdent([]string{"a", "b", "c"})
		assert.Contains(t, result, `"a"`)
		assert.Contains(t, result, `"b"`)
		assert.Contains(t, result, `"c"`)
	})

	t.Run("marshals nil value", func(t *testing.T) {
		result := jsonIdent(nil)
		assert.Equal(t, "null", result)
	})

	t.Run("handles unmarshalable value", func(t *testing.T) {
		// Channels cannot be marshaled to JSON
		result := jsonIdent(make(chan int))
		assert.Contains(t, result, "Error")
	})
}

func TestNewEnvCommand(t *testing.T) {
	t.Run("creates env command", func(t *testing.T) {
		cmd := NewEnvCommand()
		assert.NotNil(t, cmd)
		assert.Equal(t, "env", cmd.Use)
		assert.Contains(t, cmd.Short, "environment variables")
	})

	t.Run("prints environment variables", func(t *testing.T) {
		cmd := NewEnvCommand()
		out := ExecuteCmd(t, cmd)

		// Should contain APIGEAR_CONFIG_DIR
		assert.Contains(t, out, "APIGEAR_CONFIG_DIR=")

		// Should contain at least some APIGEAR_ prefixed variables
		assert.Contains(t, out, "APIGEAR_")
	})

	t.Run("formats variables with uppercase and APIGEAR prefix", func(t *testing.T) {
		cmd := NewEnvCommand()
		out := ExecuteCmd(t, cmd)

		// Variables should be uppercase with APIGEAR_ prefix
		lines := splitLines(out)
		for _, line := range lines {
			if line != "" {
				assert.Contains(t, line, "APIGEAR_")
				assert.Contains(t, line, "=")
			}
		}
	})
}

// Helper function to split output into lines
func splitLines(s string) []string {
	lines := []string{}
	current := ""
	for _, c := range s {
		if c == '\n' {
			lines = append(lines, current)
			current = ""
		} else {
			current += string(c)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}
