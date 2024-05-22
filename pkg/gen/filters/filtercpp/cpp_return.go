package filtercpp

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(prefix string, schema *model.Schema) (string, error) {
	text := ""
	switch schema.KindType {
	case model.TypeVoid:
		text = "void"
	case model.TypeString:
		text = "std::string"
	case model.TypeInt:
		text = "int"
	case model.TypeInt32:
		text = "int32_t"
	case model.TypeInt64:
		text = "int64_t"
	case model.TypeFloat:
		text = "float"
	case model.TypeFloat32:
		text = "float"
	case model.TypeFloat64:
		text = "double"
	case model.TypeBool:
		text = "bool"
	case model.TypeExtern:
		xe := parseCppExtern(schema)
		if xe.NameSpace != "" {
			prefix = fmt.Sprintf("%s::", xe.NameSpace)
		}
		text = fmt.Sprintf("%s%s", prefix, xe.Name)
	case model.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s::%s::", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		if e != nil {
			text = fmt.Sprintf("%s%sEnum", prefix, e.Name)
		}
	case model.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s::%s::", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		if s != nil {
			text = fmt.Sprintf("%s%s", prefix, s.Name)
		}
	case model.TypeInterface:
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
func cppReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("cppReturn node is nil")
	}
	return ToReturnString(prefix, &node.Schema)
}
