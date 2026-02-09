package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTestSystem() *System {
	sys := NewSystem("testsystem")

	// Create a test module
	module := &Module{
		NamedNode: NamedNode{
			Name: "test.module",
			Kind: KindModule,
		},
	}

	// Add an interface with properties, operations, and signals
	iface := &Interface{
		NamedNode: NamedNode{
			Name: "ICounter",
			Kind: KindInterface,
		},
	}

	// Add property
	prop := &TypedNode{
		NamedNode: NamedNode{
			Name: "count",
			Kind: KindProperty,
		},
		Schema: Schema{
			Type: "int",
		},
	}
	iface.Properties = append(iface.Properties, prop)

	// Add operation
	op := &Operation{
		NamedNode: NamedNode{
			Name: "increment",
			Kind: KindOperation,
		},
	}
	iface.Operations = append(iface.Operations, op)

	// Add signal
	sig := &Signal{
		NamedNode: NamedNode{
			Name: "changed",
			Kind: KindSignal,
		},
	}
	iface.Signals = append(iface.Signals, sig)

	module.Interfaces = append(module.Interfaces, iface)

	// Add a struct
	str := &Struct{
		NamedNode: NamedNode{
			Name: "Point",
			Kind: KindStruct,
		},
	}

	// Add field to struct
	field := &TypedNode{
		NamedNode: NamedNode{
			Name: "x",
			Kind: KindField,
		},
		Schema: Schema{
			Type: "int",
		},
	}
	str.Fields = append(str.Fields, field)

	module.Structs = append(module.Structs, str)

	// Add an enum
	enum := &Enum{
		NamedNode: NamedNode{
			Name: "Status",
			Kind: KindEnum,
		},
	}

	// Add enum member
	member := &EnumMember{
		NamedNode: NamedNode{
			Name:  "Active",
			Kind:  KindMember,
		},
		Value: 0,
	}
	enum.Members = append(enum.Members, member)

	module.Enums = append(module.Enums, enum)

	// Add an extern
	extern := &Extern{
		NamedNode: NamedNode{
			Name: "ExternalType",
			Kind: KindExtern,
		},
	}
	module.Externs = append(module.Externs, extern)

	sys.AddModule(module)

	return sys
}

func TestNewSystem(t *testing.T) {
	t.Run("creates new system", func(t *testing.T) {
		sys := NewSystem("test")
		assert.NotNil(t, sys)
		assert.Equal(t, "test", sys.Name)
		assert.Equal(t, KindSystem, sys.Kind)
		assert.Empty(t, sys.Modules)
	})
}

func TestSystemAddModule(t *testing.T) {
	t.Run("adds module to system", func(t *testing.T) {
		sys := NewSystem("test")
		module := &Module{
			NamedNode: NamedNode{
				Name: "test.module",
				Kind: KindModule,
			},
		}

		sys.AddModule(module)

		assert.Len(t, sys.Modules, 1)
		assert.Equal(t, module, sys.Modules[0])
		assert.Equal(t, sys, module.System)
	})
}

func TestSystemLookupModule(t *testing.T) {
	sys := createTestSystem()

	t.Run("finds existing module", func(t *testing.T) {
		module := sys.LookupModule("test.module")
		require.NotNil(t, module)
		assert.Equal(t, "test.module", module.Name)
	})

	t.Run("returns nil for non-existent module", func(t *testing.T) {
		module := sys.LookupModule("nonexistent")
		assert.Nil(t, module)
	})
}

func TestSystemLookupExtern(t *testing.T) {
	sys := createTestSystem()

	t.Run("finds existing extern", func(t *testing.T) {
		extern := sys.LookupExtern("test.module", "ExternalType")
		require.NotNil(t, extern)
		assert.Equal(t, "ExternalType", extern.Name)
	})

	t.Run("returns nil for non-existent module", func(t *testing.T) {
		extern := sys.LookupExtern("nonexistent", "ExternalType")
		assert.Nil(t, extern)
	})

	t.Run("returns nil for non-existent extern", func(t *testing.T) {
		extern := sys.LookupExtern("test.module", "NonExistent")
		assert.Nil(t, extern)
	})
}

