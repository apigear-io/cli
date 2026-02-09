package filterjava

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func ToParamString(prefix string, schema *apimodel.Schema, name string) (string, error) {
	if schema.IsArray {
		inner := schema.InnerSchema()
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("javaParam inner value error: %s", err)
		}
		return fmt.Sprintf("%s[] %s", ret, name), nil
	} else {
		ret, err := ToReturnString(prefix, schema)
		if err != nil {
			return "xxx", fmt.Errorf("javaParam inner value error: %s", err)
		}
		return fmt.Sprintf("%s %s", ret, name), nil
	}
}

func javaParam(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("javaParam node is nil")
	}
	return ToParamString(prefix, &node.Schema, node.Name)
}
