package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
	"github.com/ettle/strcase"
)

func ToVarString(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueVar node is nil")
	}
	var text string
	schema := &node.Schema
	if !schema.IsArray && schema.KindType == objmodel.TypeBool {
		text = "b"
	}
	return fmt.Sprintf("%s%s%s", text, prefix, strcase.ToPascal(node.Name)), nil
}

func ueVar(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueVar node is nil")
	}
	return ToVarString(prefix, node)
}
