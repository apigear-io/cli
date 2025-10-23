package cmd

import (
	"github.com/apigear-io/cli/pkg/app"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/apigear-io/cli/pkg/server"
	"github.com/spf13/cobra"
)

func NewServeCommand() *cobra.Command {
	var natsHost string // natsURL
	var natsPort int
	var httpAddr string
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "starts apigear server for monitoring and simulation",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info().Msg("starting streams")
			opts := server.Options{
				NatsHost: natsHost,
				NatsPort: natsPort,
				HttpAddr: httpAddr,
			}
			app.WithServer(cmd.Context(), opts, func(s *server.Server) error {
				s.NetworkManager().OnMonitorEvent(func(event *mon.Event) {
					log.Info().Str("source", event.Device).Str("type", event.Type.String()).Str("symbol", event.Symbol).Any("data", event.Data).Msg("received monitor event")
				})
				return nil
			})
			log.Info().Msg("server is running. Press Ctrl+C to stop.")
			return helper.Wait(cmd.Context(), nil)
		},
	}

	cmd.Flags().StringVarP(&natsHost, "nats-host", "n", "localhost", "nats server to connect to")
	cmd.Flags().IntVarP(&natsPort, "nats-port", "p", 4222, "nats server port")
	cmd.Flags().StringVarP(&httpAddr, "http-addr", "a", "localhost:5555", "http server address")
	return cmd
}
