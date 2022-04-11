package mon

import (
	"encoding/json"
	"fmt"
	"objectapi/pkg/logger"
	"objectapi/pkg/mon"
	"objectapi/pkg/net"

	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	var log = logger.Get()

	var addr string
	var cmd = &cobra.Command{
		Use:   "start",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			go func() {
				for event := range mon.Emitter() {
					state, err := json.Marshal(event.State)
					if err != nil {
						log.Info("error marshalling state: ", err)
					}
					params, err := json.Marshal(event.Params)
					if err != nil {
						log.Info("error marshalling params: ", err)
					}

					fmt.Printf("<- %s %s %s %s %s %s\n", event.Timestamp.Format("15:04:05"), event.DeviceId, event.Kind, event.Symbol, state, params)
				}
			}()
			s := net.NewServer()
			s.Start(addr)
		},
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", ":5555", "address to listen on")
	return cmd
}
