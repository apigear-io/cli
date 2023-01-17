package filtergo

import (
	"text/template"
)

// PopulateFuncMap fills the given FuncMap with the functions from this package.
func PopulateFuncMap(fm template.FuncMap) {
	fm["goReturn"] = goReturn
	fm["goDefault"] = goDefault
	fm["goParam"] = goParam
	fm["goParams"] = goParams
	fm["goType"] = goType
	fm["goVar"] = goVar
	fm["goVars"] = goVars
	fm["goPublicVar"] = goPublicVar
	fm["goPublicVars"] = goPublicVars
	fm["goDoc"] = goDoc
}
