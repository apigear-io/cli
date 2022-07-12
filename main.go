package main

import (
	"os"

	"github.com/apigear-io/cli/cmd"
)

func main() {
	if cmd.Run() != 0 {
		os.Exit(1)
	}
}
