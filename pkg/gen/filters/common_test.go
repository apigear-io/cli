package filters

import (
	"testing"

	"github.com/apigear-io/cli/pkg/helper"
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
				t.Errorf("IntToWordLower(%d) = %q, want %q", tt.in, got, tt.out)
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
				t.Errorf("IntToWordTitle(%d) = %q, want %q", tt.in, got, tt.out)
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
				t.Errorf("IntToWordUpper(%d) = %q, want %q", tt.in, got, tt.out)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	var tests = []struct {
		in    string
		lower string
		title string
		upper string
	}{
		{"", "", "", ""},
		{"foo", "foo", "Foo", "FOO"},
		{"fooBar", "foo_bar", "Foo_Bar", "FOO_BAR"},
		{"foo_bar", "foo_bar", "Foo_Bar", "FOO_BAR"},
		{"foo.bar", "foo_bar", "Foo_Bar", "FOO_BAR"},
		{"foo1bar", "foo1bar", "Foo1bar", "FOO1BAR"},
	}
	for _, tt := range tests {
		t.Run(tt.lower, func(t *testing.T) {
			if got := SnakeCaseLower(tt.in); got != tt.lower {
				t.Errorf("SnakeCaseLower(%q) = %q, want %q", tt.in, got, tt.lower)
			}
		})
		t.Run(tt.title, func(t *testing.T) {
			if got := SnakeTitleCase(tt.in); got != tt.title {
				t.Errorf("SnakeTitleCase(%q) = %q, want %q", tt.in, got, tt.title)
			}
		})
		t.Run(tt.upper, func(t *testing.T) {
			if got := SnakeUpperCase(tt.in); got != tt.upper {
				t.Errorf("SnakeUpperCase(%q) = %q, want %q", tt.in, got, tt.upper)
			}
		})
	}
}

func TestCamelCase(t *testing.T) {
	var tests = []struct {
		in    string
		lower string
		title string
		upper string
	}{
		{"", "", "", ""},
		{"foo", "foo", "Foo", "FOO"},
		{"fooBar", "fooBar", "FooBar", "FOOBAR"},
		{"foo_bar", "fooBar", "FooBar", "FOOBAR"},
		{"foo.bar", "fooBar", "FooBar", "FOOBAR"},
		{"foo1bar", "foo1bar", "Foo1bar", "FOO1BAR"},
	}
	for _, tt := range tests {
		t.Run(tt.lower, func(t *testing.T) {
			if got := CamelLowerCase(tt.in); got != tt.lower {
				t.Errorf("CamelLowerCase(%q) = %q, want %q", tt.in, got, tt.lower)
			}
		})
		t.Run(tt.title, func(t *testing.T) {
			if got := CamelTitleCase(tt.in); got != tt.title {
				t.Errorf("CamelTitleCase(%q) = %q, want %q", tt.in, got, tt.title)
			}
		})
		t.Run(tt.upper, func(t *testing.T) {
			if got := CamelUpperCase(tt.in); got != tt.upper {
				t.Errorf("CamelUpperCase(%q) = %q, want %q", tt.in, got, tt.upper)
			}
		})
	}
}

func TestDotCase(t *testing.T) {
	var tests = []struct {
		in    string
		lower string
		title string
		upper string
	}{
		{"", "", "", ""},
		{"foo", "foo", "Foo", "FOO"},
		{"fooBar", "foo.bar", "Foo.Bar", "FOO.BAR"},
		{"foo_bar", "foo.bar", "Foo.Bar", "FOO.BAR"},
		{"foo.bar", "foo.bar", "Foo.Bar", "FOO.BAR"},
		{"foo1bar", "foo1bar", "Foo1bar", "FOO1BAR"},
	}
	for _, tt := range tests {
		t.Run(tt.lower, func(t *testing.T) {
			if got := DotLowerCase(tt.in); got != tt.lower {
				t.Errorf("DotLowerCase(%q) = %q, want %q", tt.in, got, tt.lower)
			}
		})
		t.Run(tt.title, func(t *testing.T) {
			if got := DotTitleCase(tt.in); got != tt.title {
				t.Errorf("DotTitleCase(%q) = %q, want %q", tt.in, got, tt.title)
			}
		})
		t.Run(tt.upper, func(t *testing.T) {
			if got := DotUpperCase(tt.in); got != tt.upper {
				t.Errorf("DotUpperCase(%q) = %q, want %q", tt.in, got, tt.upper)
			}
		})
	}
}

