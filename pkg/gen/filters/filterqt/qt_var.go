package filterqt

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/iancoleman/strcase"
)

func ToVarString(node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ToVarString node is nil")
	}
	return strcase.ToLowerCamel(node.Name), nil
}

func qtVar(node *model.TypedNode) (string, error) {
	return ToVarString(node)
}
