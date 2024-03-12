package filterpy

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

func ToVarString(node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("pyVar node is nil")
	}
	return common.SnakeCaseLower(node.Name), nil
}

func pyVar(node *model.TypedNode) (string, error) {
	return ToVarString(node)
}
