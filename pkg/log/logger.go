package log

import (
	"os"
	"time"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

var (
	logger zerolog.Logger
)

type UUIDHook struct {
}

func (h UUIDHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Str("id", helper.NewUUID())
}

func init() {
	level := zerolog.InfoLevel
	debug := os.Getenv("DEBUG") == "1"
	verbose := os.Getenv("DEBUG") == "2"
	if debug {
		level = zerolog.DebugLevel
	}
	if verbose {
		level = zerolog.TraceLevel
	}
	logFile := helper.Join(cfg.ConfigDir(), "apigear.log")
	console := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Kitchen, FieldsExclude: []string{"id"}}
	multi := zerolog.MultiLevelWriter(
		console,
		NewEventLogWriter(),
		newRollingFile(logFile),
	)
	logger = zerolog.New(multi).
		With().Timestamp().Logger().Hook(&UUIDHook{}).Level(level)

	if verbose {
		logger = logger.With().Caller().Logger()
	}
	zlog.Logger = logger
}

func Debug() *zerolog.Event {
	return logger.Debug()
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Error() *zerolog.Event {
	return logger.Error()
}

func Fatal() *zerolog.Event {
	return logger.Fatal()
}

func Panic() *zerolog.Event {
	return logger.Panic()
}

func Topic(topic string) zerolog.Logger {
	return logger.With().Str("topic", topic).Logger()
}
