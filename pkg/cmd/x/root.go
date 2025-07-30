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
	}
	cmd.AddCommand(NewDocsCommand())
	cmd.AddCommand(NewJson2YamlCommand())
	cmd.AddCommand(NewYaml2JsonCommand())
	cmd.AddCommand(NewYaml2IdlCommand())
	return cmd
}

func init() {
}
