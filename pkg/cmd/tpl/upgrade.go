package tpl

import (
	"github.com/apigear-io/cli/pkg/tpl"

	"github.com/spf13/cobra"
)

func NewUpgradeCommand() *cobra.Command {
	var all bool
	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:     "upgrade [name]",
		Aliases: []string{"up"},
		Short:   "upgrade installed template",
		Long:    `upgrade installed template. If name is not specified, all installed templates will be upgraded.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				err := tpl.UpgradeTemplates(args)
				if err != nil {
					cmd.PrintErrln(err)
				}
			} else if all {
				err := tpl.UpgradeAllTemplates()
				if err != nil {
					cmd.PrintErrln(err)
				}
			} else {
				err := cmd.Usage()
				if err != nil {
					cmd.PrintErrln(err)
				}
			}
		},
	}
	cmd.Flags().BoolVarP(&all, "all", "a", false, "upgrade all installed templates")
	return cmd
}
