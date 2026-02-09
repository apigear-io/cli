package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewShowCommand(t *testing.T) {
	t.Run("creates show command", func(t *testing.T) {
		cmd := NewShowCommand()
		assert.NotNil(t, cmd)
		assert.Equal(t, "schema", cmd.Use)
		assert.Contains(t, cmd.Aliases, "s")
		assert.Contains(t, cmd.Aliases, "show")
		assert.Contains(t, cmd.Aliases, "view")
		assert.Contains(t, cmd.Short, "schema")
	})

	t.Run("has correct aliases", func(t *testing.T) {
		cmd := NewShowCommand()
		assert.Equal(t, []string{"s", "show", "view"}, cmd.Aliases)
	})

	t.Run("has long description", func(t *testing.T) {
		cmd := NewShowCommand()
		assert.Contains(t, cmd.Long, "schema")
		assert.Contains(t, cmd.Long, "module")
		assert.Contains(t, cmd.Long, "solution")
		assert.Contains(t, cmd.Long, "rules")
	})

	t.Run("has type flag", func(t *testing.T) {
		cmd := NewShowCommand()
		flag := cmd.Flags().Lookup("type")
		assert.NotNil(t, flag)
		assert.Equal(t, "t", flag.Shorthand)
		assert.Equal(t, "module", flag.DefValue)
		assert.Contains(t, flag.Usage, "Document type")
	})

	t.Run("has format flag", func(t *testing.T) {
		cmd := NewShowCommand()
		flag := cmd.Flags().Lookup("format")
		assert.NotNil(t, flag)
		assert.Equal(t, "f", flag.Shorthand)
		assert.Equal(t, "yaml", flag.DefValue)
		assert.Contains(t, flag.Usage, "format")
	})

	t.Run("type flag defaults to module", func(t *testing.T) {
		cmd := NewShowCommand()
		docType, err := cmd.Flags().GetString("type")
		assert.NoError(t, err)
		assert.Equal(t, "module", docType)
	})

	t.Run("format flag defaults to yaml", func(t *testing.T) {
		cmd := NewShowCommand()
		format, err := cmd.Flags().GetString("format")
		assert.NoError(t, err)
		assert.Equal(t, "yaml", format)
	})

	t.Run("accepts type flag", func(t *testing.T) {
		cmd := NewShowCommand()
		err := cmd.ParseFlags([]string{"--type", "solution"})
		assert.NoError(t, err)

		docType, err := cmd.Flags().GetString("type")
		assert.NoError(t, err)
		assert.Equal(t, "solution", docType)
	})

	t.Run("accepts format flag", func(t *testing.T) {
		cmd := NewShowCommand()
		err := cmd.ParseFlags([]string{"--format", "json"})
		assert.NoError(t, err)

		format, err := cmd.Flags().GetString("format")
		assert.NoError(t, err)
		assert.Equal(t, "json", format)
	})

	t.Run("accepts both flags", func(t *testing.T) {
		cmd := NewShowCommand()
		err := cmd.ParseFlags([]string{"--type", "rules", "--format", "json"})
		assert.NoError(t, err)

		docType, err := cmd.Flags().GetString("type")
		assert.NoError(t, err)
		assert.Equal(t, "rules", docType)

		format, err := cmd.Flags().GetString("format")
		assert.NoError(t, err)
		assert.Equal(t, "json", format)
	})

	t.Run("accepts short flags", func(t *testing.T) {
		cmd := NewShowCommand()
		err := cmd.ParseFlags([]string{"-t", "solution", "-f", "yaml"})
		assert.NoError(t, err)

		docType, err := cmd.Flags().GetString("type")
		assert.NoError(t, err)
		assert.Equal(t, "solution", docType)

		format, err := cmd.Flags().GetString("format")
		assert.NoError(t, err)
		assert.Equal(t, "yaml", format)
	})

	t.Run("has RunE function", func(t *testing.T) {
		cmd := NewShowCommand()
		assert.NotNil(t, cmd.RunE)
	})
}
