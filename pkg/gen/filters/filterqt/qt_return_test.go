package filterqt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test with all the types
// properties, operation params, operation return, signal params, struct fields
func TestReturn(t *testing.T) {
	t.Parallel()
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
		{"test", "Test1", "propInt32", "qint32"},
		{"test", "Test1", "propInt64", "qint64"},
		{"test", "Test1", "propFloat", "qreal"},
		{"test", "Test1", "propFloat32", "float"},
		{"test", "Test1", "propFloat64", "double"},
		{"test", "Test1", "propString", "QString"},
		{"test", "Test1", "propBoolArray", "QList<bool>"},
		{"test", "Test1", "propIntArray", "QList<int>"},
		{"test", "Test1", "propInt32Array", "QList<qint32>"},
		{"test", "Test1", "propInt64Array", "QList<qint64>"},
		{"test", "Test1", "propFloatArray", "QList<qreal>"},
		{"test", "Test1", "propFloat32Array", "QList<float>"},
		{"test", "Test1", "propFloat64Array", "QList<double>"},
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

func TestOperationReturn(t *testing.T) {
	t.Parallel()
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
		{"test", "Test3", "opInt32", "qint32"},
		{"test", "Test3", "opInt64", "qint64"},
		{"test", "Test3", "opFloat", "qreal"},
		{"test", "Test3", "opFloat32", "float"},
		{"test", "Test3", "opFloat64", "double"},
		{"test", "Test3", "opString", "QString"},
		{"test", "Test3", "opBoolArray", "QList<bool>"},
		{"test", "Test3", "opIntArray", "QList<int>"},
		{"test", "Test3", "opInt32Array", "QList<qint32>"},
		{"test", "Test3", "opInt64Array", "QList<qint64>"},
		{"test", "Test3", "opFloatArray", "QList<qreal>"},
		{"test", "Test3", "opFloat32Array", "QList<float>"},
		{"test", "Test3", "opFloat64Array", "QList<double>"},
		{"test", "Test3", "opStringArray", "QList<QString>"},
	}
	for _, sys := range syss {
		for _, tt := range propTests {
			t.Run(tt.pn, func(t *testing.T) {
				op := sys.LookupOperation(tt.mn, tt.in, tt.pn)
				assert.NotNil(t, op)
				r, err := qtReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.rt, r)
			})
		}
	}
}

func TestReturnSymbols(t *testing.T) {
	t.Parallel()
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

func TestReturnExterns(t *testing.T) {
	t.Parallel()
	table := []struct {
		module_name    string
		interface_name string
		operation_name string
		result         string
	}{
		{"test_apigear_next", "Iface1", "func1", "XType1"},
		{"test_apigear_next", "Iface1", "func3", "demoXA::XType3A"},
		{"test_apigear_next", "Iface1", "funcList", "QList<demoXA::XType3A>"},
		{"test_apigear_next", "Iface1", "funcImportedEnum", "test::Enum1::Enum1Enum"},
		{"test_apigear_next", "Iface1", "funcImportedStruct", "test::Struct1"},
	}
	syss := loadExternSystems(t)
	for _, sys := range syss {
		for _, tt := range table {
			t.Run(tt.operation_name, func(t *testing.T) {
				op := sys.LookupOperation(tt.module_name, tt.interface_name, tt.operation_name)
				assert.NotNil(t, op)
				r, err := qtReturn("", op.Return)
				assert.NoError(t, err)
				assert.Equal(t, tt.result, r)
			})
		}
	}
}
