package filtercpp

import (
	"fmt"
	"reflect"

	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(schema *model.Schema, name string) string {
	t := schema.Type
	if schema.IsArray {
		inner := *schema
		inner.IsArray = false
		return fmt.Sprintf("const std::vector<%s> &%s", ToReturnString(&inner), name)
	}
	switch t {
	case "string":
		return fmt.Sprintf("const std::string &%s", name)
	case "int":
		return fmt.Sprintf("int %s", name)
	case "float":
		return fmt.Sprintf("double %s", name)
	case "bool":
		return fmt.Sprintf("bool %s", name)
	}
	e := schema.Module.LookupEnum(t)
	if e != nil {
		return fmt.Sprintf("%s %s", e.Name, name)
	}
	s := schema.Module.LookupStruct(t)
	if s != nil {
		return fmt.Sprintf("const %s &%s", s.Name, name)
	}
	i := schema.Module.LookupInterface(t)
	if i != nil {
		return fmt.Sprintf("%s *%s", i.Name, name)
	}
	log.Fatal().Msgf("unknown type %s", t)
	return "XXX"
}

func cppParam(node reflect.Value) (reflect.Value, error) {
	p := node.Interface().(model.ITypeProvider)
	t := ToParamString(p.GetSchema(), p.GetName())
	return reflect.ValueOf(t), nil
}
