package helper

import "strings"

func JoinMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func SelectValue(m map[string]interface{}, selector string) interface{} {
	keys := strings.Split(selector, ".")
	for _, k := range keys {
		// recursively select the value
		if v, ok := m[k]; ok {
			if m, ok = v.(map[string]interface{}); !ok {
				return v
			}
		} else {
			return nil
		}
	}
	return m
}
