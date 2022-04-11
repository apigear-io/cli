package main

import (
	"fmt"
	"objectapi/cmd"
	"os"
)

func main() {
	rootCmd := cmd.NewRootCommand()
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