func TestKebabCase(t *testing.T) {
	var tests = []struct {
		in    string
		lower string
		title string
		upper string
	}{
		{"", "", "", ""},
		{"foo", "foo", "Foo", "FOO"},
		{"fooBar", "foo-bar", "Foo-Bar", "FOO-BAR"},
		{"foo_bar", "foo-bar", "Foo-Bar", "FOO-BAR"},
		{"foo.bar", "foo-bar", "Foo-Bar", "FOO-BAR"},
		{"foo1bar", "foo1bar", "Foo1bar", "FOO1BAR"},
	}
	for _, tt := range tests {
		t.Run(tt.lower, func(t *testing.T) {
			if got := KebabLowerCase(tt.in); got != tt.lower {
				t.Errorf("KebabLowerCase(%q) = %q, want %q", tt.in, got, tt.lower)
			}
		})
		t.Run(tt.title, func(t *testing.T) {
			if got := KebabTitleCase(tt.in); got != tt.title {
				t.Errorf("KebabTitleCase(%q) = %q, want %q", tt.in, got, tt.title)
			}
		})
		t.Run(tt.upper, func(t *testing.T) {
			if got := KebabUpperCase(tt.in); got != tt.upper {
				t.Errorf("KebabUpperCase(%q) = %q, want %q", tt.in, got, tt.upper)
			}
		})
	}
}

func TestPathCase(t *testing.T) {
	var tests = []struct {
		in    string
		lower string
		title string
		upper string
	}{
		{"", "", "", ""},
		{"foo", "foo", "Foo", "FOO"},
		{"fooBar", "foo/bar", "Foo/Bar", "FOO/BAR"},
		{"foo_bar", "foo/bar", "Foo/Bar", "FOO/BAR"},
		{"foo.bar", "foo/bar", "Foo/Bar", "FOO/BAR"},
		{"foo1bar", "foo1bar", "Foo1bar", "FOO1BAR"},
	}
	for _, tt := range tests {
		t.Run(tt.lower, func(t *testing.T) {
			if got := PathLowerCase(tt.in); got != tt.lower {
				t.Errorf("PathLowerCase(%q) = %q, want %q", tt.in, got, tt.lower)
			}
		})
		t.Run(tt.title, func(t *testing.T) {
			if got := PathTitleCase(tt.in); got != tt.title {
				t.Errorf("PathTitleCase(%q) = %q, want %q", tt.in, got, tt.title)
			}
		})
		t.Run(tt.upper, func(t *testing.T) {
			if got := PathUpperCase(tt.in); got != tt.upper {
				t.Errorf("PathUpperCase(%q) = %q, want %q", tt.in, got, tt.upper)
			}
		})
	}
}

func TestUpperCase(t *testing.T) {
	var tests = []struct {
		in string
		up string
	}{
		{"", ""},
		{"foo", "FOO"},
		{"fooBar", "FOOBAR"},
		{"foo_bar", "FOO_BAR"},
		{"foo.bar", "FOO.BAR"},
		{"foo1bar", "FOO1BAR"},
	}
	for _, tt := range tests {
		t.Run(tt.up, func(t *testing.T) {
			if got := UpperCase(tt.in); got != tt.up {
				t.Errorf("Upper(%q) = %q, want %q", tt.in, got, tt.up)
			}
		})
	}
}

func TestLowerCase(t *testing.T) {
	var tests = []struct {
		in string
		lw string
	}{
		{"", ""},
		{"foo", "foo"},
		{"fooBar", "foobar"},
		{"foo_bar", "foo_bar"},
		{"foo.bar", "foo.bar"},
		{"foo1bar", "foo1bar"},
	}
	for _, tt := range tests {
		t.Run(tt.lw, func(t *testing.T) {
			if got := LowerCase(tt.in); got != tt.lw {
				t.Errorf("Lower(%q) = %q, want %q", tt.in, got, tt.lw)
			}
		})
	}
}

func TestUpperFirst(t *testing.T) {
	var tests = []struct {
		in string
		up string
	}{
		{"", ""},
		{"foo", "Foo"},
		{"fooBar", "FooBar"},
		{"foo_bar", "Foo_bar"},
		{"foo.bar", "Foo.bar"},
		{"foo1bar", "Foo1bar"},
	}
	for _, tt := range tests {
		t.Run(tt.up, func(t *testing.T) {
			if got := UpperFirst(tt.in); got != tt.up {
				t.Errorf("UpperFirst(%q) = %q, want %q", tt.in, got, tt.up)
			}
		})
	}
}

func TestLowerFirst(t *testing.T) {
	var tests = []struct {
		in string
		lw string
	}{
		{"", ""},
		{"foo", "foo"},
		{"fooBar", "fooBar"},
		{"foo_bar", "foo_bar"},
		{"foo.bar", "foo.bar"},
		{"foo1bar", "foo1bar"},
	}
	for _, tt := range tests {
		t.Run(tt.lw, func(t *testing.T) {
			if got := LowerFirst(tt.in); got != tt.lw {
				t.Errorf("LowerFirst(%q) = %q, want %q", tt.in, got, tt.lw)
			}
		})
	}
}

