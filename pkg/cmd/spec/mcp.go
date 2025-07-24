package spec

import (
	"context"
	"fmt"

	"github.com/apigear-io/cli/pkg/spec"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterMCPTools(s *server.MCPServer) {
	registerSpecCheckTool(s)
	registerSpecShemaTool(s)
}

func registerSpecCheckTool(s *server.MCPServer) {
	specCheckTool := mcp.NewTool("specificationCheck",
		mcp.WithDescription("Load and validate API specification(module, solution, rules) files"),
		mcp.WithString("file", mcp.Required(), mcp.Description("Path to specification file")),
	)
	s.AddTool(specCheckTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		file, err := request.RequireString("file")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		var result string
		spec, err := spec.CheckFile(file)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		if spec.Valid() {
			result = fmt.Sprintf("valid: %s\n", file)
		} else {
			result = fmt.Sprintf("invalid: %s\n", file)
			for _, desc := range spec.Errors {
				result += fmt.Sprintf("file: %s\n", file)
				result += fmt.Sprintf("%s\n", desc.String())
			}
		}
		return mcp.NewToolResultText(string(result)), nil
	})
}

func registerSpecShemaTool(s *server.MCPServer) {
	specCheckTool := mcp.NewTool("specificationSchema",
		mcp.WithDescription("Show the schema for module, solution, rules documents in either yaml or json format"),
		mcp.WithString("type", mcp.Required(), mcp.Description("Document type (module, solution, rules)")),
		mcp.WithString("format", mcp.Required(), mcp.Description("Output schema format (yaml, json)")),
	)
	s.AddTool(specCheckTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		docType, err := request.RequireString("type")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		var documentType spec.DocumentType
		switch docType {
		case "module":
			documentType = spec.DocumentTypeModule
		case "solution":
			documentType = spec.DocumentTypeSolution
		case "rules":
			documentType = spec.DocumentTypeRules
		default:
			return mcp.NewToolResultError(fmt.Sprintf("%s is not a valid document type", docType)), nil
		}

		outputformat, err := request.RequireString("format")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		var schemaFormat spec.SchemaFormat
		switch outputformat {
		case "yaml":
			schemaFormat = spec.SchemaFormatYaml
		case "json":
			schemaFormat = spec.SchemaFormatJson
		default:
			return mcp.NewToolResultError(fmt.Sprintf("%s is not a valid output format", outputformat)), nil
		}

		schema, err := spec.ShowSchemaFile(documentType, schemaFormat)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(fmt.Sprintf("%s\n", *schema)), nil
	})
}
