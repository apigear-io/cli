package mon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonReader(t *testing.T) {
	// create a channel to receive events
	// create a reader
	events, err := ReadJsonEvents("testdata/events.ndjson")
	assert.NoError(t, err)
	assert.Equal(t, 4, len(events))
}
