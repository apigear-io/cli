package mon

import (
	"apigear/pkg/log"
	"os"

	"github.com/gocarina/gocsv"
)

// ReadCsvEvents reads events from a csv file
// and sends them to the emitter channel.
func ReadCsvEvents(fn string, emitter chan *Event) error {
	// read file line by line using scanner
	file, err := os.Open(fn)
	if err != nil {
		log.Error("failed to open file ", fn, ": ", err)
		return err
	}
	defer file.Close()
	events := []*Event{}
	err = gocsv.UnmarshalFile(file, &events)
	if err != nil {
		log.Error("failed to unmarshal file ", fn, ": ", err)
	}
	for _, event := range events {
		emitter <- event
	}
	close(emitter)
	return nil
}
