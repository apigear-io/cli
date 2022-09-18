package core

import (
	"fmt"
	"strings"
)

func SplitSymbol(symbol string) (string, string) {
	parts := strings.Split(symbol, "/")
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

func MakeSymbol(typ, member string) string {
	if member == "" {
		return typ
	}
	return fmt.Sprintf("%s/%s", typ, member)
}
