package filtercpp

import (
	"fmt"
	"objectapi/pkg/model"
	"reflect"
	"strings"
)

func ToDefaultString(t string) string {
	isArray := strings.HasSuffix(t, "[]")
	if isArray {
		t = t[:len(t)-2]
	}
	s := ""
	switch t {
	case "string":
		s = "\"\""
	case "int":
		s = "0"
	case "float":
		s = "0.0"
	case "bool":
		s = "false"
	default:
		s = t
	}
	if isArray {
		s = fmt.Sprintf("std::vector<%s>", s)
	}
	return s
}

func cppDefault(node reflect.Value) (reflect.Value, error) {
	schema := node.Interface().(*model.TypedNode).Schema
	t := ToDefaultString(schema.Type)
	return reflect.ValueOf(t), nil
}
