package filterjava

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToListDefaultString(schema *model.Schema, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToListDefaultString schema is nil")
	}
	if !schema.IsArray {
		return ToDefaultString(schema, prefix)
	}
	return "new ArrayList<>()", nil
}

func javaListDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("javaListDefault node is nil")
	}
	return ToListDefaultString(&node.Schema, prefix)
}
