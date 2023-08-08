package log

import (
	"bytes"
	"encoding/json"
	"io"
)

var (
	eventEmitter func(e map[string]interface{})
	bytesEmitter func(s string)
)

type EventLogWriter struct {
}

func NewEventLogWriter() io.Writer {
	return &EventLogWriter{}
}

func (w *EventLogWriter) Write(p []byte) (n int, err error) {
	if bytesEmitter != nil {
		bytesEmitter(string(p))
	}
	if eventEmitter != nil {
		event := map[string]interface{}{}
		d := json.NewDecoder(bytes.NewReader(p))
		d.UseNumber()
		err = d.Decode(&event)
		if err != nil {
			return 0, err
		}
		// event["id"] = uuid.New().String()
		eventEmitter(event)
	}
	return len(p), nil
}

func OnReportEvent(handler func(e map[string]interface{})) {
	eventEmitter = handler
}

func OnReportBytes(handler func(s string)) {
	bytesEmitter = handler
}
