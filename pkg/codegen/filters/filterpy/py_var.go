package filterpy

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/codegen/filters/common"
	"github.com/apigear-io/cli/pkg/objmodel"
)

func ToVarString(node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("pyVar node is nil")
	}
	return common.SnakeCaseLower(node.Name), nil
}

func pyVar(node *objmodel.TypedNode) (string, error) {
	return ToVarString(node)
}
