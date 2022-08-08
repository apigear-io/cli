package filtergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParam(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propBool", "propBool bool"},
		{"test", "Test1", "propInt", "propInt int"},
		{"test", "Test1", "propFloat", "propFloat float64"},
		{"test", "Test1", "propString", "propString string"},
		{"test", "Test1", "propBoolArray", "propBoolArray []bool"},
		{"test", "Test1", "propIntArray", "propIntArray []int"},
		{"test", "Test1", "propFloatArray", "propFloatArray []float64"},
		{"test", "Test1", "propStringArray", "propStringArray []string"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := goParam(prop, "")
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestParamSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "propEnum Enum1"},
		{"test", "Test2", "propStruct", "propStruct Struct1"},
		{"test", "Test2", "propInterface", "propInterface *Interface1"},
		{"test", "Test2", "propEnumArray", "propEnumArray []Enum1"},
		{"test", "Test2", "propStructArray", "propStructArray []Struct1"},
		{"test", "Test2", "propInterfaceArray", "propInterfaceArray []*Interface1"},
	}
	sys := loadSystem(t)
	for _, tt := range table {
		t.Run(tt.pn, func(t *testing.T) {
			prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
			assert.NotNil(t, prop)
			r, err := goParam(prop, "")
			assert.NoError(t, err)
			assert.Equal(t, tt.rt, r)
		})
	}
}

func TestParamWithErrors(t *testing.T) {
	s, err := goParam(nil, "")
	assert.Error(t, err)
	assert.Equal(t, "", s)
}
