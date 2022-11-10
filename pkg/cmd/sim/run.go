package sim

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/net"
	"github.com/apigear-io/cli/pkg/net/rpc"
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/sim/actions"
	"github.com/apigear-io/cli/pkg/spec"

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
		Short: "Run simulation server using an optional scenario file",
		Long: `Simulation server simulates the API backend. 
In its simplest form it just answers every call and all properties are set to default values. 
Using a scenario you can define additional static and scripted data and behavior.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			go handleSignal(cancel)
			log.Info().Msgf("run simulation server")
			var doc *spec.ScenarioDoc
			if len(args) == 1 {
				file := args[0]
				result, err := spec.CheckFile(file)
				if err != nil {
					return err
				}
				if !result.Valid() {
					log.Error().Msgf("scenario file is not valid: %v", result.Errors())
					for _, err := range result.Errors() {
						entry := fmt.Sprintf("%s: %s", err.Field(), err.Description())
						log.Error().Msg(entry)
					}
				}
				aDoc, err := actions.ReadScenario(file)
				if err != nil {
					log.Error().Msgf("read scenario file: %v", err)
				}
				if aDoc.Name == "" {
					aDoc.Name = file
				}
				log.Info().Msgf("run simulation from scenario %s", file)
				doc = aDoc
			}
			simu := sim.NewSimulation()
			if doc != nil {
				err := simu.LoadScenario(doc.Name, doc)
				if err != nil {
					return err
				}
				go func() {
					err = simu.PlayAllSequences()
					if err != nil {
						log.Error().Msgf("play scenario: %v", err)
					}
				}()
			}
			// start rpc server
			log.Info().Msg("start rpc hub")
			handler := net.NewSimuRpcHandler(simu)
			hub := rpc.NewHub(ctx)
			go func() {
				for req := range hub.Requests() {
					err := handler.HandleMessage(req)
					if err != nil {
						log.Error().Err(err).Msg("handle rpc request")
					}
				}
			}()
			s := net.NewHTTPServer()
			s.Router().HandleFunc("/ws", hub.ServeHTTP)
			log.Info().Msgf("rpc server ws://%s/ws", addr)
			go func() {
				err := s.Start(addr)
				if err != nil {
					log.Error().Err(err).Msg("start rpc server")
				}
			}()
			<-ctx.Done()
			log.Info().Msgf("shutting down rpc hub")
			return nil
		},
	}
	cmd.PostRun = func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("stop simulation server")
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:8081", "address to listen on")
	return cmd
}
