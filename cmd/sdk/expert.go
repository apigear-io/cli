package sdk

import (
	"objectapi/pkg/sol"
	"objectapi/pkg/spec"
	"os"

	"github.com/spf13/cobra"
)

type ExpertOptions struct {
	inputs      []string
	outputDir   string
	features    []string
	force       bool
	watch       bool
	templateDir string
}

func NewExpertCommand() *cobra.Command {
	options := &ExpertOptions{}

	cmd := &cobra.Command{
		Use:     "x",
		Aliases: []string{"expert"},
		Short:   "generate code using expert mode",
		Long:    `In expert mode you can individually set your generator options. This is helpful when you do not have a solution document.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Debugf("expert mode: %v", options)
			doc := spec.SolutionDoc{
				Schema: "apigear.solution/1.0",
				Layers: []spec.SolutionLayer{
					{
						Inputs:   options.inputs,
						Output:   options.outputDir,
						Template: options.templateDir,
						Features: options.features,
						Force:    options.force,
					},
				},
			}
			log.Debugf("solution doc: %v", doc)
			rootDir, err := os.Getwd()
			log.Debugf("rootDir: %s", rootDir)
			if err != nil {
				log.Fatalf("failed to get current directory: %s", err)
			}
			proc := sol.NewSolutionRunner(rootDir, doc)
			proc.Run()
		},
	}
	cmd.Flags().StringVarP(&options.templateDir, "template", "t", "tpl", "template directory")
	cmd.Flags().StringSliceVarP(&options.inputs, "input", "i", []string{"api"}, "input files")
	cmd.Flags().StringVarP(&options.outputDir, "output", "o", "out", "output directory")
	cmd.Flags().StringSliceVarP(&options.features, "feature", "f", []string{"core"}, "features to enable")
	cmd.Flags().BoolVarP(&options.force, "force", "", false, "force overwrite")
	cmd.Flags().BoolVarP(&options.watch, "watch", "", false, "watch for changes")
	cmd.MarkFlagRequired("input")
	cmd.MarkFlagRequired("output")
	cmd.MarkFlagRequired("template")
	return cmd
}
