package filterjava

import "text/template"

func PopulateFuncMap() template.FuncMap {
	fm := make(template.FuncMap)
	fm["javaDefault"] = javaDefault
	fm["javaReturn"] = javaReturn
	fm["javaParam"] = javaParam
	fm["javaParams"] = javaParams
	fm["javaVar"] = javaVar
	fm["javaVars"] = javaVars
	fm["javaType"] = javaType
	fm["javaExtern"] = javaExtern
	return fm
}
