package gen

import (
	"context"
	"fmt"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/sol"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterMCPTools(s *server.MCPServer) {
	registerGenerateSolutionTool(s)
	registerGenerateExpertTool(s)
}

func registerGenerateSolutionTool(s *server.MCPServer) {
	genSolutionTool := mcp.NewTool("generateSolution",
		mcp.WithDescription("Generate SDK based on a solution document"),
		mcp.WithString("solution", mcp.Required(), mcp.Description("Path to solution file")),
		mcp.WithString("force", mcp.Description("Force overwrite (true/false)")),
		mcp.WithString("watch", mcp.Description("Watch for changes (true/false). This keeps the process running.")),
	)
	s.AddTool(genSolutionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		solutionPath, err := request.RequireString("solution")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		forceEnabled := false
		if force, err := request.RequireString("force"); err == nil && force == "true" {
			forceEnabled = true
		}
		watchEnabled := false
		if watch, err := request.RequireString("watch"); err == nil && watch == "true" {
			watchEnabled = true
		}

		err = runGenerateSolution(solutionPath, watchEnabled, forceEnabled)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(fmt.Sprintf("Successfully ran code generation for solution file: %s", solutionPath)), nil
	})
}

func registerGenerateExpertTool(s *server.MCPServer) {
	options := &ExpertOptions{}

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
		options.inputs = append(options.inputs, input)

		output, err := request.RequireString("output")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		options.outputDir = output

		template, err := request.RequireString("template")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		options.templateDir = template

		if features, err := request.RequireString("features"); err == nil && features != "" {
			options.features = append(options.features, features)
		}
		options.force = false
		if force, err := request.RequireString("force"); err == nil && force == "true" {
			options.force = true
		}
		options.watch = false
		if watch, err := request.RequireString("watch"); err == nil && watch == "true" {
			options.watch = true
		}

		doc := makeSolution(options)
		if err := doc.Validate(); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid solution document: %s", err.Error())), nil
		}
		runner := sol.NewRunner()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if err := runner.RunDoc(ctx, doc.RootDir, doc); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		if options.watch {
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
