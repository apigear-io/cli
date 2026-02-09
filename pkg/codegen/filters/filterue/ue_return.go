package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
	"github.com/ettle/strcase"
)

//TODO: add test including prefix for all filters

func ToReturnString(prefix string, schema *apimodel.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToReturnString schema is nil")
	}
	moduleId := strcase.ToPascal(schema.Module.Name)
	if schema.Import != "" {
		moduleId = strcase.ToPascal(schema.Import)
	}
	var text string
	switch schema.KindType {
	case apimodel.TypeString:
		text = "FString"
	case apimodel.TypeInt:
		text = "int32"
	case apimodel.TypeInt32:
		text = "int32"
	case apimodel.TypeInt64:
		text = "int64"
	case apimodel.TypeFloat:
		text = "float"
	case apimodel.TypeFloat32:
		text = "float"
	case apimodel.TypeFloat64:
		text = "double"
	case apimodel.TypeBool:
		text = "bool"
	case apimodel.TypeVoid:
		text = "void"
	case apimodel.TypeEnum:
		text = fmt.Sprintf("%sE%s%s", prefix, moduleId, schema.Type)
	case apimodel.TypeStruct:
		text = fmt.Sprintf("%sF%s%s", prefix, moduleId, schema.Type)
	case apimodel.TypeExtern:
		text = ueExtern(schema.GetExtern()).Name
	case apimodel.TypeInterface:
		text = fmt.Sprintf("TScriptInterface<%sI%s%sInterface>", prefix, moduleId, schema.Type)
	default:
		return "xxx", fmt.Errorf("ueReturn unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		text = fmt.Sprintf("TArray<%s>", text)
	}
	return text, nil
}

func ueReturn(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueReturn called with nil node")
	}
	return ToReturnString(prefix, &node.Schema)
}
