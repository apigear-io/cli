package tpl

import (
	"github.com/apigear-io/cli/pkg/repos"

	"github.com/spf13/cobra"
)

func NewInstallCommand() *cobra.Command {
	// cmd represents the pkgInstall command
	var version string
	var cmd = &cobra.Command{
		Use:     "install [name]",
		Short:   "install template into cache by name from registry",
		Aliases: []string{"i"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			repoID := args[0]
			fixedRepoId, err := repos.GetOrInstallTemplateFromRepoID(repoID)
			cmd.Printf("using template %s\n", fixedRepoId)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
		},
	}
	cmd.Flags().StringVarP(&version, "version", "v", "latest", "template version to install")
	return cmd
}
