package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/net/rpc"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	flag.Parse()
	hub := rpc.NewHub(ctx)
	s := http.Server{Addr: *addr, Handler: hub}
	defer func() {
		err := s.Shutdown(ctx)
		if err != nil {
			log.Error().Msgf("error shutting down server: %s", err)
		}
	}()
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
		fmt.Println("server closed")
	}()
	go func() {
		for req := range hub.Requests() {
			var m rpc.Message
			err := req.AsJSON(&m)
			if err != nil {
				fmt.Printf("error: %v", err)
				continue
			}
			reply := rpc.MakeNotify("pong", m.Params)
			err = req.ReplyJSON(reply)
			if err != nil {
				fmt.Printf("error: %v", err)
				continue
			}
		}
	}()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs // wait for SIGINT or SIGTERM
}
