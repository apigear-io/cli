package filtercpp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestReturn(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propVoid", "void"},
		{"test", "Test1", "propBool", "bool"},
		{"test", "Test1", "propInt", "int"},
		{"test", "Test1", "propFloat", "float"},
		{"test", "Test1", "propString", "std::string"},
		{"test", "Test1", "propBoolArray", "std::list<bool>"},
		{"test", "Test1", "propIntArray", "std::list<int>"},
		{"test", "Test1", "propFloatArray", "std::list<float>"},
		{"test", "Test1", "propStringArray", "std::list<std::string>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := cppReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestOperationReturn(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test3", "opVoid", "void"},
		{"test", "Test3", "opBool", "bool"},
		{"test", "Test3", "opInt", "int"},
		{"test", "Test3", "opFloat", "float"},
		{"test", "Test3", "opString", "std::string"},
		{"test", "Test3", "opBoolArray", "std::list<bool>"},
		{"test", "Test3", "opIntArray", "std::list<int>"},
		{"test", "Test3", "opFloatArray", "std::list<float>"},
		{"test", "Test3", "opStringArray", "std::list<std::string>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := cppReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestReturnSymbols(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1Enum"},
		{"test", "Test2", "propStruct", "Struct1"},
		{"test", "Test2", "propInterface", "Interface1*"},
		{"test", "Test2", "propEnumArray", "std::list<Enum1Enum>"},
		{"test", "Test2", "propStructArray", "std::list<Struct1>"},
		{"test", "Test2", "propInterfaceArray", "std::list<Interface1*>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := cppReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
