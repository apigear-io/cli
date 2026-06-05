package filtercs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

// ToAsyncReturnString wraps the C# return type in a System.Threading.Tasks.Task.
// A void return becomes a non-generic Task; everything else becomes Task<T>.
// Unlike Java's CompletableFuture, C# generics accept value types directly, so
// no boxing of primitives is required.
func ToAsyncReturnString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToAsyncReturnString schema is nil")
	}
	if schema.KindType == model.TypeVoid && !schema.IsArray {
		return "Task", nil
	}
	inner, err := ToReturnString(prefix, schema)
	if err != nil {
		return "xxx", fmt.Errorf("csAsyncReturn type error: %s", err)
	}
	return fmt.Sprintf("Task<%s>", inner), nil
}

func csAsyncReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("csAsyncReturn node is nil")
	}
	return ToAsyncReturnString(prefix, &node.Schema)
}
