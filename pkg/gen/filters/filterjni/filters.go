package filterjni

import (
	"text/template"
)

// PopulateFuncMap fills the given FuncMap with the functions from this package.
func PopulateFuncMap(fm template.FuncMap) {
	fm["jniToReturnType"] = jniToReturnType
	fm["jniJavaParam"] = jniJavaParam
	fm["jniJavaParams"] = jniJavaParams
	fm["jniJavaSignatureParam"] = jniJavaSignatureParam
	fm["jniJavaSignatureParams"] = jniJavaSignatureParams
	fm["jniSignatureType"] = jniSignatureType
	fm["jniToEnvNameType"] = jniToEnvNameType
}
