package tpl

import (
	"github.com/apigear-io/cli/pkg/tpl"
	"github.com/spf13/cobra"
)

func NewInitCommand() *cobra.Command {
	var dir string
	var lang string
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize a template",
		Long:  `Initialize a template`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Printf("initializing template %s using language %s\n", dir, lang)
			return tpl.InitTemplate(dir, lang)
		},
	}
	cmd.Flags().StringVarP(&dir, "dir", "d", ".", "template directory to init")
	cmd.MarkFlagRequired("dir")
	cmd.Flags().StringVarP(&lang, "lang", "l", "cpp", "language to init [cpp, go, py, rs, ts, ue]")
	cmd.MarkFlagRequired("lang")
	return cmd
}
