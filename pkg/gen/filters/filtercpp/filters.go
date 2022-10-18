package filtercpp

import (
	"text/template"
)

// PopulateFuncMap fills the given FuncMap with the functions from this package.
func PopulateFuncMap(fm template.FuncMap) {
	fm["cppNs"] = ns
	fm["cppNsOpen"] = nsOpen
	fm["cppNsClose"] = nsClose
	fm["cppReturn"] = cppReturn
	fm["cppDefault"] = cppDefault
	fm["cppParam"] = cppParam
	fm["cppParams"] = cppParams
	fm["cppGpl"] = cppGpl
	fm["cppVar"] = cppVar
	fm["cppVars"] = cppVars
	fm["cppType"] = cppType
}
