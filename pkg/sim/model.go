package sim

// Scenario is a scenario of a simulation.
type ScenarioDoc struct {
	Name        string         `json:"name" yaml:"name"`
	Version     string         `json:"version" yaml:"version"`
	Description string         `json:"description" yaml:"description"`
	Services    []ServiceEntry `json:"services" yaml:"services"`
}

func (s ScenarioDoc) String() string {
	return s.Name
}

// LookupService returns the service entry for the given service name.
func (s ScenarioDoc) LookupService(name string) *ServiceEntry {
	for _, service := range s.Services {
		if service.Name == name {
			return &service
		}
	}
	return nil
}

// ServiceEntry is a single interface entry in a scenario.
type ServiceEntry struct {
	Name        string          `json:"name" yaml:"name"`
	Description string          `json:"description" yaml:"description"`
	Properties  []PropertyEntry `json:"properties" yaml:"properties"`
	Methods     []MethodEntry   `json:"methods" yaml:"methods"`
}

func (s ServiceEntry) String() string {
	return s.Name
}

// LookupMethod returns the method entry for the given method name.
func (s ServiceEntry) LookupMethod(name string) *MethodEntry {
	for _, method := range s.Methods {
		if method.Name == name {
			return &method
		}
	}
	return nil
}

// PropertyEntry is a named value of an interface method or return entry.
// { "name": "count", "value": 0 }
type PropertyEntry struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Value       any    `json:"value" yaml:"value"`
}

type MethodEntry struct {
	Name        string        `json:"name" yaml:"name"`
	Description string        `json:"description" yaml:"description"`
	Actions     []ActionEntry `json:"actions" yaml:"actions"`
}

type ActionParams []any

type ActionEntry map[string]ActionParams

// Params returns the parameters of the action entry.
func (a ActionEntry) Params(name string) []any {
	return a[name]
}

// Param returns the parameter at index
func (a ActionEntry) Param(index int) any {


