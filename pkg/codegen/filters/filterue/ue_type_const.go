package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
	"github.com/ettle/strcase"
)

func ToConstTypeString(prefix string, schema *apimodel.Schema) (string, error) {
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
		text = "const FString&"
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
		text = fmt.Sprintf("const %sF%s%s&", prefix, moduleId, schema.Type)
	case apimodel.TypeExtern:
		text = fmt.Sprintf("const %s&", ueExtern(schema.GetExtern()).Name)
	case apimodel.TypeInterface:
		text = fmt.Sprintf("const TScriptInterface<%sI%s%sInterface>&", prefix, moduleId, schema.Type)
	default:
		return "xxx", fmt.Errorf("ueConstType unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		switch schema.KindType {
		case apimodel.TypeString:
			text = "const TArray<FString>&"
		case apimodel.TypeInt:
			text = "const TArray<int32>&"
		case apimodel.TypeInt32:
			text = "const TArray<int32>&"
		case apimodel.TypeInt64:
			text = "const TArray<int64>&"
		case apimodel.TypeFloat:
			text = "const TArray<float>&"
		case apimodel.TypeFloat32:
			text = "const TArray<float>&"
		case apimodel.TypeFloat64:
			text = "const TArray<double>&"
		case apimodel.TypeBool:
			text = "const TArray<bool>&"
		case apimodel.TypeVoid:
			text = "const TArray<void>&"
		case apimodel.TypeEnum:
			text = fmt.Sprintf("const TArray<%sE%s%s>&", prefix, moduleId, schema.Type)
		case apimodel.TypeStruct:
			text = fmt.Sprintf("const TArray<%sF%s%s>&", prefix, moduleId, schema.Type)
		case apimodel.TypeExtern:
			text = fmt.Sprintf("const TArray<%s>&", ueExtern(schema.GetExtern()).Name)
		case apimodel.TypeInterface:
			text = fmt.Sprintf("const TArray<TScriptInterface<%sI%s%sInterface>>&", prefix, moduleId, schema.Type)
		default:
			return "xxx", fmt.Errorf("ueConstType unknown schema %s", schema.Dump())
		}
	}
	return text, nil
}

func ueConstType(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueConstType node is nil")
	}
	return ToConstTypeString(prefix, &node.Schema)
}
