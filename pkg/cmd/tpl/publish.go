package tpl

import (
	"github.com/apigear-io/cli/pkg/tpl"
	"github.com/spf13/cobra"
)

func NewPublishCommand() *cobra.Command {
	var dir string
	var cmd = &cobra.Command{
		Use:   "publish",
		Short: "Publish a template to a market place",
		Long:  `Publish a template to a market place. The template needs to be a public github repository.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Printf("publishing template %s\n", dir)
			return tpl.PublishTemplate(dir)
		},
	}
	cmd.Flags().StringVarP(&dir, "dir", "d", ".", "template directory")
	cmd.MarkFlagRequired("dir")
	return cmd
}
