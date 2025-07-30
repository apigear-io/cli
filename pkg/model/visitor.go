package model

type ModelVisitor interface {
	VisitSystem(s *System) error
	VisitModule(m *Module) error
	VisitExtern(e *Extern) error
	VisitInterface(i *Interface) error
	VisitOperation(o *Operation) error
	VisitParameter(p *TypedNode) error
	VisitSignal(s *Signal) error
	VisitStruct(s *Struct) error
	VisitEnum(e *Enum) error
	VisitEnumMember(v *EnumMember) error
	VisitTypedNode(p *TypedNode) error
}

type AcceptModelVisitor func(v ModelVisitor) error
