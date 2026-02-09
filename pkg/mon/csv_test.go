package mon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSVEvents(t *testing.T) {
	t.Run("reads events from valid CSV file", func(t *testing.T) {
		events, err := ReadCsvEvents("testdata/events.csv")
		assert.NoError(t, err)
		assert.Equal(t, 4, len(events))
	})

	t.Run("returns error for non-existent file", func(t *testing.T) {
		events, err := ReadCsvEvents("testdata/nonexistent.csv")
		assert.Error(t, err)
		assert.Nil(t, events)
	})

	t.Run("handles empty CSV file", func(t *testing.T) {
		events, err := ReadCsvEvents("testdata/empty.csv")
		// Should not error, just return empty slice
		assert.NoError(t, err)
		assert.Empty(t, events)
	})
}
