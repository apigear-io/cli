package tpl

import (
	"github.com/apigear-io/cli/pkg/tpl"

	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:     "update [template]",
		Aliases: []string{"up"},
		Short:   "Update template registry.",
		Long:    `Fetch latest template registry.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := tpl.UpdateRegistry()
			if err != nil {
				cmd.PrintErrln(err)
			}
			cmd.Printf("local template registry updated\n")
		},
	}
	return cmd
}
