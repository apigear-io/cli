package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(schema *model.Schema, name string, prefix string) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToParamString schema is nil")
	}
	t := schema.Type
	if schema.IsArray {
		inner := *schema
		inner.IsArray = false
		innerValue, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "", fmt.Errorf("ToParamString inner value error: %s", err)
		}
		return fmt.Sprintf("%s []%s", name, innerValue), nil
	}
	switch t {
	case "string":
		return fmt.Sprintf("%s string", name), nil
	case "int":
		return fmt.Sprintf("%s int64", name), nil
	case "float":
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
	return "XXX", fmt.Errorf("unknown type %s", t)
}

func goParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("goParam called with nil node")
	}
	return ToParamString(&node.Schema, node.GetName(), prefix)
}
