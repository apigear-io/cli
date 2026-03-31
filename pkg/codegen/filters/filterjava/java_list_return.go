package filterjava

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func ToListReturnString(prefix string, schema *objmodel.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToListReturnString schema is nil")
	}
	if !schema.IsArray {
		return ToReturnString(prefix, schema)
	}
	inner := schema.InnerSchema()
	elementType, err := ToReturnString(prefix, &inner)
	if err != nil {
		return "xxx", fmt.Errorf("javaListReturn element type error: %s", err)
	}
	boxed := toBoxedType(elementType)
	return fmt.Sprintf("List<%s>", boxed), nil
}

func javaListReturn(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("javaListReturn node is nil")
	}
	return ToListReturnString(prefix, &node.Schema)
}
