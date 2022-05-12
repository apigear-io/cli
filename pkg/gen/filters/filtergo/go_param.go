package filtergo

import (
	"fmt"
	"objectapi/pkg/log"
	"objectapi/pkg/model"
)

func ToParamString(schema *model.Schema, name string) string {
	t := schema.Type
	if schema.IsArray {
		inner := *schema
		inner.IsArray = false
		return fmt.Sprintf("%s []%s", name, ToReturnString(&inner))
	}
	switch t {
	case "string":
		return fmt.Sprintf("%s string", name)
	case "int":
		return fmt.Sprintf("%s int", name)
	case "float":
		return fmt.Sprintf("%s float64", name)
	case "bool":
		return fmt.Sprintf("%s bool", name)
	}
	e := schema.Module.LookupEnum(t)
	if e != nil {
		return fmt.Sprintf("%s %s", name, e.Name)
	}
	s := schema.Module.LookupStruct(t)
	if s != nil {
		return fmt.Sprintf("%s %s", name, s.Name)
	}
	i := schema.Module.LookupInterface(t)
	if i != nil {
		return fmt.Sprintf("%s *%s", name, i.Name)
	}
	log.Fatalf("unknown type %s", t)
	return "XXX"
}

func goParam(p *model.TypedNode) string {
	return ToParamString(&p.Schema, p.Name)
}
