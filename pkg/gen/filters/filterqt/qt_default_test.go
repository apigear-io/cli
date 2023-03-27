package filterqt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operations params, operation return, signal params, struct fields
func TestDefaultFromIdl(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propVoid", "void"},
		{"test", "Test1", "propBool", "false"},
		{"test", "Test1", "propInt", "0"},
		{"test", "Test1", "propInt32", "0"},
		{"test", "Test1", "propInt64", "0LL"},
		{"test", "Test1", "propFloat", "0.0f"},
		{"test", "Test1", "propFloat32", "0.0f"},
		{"test", "Test1", "propFloat64", "0.0"},
		{"test", "Test1", "propString", "QString()"},
		{"test", "Test1", "propBoolArray", "QList<bool>()"},
		{"test", "Test1", "propIntArray", "QList<int>()"},
		{"test", "Test1", "propInt32Array", "QList<qint32>()"},
		{"test", "Test1", "propInt64Array", "QList<qint64>()"},
		{"test", "Test1", "propFloatArray", "QList<qreal>()"},
		{"test", "Test1", "propFloat32Array", "QList<float>()"},
		{"test", "Test1", "propFloat64Array", "QList<double>()"},
		{"test", "Test1", "propStringArray", "QList<QString>()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := qtDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestDefaultSymbolsFromIdl(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "Enum1::Default"},
		{"test", "Test2", "propStruct", "Struct1()"},
		{"test", "Test2", "propInterface", "nullptr"},
		{"test", "Test2", "propEnumArray", "QList<Enum1::Enum1Enum>()"},
		{"test", "Test2", "propStructArray", "QList<Struct1>()"},
		{"test", "Test2", "propInterfaceArray", "QList<Interface1*>()"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := qtDefault("", prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestDefaultFromIdlWithPrefix_makesNoDifference(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test1", "propVoid", "void"},
		{"test", "Test1", "propBool", "false"},
		{"test", "Test1", "propInt", "0"},
		{"test", "Test1", "propInt32", "0"},
		{"test", "Test1", "propInt64", "0LL"},
		{"test", "Test1", "propFloat", "0.0f"},
		{"test", "Test1", "propFloat32", "0.0f"},
		{"test", "Test1", "propFloat64", "0.0"},
		{"test", "Test1", "propString", "QString()"},
		{"test", "Test1", "propBoolArray", "QList<bool>()"},
		{"test", "Test1", "propIntArray", "QList<int>()"},
		{"test", "Test1", "propInt32Array", "QList<qint32>()"},
		{"test", "Test1", "propInt64Array", "QList<qint64>()"},
		{"test", "Test1", "propFloatArray", "QList<qreal>()"},
		{"test", "Test1", "propFloat32Array", "QList<float>()"},
		{"test", "Test1", "propFloat64Array", "QList<double>()"},
		{"test", "Test1", "propStringArray", "QList<QString>()"},
	}
	prefix := "my_prefix::"
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := qtDefault(prefix, prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestDefaultSymbolsFromIdlWithPrefix(t *testing.T) {
	syss := loadTestSystems(t)
	var propTests = []struct {
		mn string
		in string
		pn string
		rt string
	}{
		{"test", "Test2", "propEnum", "my_prefix::Enum1::Default"},
		{"test", "Test2", "propStruct", "my_prefix::Struct1()"},
		{"test", "Test2", "propInterface", "nullptr"},
		{"test", "Test2", "propEnumArray", "QList<my_prefix::Enum1::Enum1Enum>()"},
		{"test", "Test2", "propStructArray", "QList<my_prefix::Struct1>()"},
		{"test", "Test2", "propInterfaceArray", "QList<my_prefix::Interface1*>()"},
	}
	prefix := "my_prefix::"
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				prop := sys.LookupProperty(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, prop)
				r, err := qtDefault(prefix, prop)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}
