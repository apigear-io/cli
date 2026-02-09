package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
	"github.com/ettle/strcase"
)

func ToVarString(node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ToVarString node is nil")
	}
	return node.Name, nil
}

func ToPublicVarString(node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ToPublicVarString node is nil")
	}
	return strcase.ToPascal(node.Name), nil
}

func goVar(node *objmodel.TypedNode) (string, error) {
	return ToVarString(node)
}

func goPublicVar(node *objmodel.TypedNode) (string, error) {
	return ToPublicVarString(node)
}
