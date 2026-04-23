package net

import "strings"

// NormalizePath ensures a path starts with a leading slash and removes any trailing slash (except for root).
func NormalizePath(path string) string {
	p := strings.TrimSpace(path)
	if p == "" || p == "/" {
		return "/"
	}
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	if len(p) > 1 && strings.HasSuffix(p, "/") {
		p = strings.TrimRight(p, "/")
		if p == "" {
			return "/"
		}
	}
	return p
}
