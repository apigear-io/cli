package mon

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
	"github.com/apigear-io/cli/pkg/log"
)

// ReadJsonEvents reads events from a json stream file
// and sends them to the emitter channel.
func ReadJsonEvents(fn string, emitter chan *Event) error {
	// read file line by line using scanner
	file, err := os.Open(fn)
	if err != nil {
		log.Error().Err(err).Msgf("failed to open file %s", fn)
		return err
	}
	defer func() {
		file.Close()
		close(emitter)
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log.Info().Msgf("send: %s", line)
		// decode each line into an event using json
		event := &Event{}
		err := json.NewDecoder(bufio.NewReader(strings.NewReader(line))).Decode(event)
		if err != nil {
			log.Error().Err(err).Msgf("failed to decode line: %s", line)
			continue
		}
		emitter <- event
	}
	if err := scanner.Err(); err != nil {
		log.Error().Err(err).Msgf("failed to read file: %s", fn)
		return err
	}
	return nil
}
