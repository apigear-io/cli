package filterjs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToVarString(node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("jsVar node is nil")
	}
	return node.Name, nil
}

func jsVar(node *model.TypedNode) (string, error) {
	return ToVarString(node)
}
