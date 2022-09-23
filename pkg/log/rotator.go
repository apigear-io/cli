package log

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

func newRollingFile(fname string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   fname,
		MaxSize:    100, // megabytes
		MaxBackups: 52,
		MaxAge:     7,     //days
		Compress:   false, // disabled by default
	}
}
