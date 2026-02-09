package filterpy

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/codegen/filters/common"
	"github.com/apigear-io/cli/pkg/apimodel"
)

func ToVarString(node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("pyVar node is nil")
	}
	return common.SnakeCaseLower(node.Name), nil
}

func pyVar(node *apimodel.TypedNode) (string, error) {
	return ToVarString(node)
}
