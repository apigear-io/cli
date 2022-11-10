package sim

import (
	"context"
	"path/filepath"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/net"
	"github.com/apigear-io/cli/pkg/net/rpc"
	"github.com/spf13/cobra"
)

type ConsoleHandler struct{}

// Very similar to message handler
func (c ConsoleHandler) HandleMessage(msg rpc.Message) error {
	log.Debug().Msgf("handle message: %+v", msg)
	switch msg.Method {
	case "simu.state":
		log.Info().Msgf("<- state: %v", msg.Params)
	case "simu.call":
		log.Info().Msgf("<- reply[%d]: %v", msg.Id, msg.Params)
	}
	return nil
}

func NewClientCommand() *cobra.Command {
	type ClientOptions struct {
		addr   string
		script string
		sleep  time.Duration
		repeat int
	}
	var options = &ClientOptions{}
	// cmd represents the simCli command
	var cmd = &cobra.Command{
		Use:   "feed",
		Short: "Feed simulation from command line",
		Long:  `Feed simulation calls using JSON documents from command line`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.script = args[0]
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			log.Debug().Msgf("run script %s", options.script)
			switch filepath.Ext(options.script) {
			case ".ndjson":
				emitter := make(chan []byte)
				writer := ConsoleHandler{}
				conn, err := rpc.Dial(ctx, options.addr)
				if err != nil {
					return err
				}
				go func() {
					net.ScanJsonDelimitedFile(options.script, options.sleep, options.repeat, emitter)
				}()
				go func() {
					for data := range emitter {
						log.Info().Msgf("-> %s", data)
						var m rpc.Message
						err := rpc.MessageFromJson(data, &m)
						if err != nil {
							log.Error().Msgf("parse message: %v", err)
							continue
						}
						err = conn.WriteJSON(m)
						if err != nil {
							log.Warn().Msgf("write message: %v", err)
						}
					}
					// wait for all messages to be sent
					log.Info().Msg("wait for all messages sent and exit...")
					time.Sleep(1 * time.Second)
					cancel()
				}()
				go func() {
					for {
						select {
						case <-ctx.Done():
							return
						default:
							var msg rpc.Message
							err := conn.ReadJSON(&msg)
							if err != nil {
								log.Warn().Msgf("read message: %v", err)
								return
							}
							err = writer.HandleMessage(msg)
							if err != nil {
								log.Warn().Msgf("handle message: %v", err)
							}
						}
					}
				}()
			}
			<-ctx.Done()
			return nil
		},
	}
	cmd.Flags().DurationVarP(&options.sleep, "sleep", "", 0, "sleep duration between messages")
	cmd.Flags().StringVarP(&options.addr, "addr", "", "ws://127.0.0.1:8081/ws", "address of the simulation server")
	cmd.Flags().IntVarP(&options.repeat, "repeat", "", 1, "number of times to repeat the script")
	return cmd
}
