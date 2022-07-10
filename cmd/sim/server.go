package sim

import (
	"apigear/pkg/log"
	"apigear/pkg/net"
	"apigear/pkg/net/rpc"
	"apigear/pkg/sim"
	"fmt"

	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	var addr string

	// cmd represents the simSvr command
	var cmd = &cobra.Command{
		Use:   "server [scenario to run]",
		Args:  cobra.MaximumNArgs(1),
		Short: "Runs the simulation server using am optional scenario file",
		Long: `The simulation server simulates the service backend. 
In its simplest form it just answers every call and all properties are set to default values. 
Using a scenario you can define additional static and scripted data and behavior.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Debug("run simulation server")
			var scenario *sim.ScenarioDoc
			if len(args) == 1 {
				file := args[0]
				doc, err := sim.ReadScenario(file)
				if err != nil {
					return fmt.Errorf("failed to read scenario file %s: %v", file, err)
				}
				log.Info("run simulation server from scenario ", file)
				scenario = doc
			}
			simu := sim.NewSimulation()
			simu.AddScenario(scenario)
			// start rpc server
			log.Debugf("start rpc hub")
			handler := net.NewSimuRpcHandler(simu)
			hub := rpc.NewHub(handler)
			s := net.NewHTTPServer()
			s.Router().HandleFunc("/ws/", hub.HandleWebsocketRequest)
			log.Debugf("handle ws rpc server on %s/ws/", addr)
			return s.Start(addr)
		},
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", ":5555", "address to listen on")
	return cmd
}
