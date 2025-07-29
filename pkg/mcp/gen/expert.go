package gen

import (
	"context"
	"fmt"

	"github.com/apigear-io/cli/pkg/cmd/gen"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/sol"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerGenerateExpertTool(s *server.MCPServer) {
	options := &gen.ExpertOptions{}

	genExpertTool := mcp.NewTool("generateExpert",
		mcp.WithDescription("Generate code using expert mode with individual generator options"),
		mcp.WithString("input", mcp.Required(), mcp.Description("Input module files (comma-separated)")),
		mcp.WithString("output", mcp.Required(), mcp.Description("Output directory")),
		mcp.WithString("template", mcp.Required(), mcp.Description("Template directory")),
		mcp.WithString("features", mcp.Description("Features to enable (comma-separated, defaults to 'all')")),
		mcp.WithString("force", mcp.Description("Force overwrite (true/false)")),
		mcp.WithString("watch", mcp.Description("Watch for changes (true/false). This keeps the process running.")),
	)
	s.AddTool(genExpertTool, func(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		input, err := request.RequireString("input")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		options.Inputs = append(options.Inputs, input)

		output, err := request.RequireString("output")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		options.OutputDir = output

		template, err := request.RequireString("template")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		options.TemplateDir = template

		if features, err := request.RequireString("features"); err == nil && features != "" {
			options.Features = append(options.Features, features)
		}
		options.Force = false
		if force, err := request.RequireString("force"); err == nil && force == "true" {
			options.Force = true
		}
		options.Watch = false
		if watch, err := request.RequireString("watch"); err == nil && watch == "true" {
			options.Watch = true
		}

		doc := gen.MakeSolution(options)
		if err := doc.Validate(); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid solution document: %s", err.Error())), nil
		}
		runner := sol.NewRunner()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if err := runner.RunDoc(ctx, doc.RootDir, doc); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		if options.Watch {
			err := runner.WatchDoc(ctx, doc.RootDir, doc)
			if err != nil {
				cancel()
				return mcp.NewToolResultError(fmt.Sprintf("error watching solution file: %s", err.Error())), nil
			}
			helper.WaitForInterrupt(cancel)
		}
		return mcp.NewToolResultText("Successfully ran code generation with expert options"), nil
	})
}
