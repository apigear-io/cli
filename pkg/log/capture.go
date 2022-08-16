package log

import "github.com/sirupsen/logrus"

type logCapture struct {
	*logrus.Logger
}

func NewLogCapture(logger *logrus.Logger) *logCapture {
	return &logCapture{
		Logger: logger,
	}
}

func (w *logCapture) Write(p []byte) (n int, err error) {
	w.Debug(string(p))
	return len(p), nil
}
