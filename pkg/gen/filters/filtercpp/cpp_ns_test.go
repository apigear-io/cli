package filtercpp

import (
	"reflect"
	"testing"

	"github.com/apigear-io/cli/pkg/model"

	"github.com/stretchr/testify/assert"
)

func TestNSOpen(t *testing.T) {
	table := []struct {
		in  string
		out string
	}{
		{"a", "namespace a {"},
		{"a.b", "namespace a { namespace b {"},
		{"a.b.c", "namespace a { namespace b { namespace c {"},
	}
	for _, tt := range table {
		t.Run(tt.in, func(t *testing.T) {
			m := model.NewModule(tt.in, "1.0")
			r, err := nsOpen(reflect.ValueOf(m))
			assert.NoError(t, err)
			assert.Equal(t, tt.out, r.String())
		})
	}
}

func TestNSClose(t *testing.T) {
	table := []struct {
		in  string
		out string
	}{
		{"a", "} // namespace a"},
		{"a.b", "} } // namespace a::b"},
		{"a.b.c", "} } } // namespace a::b::c"},
	}
	for _, tt := range table {
		m := model.NewModule(tt.in, "1.0")
		r, err := nsClose(reflect.ValueOf(m))
		assert.NoError(t, err)
		assert.Equal(t, tt.out, r.String())
	}
}

func TestNS(t *testing.T) {
	table := []struct {
		in  string
		out string
	}{
		{"a", "a"},
		{"a.b", "a::b"},
		{"a.b.c", "a::b::c"},
	}
	for _, tt := range table {
		m := model.NewModule(tt.in, "1.0")
		r, err := ns(reflect.ValueOf(m))
		assert.NoError(t, err)
		assert.Equal(t, tt.out, r.String())
	}
}
