package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Configure sets up zerolog for the application.
func Configure(verbose bool) {
	level := zerolog.InfoLevel
	if verbose {
		level = zerolog.DebugLevel
	}

	zerolog.TimeFieldFormat = time.RFC3339
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly}
	logger := zerolog.New(output).With().Timestamp().Logger().Level(level)
	log.Logger = logger
}
