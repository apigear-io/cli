package filterjs

import (
	"text/template"
)

// PopulateFuncMap fills the given FuncMap with the functions from this package.
func PopulateFuncMap(fm template.FuncMap) {
	fm["jsReturn"] = jsReturn
	fm["jsDefault"] = jsDefault
	fm["jsParam"] = jsParam
	fm["jsParams"] = jsParams
	fm["jsVar"] = jsVar
	fm["jsVars"] = jsVars
	fm["jsType"] = jsType
}
