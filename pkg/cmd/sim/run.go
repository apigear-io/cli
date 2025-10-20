package sim

import (
	"context"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/tasks"
	"github.com/nats-io/nats.go"

	"github.com/spf13/cobra"
)

func NewRunCommand() *cobra.Command {
	var fn string
	var natsServer string
	var watch bool

	// cmd represents the simSvr command
	var cmd = &cobra.Command{
		Use:     "run",
		Aliases: []string{"r"},
		Args:    cobra.ExactArgs(1),
		Short:   "Run simulation server using an optional scenario file",
		Long: `Simulation server simulates the API backend. 
In its simplest form it just answers every call and all properties are set to default values. 
Using a scenario you can define additional static and scripted data and behavior.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			absFile, err := filepath.Abs(args[0])
			if err != nil {
				return err
			}

			sim.WithClient(cmd.Context(), natsServer, func(ctx context.Context, client *sim.Client) error {
				taskManager := tasks.NewTaskManager()
				taskName := "sim-script"

				// Create task function that runs the script
				taskFunc := func(ctx context.Context) error {
					resp, err := client.RunScript(absFile)
					if err != nil {
						log.Error().Err(err).Msg("failed to run script")
						return err
					}
					if resp.Error != "" {
						log.Error().Err(err).Str("error", resp.Error).Msg("failed to run script")
						return err
					}
					log.Info().Str("file", absFile).Msg("script executed")
					return nil
				}

				// Register the task
				taskManager.Register(taskName, map[string]interface{}{
					"script_file": absFile,
					"function":    fn,
				}, taskFunc)

				if watch {
					log.Info().Str("file", absFile).Msg("watching script file")
					// Use task manager's watch functionality
					if err := taskManager.Watch(ctx, taskName, absFile); err != nil {
						return err
					}
				} else {
					// Run once without watching
					if err := taskManager.Run(ctx, taskName); err != nil {
						return err
					}
				}
				return helper.Wait(ctx, nil)
			})
			return nil
		},
	}
	cmd.Flags().StringVar(&fn, "fn", "main", "function to run")
	cmd.Flags().StringVar(&natsServer, "nats-server", nats.DefaultURL, "nats server url")
	cmd.Flags().BoolVar(&watch, "watch", false, "watch for changes in the script file")
	return cmd
}
