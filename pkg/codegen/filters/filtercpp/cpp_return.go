package filtercpp

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/codegen/filters/common"
	"github.com/apigear-io/cli/pkg/objmodel"
)

func ToReturnString(prefix string, schema *objmodel.Schema) (string, error) {
	text := ""
	switch schema.KindType {
	case objmodel.TypeVoid:
		text = "void"
	case objmodel.TypeString:
		text = "std::string"
	case objmodel.TypeInt:
		text = "int"
	case objmodel.TypeInt32:
		text = "int32_t"
	case objmodel.TypeInt64:
		text = "int64_t"
	case objmodel.TypeFloat:
		text = "float"
	case objmodel.TypeFloat32:
		text = "float"
	case objmodel.TypeFloat64:
		text = "double"
	case objmodel.TypeBool:
		text = "bool"
	case objmodel.TypeExtern:
		xe := parseCppExtern(schema)
		if xe.NameSpace != "" {
			prefix = fmt.Sprintf("%s::", xe.NameSpace)
		} else {
			prefix = "" // Externs should not be prefixed with any other prefix than given in extern info.
		}
		text = fmt.Sprintf("%s%s", prefix, xe.Name)
	case objmodel.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s::%s::", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		if e != nil {
			text = fmt.Sprintf("%s%sEnum", prefix, e.Name)
		}
	case objmodel.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s::%s::", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		if s != nil {
			text = fmt.Sprintf("%s%s", prefix, s.Name)
		}
	case objmodel.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s::%s::", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		if i != nil {
			text = fmt.Sprintf("%s%s*", prefix, i.Name)
		}
	}
	if schema.IsArray {
		return fmt.Sprintf("std::list<%s>", text), nil
	}
	return text, nil
}

// cast value to TypedNode and deduct the cpp return type
func cppReturn(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("cppReturn node is nil")
	}
	return ToReturnString(prefix, &node.Schema)
}
