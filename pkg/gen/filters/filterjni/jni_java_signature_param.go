package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

func makeFullTypeName(module string, typename string) string {
	var camelModuleName = common.CamelLowerCase(module)
	packageName := camelModuleName + "/" + camelModuleName + "_api"
	var text = "L" + packageName + "/" + typename + ";"
	return text
}

func jniSignatureType(node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("jniSignatureType node is nil")
	}

	var text string
	switch node.Schema.KindType {
	case model.TypeString:
		text = "Ljava/lang/String;"
	case model.TypeInt:
		text = "I"
	case model.TypeInt32:
		text = "I"
	case model.TypeInt64:
		text = "J"
	case model.TypeFloat:
		text = "F"
	case model.TypeFloat32:
		text = "F"
	case model.TypeFloat64:
		text = "D"
	case model.TypeBool:
		text = "Z"
	case model.TypeVoid:
		text = "V"
	// enums are expected to passed as integers
	case model.TypeEnum:
		e := node.Schema.LookupEnum(node.Schema.Import, node.Schema.Type)
		if e != nil {
			text = makeFullTypeName(e.Module.Name, e.Name)
		} else {
			return "xxx", fmt.Errorf("ToSignatureType interface not found %s", node.Schema.Dump())
		}
	case model.TypeStruct:
		s := node.Schema.LookupStruct(node.Schema.Import, node.Schema.Type)
		if s != nil {
			text = makeFullTypeName(s.Module.Name, s.Name)
		} else {
			return "xxx", fmt.Errorf("ToSignatureType interface not found %s", node.Schema.Dump())
		}
	case model.TypeExtern:
		return "xxx", fmt.Errorf("ToSignatureType TypeExtern not supported yet %s", node.Schema.Dump())
	case model.TypeInterface:
		i := node.Schema.LookupInterface(node.Schema.Import, node.Schema.Type)
		if i != nil {
			text = makeFullTypeName(i.Module.Name, i.Name)
		} else {
			return "xxx", fmt.Errorf("ToSignatureType interface not found %s", node.Schema.Dump())
		}
	default:
		return "xxx", fmt.Errorf("jniJavaSignatureParam unknown schema %s", node.Schema.Dump())
	}
	if node.Schema.IsArray {
		text = fmt.Sprintf("[%s", text)
	}
	return text, nil
}

func jniJavaSignatureParam(node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("jniJavaSignatureParam called with nil nodes")
	}
	return jniSignatureType(node)
}
