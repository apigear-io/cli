package mon

import (
	"encoding/json"
	"os"

	"github.com/gocarina/gocsv"
)

// ReadCsvEvents reads events from a csv file
// and sends them to the emitter channel.
func ReadCsvEvents(fn string) ([]Event, error) {
	log.Debug().Msgf("read csv events from %s", fn)
	var events []Event
	// read file line by line using scanner
	file, err := os.Open(fn)
	if err != nil {
		log.Error().Err(err).Msgf("open file %s", fn)
		return nil, err
	}
	defer file.Close()

	// need to have a intermediate struct to unmarshal csv
	type Row struct {
		Type   string `csv:"type"`
		Symbol string `csv:"symbol"`
		Data   string `csv:"data"`
	}
	rows := []*Row{}
	err = gocsv.UnmarshalFile(file, &rows)
	if err != nil {
		log.Error().Err(err).Msgf("unmarshal file %s", fn)
	}

	// send events to emitter channel
	for _, row := range rows {
		// unmarshal csv data into json for event payload
		data := Payload{}
		if row.Data != "" {
			err = json.Unmarshal([]byte(row.Data), &data)
			if err != nil {
				log.Error().Err(err).Msg("unmarshal data")
				continue
			}
		}
		// create event and send to emitter channel
		evt := Event{
			Type:   ParseEventType(row.Type),
			Symbol: row.Symbol,
			Data:   data,
		}
		events = append(events, evt)
	}
	return events, nil
}
