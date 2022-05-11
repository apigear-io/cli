package filtercpp

import (
	"text/template"
)

// PopulateFuncMap fills the given FuncMap with the functions from this package.
func PopulateFuncMap(fm template.FuncMap) {
	fm["cpp_ns"] = ns
	fm["cpp_ns_open"] = nsOpen
	fm["cpp_ns_close"] = nsClose
	fm["cpp_return"] = cppReturn
	fm["cpp_default"] = cppDefault
	fm["cpp_param"] = cppParam
	fm["cpp_params"] = cppParams
}
