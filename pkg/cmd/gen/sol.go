package gen

import (
	"context"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sol"
	"github.com/apigear-io/cli/pkg/spec"
	"github.com/apigear-io/cli/pkg/tasks"

	"github.com/spf13/cobra"
)

func NewSolutionCommand() *cobra.Command {
	var source string
	var watch bool
	var force bool
	var cmd = &cobra.Command{
		Use:     "solution [solution-file]",
		Short:   "Generate SDK using a solution document",
		Aliases: []string{"sol", "s"},
		Args:    cobra.ExactArgs(1),
		Long: `A solution is a yaml document which describes different layers. 
Each layer defines the input module files, output directory and the features to enable, 
as also the other options. To create a demo module or solution use the 'project create' command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info().Msgf("generating solution %s", args[0])
			source = args[0]
			return RunGenerateSolution(source, watch, force)
		},
	}
	cmd.Flags().BoolVarP(&watch, "watch", "", false, "watch solution file for changes")
	cmd.Flags().BoolVarP(&force, "force", "", false, "force overwrite")
	return cmd
}

func RunGenerateSolution(solutionPath string, watch bool, force bool) error {
	result, err := spec.CheckFileAndType(solutionPath, spec.DocumentTypeSolution)
	if err != nil {
		return err
	}
	if !result.Valid() {
		for _, err := range result.Errors {
			log.Warn().Msgf("source %s at %s error %s", solutionPath, err.Field, err.Description)
		}
		return nil
	}
	runner := sol.NewRunner()
	runner.OnTask(func(evt *tasks.TaskEvent) {
		log.Debug().Msgf("[%s] task %s: %v", evt.State, evt.Name, evt.Meta)
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if watch {
		err := runner.WatchSource(ctx, solutionPath, force)
		if err != nil {
			log.Error().Err(err).Msg("watching solution file")
			cancel()
		}
		helper.WaitForInterrupt(cancel)
	} else {
		err = runner.RunSource(ctx, solutionPath, force)
		if err != nil {
			return err
		}
	}
	return nil
}
