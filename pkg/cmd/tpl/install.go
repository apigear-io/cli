package tpl

import (
	"github.com/apigear-io/cli/pkg/tpl"

	"github.com/spf13/cobra"
)

func NewInstallCommand() *cobra.Command {
	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:     "install [name]",
		Short:   "install template",
		Long:    `install template from registry using a name`,
		Aliases: []string{"i"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				for _, name := range args {
					err := tpl.InstallTemplate(name)
					if err != nil {
						cmd.PrintErrln(err)
					}
				}
			}
		},
	}
	return cmd
}
