package filtercpp

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

func ToTypeRefString(prefix string, schema *model.Schema) (string, error) {
	if schema.IsArray {
		inner := schema.InnerSchema()
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", err
		}
		return fmt.Sprintf("const std::list<%s>&", ret), nil
	}
	text := ""
	switch schema.Type {
	case "void":
		text = "void"
	case "string":
		text = "const std::string&"
	case "int":
		text = "int"
	case "int32":
		text = "int32_t"
	case "int64":
		text = "int64_t"
	case "float":
		text = "float"
	case "float32":
		text = "float"
	case "float64":
		text = "double"
	case "bool":
		text = "bool"
	default:
		if schema.GetExtern() != nil {
			xe := parseCppExtern(schema)
			if xe.NameSpace != "" {
				prefix = fmt.Sprintf("%s::", xe.NameSpace)
			}
			text = fmt.Sprintf("const %s%s&", prefix, xe.Name)
		}
		if schema.Import != "" {
			prefix = fmt.Sprintf("%s::%s::", common.CamelTitleCase(schema.System().Name), common.CamelTitleCase(schema.Import))
		}
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e != nil {
			text = fmt.Sprintf("%s%sEnum", prefix, e.Name)
		}
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s != nil {
			text = fmt.Sprintf("const %s%s&", prefix, s.Name)
		}
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i != nil {
			text = fmt.Sprintf("%s%s*", prefix, i.Name)
		}
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
