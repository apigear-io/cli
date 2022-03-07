package gen

import (
	"fmt"

	"github.com/spf13/cobra"
)

type GenExpertOptions struct {
	inputs      []string
	outputDir   string
	features    []string
	force       bool
	watch       bool
	templateDir string
}

func NewGenExpertCommand() *cobra.Command {
	o := GenExpertOptions{}

	var cmd = &cobra.Command{
		Use:     "x",
		Aliases: []string{"expert"},
		Short:   "generate code using expert mode",
		Long:    `In expert mode you can individually set your generator options. This is helpful when you do not have a solution document.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("gen %s %s %s %s %t %t\n", o.templateDir, o.inputs, o.outputDir, o.features, o.force, o.watch)
		},
	}
	cmd.Flags().StringVarP(&o.templateDir, "template", "t", "tpl", "template directory")
	cmd.Flags().StringSliceVarP(&o.inputs, "input", "i", []string{"api"}, "input files")
	cmd.Flags().StringVarP(&o.outputDir, "output", "o", "out", "output directory")
	cmd.Flags().StringSliceVarP(&o.features, "feature", "f", []string{"core"}, "features to enable")
	cmd.Flags().BoolVarP(&o.force, "force", "", false, "force overwrite")
	cmd.Flags().BoolVarP(&o.watch, "watch", "", false, "watch for changes")
	return cmd
}
