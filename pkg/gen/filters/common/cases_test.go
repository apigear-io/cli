package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		{"fooBar_", "foo_bar_", "Foo_Bar_", "FOO_BAR_"},
	}
	for _, tt := range tests {
		t.Run(tt.lower, func(t *testing.T) {
			assert.Equal(t, tt.lower, SnakeCaseLower(tt.in))
		})
		t.Run(tt.title, func(t *testing.T) {
			assert.Equal(t, tt.title, SnakeTitleCase(tt.in))
		})
		t.Run(tt.upper, func(t *testing.T) {
			assert.Equal(t, tt.upper, SnakeUpperCase(tt.in))
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
		{"fooBar_", "fooBar_", "FooBar_", "FOOBAR_"},
	}
	for _, tt := range tests {
		t.Run(tt.lower, func(t *testing.T) {
			assert.Equal(t, tt.lower, CamelLowerCase(tt.in))
		})
		t.Run(tt.title, func(t *testing.T) {
			assert.Equal(t, tt.title, CamelTitleCase(tt.in))
		})
		t.Run(tt.upper, func(t *testing.T) {
			assert.Equal(t, tt.upper, CamelUpperCase(tt.in))
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
		{"fooBar_", "foo.bar_", "Foo.Bar_", "FOO.BAR_"},
	}
	for _, tt := range tests {
		t.Run(tt.lower, func(t *testing.T) {
			assert.Equal(t, tt.lower, DotLowerCase(tt.in))
		})
		t.Run(tt.title, func(t *testing.T) {
			assert.Equal(t, tt.title, DotTitleCase(tt.in))
		})
		t.Run(tt.upper, func(t *testing.T) {
			assert.Equal(t, tt.upper, DotUpperCase(tt.in))
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
		{"fooBar_", "foo-bar_", "Foo-Bar_", "FOO-BAR_"},
	}
	for _, tt := range tests {
		t.Run(tt.lower, func(t *testing.T) {
			assert.Equal(t, tt.lower, KebabLowerCase(tt.in))
		})
		t.Run(tt.title, func(t *testing.T) {
			assert.Equal(t, tt.title, KebabTitleCase(tt.in))
		})
		t.Run(tt.upper, func(t *testing.T) {
			assert.Equal(t, tt.upper, KebabUpperCase(tt.in))
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
		{"fooBar_", "foo/bar_", "Foo/Bar_", "FOO/BAR_"},
	}
	for _, tt := range tests {
		t.Run(tt.lower, func(t *testing.T) {
			assert.Equal(t, tt.lower, PathLowerCase(tt.in))
		})
		t.Run(tt.title, func(t *testing.T) {
			assert.Equal(t, tt.title, PathTitleCase(tt.in))
		})
		t.Run(tt.upper, func(t *testing.T) {
			assert.Equal(t, tt.upper, PathUpperCase(tt.in))
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
		{"fooBar_", "FOOBAR_"},
	}
	for _, tt := range tests {
		t.Run(tt.up, func(t *testing.T) {
			assert.Equal(t, tt.up, UpperCase(tt.in))
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
		{"fooBar_", "foobar_"},
	}
	for _, tt := range tests {
		t.Run(tt.lw, func(t *testing.T) {
			assert.Equal(t, tt.lw, LowerCase(tt.in))
		})
	}
}
