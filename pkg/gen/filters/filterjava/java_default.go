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
			return fmt.Sprintf("new %s%s[]{}", prefix, schema.Type), nil
		case model.TypeStruct:
			symbol := schema.GetStruct()
			text = fmt.Sprintf("new %s%s[]{}", prefix, symbol.Name)
		case model.TypeExtern:
			xe := parseJavaExtern(schema)
			if xe.Default != "" {
				text = xe.Default
			} else {
				text = fmt.Sprintf("%s%s()", prefix, xe.Name)
			}
		case model.TypeInterface:
			symbol := schema.GetInterface()
			text = fmt.Sprintf("new %s%s[]{}", prefix, symbol.Name)
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
			symbol := schema.GetEnum()
			member := symbol.Members[0]
			// upper case first letter
			// TODO: EnumValues: using camel-cases for enum values: strcase.ToCamel(member.Name)
			text = fmt.Sprintf("%s%s.%s", prefix, symbol.Name, common.CamelTitleCase(member.Name))
		case model.TypeStruct:
			symbol := schema.GetStruct()
			text = fmt.Sprintf("new %s%s()", prefix, symbol.Name)
		case model.TypeExtern:
			xe := parseJavaExtern(schema)
			text = fmt.Sprintf("new %s%s()", prefix, xe.Name)
		case model.TypeInterface:
			symbol := schema.GetInterface()
			text = fmt.Sprintf("new %s%s()", prefix, symbol.Name)
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
