package filtercpp

import (
	"fmt"
	"reflect"

	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(schema *model.Schema, name string) (string, error) {
	t := schema.Type
	if schema.IsArray {
		inner := *schema
		inner.IsArray = false
		ret, err := ToReturnString(&inner)
		if err != nil {
			return "", fmt.Errorf("ToParamString inner value error: %s", err)
		}
		return fmt.Sprintf("const std::vector<%s> &%s", ret, name), nil
	}
	switch t {
	case "string":
		return fmt.Sprintf("const std::string &%s", name), nil
	case "int":
		return fmt.Sprintf("int %s", name), nil
	case "float":
		return fmt.Sprintf("double %s", name), nil
	case "bool":
		return fmt.Sprintf("bool %s", name), nil
	}
	e := schema.Module.LookupEnum(t)
	if e != nil {
		return fmt.Sprintf("%s %s", e.Name, name), nil
	}
	s := schema.Module.LookupStruct(t)
	if s != nil {
		return fmt.Sprintf("const %s &%s", s.Name, name), nil
	}
	i := schema.Module.LookupInterface(t)
	if i != nil {
		return fmt.Sprintf("%s *%s", i.Name, name), nil
	}
	return "XXX", fmt.Errorf("ToParamString: unknown type %s", t)
}

func cppParam(node reflect.Value) (reflect.Value, error) {
	p := node.Interface().(model.ITypeProvider)
	t, err := ToParamString(p.GetSchema(), p.GetName())
	return reflect.ValueOf(t), err
}
