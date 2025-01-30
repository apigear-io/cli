package sim

import (
	"os"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/nats-io/nats.go"

	"github.com/spf13/cobra"
)

func NewRunCommand() *cobra.Command {
	var natsURL string
	var script string
	var fn string

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
			return nil
		},
	}
	cmd.Flags().StringVar(&natsURL, "nats-url", nats.DefaultURL, "nats server to connect to")
	cmd.Flags().StringVar(&script, "script", "", "script to run")
	cmd.Flags().StringVar(&fn, "fn", "main", "function to run")
	return cmd
}
