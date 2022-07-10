package main

import (
	"apigear/cmd"
	"os"
)

func main() {
	if cmd.Run() != 0 {
		os.Exit(1)
	}
}
