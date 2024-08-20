package tpl

import "github.com/spf13/cobra"

func NewLintCommand() *cobra.Command {
	var dir string
	var cmd = &cobra.Command{
		Use:   "lint",
		Short: "Lint a template",
		Long:  `Lint a template`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("lint template")
		},
	}
	cmd.Flags().StringVarP(&dir, "dir", "d", ".", "template directory")
	cmd.MarkFlagRequired("dir")
	return cmd
}