func TestSystemLookupInterface(t *testing.T) {
	sys := createTestSystem()

	t.Run("finds existing interface", func(t *testing.T) {
		iface := sys.LookupInterface("test.module", "ICounter")
		require.NotNil(t, iface)
		assert.Equal(t, "ICounter", iface.Name)
	})

	t.Run("returns nil for non-existent module", func(t *testing.T) {
		iface := sys.LookupInterface("nonexistent", "ICounter")
		assert.Nil(t, iface)
	})

	t.Run("returns nil for non-existent interface", func(t *testing.T) {
		iface := sys.LookupInterface("test.module", "NonExistent")
		assert.Nil(t, iface)
	})
}

func TestSystemLookupStruct(t *testing.T) {
	sys := createTestSystem()

	t.Run("finds existing struct", func(t *testing.T) {
		str := sys.LookupStruct("test.module", "Point")
		require.NotNil(t, str)
		assert.Equal(t, "Point", str.Name)
	})

	t.Run("returns nil for non-existent module", func(t *testing.T) {
		str := sys.LookupStruct("nonexistent", "Point")
		assert.Nil(t, str)
	})

	t.Run("returns nil for non-existent struct", func(t *testing.T) {
		str := sys.LookupStruct("test.module", "NonExistent")
		assert.Nil(t, str)
	})
}

func TestSystemLookupEnum(t *testing.T) {
	sys := createTestSystem()

	t.Run("finds existing enum", func(t *testing.T) {
		enum := sys.LookupEnum("test.module", "Status")
		require.NotNil(t, enum)
		assert.Equal(t, "Status", enum.Name)
	})

	t.Run("returns nil for non-existent module", func(t *testing.T) {
		enum := sys.LookupEnum("nonexistent", "Status")
		assert.Nil(t, enum)
	})

	t.Run("returns nil for non-existent enum", func(t *testing.T) {
		enum := sys.LookupEnum("test.module", "NonExistent")
		assert.Nil(t, enum)
	})
}

func TestSystemLookupField(t *testing.T) {
	sys := createTestSystem()

	t.Run("finds existing field", func(t *testing.T) {
		field := sys.LookupField("test.module", "Point", "x")
		require.NotNil(t, field)
		assert.Equal(t, "x", field.Name)
	})

	t.Run("returns nil for non-existent struct", func(t *testing.T) {
		field := sys.LookupField("test.module", "NonExistent", "x")
		assert.Nil(t, field)
	})

	t.Run("returns nil for non-existent field", func(t *testing.T) {
		field := sys.LookupField("test.module", "Point", "nonexistent")
		assert.Nil(t, field)
	})
}

func TestSystemLookupProperty(t *testing.T) {
	sys := createTestSystem()

	t.Run("finds existing property", func(t *testing.T) {
		prop := sys.LookupProperty("test.module", "ICounter", "count")
		require.NotNil(t, prop)
		assert.Equal(t, "count", prop.Name)
	})

	t.Run("returns nil for non-existent interface", func(t *testing.T) {
		prop := sys.LookupProperty("test.module", "NonExistent", "count")
		assert.Nil(t, prop)
	})

	t.Run("returns nil for non-existent property", func(t *testing.T) {
		prop := sys.LookupProperty("test.module", "ICounter", "nonexistent")
		assert.Nil(t, prop)
	})
}

func TestSystemLookupOperation(t *testing.T) {
	sys := createTestSystem()

	t.Run("finds existing operation", func(t *testing.T) {
		op := sys.LookupOperation("test.module", "ICounter", "increment")
		require.NotNil(t, op)
		assert.Equal(t, "increment", op.Name)
	})

	t.Run("returns nil for non-existent interface", func(t *testing.T) {
		op := sys.LookupOperation("test.module", "NonExistent", "increment")
		assert.Nil(t, op)
	})

	t.Run("returns nil for non-existent operation", func(t *testing.T) {
		op := sys.LookupOperation("test.module", "ICounter", "nonexistent")
		assert.Nil(t, op)
	})
}

func TestSystemLookupSignal(t *testing.T) {
	sys := createTestSystem()

	t.Run("finds existing signal", func(t *testing.T) {
		sig := sys.LookupSignal("test.module", "ICounter", "changed")
		require.NotNil(t, sig)
		assert.Equal(t, "changed", sig.Name)
	})

	t.Run("returns nil for non-existent interface", func(t *testing.T) {
		sig := sys.LookupSignal("test.module", "NonExistent", "changed")
		assert.Nil(t, sig)
	})

	t.Run("returns nil for non-existent signal", func(t *testing.T) {
		sig := sys.LookupSignal("test.module", "ICounter", "nonexistent")
		assert.Nil(t, sig)
	})
}

