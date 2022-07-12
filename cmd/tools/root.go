package tools

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tools",
		Aliases: []string{"t"},
		Short:   "assorted tool commands",
		Long:    `General purpose tools used for various reasons.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
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
