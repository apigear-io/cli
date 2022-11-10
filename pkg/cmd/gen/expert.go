package gen

import (
	"fmt"
	"os"
	"sync"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sol"
	"github.com/apigear-io/cli/pkg/spec"

	"github.com/spf13/cobra"
)

func Must(err error) {
	if err != nil {
		log.Fatal().Err(err).Msg("parse command line")
	}
}

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
		Use:     "expert",
		Aliases: []string{"x"},
		Short:   "Generate code using expert mode",
		Long:    `in expert mode you can individually set your generator options. This is helpful when you do not have a solution document.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			doc := makeSolution(options)
			err := doc.Validate()
			if err != nil {
				return fmt.Errorf("invalid solution document: %w", err)
			}
			runner := sol.NewRunner()

			if options.watch {
				// TODO: how to watch from a document and not from a file?
				var wg sync.WaitGroup
				wg.Add(1)
				done, err := runner.StartWatch(doc.RootDir, doc)
				if err != nil {
					return err
				}
				wg.Wait()
				done <- true
			} else {
				err := runner.RunDoc(doc.RootDir, doc)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&options.templateDir, "template", "t", "tpl", "template directory")
	cmd.Flags().StringSliceVarP(&options.inputs, "input", "i", []string{"api"}, "input files")
	cmd.Flags().StringVarP(&options.outputDir, "output", "o", "out", "output directory")
	cmd.Flags().StringSliceVarP(&options.features, "feature", "f", []string{"core"}, "features to enable")
	cmd.Flags().BoolVarP(&options.force, "force", "", false, "force overwrite")
	cmd.Flags().BoolVarP(&options.watch, "watch", "", false, "watch for changes")
	Must(cmd.MarkFlagRequired("input"))
	Must(cmd.MarkFlagRequired("output"))
	Must(cmd.MarkFlagRequired("template"))
	return cmd
}

func makeSolution(options *ExpertOptions) *spec.SolutionDoc {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("get current working directory")
	}
	return &spec.SolutionDoc{
		Schema:  "apigear.solution/1.0",
		RootDir: rootDir,
		Layers: []*spec.SolutionLayer{
			{
				Inputs:   options.inputs,
				Output:   options.outputDir,
				Template: options.templateDir,
				Features: options.features,
				Force:    options.force,
			},
		},
	}
}
