package mon

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSVEvents(t *testing.T) {
	// create a channel to receive events
	emitter := make(chan *Event)
	// create a reader
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		// start the reader
		// read events from the channel
		events := 0
		for range emitter {
			events++
		}
		assert.Equal(t, 4, events)
		wg.Done()
	}()
	go func() {
		err := ReadCsvEvents("testdata/events.csv", emitter)
		assert.NoError(t, err)
		wg.Done()
	}()
	wg.Wait()

}
