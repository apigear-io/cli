package mon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewClientCommand(t *testing.T) {
	t.Run("creates feed command", func(t *testing.T) {
		cmd := NewClientCommand()
		assert.NotNil(t, cmd)
		assert.Equal(t, "feed", cmd.Use)
		assert.Contains(t, cmd.Short, "Feed")
		assert.Contains(t, cmd.Short, "script")
	})

	t.Run("has long description", func(t *testing.T) {
		cmd := NewClientCommand()
		assert.Contains(t, cmd.Long, "Feeds")
		assert.Contains(t, cmd.Long, "API calls")
		assert.Contains(t, cmd.Long, "monitor")
	})

	t.Run("requires exactly one argument", func(t *testing.T) {
		cmd := NewClientCommand()
		assert.NotNil(t, cmd.Args)

		// Test with no arguments
		err := cmd.Args(cmd, []string{})
		assert.Error(t, err)

		// Test with one argument (should pass)
		err = cmd.Args(cmd, []string{"script.json"})
		assert.NoError(t, err)

		// Test with two arguments
		err = cmd.Args(cmd, []string{"script1.json", "script2.json"})
		assert.Error(t, err)
	})

	t.Run("has url flag", func(t *testing.T) {
		cmd := NewClientCommand()
		flag := cmd.Flags().Lookup("url")
		assert.NotNil(t, flag)
		assert.Equal(t, "http://localhost:5555/monitor/123", flag.DefValue)
		assert.Contains(t, flag.Usage, "monitor server")
	})

	t.Run("has repeat flag", func(t *testing.T) {
		cmd := NewClientCommand()
		flag := cmd.Flags().Lookup("repeat")
		assert.NotNil(t, flag)
		assert.Equal(t, "1", flag.DefValue)
		assert.Contains(t, flag.Usage, "repeat")
	})

	t.Run("has sleep flag", func(t *testing.T) {
		cmd := NewClientCommand()
		flag := cmd.Flags().Lookup("sleep")
		assert.NotNil(t, flag)
		assert.Equal(t, "0s", flag.DefValue)
		assert.Contains(t, flag.Usage, "sleep")
	})

	t.Run("url flag defaults to localhost:5555", func(t *testing.T) {
		cmd := NewClientCommand()
		url, err := cmd.Flags().GetString("url")
		assert.NoError(t, err)
		assert.Equal(t, "http://localhost:5555/monitor/123", url)
	})

	t.Run("repeat flag defaults to 1", func(t *testing.T) {
		cmd := NewClientCommand()
		repeat, err := cmd.Flags().GetInt("repeat")
		assert.NoError(t, err)
		assert.Equal(t, 1, repeat)
	})

	t.Run("sleep flag defaults to 0", func(t *testing.T) {
		cmd := NewClientCommand()
		sleep, err := cmd.Flags().GetDuration("sleep")
		assert.NoError(t, err)
		assert.Equal(t, time.Duration(0), sleep)
	})

	t.Run("accepts url flag", func(t *testing.T) {
		cmd := NewClientCommand()
		err := cmd.ParseFlags([]string{"--url", "http://example.com:8080"})
		assert.NoError(t, err)

		url, err := cmd.Flags().GetString("url")
		assert.NoError(t, err)
		assert.Equal(t, "http://example.com:8080", url)
	})

	t.Run("accepts repeat flag", func(t *testing.T) {
		cmd := NewClientCommand()
		err := cmd.ParseFlags([]string{"--repeat", "5"})
		assert.NoError(t, err)

		repeat, err := cmd.Flags().GetInt("repeat")
		assert.NoError(t, err)
		assert.Equal(t, 5, repeat)
	})

	t.Run("accepts sleep flag", func(t *testing.T) {
		cmd := NewClientCommand()
		err := cmd.ParseFlags([]string{"--sleep", "100ms"})
		assert.NoError(t, err)

		sleep, err := cmd.Flags().GetDuration("sleep")
		assert.NoError(t, err)
		assert.Equal(t, 100*time.Millisecond, sleep)
	})

	t.Run("accepts all flags", func(t *testing.T) {
		cmd := NewClientCommand()
		err := cmd.ParseFlags([]string{
			"--url", "http://test.com:9999",
			"--repeat", "10",
			"--sleep", "50ms",
		})
		assert.NoError(t, err)

		url, _ := cmd.Flags().GetString("url")
		assert.Equal(t, "http://test.com:9999", url)

		repeat, _ := cmd.Flags().GetInt("repeat")
		assert.Equal(t, 10, repeat)

		sleep, _ := cmd.Flags().GetDuration("sleep")
		assert.Equal(t, 50*time.Millisecond, sleep)
	})

	t.Run("accepts script argument", func(t *testing.T) {
		cmd := NewClientCommand()
		cmd.SetArgs([]string{"test.json"})
		err := cmd.ParseFlags([]string{})
		assert.NoError(t, err)
	})

	t.Run("has RunE function", func(t *testing.T) {
		cmd := NewClientCommand()
		assert.NotNil(t, cmd.RunE)
	})
}
