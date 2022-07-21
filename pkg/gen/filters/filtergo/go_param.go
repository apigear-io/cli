package filtergo

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(schema *model.Schema, name string, prefix string) string {
	t := schema.Type
	if schema.IsArray {
		inner := *schema
		inner.IsArray = false
		return fmt.Sprintf("%s []%s", name, ToReturnString(&inner, prefix))
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
		return fmt.Sprintf("%s %s%s", name, prefix, e.Name)
	}
	s := schema.Module.LookupStruct(t)
	if s != nil {
		return fmt.Sprintf("%s %s%s", name, prefix, s.Name)
	}
	i := schema.Module.LookupInterface(t)
	if i != nil {
		return fmt.Sprintf("%s *%s%s", name, prefix, i.Name)
	}
	log.Fatalf("unknown type %s", t)
	return "XXX"
}

func goParam(node *model.TypedNode, prefix string) string {
	if node == nil {
		log.Warnf("goParam called with nil node")
		return ""
	}
	return ToParamString(&node.Schema, node.GetName(), prefix)
}
