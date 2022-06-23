package main

import (
	"apigear/cmd"
	"fmt"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

type VersionInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func main() {
	rootCmd := cmd.NewRootCommand()
	rootCmd.Version = fmt.Sprintf("%s-%s-%s", version, commit, date)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
