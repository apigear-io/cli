package filterrust

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(prefix string, schema *model.Schema, node *model.TypedNode) (string, error) {
	t := schema.Type
	name, err := ToVarString(node)
	if err != nil {
		return "xxx", fmt.Errorf("ToParamString inner value error: %s", err)
	}
	if schema.IsArray {
		inner := *schema
		inner.IsArray = false
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("ToParamString inner value error: %s", err)
		}
		return fmt.Sprintf("%s: &Vec<%s>", name, ret), nil
	}
	switch t {
	case "string":
		return fmt.Sprintf("%s: &str", name), nil
	case "int":
		return fmt.Sprintf("%s: i32", name), nil
	case "int32":
		return fmt.Sprintf("%s: i32", name), nil
	case "int64":
		return fmt.Sprintf("%s: i64", name), nil
	case "float":
		return fmt.Sprintf("%s: f32", name), nil
	case "float32":
		return fmt.Sprintf("%s: f32", name), nil
	case "float64":
		return fmt.Sprintf("%s: f64", name), nil
	case "bool":
		return fmt.Sprintf("%s: bool", name), nil
	}
	e := schema.Module.LookupEnum(t)
	if e != nil {
		return fmt.Sprintf("%s: %sEnum", name, e.Name), nil
	}
	s := schema.Module.LookupStruct(t)
	if s != nil {
		return fmt.Sprintf("%s: &%s", name, s.Name), nil
	}
	i := schema.Module.LookupInterface(t)
	if i != nil {
		return fmt.Sprintf("%s: &%s", name, i.Name), nil
	}
	return "xxx", fmt.Errorf("ToParamString: unknown type %s", t)
}

func rustParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("rustParam node is nil")
	}
	return ToParamString(prefix, &node.Schema, node)
}
