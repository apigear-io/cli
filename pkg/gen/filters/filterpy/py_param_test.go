package filterpy

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
		{"test", "Test1", "propBool", "propBool: bool"},
		{"test", "Test1", "propInt", "propInt: int"},
		{"test", "Test1", "propFloat", "propFloat: float"},
		{"test", "Test1", "propString", "propString: str"},
		{"test", "Test1", "propBoolArray", "propBoolArray: list[bool]"},
		{"test", "Test1", "propIntArray", "propIntArray: list[int]"},
		{"test", "Test1", "propFloatArray", "propFloatArray: list[float]"},
		{"test", "Test1", "propStringArray", "propStringArray: list[str]"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := pyParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestParamSymbols(t *testing.T) {
	table := []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "propEnum: Enum1"},
		{"test", "Test2", "propStruct", "propStruct: Struct1"},
		{"test", "Test2", "propInterface", "propInterface: Interface1"},
		{"test", "Test2", "propEnumArray", "propEnumArray: list[Enum1]"},
		{"test", "Test2", "propStructArray", "propStructArray: list[Struct1]"},
		{"test", "Test2", "propInterfaceArray", "propInterfaceArray: list[Interface1]"},
	}
	syss := loadTestSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := pyParam("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
