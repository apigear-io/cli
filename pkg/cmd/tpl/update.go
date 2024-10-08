package tpl

import (
	"github.com/apigear-io/cli/pkg/repos"

	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "update the template registry",
		Run: func(cmd *cobra.Command, args []string) {
			err := repos.Registry.Update()
			if err != nil {
				cmd.PrintErrln(err)
			}
			cmd.Println("template registry updated")
		},
	}
	return cmd
}
