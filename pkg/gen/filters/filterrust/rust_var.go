package filterrust

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

func ToVarString(node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ToVarString node is nil")
	}
	return common.SnakeCaseLower(node.Name), nil
}

func rustVar(node *model.TypedNode) (string, error) {
	return ToVarString(node)
}
