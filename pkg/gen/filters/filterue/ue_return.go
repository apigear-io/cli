package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/iancoleman/strcase"
)

//TODO: add test including prefix for all filters

func ToReturnString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToReturnString schema is nil")
	}
	moduleId := strcase.ToCamel(schema.Module.Name)
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
		text = fmt.Sprintf("TArray<%s>", text)
	}
	return text, nil
}

func ueReturn(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("goReturn node is nil")
	}
	return ToReturnString(prefix, &node.Schema)
}
