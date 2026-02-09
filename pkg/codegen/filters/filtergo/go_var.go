package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
	"github.com/ettle/strcase"
)

func ToVarString(node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ToVarString node is nil")
	}
	return node.Name, nil
}

func ToPublicVarString(node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ToPublicVarString node is nil")
	}
	return strcase.ToPascal(node.Name), nil
}

func goVar(node *apimodel.TypedNode) (string, error) {
	return ToVarString(node)
}

func goPublicVar(node *apimodel.TypedNode) (string, error) {
	return ToPublicVarString(node)
}
