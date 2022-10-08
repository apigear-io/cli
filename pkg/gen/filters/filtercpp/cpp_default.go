package filtercpp

import (
	"fmt"
	"reflect"

	"github.com/apigear-io/cli/pkg/model"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(schema *model.Schema) (string, error) {
	t := schema.Type
	text := ""
	switch t {
	case "string":
		text = "std::string()"
	case "int":
		text = "0"
	case "float":
		text = "0.0"
	case "bool":
		text = "false"
	default:
		if schema.Module == nil {
			return "xxx", fmt.Errorf("schema.Module is nil")
		}
		e := schema.Module.LookupEnum(t)
		if e != nil {
			text = fmt.Sprintf("%s::%s", e.Name, e.Members[0].Name)
		}
		s := schema.Module.LookupStruct(t)
		if s != nil {
			text = fmt.Sprintf("%s()", s.Name)
		}
		i := schema.Module.LookupInterface(t)
		if i != nil {
			text = "nullptr"
		}
	}
	if schema.IsArray {
		inner := model.Schema{Type: t, Module: schema.Module}
		ret, err := ToReturnString(&inner)
		if err != nil {
			return "xxx", fmt.Errorf("ToDefaultString inner value error: %s", err)
		}
		text = fmt.Sprintf("std::vector<%s>()", ret)
	}
	return text, nil
}

// cppDefault returns the default value for a type
func cppDefault(node reflect.Value) (reflect.Value, error) {
	p := node.Interface().(model.ITypeProvider)
	t, err := ToDefaultString(p.GetSchema())
	return reflect.ValueOf(t), err
}
