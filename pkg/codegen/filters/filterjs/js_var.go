package filterjs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func ToVarString(node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("jsVar node is nil")
	}
	return node.Name, nil
}

func jsVar(node *apimodel.TypedNode) (string, error) {
	return ToVarString(node)
}
