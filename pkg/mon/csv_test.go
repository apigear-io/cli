package mon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSVEvents(t *testing.T) {
	events, err := ReadCsvEvents("testdata/events.csv")
	assert.NoError(t, err)
	assert.Equal(t, 4, len(events))
}
