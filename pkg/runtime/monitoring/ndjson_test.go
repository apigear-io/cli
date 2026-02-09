package monitoring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonReader(t *testing.T) {
	t.Run("reads events from valid NDJSON file", func(t *testing.T) {
		// create a channel to receive events
		// create a reader
		events, err := ReadJsonEvents("testdata/events.ndjson")
		assert.NoError(t, err)
		assert.Equal(t, 4, len(events))
	})

	t.Run("returns error for non-existent file", func(t *testing.T) {
		events, err := ReadJsonEvents("testdata/nonexistent.ndjson")
		assert.Error(t, err)
		assert.Nil(t, events)
	})

	t.Run("handles empty NDJSON file", func(t *testing.T) {
		events, err := ReadJsonEvents("testdata/empty.ndjson")
		assert.NoError(t, err)
		assert.Empty(t, events)
	})

	t.Run("returns error for invalid JSON line", func(t *testing.T) {
		events, err := ReadJsonEvents("testdata/invalid.ndjson")
		assert.Error(t, err)
		assert.Nil(t, events)
	})
}
