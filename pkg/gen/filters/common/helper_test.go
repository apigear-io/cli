package common

import (
	"testing"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/stretchr/testify/assert"
)

func TestIntToWordLower(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		in   int
		pre  string
		post string
		out  string
	}{
		{0, "", "", ""},
		{1, "", "", "one"},
		{2, "", "", "two"},
		{3, "", "", "three"},
		{1, "pre", "", "preone"},
		{1, "", "post", "onepost"},
		{2, "", "post", "twoposts"},
		{3, "pre", "post", "prethreeposts"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			if got := IntToWordLower(tt.in, tt.pre, tt.post); got != tt.out {
				t.Errorf("IntToWordLower(%d) = %q, want %q", tt.in, got, tt.out)
			}
		})
	}
}

func TestIntToWordCamel(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		in   int
		pre  string
		post string
		out  string
	}{
		{0, "", "", ""},
		{1, "", "", "One"},
		{2, "", "", "Two"},
		{3, "", "", "Three"},
		{1, "pre", "", "preOne"},
		{1, "", "post", "Onepost"},
		{2, "", "post", "Twoposts"},
		{3, "pre", "post", "preThreeposts"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			if got := IntToWordTitle(tt.in, tt.pre, tt.post); got != tt.out {
				t.Errorf("IntToWordTitle(%d) = %q, want %q", tt.in, got, tt.out)
			}
		})
	}
}

func TestIntToWordUpper(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		in   int
		pre  string
		post string
		out  string
	}{
		{0, "", "", ""},
		{1, "", "", "ONE"},
		{2, "", "", "TWO"},
		{3, "", "", "THREE"},
		{1, "pre", "", "preONE"},
		{2, "", "post", "TWOposts"},
		{3, "pre", "post", "preTHREEposts"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			if got := IntToWordUpper(tt.in, tt.pre, tt.post); got != tt.out {
				t.Errorf("IntToWordUpper(%d) = %q, want %q", tt.in, got, tt.out)
			}
		})
	}
}

func TestPluralize(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		in  string
		n   int
		out string
	}{
		{"", 0, ""},
		{"", 1, ""},
		{"", 2, ""},
		{"foo", 0, "foo"},
		{"foo", 1, "foo"},
		{"foo", 2, "foos"},
		{"foo", 3, "foos"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			if got := Pluralize(tt.in, tt.n); got != tt.out {
				t.Errorf("Pluralize(%q) = %q, want %q", tt.in, got, tt.out)
			}
		})
	}
}

func TestAbbreviate(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		in  string
		out string
	}{
		{"", ""},
		{"foo", "F"},
		{"hello.world", "HW"},
		{"hEllo.worLd", "HEWL"},
		{"HEllo.worLd", "HEWL"},
		{"HEllo.worLD", "HEWL"},
		{"HEllo.worLD2", "HEWL2"},
		{"HELlo.worLD", "HLWL"},
		{"hello_world", "HW"},
		{"heLlo_wOrld", "HLWO"},
		{"hello world", "HW"},
		{"hello worlD", "HWD"},
		{"hello.2world", "H2"},
		{"1hello.world", "W"},
		{"1hello.2world", ""},
		{"1hELlo.2world", "EL2"},
		{"1hELLlo.2world", "EL2"},
		{"1hELLlo.2woRld", "EL2R"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			if got := helper.Abbreviate(tt.in); got != tt.out {
				t.Errorf("Abbreviate(%q) = %q, want %q", tt.in, got, tt.out)
			}
		})
	}
}

func TestCollectFields(t *testing.T) {
	t.Parallel()
	var listOfStructs = []struct {
		a string
		b string
	}{
		{"foo1", "goo1"},
		{"foo2", "goo2"},
		{"foo3", "goo3"},
	}
	var listA = []string{"foo1", "foo2", "foo3"}
	var listB = []string{"goo1", "goo2", "goo3"}
	t.Run("getListOfFields", func(t *testing.T) {
		resultA, err := CollectFields(listOfStructs, "a")
		assert.Equal(t, listA, resultA)
		assert.NoError(t, err)
		resultB, err := CollectFields(listOfStructs, "b")
		assert.Equal(t, listB, resultB)
		assert.NoError(t, err)
	})
}

func TestNoFieldCollectFields(t *testing.T) {
	t.Parallel()
	var listOfStructs = []struct {
		a string
		b string
	}{
		{"foo1", "goo1"},
	}
	var emptyList = []string{}
	t.Run("getListOfFields", func(t *testing.T) {
		result, err := CollectFields(listOfStructs, "c")
		assert.Equal(t, emptyList, result)
		assert.Error(t, err)
	})
}

func TestUnique(t *testing.T) {
	t.Parallel()

	var inputList = []string{"abc", "<def>", "<ghi>", "ab::c", "abc", "ghi", "<ghi>"}
	var noDuplicatesList = []string{"<def>", "<ghi>", "ab::c", "abc", "ghi"}

	t.Run("getListOfFields", func(t *testing.T) {
		assert.Equal(t, noDuplicatesList, Unique(inputList))
	})
}
