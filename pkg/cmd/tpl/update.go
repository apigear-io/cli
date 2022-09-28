package tpl

import (
	"github.com/apigear-io/cli/pkg/tpl"

	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "update template registry",
		Long:  `update registry from remote source.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := tpl.UpdateRegistry()
			if err != nil {
				cmd.PrintErrln(err)
			}
			cmd.Println("template registry updated")
		},
	}
	return cmd
}
