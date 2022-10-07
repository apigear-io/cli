package log

import (
	"os"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

var (
	Debug = zlog.Debug
	Info  = zlog.Info
	Warn  = zlog.Warn
	Error = zlog.Error
	Fatal = zlog.Fatal
	Panic = zlog.Panic
)

func init() {
	level := zerolog.InfoLevel
	debug := os.Getenv("DEBUG") == "1"
	verbose := os.Getenv("VERBOSE") == "1"
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
	Debug().Msgf("log level: %s", level)
}
