package cmd

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/log"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

type VersionInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func Run() int {
	rootCmd := NewRootCommand()
	rootCmd.Version = fmt.Sprintf("%s-%s-%s", version, commit, date)

	if err := rootCmd.Execute(); err != nil {
		log.Warn(err)
		return -1
	}
	return 0
}
