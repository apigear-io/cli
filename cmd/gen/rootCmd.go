package gen

import (
	"github.com/spf13/cobra"
)

// genCmd represents the generate command
var RootCmd = &cobra.Command{
	Use:     "gen",
	Aliases: []string{"generate", "g"},
	Short:   "generates code",
	Long:    `provides a set of commands to generate code`,
}

func init() {
	RootCmd.AddCommand(NewGenExpertCommand())
	RootCmd.AddCommand(NewGenSolution())
}
