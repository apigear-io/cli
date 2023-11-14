package common

import "testing"

func TestSnakeCase(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
