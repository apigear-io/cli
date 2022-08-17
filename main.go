package main

import (
	"fmt"
	"os"

	"github.com/apigear-io/cli/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Printf("apigear-cli version %s, commit %s, built at %s\n", version, commit, date)
	if cmd.Run() != 0 {
		os.Exit(1)
	}
}
