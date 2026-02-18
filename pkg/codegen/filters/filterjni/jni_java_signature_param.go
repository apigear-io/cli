package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/codegen/filters/common"
	"github.com/apigear-io/cli/pkg/codegen/filters/filterjava"
	"github.com/apigear-io/cli/pkg/objmodel"
)

func makeFullTypeName(module string, typename string) string {
	var camelModuleName = common.CamelLowerCase(module)
	packageName := camelModuleName + "/" + camelModuleName + "_api"
	var text = "L" + packageName + "/" + common.CamelTitleCase(typename) + ";"
	return text
}

func jniSignatureType(node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("jniSignatureType node is nil")
	}

	var text string
	switch node.KindType {
	case objmodel.TypeString:
		text = "Ljava/lang/String;"
	case objmodel.TypeInt:
		text = "I"
	case objmodel.TypeInt32:
		text = "I"
	case objmodel.TypeInt64:
		text = "J"
	case objmodel.TypeFloat:
		text = "F"
	case objmodel.TypeFloat32:
		text = "F"
	case objmodel.TypeFloat64:
		text = "D"
	case objmodel.TypeBool:
		text = "Z"
	case objmodel.TypeVoid:
		text = "V"
	// enums are expected to passed as integers
	case objmodel.TypeEnum:
		e := node.LookupEnum(node.Import, node.Type)
		if e != nil {
			text = makeFullTypeName(e.Module.Name, e.Name)
		} else {
			return "xxx", fmt.Errorf("ToSignatureType interface not found %s", node.Dump())
		}
	case objmodel.TypeStruct:
		s := node.LookupStruct(node.Import, node.Type)
		if s != nil {
			text = makeFullTypeName(s.Module.Name, s.Name)
		} else {
			return "xxx", fmt.Errorf("ToSignatureType interface not found %s", node.Dump())
		}
	case objmodel.TypeExtern:
		xe := filterjava.MakeJavaExtern(&node.Schema)
		if xe.Package != "" {
			java_module := xe.Package
			java_module = common.Replace(java_module, ".", "/")
			text = "L" + java_module + "/" + xe.Name + ";"
		} else {
			text = "L" + xe.Name + ";"
		}
	case objmodel.TypeInterface:
		i := node.LookupInterface(node.Import, node.Type)
		if i != nil {
			name := "I" + i.Name
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

func jniJavaSignatureParam(node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("jniJavaSignatureParam called with nil nodes")
	}
	return jniSignatureType(node)
}
