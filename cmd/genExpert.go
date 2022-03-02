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

func newGenExportCmd() *cobra.Command {
	var inputs []string
	var outputDir string
	var features []string
	var force bool
	var watch bool
	var templateDir string

	var cmd = &cobra.Command{
		Use:     "expert",
		Aliases: []string{"x"},
		Short:   "generate code using expert mode",
		Long:    `In expert mode you can individually set your generator options. This is helpful when you do not have a solution document.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("gen %s %s %s %s %t %t\n", templateDir, inputs, outputDir, features, force, watch)
		},
	}
	cmd.Flags().StringVarP(&templateDir, "template", "t", "tpl", "template directory")
	cmd.Flags().StringSliceVarP(&inputs, "input", "i", []string{"api"}, "input files")
	cmd.Flags().StringVarP(&outputDir, "output", "o", "out", "output directory")
	cmd.Flags().StringSliceVarP(&features, "feature", "f", []string{"core"}, "features to enable")
	cmd.Flags().BoolVarP(&force, "force", "", false, "force overwrite")
	cmd.Flags().BoolVarP(&watch, "watch", "", false, "watch for changes")
	return cmd
}

func init() {
	genCmd.AddCommand(newGenExportCmd())
}
