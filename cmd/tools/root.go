package tools

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tools",
		Aliases: []string{"t"},
		Short:   "assorted tools",
		Long:    `General purpose tools used for various reasons.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
	cmd.AddCommand(NewCheckCommand())
	cmd.AddCommand(NewDocsCommand())
	cmd.AddCommand(NewJson2YamlCommand())
	cmd.AddCommand(NewYaml2JsonCommand())
	return cmd
}

func init() {
}
