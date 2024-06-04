package filterpy

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

func ToReturnString(schema *model.Schema, prefix string) (string, error) {
	text := ""
	switch schema.KindType {
	case model.TypeString:
		text = "str"
	case model.TypeInt:
		text = "int"
	case model.TypeInt32:
		text = "int32"
	case model.TypeInt64:
		text = "int64"
	case model.TypeFloat:
		text = "float"
	case model.TypeFloat32:
		text = "float32"
	case model.TypeFloat64:
		text = "float64"
	case model.TypeBool:
		text = "bool"
	case model.TypeExtern:
		x := schema.LookupExtern(schema.Import, schema.Type)
		if x == nil {
			return "xxx", fmt.Errorf("pyReturn extern not found: %s", schema.Dump())
		}
		xe := parsePyExtern(schema)
		if xe.Import != "" {
			prefix = fmt.Sprintf("%s.", xe.Import)
		}
		text = fmt.Sprintf("%s%s", prefix, xe.Name)
	case model.TypeEnum:
		e := schema.LookupEnum("", schema.Type)
		e_imported := schema.LookupEnum(schema.Import, schema.Type)
		if e == nil && e_imported == nil {
			return "xxx", fmt.Errorf("pyReturn enum not found: %s", schema.Dump())
		}
		// if enum is local it is found both as e and e_imported
		ident := common.CamelTitleCase(e_imported.Name)
		if e == nil {
			prefix = fmt.Sprintf("%s.api.", e_imported.Module.Name)
		}
		text = fmt.Sprintf("%s%s", prefix, ident)
	case model.TypeStruct:
		s := schema.LookupStruct("", schema.Type)
		s_imported := schema.LookupStruct(schema.Import, schema.Type)
		if s == nil && s_imported == nil {
			return "xxx", fmt.Errorf("pyReturn struct not found: %s", schema.Dump())
		}
		// if struct is local it is found both as s and s_imported
		ident := common.CamelTitleCase(s_imported.Name)
		if s == nil {
			prefix = fmt.Sprintf("%s.api.", s_imported.Module.Name)
		}
		text = fmt.Sprintf("%s%s", prefix, ident)
	case model.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("pyReturn interface not found: %s", schema.Dump())
		}
		ident := common.CamelTitleCase(i.Name)
		text = fmt.Sprintf("%s%s", prefix, ident)
	case model.TypeVoid:
		text = "None"
	default:
		return "xxx", fmt.Errorf("pyReturn unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		text = fmt.Sprintf("list[%s]", text)
	}
	return text, nil
}

// cast value to TypedNode and deduct the py return type
func pyReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("pyReturn called with nil node")
	}
	return ToReturnString(&node.Schema, prefix)
}
