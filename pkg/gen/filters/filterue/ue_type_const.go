package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/iancoleman/strcase"
)

func ToConstTypeString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToReturnString schema is nil")
	}
	moduleId := strcase.ToCamel(schema.Module.Name)
	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "const FString&"
	case model.TypeInt:
		text = "int32"
	case model.TypeFloat:
		text = "float"
	case model.TypeBool:
		text = "bool"
	case model.TypeEnum:
		text = fmt.Sprintf("const %sE%s%s&", prefix, moduleId, schema.Type)
	case model.TypeStruct:
		text = fmt.Sprintf("const %sF%s%s&", prefix, moduleId, schema.Type)
	case model.TypeInterface:
		text = fmt.Sprintf("%sF%s%s*", prefix, moduleId, schema.Type)
	case model.TypeVoid:
		text = ""
	default:
		return "", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
	}
	if schema.IsArray {
		switch schema.KindType {
		case model.TypeInt, model.TypeFloat, model.TypeBool:
			text = fmt.Sprintf("const TArray<%s>&", text)
		case model.TypeString:
			text = "const TArray<FString>&"
		case model.TypeEnum:
			text = fmt.Sprintf("const TArray<%sE%s%s>&", prefix, moduleId, schema.Type)
		case model.TypeStruct:
			text = fmt.Sprintf("const TArray<%sF%s%s>&", prefix, moduleId, schema.Type)
		case model.TypeInterface:
			text = fmt.Sprintf("const TArray<%sF%s%s*>&", prefix, moduleId, schema.Type)
		default:
			return "", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
		}
	}
	return text, nil
}

func ueConstType(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("goReturn node is nil")
	}
	return ToConstTypeString(prefix, &node.Schema)
}
