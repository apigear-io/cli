package log

import (
	"time"

	"github.com/sirupsen/logrus"
)

type ReportHook struct {
	LogLevels []logrus.Level
}

func (h *ReportHook) Levels() []logrus.Level {
	return h.LogLevels
}

func (h *ReportHook) Fire(entry *logrus.Entry) error {
	var topic string
	if val, ok := entry.Data["topic"]; ok {
		topic = val.(string)
	}
	reportEntry(&ReportEntry{
		Level:   entry.Level.String(),
		Message: entry.Message,
		Time:    entry.Time,
		Topic:   topic,
	})
	return nil
}

func NewReportHook() *ReportHook {
	return &ReportHook{
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.WarnLevel,
			logrus.ErrorLevel,
			logrus.PanicLevel,
			logrus.FatalLevel,
		},
	}
}

type ReportEntry struct {
	Level   string    `json:"level"`
	Topic   string    `json:"topic"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

var emitter = make(chan *ReportEntry)

func reportEntry(entry *ReportEntry) {
	select {
	case emitter <- entry:
		// if we can send the entry, we are good
	default:
		// if we can't send the entry, we drop it
	}
}

func Emitter() <-chan *ReportEntry {
	return emitter
}
