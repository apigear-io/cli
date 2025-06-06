package tpl

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/tpl"
	"github.com/spf13/cobra"
)

func NewPublishCommand() *cobra.Command {
	var dir string
	var cmd = &cobra.Command{
		Use:   "publish",
		Short: "publish a template to a template registry (TBD)",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Printf("publishing template %s to the registry\n", dir)
			return tpl.PublishTemplate(dir)
		},
	}
	cmd.Flags().StringVarP(&dir, "dir", "d", ".", "template directory")
	err := cmd.MarkFlagRequired("dir")
	if err != nil {
		log.Error().Err(err).Msg("failed to mark flag required")
	}
	return cmd
}
