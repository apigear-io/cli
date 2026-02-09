package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/apimodel"
)

func ToType(schema *apimodel.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToType schema is nil")
	}

	var text string
	switch schema.KindType {
	case apimodel.TypeString:
		text = "jstring"
	case apimodel.TypeInt:
		text = "jint"
	case apimodel.TypeInt32:
		text = "jint"
	case apimodel.TypeInt64:
		text = "jlong"
	case apimodel.TypeFloat:
		text = "jfloat"
	case apimodel.TypeFloat32:
		text = "jfloat"
	case apimodel.TypeFloat64:
		text = "jdouble"
	case apimodel.TypeBool:
		text = "jboolean"
	case apimodel.TypeVoid:
		text = "void"
	// enums are expected to passed as integers
	case apimodel.TypeEnum:
		text = "jobject"
	case apimodel.TypeStruct:
		text = "jobject"
	case apimodel.TypeExtern:
		text = "jobject"
	case apimodel.TypeInterface:
		text = "jobject"
	default:
		return "xxx", fmt.Errorf("jniToReturnType unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		if schema.KindType == apimodel.TypeString {
			text = "jobject"
		}
		text = fmt.Sprintf("%sArray", text)
	}
	return text, nil
}

func jniToReturnType(node *apimodel.TypedNode) (string, error) {
	return ToType(&node.Schema)
}
