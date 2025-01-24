package cmd

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/apigear-io/cli/pkg/net"
	"github.com/apigear-io/cli/pkg/sim"
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
			simman := sim.GetManager()
			netman := net.GetManager()
			err := netman.Start(&net.Options{
				NatsHost:           natsHost,
				NatsPort:           natsPort,
				HttpAddr:           httpAddr,
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
			return netman.Wait(cmd.Context())
		},
	}

	cmd.Flags().StringVarP(&natsHost, "nats-host", "n", "localhost", "nats server to connect to")
	cmd.Flags().IntVarP(&natsPort, "nats-port", "p", 4222, "nats server port")
	cmd.Flags().StringVarP(&httpAddr, "http-addr", "a", "localhost:5555", "http server address")
	return cmd
}
