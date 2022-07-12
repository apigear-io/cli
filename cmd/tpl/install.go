package tpl

import (
	"github.com/apigear-io/cli/pkg/tpl"

	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var url string
	var name string

	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:   "install [source]",
		Short: "Download and install a template from a source",
		Long:  `Download and install a template from a git repository.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("install %s\n", args[0])
			name = args[0]
			err := tpl.GetTemplate(name, url)
			if err != nil {
				cmd.PrintErrln(err)
			}
		},
	}
	cmd.Flags().StringVarP(&url, "url", "u", "", "url of the template repository to install")
	return cmd
}
