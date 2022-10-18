package filterpy

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/ettle/strcase"
)

func ToVarString(node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ToVarString node is nil")
	}
	return strcase.ToSnake(node.Name), nil
}

func pyVar(node *model.TypedNode) (string, error) {
	return ToVarString(node)
}
