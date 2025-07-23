package filterue

import (
	"text/template"
)

// PopulateFuncMap fills the given FuncMap with the functions from this package.
func PopulateFuncMap(fm template.FuncMap) {
	fm["ueParam"] = ueParam
	fm["ueParams"] = ueParams
	fm["ueReturn"] = ueReturn
	fm["ueDefault"] = ueDefault
	fm["ueTestValue"] = ueTestValue
	fm["ueConstType"] = ueConstType
	fm["ueType"] = ueType
	fm["ueVar"] = ueVar
	fm["ueVars"] = ueVars
	fm["ueIsStdSimpleType"] = ueIsStdSimpleType
	fm["ueExtern"] = ueExtern
	fm["ueJavaPath"] = ueJavaPath
	fm["ueJavaPckgName"] = ueJavaPckgName
	fm["ueJniJavaParam"] = ueJniJavaParam
	fm["ueJniJavaParams"] = ueJniJavaParams
	fm["ueJniClassPathPrefix"] = ueJniClassPathPrefix
	fm["javaOnSingalStyle"] = javaOnSingalStyle
	fm["javaOnPropertyChangedStyle"] = javaOnPropertyChangedStyle
	fm["ueJniToReturnType"] = ueToReturnType
	fm["jniNameSetProperty"] = jniNameSetProperty
	fm["jniNameGetProperty"] = jniNameGetProperty
	fm["jniNameOperation"] = jniNameOperation
	fm["ueJniSignatureType"] = ueJniSignatureType
	fm["ueJniJavaSignatureParams"] = ueJniJavaSignatureParams
	fm["ueToEnvNameType"] = ueToEnvNameType
	fm["ueJniJavaSignatureParam"] = ueJniJavaSignatureParam
	fm["ueJniEmptyReturn"] = ueJniEmptyReturn

}
