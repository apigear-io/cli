package sim

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/net"
	"github.com/apigear-io/cli/pkg/net/rpc"
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/sim/actions"
	"github.com/apigear-io/cli/pkg/spec"

	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	var addr string

	// cmd represents the simSvr command
	var cmd = &cobra.Command{
		Use:   "run [scenario to run]",
		Args:  cobra.MaximumNArgs(1),
		Short: "Runs the simulation server using am optional scenario file",
		Long: `The simulation server simulates the API backend. 
In its simplest form it just answers every call and all properties are set to default values. 
Using a scenario you can define additional static and scripted data and behavior.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Debugln("run simulation server")
			log.OnReport(func(entry *log.ReportEntry) {
				cmd.Print(entry.Message)
			})
			var scenario *spec.ScenarioDoc
			if len(args) == 1 {
				file := args[0]
				doc, err := actions.ReadScenario(file)
				if err != nil {
					return fmt.Errorf("failed to read scenario file %s: %v", file, err)
				}
				log.Infof("run simulation from scenario %s\n", file)
				scenario = doc
			}
			simu := sim.NewSimulation()
			err := simu.LoadScenario(scenario)
			if err != nil {
				return fmt.Errorf("failed to load scenario: %v", err)
			}
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
	cmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:5555", "address to listen on")
	return cmd
}
