package mon

import (
	"encoding/json"
	"io"

	"github.com/apigear-io/cli/pkg/helper"
)

// ReadJsonEvents reads monitor events from a json stream file
func ReadJsonEvents(fn string) ([]Event, error) {
	scanner := helper.NewNDJSONScanner(0, 1)
	var events []Event

	err := scanner.ScanFile(fn, func(line []byte) error {
		var event Event
		if err := json.Unmarshal(line, &event); err != nil {
			return err
		}
		events = append(events, event)
		return nil
	})
	if err != nil && err != io.EOF {
		return nil, err
	}
	return events, nil
}
