package model_test

import (
	"testing"

	"github.com/apigear-io/cli/pkg/idl"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/stretchr/testify/assert"
)

const IDL = `module TestModule 1.0

enum TestEnum {
	Value1 = 0
	Value2 = 1
}

struct TestStruct {
	Field1: string
	Field2: int
}

interface TestInterface {
	property1: string
	property2: int
	method1(param1: string): void
	method2(param1: int, param2: TestEnum): TestStruct
	signal TestSignal(param1: string, param2: int): void
}
`

type MockMessage struct {
	Name string
	Kind string
}
type MockVisitor struct {
	visited []model.NamedNode
}

func (v *MockVisitor) VisitTypedNode(node *model.TypedNode) error {
	v.visited = append(v.visited, node.NamedNode)
	return nil
}

func (v *MockVisitor) VisitSignal(node *model.Signal) error {
	v.visited = append(v.visited, node.NamedNode)
	return nil
}

func (v *MockVisitor) VisitOperation(node *model.Operation) error {
	v.visited = append(v.visited, node.NamedNode)
	return nil
}

func (v *MockVisitor) VisitSystem(s *model.System) error {
	v.visited = append(v.visited, s.NamedNode)
	return nil
}

func (v *MockVisitor) VisitModule(m *model.Module) error {
	v.visited = append(v.visited, m.NamedNode)
	return nil
}

func (v *MockVisitor) VisitExtern(e *model.Extern) error {
	v.visited = append(v.visited, e.NamedNode)
	return nil
}

func (v *MockVisitor) VisitInterface(i *model.Interface) error {
	v.visited = append(v.visited, i.NamedNode)
	return nil
}

func (v *MockVisitor) VisitStruct(s *model.Struct) error {
	v.visited = append(v.visited, s.NamedNode)
	return nil
}

func (v *MockVisitor) VisitEnum(e *model.Enum) error {
	v.visited = append(v.visited, e.NamedNode)
	return nil
}

func (v *MockVisitor) VisitEnumMember(m *model.EnumMember) error {
	v.visited = append(v.visited, m.NamedNode)
	return nil
}

func (v *MockVisitor) VisitParameter(p *model.TypedNode) error {
	v.visited = append(v.visited, p.NamedNode)
	return nil
}

func TestVisitor(t *testing.T) {
	// Create a mock visitor
	system := model.NewSystem("TestSystem")
	p := idl.NewParser(system)
	err := p.ParseString(IDL)
	assert.NoError(t, err)
	module := system.LookupModule("TestModule")
	assert.NotNil(t, module)

	// Create a mock visitor
	mockVisitor := &MockVisitor{}
	system.AcceptModelVisitor(mockVisitor)
	if err != nil {
		t.Errorf("AcceptModelVisitor returned an error: %v", err)
	}
	// Check if all nodes were visited
	assert.NotEmpty(t, mockVisitor.visited)
	// Check if specific nodes were visited
	visitedNames := make(map[string]bool)
	for _, node := range mockVisitor.visited {
		visitedNames[node.Name] = true
	}
	assert.True(t, visitedNames["TestSystem"])
	assert.True(t, visitedNames["TestModule"])
	assert.True(t, visitedNames["TestEnum"])
	assert.True(t, visitedNames["TestStruct"])
	assert.True(t, visitedNames["TestInterface"])
	assert.True(t, visitedNames["TestSignal"])
	assert.True(t, visitedNames["Field1"])
	assert.True(t, visitedNames["Field2"])
	assert.True(t, visitedNames["property1"])
	assert.True(t, visitedNames["property2"])
	assert.True(t, visitedNames["method1"])
	assert.True(t, visitedNames["method2"])
	assert.True(t, visitedNames["Value1"])
	assert.True(t, visitedNames["Value2"])
	assert.True(t, visitedNames["param1"])
	assert.True(t, visitedNames["param2"])

}
