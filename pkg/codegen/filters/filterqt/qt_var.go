package filterqt

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func ToVarString(node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("qtVar node is nil")
	}
	return node.Name, nil
}

func qtVar(node *apimodel.TypedNode) (string, error) {
	return ToVarString(node)
}
