package main

import (
	"os"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/cmd"
	"github.com/apigear-io/cli/pkg/log"
)

var (
	version = "0.0.0"
	commit  = "none"
	date    = "unknown"
)

func main() {
	log.SentryInit(log.CLI_DSN)
	log.SentryCaptureArgs()
	cfg.SetBuildInfo(version, commit, date)
	log.Debug().Msgf("version: %s-%s-%s", version, commit, date)
	defer func() {
		log.SentryRecover()
		log.SentryFlush()
	}()
	err := cmd.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("cli")
		log.SentryCaptureError(err)
		os.Exit(1)
	}
}
