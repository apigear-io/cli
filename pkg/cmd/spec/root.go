package spec

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// cmd represents the sim command
	var cmd = &cobra.Command{
		Use:     "spec",
		Aliases: []string{"simulate", "s"},
		Short:   "Load and validate files",
		Long:    `Specification defines the file formats used inside apigear`,
	}
	cmd.AddCommand(NewCheckCommand())
	return cmd
}
