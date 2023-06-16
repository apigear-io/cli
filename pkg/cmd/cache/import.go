package cache

import (
	"github.com/apigear-io/cli/pkg/repos"

	"github.com/spf13/cobra"
)

func NewImportCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "import [git-url] [version]",
		Short: "import template from git-url and version",
		Long:  `import template from a git-url`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			version := args[1]
			cmd.Printf("importing template from %s\n", url)
			fqn, err := repos.Cache.Install(url, version)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			cmd.Printf("imported template as %s\n", fqn)
		},
	}
	return cmd
}
