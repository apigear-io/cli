package filterjs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func ToVarString(node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("jsVar node is nil")
	}
	return node.Name, nil
}

func jsVar(node *objmodel.TypedNode) (string, error) {
	return ToVarString(node)
}
