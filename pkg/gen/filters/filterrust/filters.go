package filterrust

import (
	"text/template"
)

// PopulateFuncMap fills the given FuncMap with the functions from this package.
func PopulateFuncMap(fm template.FuncMap) {
	fm["rustNs"] = ns
	fm["rustNsOpen"] = nsOpen
	fm["rustNsClose"] = nsClose
	fm["rustReturn"] = rustReturn
	fm["rustDefault"] = rustDefault
	fm["rustParam"] = rustParam
	fm["rustParams"] = rustParams
	fm["rustVar"] = rustVar
	fm["rustVars"] = rustVars
	fm["rustType"] = rustType
	fm["rustTypeRef"] = rustTypeRef
}
