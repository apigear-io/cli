package log

import (
	"os"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func Topic(topic string) zerolog.Logger {
	return zlog.With().Str("topic", topic).Logger()
}

func init() {
	level := zerolog.InfoLevel
	debug := os.Getenv("DEBUG") != ""
	verbose := os.Getenv("VERBOSE") != ""
	if debug {
		level = zerolog.DebugLevel
	}
	if verbose {
		level = zerolog.TraceLevel
	}

	logFile := helper.Join(config.ConfigDir, "apigear.log")
	multi := zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{Out: os.Stderr},
		NewReportWriter("info"),
		newRollingFile(logFile),
	)
	zlog.Logger = zlog.Output(multi).Level(level)
	if verbose {
		zlog.Logger = zlog.Logger.With().Caller().Logger()
	}
}

var Debug = zlog.Debug
var Info = zlog.Info
var Warn = zlog.Warn
var Error = zlog.Error
var Fatal = zlog.Fatal
var Panic = zlog.Panic
