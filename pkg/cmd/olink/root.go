package olink

import (
	"github.com/apigear-io/objectlink-core-go/cli"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// cmd represents the olink command
	var cmd = &cobra.Command{
		Use:     "olink",
		Aliases: []string{"ol"},
		Short:   "Start an ObjectLink REPL to test the olink protocol",
		Long: `The olink command starts an interactive REPL (Read-Eval-Print Loop) for testing 
the ObjectLink protocol. It provides commands to connect to servers, link to objects,
invoke methods, set properties, and observe signals.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Run the ObjectLink CLI REPL
			cli.Run()
			return nil
		},
	}
	return cmd
}
