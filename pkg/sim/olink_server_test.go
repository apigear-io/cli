package sim

import (
	"slices"

	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

type MockEngineServer struct {
	sources []remote.IObjectSource
}

var _ IOlinkServer = (*MockEngineServer)(nil)

func (m *MockEngineServer) RegisterSource(source remote.IObjectSource) {
	m.sources = append(m.sources, source)
}
func (m *MockEngineServer) UnregisterSource(source remote.IObjectSource) {
	// remove the source
	for i, s := range m.sources {
		if s == source {
			m.sources = slices.Delete(m.sources, i, 1)
		}
	}
}
