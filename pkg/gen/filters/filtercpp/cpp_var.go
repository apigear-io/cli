package filtercpp

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/ettle/strcase"
)

func ToVarString(node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ToVarString node is nil")
	}
	return strcase.ToCamel(node.Name), nil
}

func cppVar(node *model.TypedNode) (string, error) {
	return ToVarString(node)
}
