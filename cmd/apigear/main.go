package main

import (
	"fmt"
	"os"

	"github.com/apigear-io/cli/pkg/cmd"
	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
)

var (
	version = "0.0.0"
	commit  = "none"
	date    = "unknown"
)

func main() {
	helper.SentryInit(helper.CLI_DSN)
	config.Set(config.KeyVersion, version)
	config.Set(config.KeyCommit, commit)
	config.Set(config.KeyDate, date)
	log.Debug().Msgf("version: %s-%s-%s", version, commit, date)
	defer helper.SentryFlush()
	if cmd.Run() != 0 {
		helper.SentryCaptureError(fmt.Errorf("cmd.Run() != 0"))
		os.Exit(1)
	}
	helper.SentryCaptureMessage("cmd.Run() == 0")
}
