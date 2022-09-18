package mon

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

// ReadJsonEvents reads events from a json stream file
// and sends them to the emitter channel.
func ReadJsonEvents(fn string, emitter chan *Event) error {
	// read file line by line using scanner
	file, err := os.Open(fn)
	if err != nil {
		log.Error("failed to open file ", fn, ": ", err)
		return err
	}
	defer func() {
		file.Close()
		close(emitter)
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log.Infof("send: %s", line)
		// decode each line into an event using json
		event := &Event{}
		err := json.NewDecoder(bufio.NewReader(strings.NewReader(line))).Decode(event)
		if err != nil {
			log.Error("failed to decode line: ", line, ": ", err)
			continue
		}
		emitter <- event
	}
	if err := scanner.Err(); err != nil {
		log.Error("failed to read file: ", fn, ": ", err)
		return err
	}
	return nil
}
