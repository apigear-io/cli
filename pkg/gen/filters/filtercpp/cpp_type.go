package filtercpp

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

var cppType = cppReturn

func ToTypeRefString(prefix string, schema *model.Schema) (string, error) {
	t := schema.Type
	text := ""
	switch t {
	case "void":
		text = "void"
	case "string":
		text = "const std::string&"
	case "int":
		text = "int"
	case "float":
		text = "float"
	case "bool":
		text = "bool"
	default:
		if schema.Module == nil {
			return "xxx", fmt.Errorf("schema.Module is nil")
		}
		e := schema.Module.LookupEnum(t)
		if e != nil {
			text = fmt.Sprintf("%s%sEnum", prefix, e.Name)
		}
		s := schema.Module.LookupStruct(t)
		if s != nil {
			text = fmt.Sprintf("const %s%s&", prefix, s.Name)
		}
		i := schema.Module.LookupInterface(t)
		if i != nil {
			text = fmt.Sprintf("%s%s*", prefix, i.Name)
		}
	}
	if schema.IsArray {
		schema.IsArray = false
		inner, err := ToReturnString(prefix, schema)
		if err != nil {
			return "xxx", err
		}
		text = fmt.Sprintf("const std::list<%s>&", inner)
	}
	return text, nil
}

// cast value to TypedNode and deduct the cpp return type
func cppTypeRef(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("cppTypeRef node is nil")
	}
	return ToTypeRefString(prefix, &node.Schema)
}
