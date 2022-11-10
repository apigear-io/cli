package log

import (
	"os"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/rs/zerolog"
)

var (
	logger zerolog.Logger
)

func init() {
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	level := zerolog.InfoLevel
	debug := os.Getenv("DEBUG") == "1"
	verbose := os.Getenv("DEBUG") == "2"
	if debug || verbose {
		level = zerolog.DebugLevel
	}
	if verbose {
		level = zerolog.TraceLevel
	}
	logFile := helper.Join(cfg.ConfigDir(), "apigear.log")
	multi := zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"},
		NewReportWriter(),
		newRollingFile(logFile),
	)
	logger = logger.Output(multi).Level(level)
	if verbose {
		logger = logger.With().Caller().Logger()
	}
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
