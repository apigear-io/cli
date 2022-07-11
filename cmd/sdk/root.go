package sdk

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// genCmd represents the generate command
	var cmd = &cobra.Command{
		Use:     "sdk",
		Aliases: []string{"gen", "g"},
		Short:   "generates sdk",
		Long:    `provides a set of commands to generate code`,
	}
	cmd.AddCommand(NewExpertCommand(), NewSolutionCommand())
	return cmd
}
