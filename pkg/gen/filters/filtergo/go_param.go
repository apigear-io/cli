package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(prefix string, schema *model.Schema, name string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToParamString schema is nil")
	}
	t := schema.Type
	if schema.IsArray {
		inner := *schema
		inner.IsArray = false
		innerValue, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("ToParamString inner value error: %s", err)
		}
		return fmt.Sprintf("%s []%s", name, innerValue), nil
	}
	switch t {
	case "string":
		return fmt.Sprintf("%s string", name), nil
	case "int":
		return fmt.Sprintf("%s int32", name), nil
	case "int32":
		return fmt.Sprintf("%s int32", name), nil
	case "int64":
		return fmt.Sprintf("%s int64", name), nil
	case "float":
		return fmt.Sprintf("%s float32", name), nil
	case "float32":
		return fmt.Sprintf("%s float32", name), nil
	case "float64":
		return fmt.Sprintf("%s float64", name), nil
	case "bool":
		return fmt.Sprintf("%s bool", name), nil
	}
	e := schema.Module.LookupEnum(t)
	if e != nil {
		return fmt.Sprintf("%s %s%s", name, prefix, e.Name), nil
	}
	s := schema.Module.LookupStruct(t)
	if s != nil {
		return fmt.Sprintf("%s %s%s", name, prefix, s.Name), nil
	}
	i := schema.Module.LookupInterface(t)
	if i != nil {
		return fmt.Sprintf("%s *%s%s", name, prefix, i.Name), nil
	}
	return "xxx", fmt.Errorf("unknown type %s", t)
}

func goParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("goParam called with nil node")
	}
	return ToParamString(prefix, &node.Schema, node.GetName())
}
