package sim

import (
	"fmt"
	"strings"
)

// Symbol is a symbol in the form of <interface>/<resource>.
type Symbol string

// String returns the string representation of a symbol.
func (s Symbol) String() string {
	return string(s)
}

// Split splits a symbol into its path and resource parts.
// A symbol is in the form of <interface>/<resource>.
// A an interface can be prefixed with a module name using the dot (".") separator.
func (s Symbol) Split() (string, string) {
	parts := strings.Split(string(s), "/")
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

// Resource returns the resource part of a symbol.
func (s Symbol) Resource() string {
	_, resource := s.Split()
	return resource
}

// Path returns the path part of a symbol.
func (s Symbol) Path() string {
	path, _ := s.Split()
	return path
}

// MakeSymbol creates a symbol from a path and resource.
func MakeSymbol(path, resource string) Symbol {
	if resource == "" {
		return Symbol(path)
	}
	return Symbol(fmt.Sprintf("%s/%s", path, resource))
}
