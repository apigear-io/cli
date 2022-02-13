package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// template dir
var templateDir string
var inputs []string
var outputDir string
var features []string
var force bool
var watch bool

// genCmd represents the generate command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generates code",
	Long:  `generates code from a set of api modules and templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gen %s %s %s %s %t %t\n", templateDir, inputs, outputDir, features, force, watch)
	},
}

func init() {
	// add template dir flag
	genCmd.Flags().StringVarP(&templateDir, "template", "t", "tpl", "template directory")
	// add inputs flag
	genCmd.Flags().StringSliceVarP(&inputs, "input", "i", []string{"api"}, "input files")
	// add output dir flag
	genCmd.Flags().StringVarP(&outputDir, "output", "o", "out", "output directory")
	// add features flag
	genCmd.Flags().StringSliceVarP(&features, "feature", "f", []string{"core"}, "features to enable")
	// add force flag
	genCmd.Flags().BoolVarP(&force, "force", "", false, "force overwrite")
	// add watch flag
	genCmd.Flags().BoolVarP(&watch, "watch", "", false, "watch for changes")

	rootCmd.AddCommand(genCmd)

}
