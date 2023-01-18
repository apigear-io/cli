package filtergo

import (
	"strings"

	"github.com/apigear-io/cli/pkg/model"
)

func formatDoc(doc string) string {

	sb := strings.Builder{}
	lines := strings.Split(doc, "\n")

	// for each line
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		sb.WriteString("// ")
		sb.WriteString(line)
		sb.WriteString("\n")
	}
	return sb.String()
}

func goDoc(node *model.NamedNode) (string, error) {
	return formatDoc(node.Description), nil
}
