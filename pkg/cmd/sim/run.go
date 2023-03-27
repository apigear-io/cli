package sim

import (
	"context"
	"encoding/json"
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
			simu := sim.NewSimulation()
			simu.OnEvent(func(event *core.SimuEvent) {
				kwargs, _ := json.Marshal(event.KWArgs)
				cmd.Printf("-> %s %s %s %s %s %s\n", event.Timestamp.Format("15:04:05"), event.Type, event.Symbol, event.Name, event.Args, kwargs)
			})
			tm := tasks.NewTaskManager()
			if len(args) == 1 {
				source := args[0]
				tm.On(func(evt *tasks.TaskEvent) {
					log.Debug().Msgf("[%s] task %s: %v", evt.State, evt.Name, evt.Meta)
				})
				run := func(ctx context.Context) error {
					return runScenarioFile(ctx, source, simu)
				}
				meta := map[string]interface{}{
					"scenario": source,
				}
				tm.Register(source, meta, run)
				tm.Run(ctx, source)
				err := tm.Watch(ctx, source, source)
				if err != nil {
					log.Error().Err(err).Str("scenario", source).Msg("run scenario")
				}
				// We need to cancel all tasks on exit
			}
			hub := net.NewSimuHub(ctx, simu)
			s := net.NewHTTPServer()
			s.Router().HandleFunc("/ws", hub.ServeHTTP)
			log.Info().Msgf("simulation server listens on ws://%s/ws", addr)

			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
			defer func() {
				signal.Stop(signalChan)
			}()

			go func() {
				select {
				case <-signalChan:
					log.Debug().Msg("stop simulation server")
					cancel()
					s.Stop()
				case <-ctx.Done():
				}
			}()
			return s.Start(addr)

		},
	}
	cmd.PostRun = func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("stop simulation server")
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:4333", "address to listen on")
	return cmd
}

func runScenarioFile(ctx context.Context, source string, simu *sim.Simulation) error {
	log.Debug().Msgf("run scenario file %s", source)
	doc, err := ReadScenario(source)
	if err != nil {
		return err
	}
	err = simu.LoadScenario(doc.Name, doc)
	if err != nil {
		return err
	}
	err = simu.PlayAllSequences(ctx)
	if err != nil {
		return err
	}
	return nil
}
