package tpl

import (
	"github.com/apigear-io/cli/pkg/repos"

	"github.com/spf13/cobra"
)

func NewCleanCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "clean",
		Short: "clean all templates from the local cache",
		Run: func(cmd *cobra.Command, _ []string) {
			err := repos.Cache.Clean()
			if err != nil {
				cmd.PrintErrln(err)
			} else {
				cmd.Println("template cache cleaned")
			}
		},
	}
	return cmd
}
