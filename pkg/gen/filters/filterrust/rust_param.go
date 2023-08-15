package filterrust

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(prefix string, schema *model.Schema, name string) (string, error) {
	t := schema.Type
	if schema.IsArray {
		inner := *schema
		inner.IsArray = false
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("ToParamString inner value error: %s", err)
		}
		return fmt.Sprintf("%s: &[%s]", common.SnakeCaseLower(name), ret), nil
	}
	switch t {
	case "string":
		return fmt.Sprintf("%s: &String", common.SnakeCaseLower(name)), nil
	case "int":
		return fmt.Sprintf("%s: i32", common.SnakeCaseLower(name)), nil
	case "int32":
		return fmt.Sprintf("%s: i32", common.SnakeCaseLower(name)), nil
	case "int64":
		return fmt.Sprintf("%s: i64", common.SnakeCaseLower(name)), nil
	case "float":
		return fmt.Sprintf("%s: f32", common.SnakeCaseLower(name)), nil
	case "float32":
		return fmt.Sprintf("%s: f32", common.SnakeCaseLower(name)), nil
	case "float64":
		return fmt.Sprintf("%s: f64", common.SnakeCaseLower(name)), nil
	case "bool":
		return fmt.Sprintf("%s: bool", common.SnakeCaseLower(name)), nil
	}
	e := schema.Module.LookupEnum(t)
	if e != nil {
		return fmt.Sprintf("%s: %sEnum", common.SnakeCaseLower(name), e.Name), nil
	}
	s := schema.Module.LookupStruct(t)
	if s != nil {
		return fmt.Sprintf("%s: &%s", common.SnakeCaseLower(name), s.Name), nil
	}
	i := schema.Module.LookupInterface(t)
	if i != nil {
		return fmt.Sprintf("%s: &%s", common.SnakeCaseLower(name), i.Name), nil
	}
	return "xxx", fmt.Errorf("ToParamString: unknown type %s", t)
}

func rustParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("rustParam node is nil")
	}
	return ToParamString(prefix, &node.Schema, node.Name)
}
