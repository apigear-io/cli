package filterjava

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/codegen/filters/common"
	"github.com/apigear-io/cli/pkg/objmodel"
)

func ToDefaultString(schema *objmodel.Schema, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToDefaultString schema is nil")
	}
	var text string
	if schema.IsArray {
		switch schema.KindType {
		case objmodel.TypeString:
			text = "new String[]{}"
		case objmodel.TypeInt:
			text = "new int[]{}"
		case objmodel.TypeInt32:
			text = "new int[]{}"
		case objmodel.TypeInt64:
			text = "new long[]{}"
		case objmodel.TypeFloat:
			text = "new float[]{}"
		case objmodel.TypeFloat32:
			text = "new float[]{}"
		case objmodel.TypeFloat64:
			text = "new double[]{}"
		case objmodel.TypeBool:
			text = "new boolean[]{}"
		case objmodel.TypeEnum:
			e_local := schema.LookupEnum("", schema.Type)
			e_imported := schema.LookupEnum(schema.Import, schema.Type)
			if e_local == nil && e_imported == nil {
				return "xxx", fmt.Errorf("javaDefault enum not found: %s", schema.Dump())
			}
			if e_local == nil {
				prefix = fmt.Sprintf("%s.%s_api.", common.CamelLowerCase(e_imported.Module.Name), common.CamelLowerCase(e_imported.Module.Name))
			}
			return fmt.Sprintf("new %s%s[]{}", prefix, common.CamelTitleCase(e_imported.Name)), nil
		case objmodel.TypeStruct:
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
		case objmodel.TypeExtern:
			xe := parseJavaExtern(schema)
			var java_module string
			java_module = ""
			if xe.Package != "" {
				java_module = fmt.Sprintf("%s.", xe.Package)
			}
			text = fmt.Sprintf("new %s%s[]{}", java_module, xe.Name)
		case objmodel.TypeInterface:
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
		case objmodel.TypeString:
			text = "new String()"
		case objmodel.TypeInt:
			text = "0"
		case objmodel.TypeInt32:
			text = "0"
		case objmodel.TypeInt64:
			text = "0L"
		case objmodel.TypeFloat:
			text = "0.0f"
		case objmodel.TypeFloat32:
			text = "0.0f"
		case objmodel.TypeFloat64:
			text = "0.0"
		case objmodel.TypeBool:
			text = "false"
		case objmodel.TypeEnum:
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
		case objmodel.TypeStruct:
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
		case objmodel.TypeExtern:
			xe := parseJavaExtern(schema)
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
		case objmodel.TypeInterface:
			i_local := schema.LookupInterface("", schema.Type)
			i_imported := schema.LookupInterface(schema.Import, schema.Type)
			if i_local == nil && i_imported == nil {
				return "xxx", fmt.Errorf("javaDefault interface not found: %s", schema.Dump())
			}
			// if interface is local it is found both as s_local and s_imported
			text = "null"
		default:
			return "xxx", fmt.Errorf("javaDefault unknown schema %s", schema.Dump())
		}
	}
	return text, nil
}

func javaDefault(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("javaDefault node is nil")
	}
	return ToDefaultString(&node.Schema, prefix)
}
