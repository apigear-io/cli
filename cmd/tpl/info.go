package tpl

import (
	"os"

	"github.com/apigear-io/cli/pkg/tpl"
	"github.com/spf13/cobra"
)

func NewOpenCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info",
		Short: "Shows information about a template.",
		Long:  `Shows the information and local path of the named template.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			t, err := tpl.GetInfo(name)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(-1)
			}
			cmd.Println("Template Information:")
			cmd.Printf("  Name:\t%s\n", t.Name)
			cmd.Printf("  Path:\t%s\n", t.Path)
			cmd.Printf("  Source:\t%s\n", t.URL)
			cmd.Printf("  Commit:\t%s\n", t.Commit)

		},
	}
	return cmd
}
