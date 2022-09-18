package log

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

func newLogFileRotator(filename string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    100, // megabytes
		MaxBackups: 52,
		MaxAge:     7,     //days
		Compress:   false, // disabled by default
	}
}
