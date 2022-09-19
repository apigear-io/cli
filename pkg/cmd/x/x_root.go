package x

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "x",
		Aliases: []string{"experimental"},
		Short:   "Experimental commands",
		Long:    `Command which are under development or experimental`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
	cmd.AddCommand(NewDocsCommand())
	cmd.AddCommand(NewJson2YamlCommand())
	cmd.AddCommand(NewYaml2JsonCommand())
	return cmd
}

func init() {
}
