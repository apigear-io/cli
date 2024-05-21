package common

import (
	"fmt"
	"slices"
	"reflect"

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

func MakeListOfFields[T any](inputList []T, fieldName string) ([]string, error) {
	list := []string {}
	for _, element := range inputList {
	    r := reflect.ValueOf(element)
		reflectValue := reflect.Indirect(r).FieldByName(fieldName)
		if !reflectValue.IsValid() {
			return list, fmt.Errorf("given struct %T has no field %s ",  element, fieldName)
		}
		value := reflectValue.String()
		list = append(list, value)
	}
	return list, nil
}

func RemoveDuplicates (inputList []string) []string {
	list := inputList
    slices.Sort(list) 
    list = slices.Compact(list)
	return list
}
