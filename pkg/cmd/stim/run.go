package stim

import (
	"context"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/tasks"

	"github.com/spf13/cobra"
)

func NewRunCommand() *cobra.Command {
	var fn string
	var watch bool

	// cmd represents the simSvr command
	var cmd = &cobra.Command{
		Use:     "run",
		Aliases: []string{"r"},
		Args:    cobra.ExactArgs(1),
		Short:   "Run stimulation script using an optional scenario file",
		Long:    `Stimulation script runs scripted calls to a service backend.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			simman := sim.NewManager(sim.ManagerOptions{})

			scriptFile := args[0]

			cwd, err := os.Getwd()
			if err != nil {
				log.Error().Err(err).Msg("failed to get current working directory")
				return err
			}

			absFile := filepath.Clean(filepath.Join(cwd, scriptFile))

			// Create task manager and register sim task
			taskManager := tasks.NewTaskManager()
			taskName := "stim-script"

			// Create task function that runs the script
			taskFunc := func(ctx context.Context) error {
				return runScript(simman, absFile, fn)
			}

			// Register the task
			taskManager.Register(taskName, map[string]interface{}{
				"script_file": absFile,
				"function":    fn,
			}, taskFunc)

			ctx := cmd.Context()

			if watch {
				log.Info().Str("file", absFile).Msg("watching script file")
				// Use task manager's watch functionality
				return taskManager.Watch(ctx, taskName, absFile)
			} else {
				// Run once without watching
				return taskManager.Run(ctx, taskName)
			}
		},
	}
	cmd.Flags().StringVar(&fn, "fn", "main", "function to run")
	cmd.Flags().BoolVar(&watch, "watch", false, "watch for changes in the script file")
	return cmd
}

func runScript(sm *sim.Manager, absFile string, fn string) error {
	log.Info().Str("script", absFile).Msg("load script file into simulation")
	content, err := os.ReadFile(absFile)
	if err != nil {
		log.Error().Err(err).Msg("failed to read script file")
		return err
	}
	script := sim.NewScript(absFile, string(content))
	sm.ScriptRun(script)
	if fn != "" {
		log.Info().Str("function", fn).Msg("run world function")
		sm.FunctionRun(fn, nil)
	}
	// Return immediately after running the script
	// Don't block here - the TaskManager will handle the lifecycle
	return nil
}
