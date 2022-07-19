package mon

import (
	"encoding/json"
	"fmt"

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
		Long:  `A monitor server runs on a HTTP port and listens for API calls.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			log.Debugf("start server on %s", addr)
			go func() {
				for event := range mon.Emitter() {
					data, err := json.Marshal(event.Data)
					if err != nil {
						log.Info("error marshalling data: ", err)
					}
					cmd.Printf("-> %s %s %s %s %s\n", event.Timestamp.Format("15:04:05"), event.Source, event.Type, event.Symbol, data)
				}
			}()
			s := net.NewHTTPServer()
			s.Router().Post("/monitor/{source}/", net.HandleMonitorRequest)
			log.Infof("handle monitor request on %s/monitor/{source}", addr)
			err := s.Start(addr)
			if err != nil {
				return fmt.Errorf("error starting server: %s", err)
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", ":5555", "address to listen on")
	return cmd
}
