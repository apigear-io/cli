package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
	"github.com/ettle/strcase"
)

func ToConstTypeString(prefix string, schema *objmodel.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToReturnString schema is nil")
	}
	moduleId := strcase.ToPascal(schema.Module.Name)
	if schema.Import != "" {
		moduleId = strcase.ToPascal(schema.Import)
	}
	var text string
	switch schema.KindType {
	case objmodel.TypeString:
		text = "const FString&"
	case objmodel.TypeInt:
		text = "int32"
	case objmodel.TypeInt32:
		text = "int32"
	case objmodel.TypeInt64:
		text = "int64"
	case objmodel.TypeFloat:
		text = "float"
	case objmodel.TypeFloat32:
		text = "float"
	case objmodel.TypeFloat64:
		text = "double"
	case objmodel.TypeBool:
		text = "bool"
	case objmodel.TypeVoid:
		text = "void"
	case objmodel.TypeEnum:
		text = fmt.Sprintf("%sE%s%s", prefix, moduleId, schema.Type)
	case objmodel.TypeStruct:
		text = fmt.Sprintf("const %sF%s%s&", prefix, moduleId, schema.Type)
	case objmodel.TypeExtern:
		text = fmt.Sprintf("const %s&", ueExtern(schema.GetExtern()).Name)
	case objmodel.TypeInterface:
		text = fmt.Sprintf("const TScriptInterface<%sI%s%sInterface>&", prefix, moduleId, schema.Type)
	default:
		return "xxx", fmt.Errorf("ueConstType unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		switch schema.KindType {
		case objmodel.TypeString:
			text = "const TArray<FString>&"
		case objmodel.TypeInt:
			text = "const TArray<int32>&"
		case objmodel.TypeInt32:
			text = "const TArray<int32>&"
		case objmodel.TypeInt64:
			text = "const TArray<int64>&"
		case objmodel.TypeFloat:
			text = "const TArray<float>&"
		case objmodel.TypeFloat32:
			text = "const TArray<float>&"
		case objmodel.TypeFloat64:
			text = "const TArray<double>&"
		case objmodel.TypeBool:
			text = "const TArray<bool>&"
		case objmodel.TypeVoid:
			text = "const TArray<void>&"
		case objmodel.TypeEnum:
			text = fmt.Sprintf("const TArray<%sE%s%s>&", prefix, moduleId, schema.Type)
		case objmodel.TypeStruct:
			text = fmt.Sprintf("const TArray<%sF%s%s>&", prefix, moduleId, schema.Type)
		case objmodel.TypeExtern:
			text = fmt.Sprintf("const TArray<%s>&", ueExtern(schema.GetExtern()).Name)
		case objmodel.TypeInterface:
			text = fmt.Sprintf("const TArray<TScriptInterface<%sI%s%sInterface>>&", prefix, moduleId, schema.Type)
		default:
			return "xxx", fmt.Errorf("ueConstType unknown schema %s", schema.Dump())
		}
	}
	return text, nil
}

func ueConstType(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueConstType node is nil")
	}
	return ToConstTypeString(prefix, &node.Schema)
}
