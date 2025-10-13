package filterjava

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

func ToDefaultString(schema *model.Schema, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToDefaultString schema is nil")
	}
	var text string
	if schema.IsArray {
		switch schema.KindType {
		case model.TypeString:
			text = "new String[]{}"
		case model.TypeInt:
			text = "new int[]{}"
		case model.TypeInt32:
			text = "new int[]{}"
		case model.TypeInt64:
			text = "new long[]{}"
		case model.TypeFloat:
			text = "new float[]{}"
		case model.TypeFloat32:
			text = "new float[]{}"
		case model.TypeFloat64:
			text = "new double[]{}"
		case model.TypeBool:
			text = "new boolean[]{}"
		case model.TypeEnum:
			e_local := schema.LookupEnum("", schema.Type)
			e_imported := schema.LookupEnum(schema.Import, schema.Type)
			if e_local == nil && e_imported == nil {
				return "xxx", fmt.Errorf("javaDefault enum not found: %s", schema.Dump())
			}
			if e_local == nil {
				prefix = fmt.Sprintf("%s.%s_api.", common.CamelLowerCase(e_imported.Module.Name), common.CamelLowerCase(e_imported.Module.Name))
			}
			return fmt.Sprintf("new %s%s[]{}", prefix, common.CamelTitleCase(e_imported.Name)), nil
		case model.TypeStruct:
			s_local := schema.LookupStruct("", schema.Type)
			s_imported := schema.LookupStruct(schema.Import, schema.Type)
			if s_local == nil && s_imported == nil {
				return "xxx", fmt.Errorf("javaDefault struct not found: %s", schema.Dump())
			}
			// if struct is local it is found both as s_local and s_imported
			if s_local == nil {
				prefix = fmt.Sprintf("%s.%s_api.", common.CamelLowerCase(s_imported.Module.Name), common.CamelLowerCase(s_imported.Module.Name))
			}
			text = fmt.Sprintf("new %s%s[]{}", prefix, common.CamelTitleCase(s_imported.Name))
		case model.TypeExtern:
			xe := parseJavaExtern(schema)
			var java_module string
			java_module = ""
			if xe.Package != "" {
				java_module = fmt.Sprintf("%s.", xe.Package)
			}
			text = fmt.Sprintf("new %s%s[]{}", java_module, xe.Name)
		case model.TypeInterface:
			i_local := schema.LookupInterface("", schema.Type)
			i_imported := schema.LookupInterface(schema.Import, schema.Type)
			if i_local == nil && i_imported == nil {
				return "xxx", fmt.Errorf("javaDefault interface not found: %s", schema.Dump())
			}
			// if interface is local it is found both as s_local and s_imported
			if i_local == nil {
				prefix = fmt.Sprintf("%s.%s_api.", common.CamelLowerCase(i_imported.Module.Name), common.CamelLowerCase(i_imported.Module.Name))
			}
			text = fmt.Sprintf("new %sI%s[]{}", prefix, common.CamelTitleCase(i_imported.Name))
		default:
			return "xxx", fmt.Errorf("javaDefault unknown schema %s", schema.Dump())
		}
	} else {
		switch schema.KindType {
		case model.TypeString:
			text = "new String()"
		case model.TypeInt:
			text = "0"
		case model.TypeInt32:
			text = "0"
		case model.TypeInt64:
			text = "0L"
		case model.TypeFloat:
			text = "0.0f"
		case model.TypeFloat32:
			text = "0.0f"
		case model.TypeFloat64:
			text = "0.0"
		case model.TypeBool:
			text = "false"
		case model.TypeEnum:
			e_local := schema.LookupEnum("", schema.Type)
			e_imported := schema.LookupEnum(schema.Import, schema.Type)
			if e_local == nil && e_imported == nil {
				return "xxx", fmt.Errorf("javaDefault enum not found: %s", schema.Dump())
			}
			// if enum is local it is found both as e_local and e_imported
			name := common.CamelTitleCase(e_imported.Name)
			member := common.CamelTitleCase(e_imported.Members[0].Name)
			if e_local == nil {
				prefix = fmt.Sprintf("%s.%s_api.", common.CamelLowerCase(e_imported.Module.Name), common.CamelLowerCase(e_imported.Module.Name))
			}
			text = fmt.Sprintf("%s%s.%s", prefix, name, member)
		case model.TypeStruct:
			s_local := schema.LookupStruct("", schema.Type)
			s_imported := schema.LookupStruct(schema.Import, schema.Type)
			if s_local == nil && s_imported == nil {
				return "xxx", fmt.Errorf("javaDefault struct not found: %s", schema.Dump())
			}
			// if struct is local it is found both as s_local and s_imported
			if s_local == nil {
				prefix = fmt.Sprintf("%s.%s_api.", common.CamelLowerCase(s_imported.Module.Name), common.CamelLowerCase(s_imported.Module.Name))
			}
			text = fmt.Sprintf("new %s%s()", prefix, s_imported.Name)
		case model.TypeExtern:
			xe := parseJavaExtern(schema)
			text = fmt.Sprintf("new %s()", xe.Name)
			if xe.Default != "" {
				text = xe.Default
			} else {
				var java_module string
				java_module = ""
				if xe.Package != "" {
					java_module = fmt.Sprintf("%s.", xe.Package)
				}
				text = fmt.Sprintf("new %s%s()", java_module, xe.Name)
			}
		case model.TypeInterface:
			i_local := schema.LookupInterface("", schema.Type)
			i_imported := schema.LookupInterface(schema.Import, schema.Type)
			if i_local == nil && i_imported == nil {
				return "xxx", fmt.Errorf("javaDefault interface not found: %s", schema.Dump())
			}
			// if interface is local it is found both as s_local and s_imported
			if i_local == nil {
				prefix = fmt.Sprintf("%s.%s_impl.", common.CamelLowerCase(i_imported.Module.Name), common.CamelLowerCase(i_imported.Module.Name))
			}
			text = "null"
		default:
			return "xxx", fmt.Errorf("javaDefault unknown schema %s", schema.Dump())
		}
	}
	return text, nil
}

func javaDefault(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("javaDefault node is nil")
	}
	return ToDefaultString(&node.Schema, prefix)
}
