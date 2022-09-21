package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/iancoleman/strcase"
)

func ToVarString(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("ToVarString node is nil")
	}
	var text string
	schema := &node.Schema
	if !schema.IsArray && schema.KindType == model.TypeBool {
		text = "b"
	}
	return fmt.Sprintf("%s%s%s", text, prefix, strcase.ToCamel(node.Name)), nil
}

func ueVar(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("ueType node is nil")
	}
	return ToVarString(prefix, node)
}
