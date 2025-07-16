package sim

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/apigear-io/cli/pkg/net"
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/tasks"

	"github.com/spf13/cobra"
)

func NewRunCommand() *cobra.Command {
	var fn string
	var port int
	var noServe bool
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
			netman := net.NewManager()
			if err := netman.Start(&net.Options{
				NatsListen:   false,
				HttpAddr:     fmt.Sprintf("localhost:%d", port),
				HttpDisabled: noServe,
			}); err != nil {
				return err
			}
			netman.OnMonitorEvent(func(event *mon.Event) {
				log.Info().Str("source", event.Source).Str("type", event.Type.String()).Str("symbol", event.Symbol).Any("data", event.Data).Msg("received monitor event")
			})
			var simman *sim.Manager
			if !noServe {
				simman = sim.NewManager(sim.ManagerOptions{})
				simman.Start(netman)
			} else {
				simman = sim.NewManager(sim.ManagerOptions{})
			}

			scriptFile := args[0]

			cwd, err := os.Getwd()
			if err != nil {
				log.Error().Err(err).Msg("failed to get current working directory")
				return err
			}

			absFile := filepath.Clean(filepath.Join(cwd, scriptFile))

			// Create task manager and register sim task
			taskManager := tasks.NewTaskManager()
			taskName := "sim-script"
			
			// Create task function that runs the script
			taskFunc := func(ctx context.Context) error {
				return runScript(ctx, simman, netman, absFile, fn)
			}
			
			// Register the task
			taskManager.Register(taskName, map[string]interface{}{
				"script_file": absFile,
				"function": fn,
			}, taskFunc)
			
			ctx := cmd.Context()
			
			if watch {
				log.Info().Str("file", absFile).Msg("watching script file")
				// Use task manager's watch functionality
				if err := taskManager.Watch(ctx, taskName, absFile); err != nil {
					return err
				}
				return netman.Wait(ctx)
			} else {
				// Run once without watching
				if err := taskManager.Run(ctx, taskName); err != nil {
					return err
				}
				return netman.Wait(ctx)
			}
		},
	}
	cmd.Flags().StringVar(&fn, "fn", "main", "function to run")
	cmd.Flags().IntVar(&port, "port", 5555, "protocol server port")
	cmd.Flags().BoolVar(&noServe, "no-serve", false, "disable protocol server")
	cmd.Flags().BoolVar(&watch, "watch", false, "watch for changes in the script file")
	return cmd
}

func runScript(ctx context.Context, sm *sim.Manager, nm *net.NetworkManager, absFile string, fn string) error {
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
