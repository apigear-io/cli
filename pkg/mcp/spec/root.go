package spec

import (
	"github.com/mark3labs/mcp-go/server"
)

func RegisterMCPTools(s *server.MCPServer) {
	registerSpecCheckTool(s)
	registerSpecShemaTool(s)
}
