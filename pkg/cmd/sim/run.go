package sim

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/apigear-io/cli/pkg/net"
	"github.com/apigear-io/cli/pkg/sim"

	"github.com/spf13/cobra"
)

func NewRunCommand() *cobra.Command {
	var fn string
	var port int
	var noServe bool

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
			cwd, err := os.Getwd()
			if err != nil {
				log.Error().Err(err).Msg("failed to get current working directory")
				return err
			}
			netman := net.NewManager()
			netman.Start(&net.Options{
				NatsListen:   false,
				HttpAddr:     fmt.Sprintf("localhost:%d", port),
				HttpDisabled: noServe,
			})
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

			absFile := filepath.Clean(filepath.Join(cwd, scriptFile))

			log.Info().Str("script", absFile).Msg("load script file into simulation")
			content, err := os.ReadFile(absFile)
			if err != nil {
				log.Error().Err(err).Msg("failed to read script file")
				return err
			}
			script := sim.NewScript(absFile, string(content))
			simman.ScriptRun(script)
			if fn != "" {
				log.Info().Str("function", fn).Msg("run world function")
				simman.FunctionRun(fn, nil)
			}
			// wait for server
			err = netman.Wait(cmd.Context())
			if err != nil {
				log.Error().Err(err).Msg("failed to wait for services")
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&fn, "fn", "main", "function to run")
	cmd.Flags().IntVar(&port, "port", 5555, "protocol server port")
	cmd.Flags().BoolVar(&noServe, "no-serve", false, "disable protocol server")
	cmd.MarkFlagRequired("script")
	return cmd
}
