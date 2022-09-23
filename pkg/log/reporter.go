package log

import (
	"bytes"
	"encoding/json"
	"io"
	"sync"
	"time"
)

var (
	emitter func(*ReportEvent)
	mu      = sync.Mutex{}
)

type ReportEvent struct {
	Level   string    `json:"level"`
	Topic   string    `json:"topic"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

type ReportWriter struct {
	Level string
}

func NewReportWriter(level string) io.Writer {
	return &ReportWriter{Level: level}
}

func (w *ReportWriter) Write(p []byte) (n int, err error) {
	var event ReportEvent
	d := json.NewDecoder(bytes.NewReader(p))
	d.UseNumber()
	err = d.Decode(&event)
	if err != nil {
		return 0, err
	}
	if emitter != nil {
		mu.Lock()
		emitter(&event)
		mu.Unlock()
	}
	return len(p), nil
}

func OnReport(handler func(*ReportEvent)) {
	emitter = handler
}
