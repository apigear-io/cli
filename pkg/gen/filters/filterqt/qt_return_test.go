package filterqt

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
		{"test", "Test1", "propFloat", "qreal"},
		{"test", "Test1", "propString", "QString"},
		{"test", "Test1", "propBoolArray", "QList<bool>"},
		{"test", "Test1", "propIntArray", "QList<int>"},
		{"test", "Test1", "propFloatArray", "QList<qreal>"},
		{"test", "Test1", "propStringArray", "QList<QString>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := qtReturn("", prop)
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
		{"test", "Test2", "propEnum", "Enum1::Enum1Enum"},
		{"test", "Test2", "propStruct", "Struct1"},
		{"test", "Test2", "propInterface", "Interface1*"},
		{"test", "Test2", "propEnumArray", "QList<Enum1::Enum1Enum>"},
		{"test", "Test2", "propStructArray", "QList<Struct1>"},
		{"test", "Test2", "propInterfaceArray", "QList<Interface1*>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := qtReturn("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
