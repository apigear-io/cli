package log

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
)

var (
	eventEmitter func(*ReportEvent)
	bytesEmitter func(s string)
)

type ReportEvent struct {
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error,omitempty"`
}

type ReportWriter struct {
}

func NewReportWriter() io.Writer {
	return &ReportWriter{}
}

func (w *ReportWriter) Write(p []byte) (n int, err error) {
	if bytesEmitter != nil {
		bytesEmitter(string(p))
	}
	if eventEmitter != nil {
		var event ReportEvent
		d := json.NewDecoder(bytes.NewReader(p))
		d.UseNumber()
		err = d.Decode(&event)
		if err != nil {
			return 0, err
		}
		eventEmitter(&event)
	}
	return len(p), nil
}

func OnReportEvent(handler func(*ReportEvent)) {
	eventEmitter = handler
}

func OnReportBytes(handler func(s string)) {
	bytesEmitter = handler
}
