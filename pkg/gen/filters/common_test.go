package filters

import (
	"testing"
)

func TestIntToWordLower(t *testing.T) {
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
				t.Errorf("IntToWord(%d) = %q, want %q", tt.in, got, tt.out)
			}
		})
	}
}

func TestIntToWordCamel(t *testing.T) {
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
				t.Errorf("IntToWord(%d) = %q, want %q", tt.in, got, tt.out)
			}
		})
	}
}

func TestIntToWordUpper(t *testing.T) {
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
				t.Errorf("IntToWord(%d) = %q, want %q", tt.in, got, tt.out)
			}
		})
	}
}
