package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/iancoleman/strcase"
)

func ToVarString(node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("ToVarString node is nil")
	}
	return strcase.ToLowerCamel(node.Name), nil
}

func ToPublicVarString(node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("ToPublicVarString node is nil")
	}
	return strcase.ToCamel(node.Name), nil
}

func goVar(node *model.TypedNode) (string, error) {
	return ToVarString(node)
}

func goPublicVar(node *model.TypedNode) (string, error) {
	return ToPublicVarString(node)
}
