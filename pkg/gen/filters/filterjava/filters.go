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
	fm["javaTestValue"] = javaTestValue
	fm["javaElementType"] = javaElementType
	fm["javaListReturn"] = javaListReturn
	fm["javaListType"] = javaListType
	fm["javaListParam"] = javaListParam
	fm["javaListParams"] = javaListParams
	fm["javaListDefault"] = javaListDefault
	fm["javaListAsyncReturn"] = javaListAsyncReturn
}
