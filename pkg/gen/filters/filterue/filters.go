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
	fm["ueConstType"] = ueConstType
	fm["ueType"] = ueType
	fm["ueVar"] = ueVar
	fm["ueVars"] = ueVars
	fm["ueIsStdSimpleType"] = ueIsStdSimpleType
	fm["ueExtern"] = ueExtern
}
