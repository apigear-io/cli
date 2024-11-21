package sim

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/net"
	"github.com/apigear-io/cli/pkg/simjs"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func runNatsServer(port int) (*net.NatsServer, error) {
	s, err := net.NewNatsServer(net.NatsServerOptions{Port: port})
	if err != nil {
		return nil, err
	}
	log.Info().Str("url", s.ClientURL()).Msg("simulation server started")
	err = s.Start()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func runSimService(s *net.NatsServer) error {
	nc, err := nats.Connect(s.ClientURL())
	if err != nil {
		return err
	}
	defer nc.Close()
	conn, err := simjs.NewConn(nc)
	if err != nil {
		return err
	}
	u := simjs.NewUniverse(conn)
	simjs.NewService(conn, u)
	return nil
}

func waitForSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
}

func NewServerCommand() *cobra.Command {
	var port int

	var cmd = &cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "Run simulation server",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := runNatsServer(port)
			if err != nil {
				return err
			}
			err = runSimService(s)
			if err != nil {
				return err
			}
			waitForSignal()
			return s.Stop()
		},
	}
	cmd.Flags().IntVarP(&port, "port", "p", nats.DefaultPort, "port to listen on")
	return cmd
}
