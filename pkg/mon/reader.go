package mon

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

// ReadJsonEvents reads events from a json stream file
// and sends them to the emitter channel.
func ReadJsonEvents(fn string) ([]*Event, error) {
	events := []*Event{}
	// read file line by line using scanner
	file, err := os.Open(fn)
	if err != nil {
		log.Error().Err(err).Msgf("open file %s", fn)
		return nil, err
	}
	defer func() {
		file.Close()
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log.Info().Msgf("send: %s", line)
		// decode each line into an event using json
		event := &Event{}
		err := json.NewDecoder(bufio.NewReader(strings.NewReader(line))).Decode(event)
		if err != nil {
			log.Error().Err(err).Msgf("decode line: %s", line)
			continue
		}
		events = append(events, event)
	}
	if err := scanner.Err(); err != nil {
		log.Error().Err(err).Msgf("read file: %s", fn)
		return nil, err
	}
	return events, nil
}
