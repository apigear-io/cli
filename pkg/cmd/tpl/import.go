package tpl

import (
	"github.com/apigear-io/cli/pkg/tpl"

	"github.com/spf13/cobra"
)

func NewImportCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "import [git-url]",
		Short: "import template",
		Long:  `import template from a git-url`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				for _, dst := range args {
					cmd.Printf("importing template from %s\n", dst)
					vcs, err := tpl.ImportTemplate(dst)
					if err != nil {
						cmd.PrintErrln(err)
						continue
					}
					cmd.Printf("imported template %s\n", vcs.FullName)
				}
			}
		},
	}
	return cmd
}
