package cmd

import (
	"github.com/spf13/cobra"
)

// genCmd represents the generate command
var genCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g", "gen"},
	Short:   "generates code",
	Long:    `provides a set of commands to generate code`,
}

func init() {
	rootCmd.AddCommand(genCmd)
}
