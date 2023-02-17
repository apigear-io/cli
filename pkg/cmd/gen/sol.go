package gen

import (
	"context"
	"fmt"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sol"
	"github.com/apigear-io/cli/pkg/spec"

	"github.com/spf13/cobra"
)

func NewSolutionCommand() *cobra.Command {
	var file string
	var watch bool
	var exec string
	var cmd = &cobra.Command{
		Use:     "solution [solution-file]",
		Short:   "Generate SDK using a solution document",
		Aliases: []string{"sol", "s"},
		Args:    cobra.ExactArgs(1),
		Long: `A solution is a yaml document which describes different layers. 
Each layer defines the input module files, output directory and the features to enable, 
as also the other options. To create a demo module or solution use the 'project create' command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			file = args[0]
			result, err := spec.CheckFileAndType(file, spec.DocumentTypeSolution)
			if err != nil {
				return err
			}
			if !result.Valid() {
				for _, err := range result.Errors() {
					entry := err.Field() + ": " + err.Description()
					cmd.Println(entry)
				}
				return fmt.Errorf("solution file is not valid")
			}
			doc, err := sol.ReadSolutionDoc(file)
			if err != nil {
				return err
			}
			runner := sol.NewRunner()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			err = runner.RunDoc(ctx, file, doc)
			if err != nil {
				return err
			}

			if watch {
				err := runner.StartWatch(ctx, file, doc)
				if err != nil {
					log.Error().Err(err).Msg("watching solution file")
					cancel()
				}
				go helper.WaitForInterrupt(cancel)
			}
			return nil
		},
	}
	cmd.Flags().BoolVarP(&watch, "watch", "", false, "watch solution file for changes")
	cmd.Flags().StringVarP(&exec, "exec", "", "", "execute a command after generation")
	return cmd
}
