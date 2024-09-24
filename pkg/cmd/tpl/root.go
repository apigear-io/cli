package tpl

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	// cmd represents the tpl command
	cmd := &cobra.Command{
		Use:     "template",
		Aliases: []string{"tpl", "t"},
		Short:   "template management",
		Long:    `template management`,
	}
	cmd.AddCommand(NewInitCommand())
	cmd.AddCommand(NewLintCommand())
	cmd.AddCommand(NewInfoCommand())
	cmd.AddCommand(NewPublishCommand())
	return cmd
}
