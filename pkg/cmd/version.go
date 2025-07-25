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
			cmd.Println(RetrieveVersion())
		},
	}
	return cmd
}

func RetrieveVersion() string {
	bi := cfg.GetBuildInfo("cli")
	version := fmt.Sprintf("%s-%s-%s", bi.Version, bi.Commit, bi.Date)
	return version
}