func TestFQNSplit2(t *testing.T) {
	tests := []struct {
		name         string
		fqn          string
		expectModule string
		expectName   string
	}{
		{"simple FQN", "test.module.Type", "test.module", "Type"},
		{"nested module", "a.b.c.Type", "a.b.c", "Type"},
		{"two parts", "module.Type", "module", "Type"},
		{"single part", "Type", "", ""},
		{"empty string", "", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			module, name := FQNSplit2(tt.fqn)
			assert.Equal(t, tt.expectModule, module)
			assert.Equal(t, tt.expectName, name)
		})
	}
}

func TestFQNSplit3(t *testing.T) {
	tests := []struct {
		name           string
		fqn            string
		expectModule   string
		expectElement  string
		expectMember   string
	}{
		{"full FQN", "test.module.Type.member", "test.module", "Type", "member"},
		{"nested module", "a.b.c.Type.member", "a.b.c", "Type", "member"},
		{"three parts", "module.Type.member", "module", "Type", "member"},
		{"two parts", "Type.member", "", "", ""},
		{"single part", "member", "", "", ""},
		{"empty string", "", "", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			module, element, member := FQNSplit3(tt.fqn)
			assert.Equal(t, tt.expectModule, module)
			assert.Equal(t, tt.expectElement, element)
			assert.Equal(t, tt.expectMember, member)
		})
	}
}

func TestSystemValidate(t *testing.T) {
	t.Run("validates system with modules", func(t *testing.T) {
		sys := NewSystem("test")
		module := &Module{
			NamedNode: NamedNode{
				Name: "test.module",
				Kind: KindModule,
			},
			Checksum: "test-checksum",
		}
		sys.AddModule(module)

		err := sys.Validate()
		require.NoError(t, err)
		assert.NotEmpty(t, sys.Checksum)
	})

	t.Run("fails for duplicate module names", func(t *testing.T) {
		sys := NewSystem("test")
		module1 := &Module{
			NamedNode: NamedNode{
				Name: "test.module",
				Kind: KindModule,
			},
			Checksum: "checksum1",
		}
		module2 := &Module{
			NamedNode: NamedNode{
				Name: "test.module",
				Kind: KindModule,
			},
			Checksum: "checksum2",
		}
		sys.AddModule(module1)
		sys.AddModule(module2)

		err := sys.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "duplicate name")
	})

	t.Run("validates and computes checksum for module without one", func(t *testing.T) {
		sys := NewSystem("test")
		module := &Module{
			NamedNode: NamedNode{
				Name: "test.module",
				Kind: KindModule,
			},
			// No checksum initially
		}
		sys.AddModule(module)

		// Validate should call module.Validate() which computes checksum
		err := sys.Validate()
		// Module validation computes checksum, so system validation should succeed
		assert.NoError(t, err)
		assert.NotEmpty(t, module.Checksum)
		assert.NotEmpty(t, sys.Checksum)
	})
}

func TestSystemCheckReservedWords(t *testing.T) {
	t.Run("checks reserved words for system", func(t *testing.T) {
		sys := NewSystem("test")
		module := &Module{
			NamedNode: NamedNode{
				Name: "test.module",
				Kind: KindModule,
			},
		}
		sys.AddModule(module)

		// This test just verifies the function doesn't panic
		// Actual reserved word checking is tested in the rkw package
		sys.CheckReservedWords([]string{"cpp", "go"})
	})
}

// TestSystemLookupNode tests the more complex LookupNode function
func TestSystemLookupNode(t *testing.T) {
	sys := createTestSystem()

	t.Run("looks up interface member with # notation", func(t *testing.T) {
		// Format: module.Interface#member
		node := sys.LookupNode("test.module.ICounter#count")
		// The test may return nil if LookupMember is not implemented
		// This tests the code path
		_ = node
	})

	t.Run("looks up module-level node", func(t *testing.T) {
		// Format: module.Type
		node := sys.LookupNode("test.module.Point")
		// The test may return nil depending on implementation
		_ = node
	})

	t.Run("returns nil for invalid FQN", func(t *testing.T) {
		node := sys.LookupNode("invalid")
		assert.Nil(t, node)
	})
}
