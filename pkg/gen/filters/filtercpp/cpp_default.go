package filtercpp

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(prefix string, schema *model.Schema) (string, error) {
	text := ""
	switch schema.Type {
	case "void":
		text = "void"
	case "string":
		text = "std::string()"
	case "int", "int32":
		text = "0"
	case "int64":
		text = "0LL"
	case "float", "float32":
		text = "0.0f"
	case "float64":
		text = "0.0"
	case "bool":
		text = "false"
	default:
		if schema.Module == nil {
			return "xxx", fmt.Errorf("schema.Module is nil")
		}
		e := schema.Module.LookupEnum(schema.Import, schema.Type)
		if e != nil {
			text = fmt.Sprintf("%sEnum::%s", e.Name, e.Members[0].Name)
		}
		s := schema.Module.LookupStruct(schema.Import, schema.Type)
		if s != nil {
			text = fmt.Sprintf("%s()", s.Name)
		}
		i := schema.Module.LookupInterface(schema.Import, schema.Type)
		if i != nil {
			text = "nullptr"
		}
	}
	if schema.IsArray {
		inner := schema.InnerSchema()
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("ToDefaultString inner value error: %s", err)
		}
		text = fmt.Sprintf("std::list<%s>()", ret)
	}
	return text, nil
}

// cppDefault returns the default value for a type
func cppDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("cppDefault node is nil")
	}
	return ToDefaultString(prefix, &node.Schema)
}
