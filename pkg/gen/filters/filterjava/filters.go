package filterjava

import "text/template"

func PopulateFuncMap(fm template.FuncMap) {
	fm["javaDefault"] = javaDefault
	fm["javaReturn"] = javaReturn
	fm["javaParam"] = javaParam
	fm["javaParams"] = javaParams
	fm["javaVar"] = javaVar
	fm["javaVars"] = javaVars
	fm["javaType"] = javaType
	fm["javaExtern"] = javaExtern
	fm["javaAsyncReturn"] = javaAsyncReturn
	fm["javaElementType"] = javaElementType
}
