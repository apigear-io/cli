package filtercpp

import (
	"fmt"
	"objectapi/pkg/model"
	"reflect"
	"strings"
)

func ToReturnString(t string) string {
	isArray := strings.HasSuffix(t, "[]")
	if isArray {
		t = t[:len(t)-2]
	}
	s := ""
	switch t {
	case "string":
		s = "std::string"
	case "int":
		s = "int"
	case "float":
		s = "double"
	case "bool":
		s = "bool"
	default:
		s = t
	}
	if isArray {
		s = fmt.Sprintf("std::vector<%s>", s)
	}
	return s
}

// cast value to TypedNode and deduct the cpp return type
func cppReturn(node reflect.Value) (reflect.Value, error) {
	p := node.Interface().(model.ISchemaProvider)
	t := ToReturnString(p.GetSchema().Type)
	return reflect.ValueOf(t), nil
}
