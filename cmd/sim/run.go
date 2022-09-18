package sim

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/net"
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/sim/actions"
	"github.com/apigear-io/cli/pkg/spec"
	"github.com/apigear-io/wsrpc/rpc"

	"github.com/spf13/cobra"
)

func handleSignal(cancel func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs
	cancel()
}

func NewServerCommand() *cobra.Command {
	var addr string

	// cmd represents the simSvr command
	var cmd = &cobra.Command{
		Use:   "run [scenario to run]",
		Args:  cobra.MaximumNArgs(1),
		Short: "Runs the simulation server using an optional scenario file",
		Long: `The simulation server simulates the API backend. 
In its simplest form it just answers every call and all properties are set to default values. 
Using a scenario you can define additional static and scripted data and behavior.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			go handleSignal(cancel)
			log.Infoln("run simulation server")
			log.OnReport(func(entry *log.ReportEntry) {
				cmd.Println(entry.Message)
			})
			var doc *spec.ScenarioDoc
			if len(args) == 1 {
				file := args[0]
				result, err := spec.CheckFile(file)
				if err != nil {
					return err
				}
				if !result.Valid() {
					fmt.Printf("scenario file %s is not valid\n", file)
					for _, err := range result.Errors() {
						entry := fmt.Sprintf("%s: %s", err.Field(), err.Description())
						fmt.Printf("- %s\n", entry)
					}
					return nil
				}
				aDoc, err := actions.ReadScenario(file)
				if err != nil {
					return fmt.Errorf("failed to read scenario file %s: %v", file, err)
				}
				if aDoc.Name == "" {
					aDoc.Name = file
				}
				log.Infof("run simulation from scenario %s", file)
				doc = aDoc
			}
			simu := sim.NewSimulation()
			if doc != nil {
				err := simu.LoadScenario(doc.Name, doc)
				if err != nil {
					return fmt.Errorf("failed to load scenario: %v", err)
				}
				go func() {
					err = simu.PlayAllSequences()
					if err != nil {
						log.Errorf("failed to play scenario: %v", err)
					}
				}()
			}
			// start rpc server
			log.Info("start rpc hub")
			handler := net.NewSimuRpcHandler(simu)
			hub := rpc.NewHub(ctx)
			go func() {
				for req := range hub.Requests() {
					err := handler.HandleMessage(req)
					if err != nil {
						log.Error(err)
					}
				}
			}()
			s := net.NewHTTPServer()
			s.Router().HandleFunc("/ws", hub.ServeHTTP)
			log.Infof("rpc server ws://%s/ws", addr)
			go func() {
				err := s.Start(addr)
				if err != nil {
					log.Error(err)
				}
			}()
			<-ctx.Done()
			log.Infof("shutting down rpc hub")
			return nil
		},
	}
	cmd.PostRun = func(cmd *cobra.Command, args []string) {
		log.Debugln("stop simulation server")
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:8081", "address to listen on")
	return cmd
}
