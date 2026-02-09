package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
	"github.com/ettle/strcase"
)

func ToTypeString(prefix string, schema *apimodel.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ueType schema is nil")
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
		return "xxx", fmt.Errorf("ueType unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		switch schema.KindType {
		case apimodel.TypeString:
			text = "TArray<FString>"
		case apimodel.TypeInt:
			text = "TArray<int32>"
		case apimodel.TypeInt32:
			text = "TArray<int32>"
		case apimodel.TypeInt64:
			text = "TArray<int64>"
		case apimodel.TypeFloat:
			text = "TArray<float>"
		case apimodel.TypeFloat32:
			text = "TArray<float>"
		case apimodel.TypeFloat64:
			text = "TArray<double>"
		case apimodel.TypeBool:
			text = "TArray<bool>"
		case apimodel.TypeEnum:
			text = fmt.Sprintf("TArray<%sE%s%s>", prefix, moduleId, schema.Type)
		case apimodel.TypeStruct:
			text = fmt.Sprintf("TArray<%sF%s%s>", prefix, moduleId, schema.Type)
		case apimodel.TypeExtern:
			text = fmt.Sprintf("TArray<%s>", ueExtern(schema.GetExtern()).Name)
		case apimodel.TypeInterface:
			text = fmt.Sprintf("TArray<TScriptInterface<%sI%s%sInterface>>", prefix, moduleId, schema.Type)
		default:
			return "xxx", fmt.Errorf("ueType unknown array schema %s", schema.Dump())
		}
	}
	return text, nil
}

func ueType(prefix string, node *apimodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueType node is nil")
	}
	return ToTypeString(prefix, &node.Schema)
}
