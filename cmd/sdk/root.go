package sdk

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// genCmd represents the generate command
	var cmd = &cobra.Command{
		Use:     "sdk",
		Aliases: []string{"gen", "g"},
		Short:   "SDK code generation",
		Long:    `Code generation using templates for SDK creation.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	cmd.AddCommand(NewExpertCommand(), NewSolutionCommand())
	return cmd
}
