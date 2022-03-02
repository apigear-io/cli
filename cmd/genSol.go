/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newGenSolution() *cobra.Command {
	// genSolCmd represents the genRun command
	var cmd = &cobra.Command{
		Use:     "solution [file to run]",
		Short:   "generate code using a solution",
		Aliases: []string{"sol", "s"},
		Long: `A solution is a yaml document which describes different layers. 
Each layer defines the input module files, output directory and the features to enable, 
as also the other options. To create a demo module or solution use the 'project create' command.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("genRun called", args[0])
		},
	}

	return cmd
}

func init() {
	genCmd.AddCommand(newGenSolution())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genRunCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genRunCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
