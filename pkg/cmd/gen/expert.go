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
	Inputs      []string
	OutputDir   string
	Features    []string
	Force       bool
	Watch       bool
	TemplateDir string
}

func NewExpertCommand() *cobra.Command {
	options := &ExpertOptions{}

	cmd := &cobra.Command{
		Use:     "expert",
		Aliases: []string{"x"},
		Short:   "Generate code using expert mode",
		Long:    `in expert mode you can individually set your generator options. This is helpful when you do not have a solution document.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			doc := MakeSolution(options)
			if err := doc.Validate(); err != nil {
				return fmt.Errorf("invalid solution document: %w", err)
			}
			runner := sol.NewRunner()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			if err := runner.RunDoc(ctx, doc.RootDir, doc); err != nil {
				return err
			}

			if options.Watch {
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
	cmd.Flags().StringVarP(&options.TemplateDir, "template", "t", "tpl", "template directory")
	cmd.Flags().StringSliceVarP(&options.Inputs, "input", "i", []string{"apigear"}, "input files")
	cmd.Flags().StringVarP(&options.OutputDir, "output", "o", "out", "output directory")
	cmd.Flags().StringSliceVarP(&options.Features, "features", "f", []string{"all"}, "features to enable")
	cmd.Flags().BoolVarP(&options.Force, "force", "", false, "force overwrite")
	cmd.Flags().BoolVarP(&options.Watch, "watch", "", false, "watch for changes")
	Must(cmd.MarkFlagRequired("input"))
	Must(cmd.MarkFlagRequired("output"))
	Must(cmd.MarkFlagRequired("template"))
	return cmd
}

func MakeSolution(options *ExpertOptions) *spec.SolutionDoc {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("get current working directory")
	}
	return &spec.SolutionDoc{
		RootDir: rootDir,
		Targets: []*spec.SolutionTarget{
			{
				Inputs:   options.Inputs,
				Output:   options.OutputDir,
				Template: options.TemplateDir,
				Features: options.Features,
				Force:    options.Force,
			},
		},
	}
}
