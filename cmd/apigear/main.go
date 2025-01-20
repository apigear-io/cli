// Main entry point for apigear cli tool
// It initializes sentry, runs the command and recovers from panics
// Version, commit and date are set by the build process
package main

import (
	"os"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/cmd"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

var (
	version = "0.0.0"
	commit  = "none"
	date    = "unknown"
)

// main entry point for apigear cli tool
func main() {
	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	if os.Getenv("DEBUG") == "1" {
		zlog.Logger = zlog.Logger.Level(zerolog.DebugLevel).With().Logger()
	}
	if os.Getenv("DEBUG") == "2" {
		zlog.Logger = zlog.Logger.Level(zerolog.TraceLevel).With().Caller().Logger()
	}

	cfg.SetBuildInfo("cli", cfg.BuildInfo{
		Version: version,
		Commit:  commit,
		Date:    date,
	})

	log.Debug().Msgf("version: %s-%s-%s", version, commit, date)
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
