package tpl

import (
	"fmt"

	"github.com/spf13/cobra"
)

type GetOptions struct {
	template string
}

func NewGetCommand() *cobra.Command {
	var options = &GetOptions{}

	// cmd represents the pkgInstall command
	var cmd = &cobra.Command{
		Use:   "get [template]",
		Short: "download and install a template",
		Long:  `Download an dinstall a template from a git repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get template")
		},
	}
	cmd.Flags().StringVarP(&options.template, "template", "t", "", "the name of the template to install")
	return cmd
}
