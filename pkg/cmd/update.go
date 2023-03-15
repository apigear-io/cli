package cmd

import (
	"context"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/apigear-io/cli/pkg/up"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var force bool
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "update the program",
		Long:  `check and update the program to the latest version`,
		Run: func(cmd *cobra.Command, args []string) {
			repo := "apigear-io/cli"
			version := cfg.BuildVersion()
			ctx := context.Background()
			u, err := up.NewUpdater(repo, version)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

			release, err := u.Check(ctx)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			if release == nil {
				cmd.Println("no new release available")
				return
			}
			cmd.Printf("New release %s available.\n", release.Version())
			cmd.Printf("See %s.\n", release.URL)
			if !force {
				result, err := pterm.DefaultInteractiveConfirm.Show("do you want to update?")
				if err != nil {
					cmd.PrintErrln(err)
					return
				}
				if !result {
					return
				}
			}
			cmd.Printf("updating to %s\n", release.Version())
			err = u.Update(ctx, release)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

		},
	}
	cmd.Flags().BoolVarP(&force, "force", "f", false, "force update")
	return cmd
}
