package mon

import (
	"bufio"
	"encoding/json"
	"os"
)

// TODO: there is already a ndjon scanner in helper package

// ReadJsonEvents reads monitor events from a json stream file
func ReadJsonEvents(fn string) ([]Event, error) {
	var events []Event
	// read file line by line using scanner
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Error().Err(err).Msgf("failed to close file %s", fn)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// decode each line into an event using json
		var event Event
		err := json.Unmarshal([]byte(line), &event)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return events, nil
}
