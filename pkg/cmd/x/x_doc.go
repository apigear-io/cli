package x

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func NewDocsCommand() *cobra.Command {
	var force bool
	cmd := &cobra.Command{
		Use:   "doc",
		Short: "exports cli docs as markdown",
		Long:  `export the cli docs as markdown document into a dir`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dir := "docs/cli"
			if len(args) > 0 {
				dir = args[0]
			}
			if force {
				err := os.MkdirAll(dir, 0755)
				if err != nil {
					log.Fatalf("failed to create dir: %v", err)
				}
			}
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				cmd.Printf("dir '%s' does not exist\n", dir)
				os.Exit(1)
			}
			cmd.Printf("exporting docs to %s\n", dir)
			err := doc.GenMarkdownTree(cmd.Root(), dir)
			if err != nil {
				log.Fatalf("error exporting docs: %v", err)
			}
		},
	}
	cmd.Flags().BoolVarP(&force, "force", "f", false, "make dir and overwrite existing files")
	return cmd
}
