package filterts

import (
	"text/template"
)

// PopulateFuncMap fills the given FuncMap with the functions from this package.
func PopulateFuncMap(fm template.FuncMap) {
	fm["tsReturn"] = tsReturn
	fm["tsDefault"] = tsDefault
	fm["tsParam"] = tsParam
	fm["tsParams"] = tsParams
	fm["tsVar"] = tsVar
	fm["tsVars"] = tsVars
	fm["tsType"] = tsType
}
