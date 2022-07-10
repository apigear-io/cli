package tpl

import (
	"apigear/pkg/tpl"

	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var url string
	var name string

	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:   "get [template]",
		Short: "download and install a template",
		Long:  `Download and install a template from a git repository.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("get %s\n", args[0])
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
