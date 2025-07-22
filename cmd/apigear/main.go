// Main entry point for apigear cli tool
// It initializes sentry, runs the command and recovers from panics
// Version, commit and date are set by the build process
package main

import (
	"os"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/cmd"
)

var (
	version = "0.0.0"
	commit  = "none"
	date    = "unknown"
)

// main entry point for apigear cli tool
func main() {
	cfg.SetBuildInfo("cli", cfg.BuildInfo{
		Version: version,
		Commit:  commit,
		Date:    date,
	})

	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
