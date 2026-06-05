package filtercs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

// ToParamString renders a single C# parameter as "<type> <name>". C# passes
// reference types and List<T> by reference already, so no const/ref qualifier
// is needed (unlike C++).
func ToParamString(prefix string, schema *model.Schema, name string) (string, error) {
	ret, err := ToReturnString(prefix, schema)
	if err != nil {
		return "xxx", fmt.Errorf("csParam type error: %s", err)
	}
	return fmt.Sprintf("%s %s", ret, name), nil
}

func csParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("csParam node is nil")
	}
	return ToParamString(prefix, &node.Schema, node.Name)
}
