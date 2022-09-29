package cmd

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/up"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "update the program",
		Long:  `check and update the program to the latest version`,
		Run: func(cmd *cobra.Command, args []string) {
			repo := "apigear-io/cli-releases"
			version := config.Get(config.KeyVersion)
			u, err := up.NewUpdater(repo, version)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			release, err := u.Check()
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			if release == nil {
				cmd.Println("no new release available")
				return
			}
			result, err := pterm.DefaultInteractiveConfirm.Show(fmt.Sprintf("do you want to update to version %s?", release.Version()))
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			if !result {
				return
			}
			err = u.Update(release)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

		},
	}
	return cmd
}
