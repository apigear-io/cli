package registry

import (
	"github.com/apigear-io/cli/pkg/repos"
	"github.com/spf13/cobra"
)

func NewInfoCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info [name]",
		Short: "display information",
		Long:  `display template cache information from given template.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			info, err := repos.Registry.Info(name)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			DisplayTemplateInfo(info)
		},
	}
	return cmd
}
