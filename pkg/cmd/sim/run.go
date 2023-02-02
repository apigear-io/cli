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
	"github.com/apigear-io/cli/pkg/sim/core"
	"github.com/apigear-io/cli/pkg/spec"

	"github.com/spf13/cobra"
)

func handleSignal(cancel func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs
	cancel()
}

func ReadScenario(file string) (*spec.ScenarioDoc, error) {
	result, err := spec.CheckFile(file)
	if err != nil {
		return nil, err
	}
	if !result.Valid() {
		log.Error().Msgf("scenario file is not valid: %v", result.Errors())
		for _, err := range result.Errors() {
			entry := fmt.Sprintf("%s: %s", err.Field(), err.Description())
			log.Error().Msg(entry)
		}
	}
	doc, err := actions.ReadScenario(file)
	if err != nil {
		return nil, err
	}
	if doc.Name == "" {
		doc.Name = file
	}
	return doc, nil
}

func StartSimuServer(ctx context.Context, addr string, simu *sim.Simulation) error {
	hub := net.NewSimuHub(ctx, simu)
	s := net.NewHTTPServer()
	s.Router().HandleFunc("/ws", hub.ServeHTTP)
	return s.Start(addr)
}

func NewServerCommand() *cobra.Command {
	var addr string

	// cmd represents the simSvr command
	var cmd = &cobra.Command{
		Use:     "run [scenario to run]",
		Aliases: []string{"r"},
		Args:    cobra.MaximumNArgs(1),
		Short:   "Run simulation server using an optional scenario file",
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
				aDoc, err := ReadScenario(args[0])
				if err != nil {
					log.Error().Msgf("read scenario file: %v", err)
				}
				doc = aDoc
			}
			simu := sim.NewSimulation()
			simu.OnEvent(func(event *core.SimuEvent) {
				objectId := event.Symbol
				if event.Name != "" {
					objectId = fmt.Sprintf("%s/%s", objectId, event.Name)
				}
				log.Info().Msgf("%s on %s: args=%+v kwargs=%+v", event.Type, objectId, event.Args, event.KWArgs)
			})
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
			log.Info().Msgf("olink server ws://%s/ws", addr)
			go func() {
				err := StartSimuServer(ctx, addr, simu)
				if err != nil {
					log.Error().Err(err).Msg("start rpc server")
				}
			}()
			<-ctx.Done()
			log.Debug().Msgf("shutting down rpc hub")
			return nil
		},
	}
	cmd.PostRun = func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("stop simulation server")
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:4333", "address to listen on")
	return cmd
}
