package cmd

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "display version information",
		Long:  `display version, commit and build-date information`,
		Run: func(cmd *cobra.Command, args []string) {
			version := cfg.BuildVersion()
			commit := cfg.BuildCommit()
			date := cfg.BuildDate()
			fmt.Printf("%s-%s-%s\n", version, commit, date)
		},
	}
	return cmd
}
