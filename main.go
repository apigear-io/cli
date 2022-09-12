package main

import (
	"os"

	"github.com/apigear-io/cli/cmd"
	"github.com/apigear-io/cli/pkg/log"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	log.Debugf("version: %s-%s-%s", version, commit, date)
	if cmd.Run() != 0 {
		os.Exit(1)
	}
}
