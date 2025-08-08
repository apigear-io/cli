package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToJniJavaParamString(schema *model.Schema, name string, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("jniJavaParam schema is nil")
	}

	t, err := ToType(schema)
	if err == nil {
		return fmt.Sprintf("%s %s%s", t, prefix, name), nil
	}

	return "xxx", fmt.Errorf("jniJavaParam: unknown schema %s", schema.Dump())
}

func jniJavaParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("jniJavaParam called with nil node")
	}
	return ToJniJavaParamString(&node.Schema, node.Name, prefix)
}
