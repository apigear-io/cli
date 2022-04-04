package main

import (
	"fmt"
	"objectapi/cmd"
	"objectapi/pkg/logger"
	"os"
)

var log = logger.Get()

func main() {
	log.Info("starting objectapi")
	rootCmd := cmd.NewRootCommand()
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
