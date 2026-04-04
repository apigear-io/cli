package filterjava

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToListAsyncReturnString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToListAsyncReturnString schema is nil")
	}
	if !schema.IsArray {
		return ToAsyncReturnString(prefix, schema)
	}
	inner := schema.InnerSchema()
	elementType, err := ToReturnString(prefix, &inner)
	if err != nil {
		return "xxx", fmt.Errorf("javaListAsyncReturn element type error: %s", err)
	}
	boxed := toBoxedType(elementType)
	return fmt.Sprintf("CompletableFuture<List<%s>>", boxed), nil
}

func javaListAsyncReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("javaListAsyncReturn node is nil")
	}
	return ToListAsyncReturnString(prefix, &node.Schema)
}
