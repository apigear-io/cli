package cmd

import (
	"fmt"
	"io/ioutil"
	"objectapi/pkg/log"
	"objectapi/pkg/spec"

	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "lint documents",
	Long:  `Lint documents.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var fn = args[0]
		doc, err := ioutil.ReadFile(fn)
		if err != nil {
			panic(err)
		}
		t, err := spec.GetDocumentType(fn)
		if err != nil {
			panic(err)
		}
		result, err := spec.LintDocumentFromString(t, doc)
		if err != nil {
			panic(err)
		}
		if result.Valid() {
			fmt.Println("valid")
		} else {
			log.Info("invalid")
			for _, desc := range result.Errors() {
				log.Info(desc.String())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lintCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lintCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lintCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
