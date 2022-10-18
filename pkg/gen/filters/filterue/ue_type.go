package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/ettle/strcase"
)

func ToTypeString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToReturnString schema is nil")
	}
	moduleId := strcase.ToPascal(schema.Module.Name)
	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "FString"
	case model.TypeInt:
		text = "int32"
	case model.TypeFloat:
		text = "float"
	case model.TypeBool:
		text = "bool"
	case model.TypeVoid:
		text = "void"
	case model.TypeEnum:
		text = fmt.Sprintf("%sE%s%s", prefix, moduleId, schema.Type)
	case model.TypeStruct:
		text = fmt.Sprintf("%sF%s%s", prefix, moduleId, schema.Type)
	case model.TypeInterface:
		text = fmt.Sprintf("%sF%s%s*", prefix, moduleId, schema.Type)
	default:
		return "xxx", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
	}
	if schema.IsArray {
		switch schema.KindType {
		case model.TypeInt, model.TypeFloat, model.TypeBool:
			text = fmt.Sprintf("TArray<%s>", text)
		case model.TypeString:
			text = "TArray<FString>"
		case model.TypeEnum:
			text = fmt.Sprintf("TArray<%sE%s%s>", prefix, moduleId, schema.Type)
		case model.TypeStruct:
			text = fmt.Sprintf("TArray<%sF%s%s>", prefix, moduleId, schema.Type)
		case model.TypeInterface:
			text = fmt.Sprintf("TArray<%sF%s%s*>", prefix, moduleId, schema.Type)
		default:
			return "xxx", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
		}
	}
	return text, nil
}

func ueType(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueType node is nil")
	}
	return ToTypeString(prefix, &node.Schema)
}
