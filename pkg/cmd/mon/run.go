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
		Use:     "run",
		Aliases: []string{"r", "start"},
		Short:   "Run the monitor server",
		Long:    `The monitor server runs on a HTTP port and listens for API calls.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			mon.Emitter.On(func(event *mon.Event) {
				data, _ := json.Marshal(event.Data)
				cmd.Printf("-> %s %s %s %s %s\n", event.Timestamp.Format("15:04:05"), event.Source, event.Type, event.Symbol, data)
			})
			s := net.NewHTTPServer()
			s.Router().Post("/monitor/{source}", net.HandleMonitorRequest)
			log.Info().Msgf("handle monitor request on %s/monitor/{source}", addr)
			return s.Start(addr)
		},
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:5555", "address to listen on")
	return cmd
}
