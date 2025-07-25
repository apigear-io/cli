package spec

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/spec"

	"github.com/spf13/cobra"
)

func NewShowCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "schema",
		Aliases: []string{"s", "show", "view"},
		Short:   "Show schema for module, solution, rules documents",
		Long:    `Show the schema for module, solutions, rules documents in either yaml or json form`,
		RunE: func(cmd *cobra.Command, args []string) error {
			docType, _ := cmd.Flags().GetString("type")
			format, _ := cmd.Flags().GetString("format")

			var documentType spec.DocumentType
			switch docType {
			case "module":
				documentType = spec.DocumentTypeModule
			case "solution":
				documentType = spec.DocumentTypeSolution
			case "rules":
				documentType = spec.DocumentTypeRules
			default:
				documentType = spec.DocumentTypeModule
			}

			var schemaFormat spec.SchemaFormat
			switch format {
			case "yaml":
				schemaFormat = spec.SchemaFormatYaml
			case "json":
				schemaFormat = spec.SchemaFormatJson
			default:
				schemaFormat = spec.SchemaFormatYaml
			}

			schema, err := spec.ShowSchemaFile(documentType, schemaFormat)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", *schema)
			return nil
		},
	}

	cmd.Flags().StringP("type", "t", "module", "Document type (module, solution, rules)")
	cmd.Flags().StringP("format", "f", "yaml", "Output schema format (yaml, json)")
	return cmd
}
