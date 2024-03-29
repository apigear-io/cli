package common

import (
	"testing"

	"github.com/apigear-io/cli/pkg/helper"
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
		{"HELlo.worLD", "HLWL"},
		{"hello_world", "HW"},
		{"heLlo_wOrld", "HLWO"},
		{"hello world", "HW"},
		{"hello worlD", "HWD"},
		{"hello.2world", "H"},
		{"1hello.world", "W"},
		{"1hello.2world", ""},
		{"1hELlo.2world", "EL"},
		{"1hELLlo.2world", "EL"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			if got := helper.Abbreviate(tt.in); got != tt.out {
				t.Errorf("Abbreviate(%q) = %q, want %q", tt.in, got, tt.out)
			}
		})
	}
}
