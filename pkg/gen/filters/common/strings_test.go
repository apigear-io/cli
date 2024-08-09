package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpperFirst(t *testing.T) {
	t.Parallel()
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
		{"fooBar_", "FooBar_"},
	}
	for _, tt := range tests {
		t.Run(tt.up, func(t *testing.T) {
			assert.Equal(t, tt.up, UpperFirst(tt.in))
		})
	}
}

func TestLowerFirst(t *testing.T) {
	t.Parallel()
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
		{"fooBar_", "fooBar_"},
	}
	for _, tt := range tests {
		t.Run(tt.lw, func(t *testing.T) {
			assert.Equal(t, tt.lw, LowerFirst(tt.in))
		})
	}
}

func TestFirstChar(t *testing.T) {
	t.Parallel()
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
		{"FooBar_", "F"},
	}
	for _, tt := range tests {
		t.Run(tt.ch, func(t *testing.T) {
			assert.Equal(t, tt.ch, FirstChar(tt.in))
		})
	}
}

func TestFirstCharUpperAndLower(t *testing.T) {
	t.Parallel()
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
		{"FooBar_", "F", "f"},
	}
	for _, tt := range tests {
		t.Run(tt.u, func(t *testing.T) {
			assert.Equal(t, tt.u, FirstCharUpper(tt.in))
		})
		t.Run(tt.l, func(t *testing.T) {
			assert.Equal(t, tt.l, FirstCharLower(tt.in))
		})
	}
}

func TestJoin(t *testing.T) {
	t.Parallel()
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
		{[]string{"foo_", "bar_"}, "", "foo_bar_"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			assert.Equal(t, tt.out, Join(tt.sep, tt.in))
		})
	}
}

func TestTrimPrefix(t *testing.T) {
	t.Parallel()
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
		{"fooBar_", "foo", "Bar_"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			assert.Equal(t, tt.out, TrimPrefix(tt.in, tt.prefix))
		})
	}
}

func TestTrimSuffix(t *testing.T) {
	t.Parallel()
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
		{"fooBar_", "Bar_", "foo"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			assert.Equal(t, tt.out, TrimSuffix(tt.in, tt.suffix))
		})
	}
}

func TestReplace(t *testing.T) {
	t.Parallel()
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
			assert.Equal(t, tt.out, Replace(tt.in, tt.old, tt.new))
		})
	}
}
