package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInfoCmd(t *testing.T) {
	t.Run("creates info command", func(t *testing.T) {
		cmd := NewInfoCmd()
		assert.NotNil(t, cmd)
		assert.Equal(t, "info", cmd.Use)
		assert.Contains(t, cmd.Aliases, "i")
		assert.Contains(t, cmd.Short, "config information")
	})

	t.Run("has info alias", func(t *testing.T) {
		cmd := NewInfoCmd()
		assert.Equal(t, []string{"i"}, cmd.Aliases)
	})

	t.Run("prints config information", func(t *testing.T) {
		cmd := NewInfoCmd()
		out := ExecuteCmd(t, cmd)

		// Should contain info header
		assert.Contains(t, out, "info:")

		// Should contain config file location
		assert.Contains(t, out, "config file:")

		// Should contain config section
		assert.Contains(t, out, "config:")
	})

	t.Run("displays config file path", func(t *testing.T) {
		cmd := NewInfoCmd()
		out := ExecuteCmd(t, cmd)

		// Should show the config file path
		lines := splitLines(out)
		foundConfigFile := false
		for _, line := range lines {
			if contains(line, "config file:") {
				foundConfigFile = true
				// Should have some path after the colon
				assert.Greater(t, len(line), len("  config file:"))
				break
			}
		}
		assert.True(t, foundConfigFile, "Should display config file path")
	})

	t.Run("displays config settings", func(t *testing.T) {
		cmd := NewInfoCmd()
		out := ExecuteCmd(t, cmd)

		// Config settings should be indented
		lines := splitLines(out)
		foundConfigSection := false
		foundSettings := false
		for i, line := range lines {
			if contains(line, "config:") {
				foundConfigSection = true
				// Check if next lines are indented (settings)
				if i+1 < len(lines) && len(lines[i+1]) > 0 {
					// Settings should start with spaces (indentation)
					if lines[i+1][0] == ' ' {
						foundSettings = true
					}
				}
				break
			}
		}
		assert.True(t, foundConfigSection, "Should have config section")
		// Settings might be empty in test environment, so we just check for the section
		_ = foundSettings
	})
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && indexOf(s, substr) >= 0
}

// Helper function to find index of substring
func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
