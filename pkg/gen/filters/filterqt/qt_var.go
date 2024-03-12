package filterqt

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToVarString(node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("qtVar node is nil")
	}
	return node.Name, nil
}

func qtVar(node *model.TypedNode) (string, error) {
	return ToVarString(node)
}
