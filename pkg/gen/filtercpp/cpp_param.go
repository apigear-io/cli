package filtercpp

import (
	"fmt"
	"log"
	"objectapi/pkg/model"
	"reflect"
	"strings"
)

func ToParamString(schema *model.Schema, name string) string {
	t := schema.Type
	isArray := strings.HasSuffix(t, "[]")
	if isArray {
		t = t[:len(t)-2] // remove the []
		inner := model.Schema{Type: t, Module: schema.Module}
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
	log.Fatalf("unknown type %s", t)
	return "XXX"
}

func cppParam(node reflect.Value) (reflect.Value, error) {
	p := node.Interface().(model.ITypeProvider)
	t := ToParamString(p.GetSchema(), p.GetName())
	return reflect.ValueOf(t), nil
}
