package sdk

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// genCmd represents the generate command
	var cmd = &cobra.Command{
		Use:     "gen",
		Aliases: []string{"generate", "g"},
		Short:   "generates code",
		Long:    `provides a set of commands to generate code`,
	}
	cmd.AddCommand(NewExpertCommand())
	cmd.AddCommand(NewSolutionCommand())
	return cmd
}
