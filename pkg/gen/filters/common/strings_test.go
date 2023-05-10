package common

import "testing"

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
