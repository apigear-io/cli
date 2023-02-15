package sim

import (
	"context"
	"fmt"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/net"
	"github.com/apigear-io/cli/pkg/sim"
	"github.com/apigear-io/cli/pkg/sim/actions"
	"github.com/apigear-io/cli/pkg/spec"
	"github.com/apigear-io/cli/pkg/tasks"

	"github.com/spf13/cobra"
)

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

func RunSimuServer(ctx context.Context, addr string, simu *sim.Simulation) error {
	hub := net.NewSimuHub(ctx, simu)
	s := net.NewHTTPServer()
	s.Router().HandleFunc("/ws", hub.ServeHTTP)
	go func() {
		err := s.Start(addr)
		if err != nil {
			log.Error().Err(err).Msg("start simulation server")
		}
	}()
	go func() {
		<-ctx.Done()
		s.Stop()
	}()
	return nil
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
			log.Info().Msgf("run simulation server")
			simu := sim.NewSimulation()

			if len(args) == 1 {
				source := args[0]
				tm := tasks.New()
				run := func(ctx context.Context) error {
					return runScenarioFile(source, simu)
				}
				tm.Register(source, run)
				err := tm.Watch(ctx, source, source)
				if err != nil {
					log.Error().Err(err).Str("scenario", source).Msg("run scenario")
				}
			}

			// start rpc server
			log.Info().Msgf("olink server ws://%s/ws", addr)
			err := RunSimuServer(ctx, addr, simu)
			if err != nil {
				log.Error().Err(err).Msg("start rpc server")
			}
			helper.WaitForInterrupt(cancel)
			return nil
		},
	}
	cmd.PostRun = func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("stop simulation server")
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:4333", "address to listen on")
	return cmd
}

func runScenarioFile(source string, simu *sim.Simulation) error {
	log.Debug().Msgf("run scenario file %s", source)
	doc, err := ReadScenario(source)
	if err != nil {
		return err
	}
	err = simu.LoadScenario(doc.Name, doc)
	if err != nil {
		return err
	}
	err = simu.PlayAllSequences(context.Background())
	if err != nil {
		return err
	}
	return nil
}
