package mon

import (
	"encoding/json"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/mon"
	"github.com/apigear-io/cli/pkg/net"

	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	var addr string
	var cmd = &cobra.Command{
		Use:   "run",
		Short: "Run the monitor server",
		Long:  `The monitor server runs on a HTTP port and listens for API calls.`,
		Run: func(cmd *cobra.Command, _ []string) {
			log.Debug().Msgf("start server on %s", addr)
			go func() {
				for event := range mon.Emitter() {
					data, err := json.Marshal(event.Data)
					if err != nil {
						log.Info().Err(err).Msg("error marshalling data: ")
					}
					cmd.Printf("-> %s %s %s %s %s\n", event.Timestamp.Format("15:04:05"), event.Source, event.Type, event.Symbol, data)
				}
			}()
			s := net.NewHTTPServer()
			s.Router().Post("/monitor/{source}/", net.HandleMonitorRequest)
			log.Info().Msgf("handle monitor request on %s/monitor/{source}", addr)
			err := s.Start(addr)
			if err != nil {
				log.Error().Msgf("failed to start server: %v", err)
			}
		},
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:5555", "address to listen on")
	return cmd
}
