package gen

import (
	"context"
	"fmt"
	"os"

	"github.com/apigear-io/cli/pkg/helper"
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
			if err := doc.Validate(); err != nil {
				return fmt.Errorf("invalid solution document: %w", err)
			}
			runner := sol.NewRunner()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			if err := runner.RunDoc(ctx, doc.RootDir, doc); err != nil {
				return err
			}

			if options.watch {
				err := runner.WatchDoc(ctx, doc.RootDir, doc)
				if err != nil {
					log.Error().Err(err).Msg("watching solution file")
					cancel()
				}
				helper.WaitForInterrupt(cancel)
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&options.templateDir, "template", "t", "tpl", "template directory")
	cmd.Flags().StringSliceVarP(&options.inputs, "input", "i", []string{"apigear"}, "input files")
	cmd.Flags().StringVarP(&options.outputDir, "output", "o", "out", "output directory")
	cmd.Flags().StringSliceVarP(&options.features, "features", "f", []string{"all"}, "features to enable")
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
		Targets: []*spec.SolutionTarget{
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
