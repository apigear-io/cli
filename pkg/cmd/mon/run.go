package mon

import (
	"encoding/json"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	var addr string
	var serve bool
	var natsURL string
	var cmd = &cobra.Command{
		Use:     "run",
		Aliases: []string{"r", "start"},
		Short:   "Run the monitor server",
		Long:    `The monitor server runs on a HTTP port and listens for API calls.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !serve {
				nc, err := nats.Connect(natsURL)
				if err != nil {
					return err
				}
				_, err = nc.Subscribe(mon.MonitorSubject+".>", func(msg *nats.Msg) {
					var event mon.Event
					err := json.Unmarshal(msg.Data, &event)
					if err != nil {
						log.Error().Msgf("unmarshal event: %v", err)
						return
					}
					log.Info().Msgf("<- %s %s %s %s %s", event.Timestamp.Format("15:04:05"), event.Source, event.Type, event.Symbol, event.Data)
				})
				if err != nil {
					log.Error().Err(err).Msg("failed to subscribe to monitor events")
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:5555", "address to listen on")
	cmd.Flags().BoolVarP(&serve, "serve", "s", false, "serve the monitor")
	cmd.Flags().StringVarP(&natsURL, "nats-url", "n", nats.DefaultURL, "nats server to connect to")
	return cmd
}
