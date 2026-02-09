package filterts

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func ToVarString(node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("tsVar node is nil")
	}
	return node.Name, nil
}

func tsVar(node *objmodel.TypedNode) (string, error) {
	return ToVarString(node)
}
