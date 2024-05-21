package filterqt

import (
	"text/template"
)

// PopulateFuncMap fills the given FuncMap with the functions from this package.
func PopulateFuncMap(fm template.FuncMap) {
	fm["qtReturn"] = qtReturn
	fm["qtDefault"] = qtDefault
	fm["qtParam"] = qtParam
	fm["qtParams"] = qtParams
	fm["qtVar"] = qtVar
	fm["qtVars"] = qtVars
	fm["qtType"] = qtType
	fm["qtNamespace"] = qtNamespace
	fm["qtExtern"] = qtExtern
	fm["qtExterns"] = qtExterns
	fm["qtMakeListOfFields_extern"] = qtMakeListOfFields_extern
}
