package gen

import (
	"github.com/mark3labs/mcp-go/server"
)

func RegisterMCPTools(s *server.MCPServer) {
	registerGenerateSolutionTool(s)
	registerGenerateExpertTool(s)
}
