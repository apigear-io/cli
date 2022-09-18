package mon

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonReader(t *testing.T) {
	// create a channel to receive events
	emitter := make(chan *Event)
	// create a reader
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		// start the reader
		// read events from the channel
		events := 0
		for range emitter {
			events++
		}
		assert.Equal(t, events, 4)
		wg.Done()
	}()
	err := ReadJsonEvents("testdata/events.ndjson", emitter)
	assert.NoError(t, err)
	wg.Wait()
}
