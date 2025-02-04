package sim

import (
	"os"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/apigear-io/cli/pkg/net"
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/nats-io/nats.go"

	"github.com/spf13/cobra"
)

func NewRunCommand() *cobra.Command {
	var natsURL string
	var script string
	var fn string
	var serve bool

	// cmd represents the simSvr command
	var cmd = &cobra.Command{
		Use:     "run",
		Aliases: []string{"r"},
		Short:   "Run simulation server using an optional scenario file",
		Long: `Simulation server simulates the API backend. 
In its simplest form it just answers every call and all properties are set to default values. 
Using a scenario you can define additional static and scripted data and behavior.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			simman := sim.GetManager()
			netman := net.GetManager()
			if serve {
				err := netman.Start(&net.Options{
					NatsHost:           "localhost",
					NatsPort:           4222,
					HttpAddr:           "localhost:5555",
					SimulationProvider: simman,
				})
				if err != nil {
					log.Error().Err(err).Msg("failed to start network manager")
					return err
				}
				_, err = simman.CreateService(netman.NatsClientURL())
				if err != nil {
					return err
				}
				netman.OnMonitorEvent(func(event *mon.Event) {
					log.Info().Str("source", event.Source).Str("type", event.Type.String()).Str("symbol", event.Symbol).Any("data", event.Data).Msg("received monitor event")
				})
			}
			// run nats server
			client, err := simman.CreateClient(natsURL)
			if err != nil {
				log.Error().Err(err).Msg("failed to create simulation client")
				return err
			}
			if script == "" {
				log.Info().Msg("running simulation server without a script")
			} else {
				log.Info().Str("script", script).Msg("load script file into simulation")
				source, err := os.ReadFile(script)
				if err != nil {
					log.Error().Err(err).Msg("failed to read simulation file")
					return err
				}
				err = client.RunScript("", model.Script{
					Name:   script,
					Source: string(source),
				})
				if err != nil {
					log.Error().Err(err).Msg("failed to run script")
				}
			}
			if fn != "" {
				log.Info().Str("function", fn).Msg("run world function")
				_, err := client.WorldCallFunction("", fn, []any{})
				if err != nil {
					log.Error().Err(err).Msg("failed to run function")
				}
			}
			if serve {
				err := netman.Wait(cmd.Context())
				if err != nil {
					log.Error().Err(err).Msg("failed to wait for services")
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&natsURL, "nats-url", nats.DefaultURL, "nats server to connect to")
	cmd.Flags().StringVar(&script, "script", "", "script to run")
	cmd.Flags().StringVar(&fn, "fn", "main", "function to run")
	cmd.Flags().BoolVar(&serve, "serve", false, "run simulation server in the foreground")
	return cmd
}
