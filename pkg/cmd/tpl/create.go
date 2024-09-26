package tpl

import (
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
	cmd.MarkFlagRequired("dir")
	cmd.Flags().StringVarP(&lang, "lang", "l", "cpp", "language to init [cpp, go, py, rs, ts, ue]")
	cmd.MarkFlagRequired("lang")
	return cmd
}
