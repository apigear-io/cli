package mon

import (
	"apigear/pkg/log"
	"apigear/pkg/mon"
	"apigear/pkg/net"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	var addr string
	var cmd = &cobra.Command{
		Use:   "run",
		Short: "run the monitor server",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Debugf("start server on %s", addr)
			go func() {
				for event := range mon.Emitter() {
					data, err := json.Marshal(event.Data)
					if err != nil {
						log.Info("error marshalling data: ", err)
					}
					fmt.Printf("-> %s %s %s %s %s\n", event.Timestamp.Format("15:04:05"), event.Source, event.Type, event.Symbol, data)
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
