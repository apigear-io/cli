package tpl

import (
	"github.com/apigear-io/cli/pkg/tpl"
	"github.com/spf13/cobra"
)

func NewInfoCommand() *cobra.Command {
	var dir string
	var cmd = &cobra.Command{
		Use:   "info",
		Short: "Display template information",
		Long:  `Display template information`,
		RunE: func(cmd *cobra.Command, args []string) error {
			info, err := tpl.Info(dir)
			if err != nil {
				return err
			}
			cmd.Println("# template info")
			cmd.Println()
			cmd.Println("## rules document")
			cmd.Println()
			cmd.Println(info.Rules)
			cmd.Println("## template files")
			cmd.Println()
			for _, file := range info.Files {
				cmd.Printf("- %s\n", file)
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&dir, "dir", "d", ".", "template directory")
	return cmd
}
