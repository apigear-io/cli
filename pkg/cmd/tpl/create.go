package tpl

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/tpl"
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
			return tpl.CreateCustomTemplate(dir, lang)
		},
	}
	cmd.Flags().StringVarP(&dir, "dir", "d", ".", "template directory to init")
	err := cmd.MarkFlagRequired("dir")
	if err != nil {
		log.Error().Err(err).Msg("failed to mark flag required")
	}
	cmd.Flags().StringVarP(&lang, "lang", "l", "cpp", "language to init [cpp, go, py, rs, ts, ue]")
	err = cmd.MarkFlagRequired("lang")
	if err != nil {
		log.Error().Err(err).Msg("failed to mark flag required")
	}
	return cmd
}
