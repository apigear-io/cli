package filterjava

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToListParamString(prefix string, schema *model.Schema, name string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToListParamString schema is nil")
	}
	if schema.IsArray {
		inner := schema.InnerSchema()
		elementType, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("javaListParam element type error: %s", err)
		}
		boxed := toBoxedType(elementType)
		return fmt.Sprintf("List<%s> %s", boxed, name), nil
	}
	return ToParamString(prefix, schema, name)
}

func javaListParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("javaListParam node is nil")
	}
	return ToListParamString(prefix, &node.Schema, node.Name)
}