func TestFirstChar(t *testing.T) {
	var tests = []struct {
		in string
		ch string
	}{
		{"", ""},
		{"foo", "f"},
		{"fooBar", "f"},
		{"foo_bar", "f"},
		{"Foo.Bar", "F"},
		{"Foo1Bar", "F"},
	}
	for _, tt := range tests {
		t.Run(tt.ch, func(t *testing.T) {
			if got := FirstChar(tt.in); got != tt.ch {
				t.Errorf("FirstChar(%q) = %q, want %q", tt.in, got, tt.ch)
			}
		})
	}
}

func TestFirstCharUpperAndLower(t *testing.T) {
	var tests = []struct {
		in string
		u  string
		l  string
	}{
		{"", "", ""},
		{"foo", "F", "f"},
		{"fooBar", "F", "f"},
		{"foo_bar", "F", "f"},
		{"Foo.Bar", "F", "f"},
		{"Foo1Bar", "F", "f"},
	}
	for _, tt := range tests {
		t.Run(tt.u, func(t *testing.T) {
			if got := FirstCharUpper(tt.in); got != tt.u {
				t.Errorf("FirstCharUpper(%q) = %q, want %q", tt.in, got, tt.u)
			}
		})
		t.Run(tt.l, func(t *testing.T) {
			if got := FirstCharLower(tt.in); got != tt.l {
				t.Errorf("FirstCharLower(%q) = %q, want %q", tt.in, got, tt.l)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	var tests = []struct {
		in  []string
		sep string
		out string
	}{
		{[]string{}, "", ""},
		{[]string{"foo"}, "", "foo"},
		{[]string{"foo", "bar"}, "", "foobar"},
		{[]string{"foo", "bar"}, " ", "foo bar"},
		{[]string{"foo", "bar"}, "-", "foo-bar"},
		{[]string{"foo", "bar"}, "_", "foo_bar"},
		{[]string{"foo", "bar"}, ".", "foo.bar"},
		{[]string{"foo", "bar"}, "1", "foo1bar"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			if got := Join(tt.sep, tt.in); got != tt.out {
				t.Errorf("Join(%q, %q) = %q, want %q", tt.in, tt.sep, got, tt.out)
			}
		})
	}
}

func TestTrimPrefix(t *testing.T) {
	var tests = []struct {
		in     string
		prefix string
		out    string
	}{
		{"", "", ""},
		{"foo", "", "foo"},
		{"foo", "foo", ""},
		{"foo", "bar", "foo"},
		{"fooBar", "foo", "Bar"},
		{"foo_bar", "foo", "_bar"},
		{"foo.bar", "foo", ".bar"},
		{"foo1bar", "foo", "1bar"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			if got := TrimPrefix(tt.in, tt.prefix); got != tt.out {
				t.Errorf("TrimPrefix(%q, %q) = %q, want %q", tt.in, tt.prefix, got, tt.out)
			}
		})
	}
}

func TestTrimSuffix(t *testing.T) {
	var tests = []struct {
		in     string
		suffix string
		out    string
	}{
		{"", "", ""},
		{"foo", "", "foo"},
		{"foo", "foo", ""},
		{"foo", "bar", "foo"},
		{"fooBar", "Bar", "foo"},
		{"foo_bar", "_bar", "foo"},
		{"foo.bar", ".bar", "foo"},
		{"foo1bar", "1bar", "foo"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			if got := TrimSuffix(tt.in, tt.suffix); got != tt.out {
				t.Errorf("TrimSuffix(%q, %q) = %q, want %q", tt.in, tt.suffix, got, tt.out)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	var tests = []struct {
		in  string
		old string
		new string
		out string
	}{
		{"", "", "", ""},
		{"foo", "", "", "foo"},
		{"foo", "foo", "", ""},
		{"foo", "bar", "", "foo"},
		{"foo", "foo", "bar", "bar"},
		{"foo", "foo", "bar", "bar"},
		{"fooBar", "foo", "bar", "barBar"},
		{"foo_bar", "foo", "bar", "bar_bar"},
		{"foo.bar", "foo", "bar", "bar.bar"},
		{"foo1bar", "foo", "bar", "bar1bar"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			if got := Replace(tt.in, tt.old, tt.new); got != tt.out {
				t.Errorf("Replace(%q, %q, %q) = %q, want %q", tt.in, tt.old, tt.new, got, tt.out)
			}
		})
	}
}

func TestPluralize(t *testing.T) {
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
