package filtercpp

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func ToVarString(node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ToVarString node is nil")
	}
	return node.Name, nil
}

func cppVar(node *objmodel.TypedNode) (string, error) {
	return ToVarString(node)
}
