package mon

import (
	"encoding/json"
	"os"

	"github.com/gocarina/gocsv"
)

// ReadCsvEvents reads events from a csv file
// and sends them to the emitter channel.
func ReadCsvEvents(fn string, emitter chan *Event) error {
	log.Debug().Msgf("read csv events from %s", fn)
	// read file line by line using scanner
	file, err := os.Open(fn)
	if err != nil {
		log.Error().Err(err).Msgf("open file %s", fn)
		return err
	}
	defer file.Close()

	// need to have a intermediate struct to unmarshal csv
	type csvEvent struct {
		Type   string `csv:"type"`
		Symbol string `csv:"symbol"`
		Data   string `csv:"data"`
	}
	events := []*csvEvent{}
	err = gocsv.UnmarshalFile(file, &events)
	if err != nil {
		log.Error().Err(err).Msgf("unmarshal file %s", fn)
	}

	// send events to emitter channel
	for _, event := range events {
		// unmarshal csv data into json for event payload
		data := Payload{}
		if event.Data != "" {
			err = json.Unmarshal([]byte(event.Data), &data)
			if err != nil {
				log.Error().Err(err).Msg("unmarshal data")
				continue
			}
		}
		// create event and send to emitter channel
		evt := Event{
			Type:   ParseEventType(event.Type),
			Symbol: event.Symbol,
			Data:   data,
		}
		emitter <- &evt
	}
	close(emitter)
	return nil
}
