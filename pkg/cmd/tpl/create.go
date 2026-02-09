package tpl

import (
	"github.com/apigear-io/cli/pkg/foundation/logging"
	"github.com/apigear-io/cli/pkg/codegen/template"
	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	var dir string
	var lang string
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "create new custom template",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Printf("create new template in %s with language %s support\n", dir, lang)
			return template.CreateCustomTemplate(dir, lang)
		},
	}
	cmd.Flags().StringVarP(&dir, "dir", "d", ".", "template directory to init")
	err := cmd.MarkFlagRequired("dir")
	if err != nil {
		logging.Error().Err(err).Msg("failed to mark flag required")
	}
	cmd.Flags().StringVarP(&lang, "lang", "l", "cpp", "language to init [cpp, go, py, rs, ts, ue]")
	err = cmd.MarkFlagRequired("lang")
	if err != nil {
		logging.Error().Err(err).Msg("failed to mark flag required")
	}
	return cmd
}
