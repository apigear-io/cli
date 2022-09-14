package tpl

import (
	"github.com/apigear-io/lib/git"
	"github.com/apigear-io/lib/tpl"

	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var url string
	var name string

	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:   "get [name]",
		Short: "Download a template from a source to the local cache.",
		Long: `Download a template from a source to the local cache. 
Templates cached can be easily used in a solutions document 
by using the template name.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			url = args[0]
			if name == "" {
				path, err := git.RepositoryNameFromGitUrl(url)
				if err != nil {
					return err
				}
				name = path
			}
			err := tpl.InstallTemplate(name, url)
			if err != nil {
				return err
			}
			cmd.Printf("template %s installed\n", name)
			return nil
		},
	}
	cmd.Flags().StringVarP(&name, "name", "", "", "name of the template repository")
	return cmd
}
