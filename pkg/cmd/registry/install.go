package registry

import (
	"github.com/apigear-io/cli/pkg/repos"

	"github.com/spf13/cobra"
)

func NewInstallCommand() *cobra.Command {
	// cmd represents the pkgInstall command
	var version string
	var cmd = &cobra.Command{
		Use:     "install [name]",
		Short:   "install template into cache",
		Long:    `install template from registry using a name`,
		Aliases: []string{"i"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			repoID := args[0]
			cmd.Printf("installing template from %s\n", repoID)
			err := repos.InstallTemplateFromRepoID(repoID)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
		},
	}
	cmd.Flags().StringVarP(&version, "version", "v", "latest", "template version to install")
	return cmd
}
