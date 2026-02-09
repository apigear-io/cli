package tpl

import (
	"github.com/apigear-io/cli/pkg/foundation/logging"
	"github.com/apigear-io/cli/pkg/codegen/template"
	"github.com/spf13/cobra"
)

func NewPublishCommand() *cobra.Command {
	var dir string
	var cmd = &cobra.Command{
		Use:   "publish",
		Short: "publish a template to a template registry (TBD)",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Printf("publishing template %s to the registry\n", dir)
			return template.PublishTemplate(dir)
		},
	}
	cmd.Flags().StringVarP(&dir, "dir", "d", ".", "template directory")
	err := cmd.MarkFlagRequired("dir")
	if err != nil {
		logging.Error().Err(err).Msg("failed to mark flag required")
	}
	return cmd
}
