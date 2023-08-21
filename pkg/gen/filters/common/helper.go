package common

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/ettle/strcase"
	"github.com/gertd/go-pluralize"
)

func NewLine() string {
	return "\n"
}

var WORDS = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
	"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty",
}

func IntToWord(i int, prefix string, postfix string, wc strcase.WordCase) string {
	if i <= 0 || i >= len(WORDS) {
		return ""
	}
	plural := Pluralize(postfix, i)
	word := strcase.ToCase(WORDS[i], wc, '\x00')
	return fmt.Sprintf("%s%s%s", prefix, word, plural)
}

func IntToWordLower(i int, prefix string, postfix string) string {
	return IntToWord(i, prefix, postfix, strcase.LowerCase)
}

func IntToWordTitle(i int, prefix string, postfix string) string {
	return IntToWord(i, prefix, postfix, strcase.TitleCase)
}

func IntToWordUpper(i int, prefix string, postfix string) string {
	return IntToWord(i, prefix, postfix, strcase.UpperCase)
}

var plural = pluralize.NewClient()

func Pluralize(s string, i int) string {
	if i <= 1 {
		return s
	}
	return plural.Plural(s)
}

func Sort(in any) ([]any, error) {
	items, err := ConvertInterfaceToSlice(in)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return items, nil
	}
	out := make([]any, len(items))
	copy(out, items)
	sort.Slice(out, func(i, j int) bool {
		return out[i].(string) < out[j].(string)
	})
	return out, nil
}

func ConvertInterfaceToSlice(input any) ([]any, error) {
	if input == nil {
		return nil, nil
	}
	items, ok := input.([]any)
	if ok {
		return items, nil
	}
	i := reflect.ValueOf(input)
	k := i.Kind()
	switch k {
	case reflect.Slice, reflect.Array:
		items = make([]any, i.Len())
		for j := 0; j < i.Len(); j++ {
			items[j] = i.Index(j).Interface()
		}
		return items, nil
	default:
		return nil, fmt.Errorf("expected []any, got %T", input)
	}

}
