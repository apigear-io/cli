package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/gen/filters/filterjava"
	"github.com/apigear-io/cli/pkg/model"
)

func makeFullTypeName(module string, typename string) string {
	var camelModuleName = common.CamelLowerCase(module)
	packageName := camelModuleName + "/" + camelModuleName + "_api"
	var text = "L" + packageName + "/" + common.CamelTitleCase(typename) + ";"
	return text
}

func jniSignatureType(node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("jniSignatureType node is nil")
	}

	var text string
	switch node.KindType {
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
		e := node.LookupEnum(node.Import, node.Type)
		if e != nil {
			text = makeFullTypeName(e.Module.Name, e.Name)
		} else {
			return "xxx", fmt.Errorf("ToSignatureType interface not found %s", node.Dump())
		}
	case model.TypeStruct:
		s := node.LookupStruct(node.Import, node.Type)
		if s != nil {
			text = makeFullTypeName(s.Module.Name, s.Name)
		} else {
			return "xxx", fmt.Errorf("ToSignatureType interface not found %s", node.Dump())
		}
	case model.TypeExtern:
		xe := filterjava.MakeJavaExtern(&node.Schema)
		var java_module string
		java_module = ""
		if xe.Package != "" {
			java_module = xe.Package
			java_module = common.Replace(java_module, ".", "/")
			text = "L" + java_module + "/" + xe.Name + ";"
		} else {
			text = "L" + xe.Name + ";"
		}
	case model.TypeInterface:
		i := node.LookupInterface(node.Import, node.Type)
		if i != nil {
			var name string
			name = "I" + i.Name
			text = makeFullTypeName(i.Module.Name, name)
		} else {
			return "xxx", fmt.Errorf("ToSignatureType interface not found %s", node.Dump())
		}
	default:
		return "xxx", fmt.Errorf("jniJavaSignatureParam unknown schema %s", node.Dump())
	}
	if node.IsArray {
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
