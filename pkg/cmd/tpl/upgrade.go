package tpl

import (
	"github.com/apigear-io/cli/pkg/tpl"

	"github.com/spf13/cobra"
)

func NewUpgradeCommand() *cobra.Command {
	var name string
	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:     "upgrade [template]",
		Aliases: []string{"up"},
		Short:   "Upgrade installed templates.",
		Long:    `Upgrade installed templates.`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name = args[0]
			err := tpl.UpgradeTemplate(name)
			if err != nil {
				cmd.PrintErrln(err)
			}
			cmd.Printf("template %s upgraded\n", name)
		},
	}
	return cmd
}
