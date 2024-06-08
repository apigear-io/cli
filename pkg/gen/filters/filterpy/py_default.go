package filterpy

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

// ToDefaultString returns the default value for a type
func ToDefaultString(schema *model.Schema, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("pyDefault schema is nil")
	}
	if schema.Module == nil {
		return "xxx", fmt.Errorf("pyDefault schema module is nil")
	}
	var text string
	if schema.IsArray {
		text = "[]"
	} else {
		switch schema.KindType {
		case model.TypeString:
			text = "\"\""
		case model.TypeInt, model.TypeInt32, model.TypeInt64:
			text = "0"
		case model.TypeFloat, model.TypeFloat32, model.TypeFloat64:
			text = "0.0"
		case model.TypeBool:
			text = "False"
		case model.TypeExtern:
			xe := parsePyExtern(schema)
			py_module := ""
			if xe.Import != "" {
				py_module = fmt.Sprintf("%s.", xe.Import)
			}
			text = fmt.Sprintf("%s%s()", py_module, xe.Name)
		case model.TypeEnum:
			e_local := schema.LookupEnum("", schema.Type)
			e_imported := schema.LookupEnum(schema.Import, schema.Type)
			if e_local == nil && e_imported == nil {
				return "xxx", fmt.Errorf("pyDefault enum not found: %s", schema.Dump())
			}
			// if enum is local it is found both as e_local and e_imported
			name := common.CamelTitleCase(e_imported.Name)
			member := common.SnakeUpperCase(e_imported.Members[0].Name)
			if e_local == nil {
				prefix = fmt.Sprintf("%s.api.", e_imported.Module.Name)
			}
			text = fmt.Sprintf("%s%s.%s", prefix, name, member)
		case model.TypeStruct:
			s_local := schema.LookupStruct("", schema.Type)
			s_imported := schema.LookupStruct(schema.Import, schema.Type)
			if s_local == nil && s_imported == nil {
				return "xxx", fmt.Errorf("pyDefault struct not found: %s", schema.Dump())
			}
			// if struct is local it is found both as s_local and s_imported
			ident := common.CamelTitleCase(s_imported.Name)
			if s_local == nil {
				prefix = fmt.Sprintf("%s.api.", s_imported.Module.Name)
			}
			text = fmt.Sprintf("%s%s()", prefix, ident)
		case model.TypeInterface:
			i := schema.LookupInterface(schema.Import, schema.Type)
			if i == nil {
				return "xxx", fmt.Errorf("pyDefault interface not found: %s", schema.Dump())
			}
			text = "None"
		case model.TypeVoid:
			text = "None"
		default:
			return "xxx", fmt.Errorf("pyDefault unknown schema %s", schema.Dump())
		}
	}
	if text == "" {
		return "xxx", fmt.Errorf("pyDefault text is empty: %s", schema.Dump())
	}
	return text, nil
}

// cppDefault returns the default value for a type
func pyDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("pyDefault called with nil node")
	}
	return ToDefaultString(&node.Schema, prefix)
}
