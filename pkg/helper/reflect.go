package helper

import "reflect"

func ToSlice(v any) []any {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		slice := make([]any, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			slice[i] = rv.Index(i).Interface()
		}
		return slice
	}
	return []any{v}
}
