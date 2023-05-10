package model

// SystemScope is used by the generator to generate code for a system
type SystemScope struct {
	// System is the root of all modules
	System *System
	// Features is the list of features that are enabled
	Features map[string]bool
	// Meta is the map of metadata
	Meta map[string]any
}

// ModuleScope is used by the generator to generate code for a module
type ModuleScope struct {
	// System is the root of all modules
	System *System
	// Module is the module that contains the interfaces, structs, and enums
	Module *Module
	// Features is the list of features that are enabled
	Features map[string]bool
	// Meta is the map of metadata
	Meta map[string]any
}

// InterfaceScope is used by the generator to generate code for an interface
type InterfaceScope struct {
	// System is the root of all modules
	System *System
	// Module is the module that contains the interfaces, structs, and enums
	Module *Module
	// Interface is the interface that contains the properties, operations and signals
	Interface *Interface
	// Features is the list of features that are enabled
	Features map[string]bool
	// Meta is the map of metadata
	Meta map[string]any
}

// StructScope is used by the generator to generate code for a struct
type StructScope struct {
	// System is the root of all modules
	System *System
	// Module is the module that contains the interfaces, structs, and enums
	Module *Module
	// Struct is the struct that contains the fields
	Struct *Struct
	// Features is the list of features that are enabled
	Features map[string]bool
	// Meta is the map of metadata
	Meta map[string]any
}

// EnumScope is used by the generator to generate code for an enum
type EnumScope struct {
	// System is the root of all modules
	System *System
	// Module is the module that contains the interfaces, structs, and enums
	Module *Module
	// Enum is the enum that contains the values
	Enum *Enum
	// Features is the list of features that are enabled
	Features map[string]bool
	// Meta is the map of metadata
	Meta map[string]any
}
