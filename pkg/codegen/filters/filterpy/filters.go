package filterpy

import (
	"text/template"
)

// PopulateFuncMap fills the given FuncMap with the functions from this package.
func PopulateFuncMap(fm template.FuncMap) {
	fm["pyReturn"] = pyReturn
	fm["pyDefault"] = pyDefault
	fm["pyParam"] = pyParam
	fm["pyParams"] = pyParams
	fm["pyFuncParams"] = pyFuncParams
	fm["pyVar"] = pyVar
	fm["pyVars"] = pyVars
	fm["pyType"] = pyType
	fm["pyExtern"] = pyExtern
	fm["pyTestValue"] = pyTestValue
}
