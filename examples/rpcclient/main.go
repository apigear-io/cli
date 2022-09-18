package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/net/rpc"
)

var addr = flag.String("addr", "ws://localhost:8080", "ws service address")

func main() {
	log.Debug("starting")
	log.Info("starting")
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, err := rpc.Dial(ctx, *addr)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				var msg rpc.Message
				err := conn.ReadJSON(&msg)
				if err != nil {
					fmt.Printf("error: %v\n", err)
					return
				}
				count, ok := msg.Params["count"].(float64)
				if !ok {
					fmt.Printf("error: %v\n", err)
					return
				}
				fmt.Printf("<- %s %v\n", msg.Method, count)
			}
		}
	}()

	go func() {
		count := 0
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				count++
				msg := rpc.MakeNotify("ping", map[string]any{"count": count})
				fmt.Printf("-> ping %d\n", count)
				err = conn.WriteJSON(msg)
			}
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs // wait for SIGINT or SIGTERM
}
