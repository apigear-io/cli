package gen

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// genCmd represents the generate command
	var cmd = &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen", "g"},
		Short:   "Generate code from APIs",
		Long:    `generate API SDKs from API descriptions using templates`,
	}
	cmd.AddCommand(NewExpertCommand(), NewSolutionCommand())
	return cmd
}
