package evt

import (
	"sync"
	"testing"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent string

func (e TestEvent) String() string {
	return string(e)
}

type TestStruct struct {
	Name string
}

const (
	PublishInt    = "pub.int"
	PublishStr    = "pub.str"
	PublishMap    = "pub.map"
	PublishStruct = "pub.struct"
	RequestInt    = "req.int"
	RequestStr    = "req.str"
	RequestMap    = "req.map"
	RequestStruct = "req.struct"
)

func setupServer(t *testing.T) (*nats.Conn, func()) {
	opts := &server.Options{
		ServerName: "apigear_server",
		DontListen: true,
	}
	server, err := server.NewServer(opts)
	assert.NoError(t, err)
	assert.NotNil(t, server)
	server.Start()
	if !server.ReadyForConnections(20 * time.Second) {
		assert.Fail(t, "nats server not ready")
	}
	nc, err := nats.Connect(server.ClientURL(), nats.InProcessServer(server))
	assert.NoError(t, err)
	assert.NotNil(t, nc)

	teardown := func() {
		nc.Drain()
		server.Shutdown()
	}
	return nc, teardown
}

type EvtTestSuit struct {
	suite.Suite
	nc       *nats.Conn
	teardown func()
	bus      IEventBus
}

func TestEvtTestSuit(t *testing.T) {
	suite.Run(t, new(EvtTestSuit))
}

func (s *EvtTestSuit) SetupSuite() {
	s.nc, s.teardown = setupServer(s.T())
	s.bus = NewNatsEventBus("test", s.nc)
	assert.NotNil(s.T(), s.bus)
}

func (s *EvtTestSuit) TearDownSuite() {
	s.teardown()
}

func (s *EvtTestSuit) SetupTest() {
}

func (s *EvtTestSuit) TearDownTest() {
}

func wrapWG(wg *sync.WaitGroup) chan struct{} {
	out := make(chan struct{})
	go func() {
		wg.Wait()
		out <- struct{}{}
	}()
	return out
}

func (s *EvtTestSuit) TestNatsEventBus_Publish() {
	rows := []struct {
		name string
		e    *Event
	}{
		{"int", NewEvent(PublishInt, 42)},
		{"str", NewEvent(PublishStr, "hello")},
		{"map", NewEvent(PublishMap, map[string]interface{}{"name": "test"})},
		{"struct", NewEvent(PublishStruct, TestStruct{Name: "test"})},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(rows))

	handleInt := func(e *Event) (*Event, error) {
		assert.Equal(s.T(), PublishInt, e.Kind)
		var value int
		err := e.Export(&value)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), 42, value)
		wg.Done()
		return nil, nil
	}

	handleStr := func(e *Event) (*Event, error) {
		assert.Equal(s.T(), PublishStr, e.Kind)
		var value string
		err := e.Export(&value)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "hello", value)
		wg.Done()
		return nil, nil
	}

	handleMap := func(e *Event) (*Event, error) {
		assert.Equal(s.T(), PublishMap, e.Kind)
		var value map[string]any
		err := e.Export(&value)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), map[string]any{"name": "test"}, value)
		wg.Done()
		return nil, nil
	}

	handleStruct := func(e *Event) (*Event, error) {
		assert.Equal(s.T(), PublishStruct, e.Kind)
		var value TestStruct
		err := e.Export(&value)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), TestStruct{Name: "test"}, value)
		wg.Done()
		return nil, nil
	}

	s.bus.Register(PublishInt, HandlerFunc(handleInt))
	s.bus.Register(PublishStr, HandlerFunc(handleStr))
	s.bus.Register(PublishMap, HandlerFunc(handleMap))
	s.bus.Register(PublishStruct, HandlerFunc(handleStruct))

	for _, row := range rows {
		err := s.bus.Publish(row.e)
		assert.NoError(s.T(), err)
	}

	select {
	case <-time.After(1 * time.Second):
		assert.Fail(s.T(), "timeout")
	case <-wrapWG(&wg):
	}
}

func (s *EvtTestSuit) TestNatsEventBus_Request() {
	rows := []struct {
		name string
		e    *Event
	}{
		{"int", NewEvent(RequestInt, 42)},
		{"str", NewEvent(RequestStr, "hello")},
		{"map", NewEvent(RequestMap, map[string]interface{}{"name": "test"})},
		{"struct", NewEvent(RequestStruct, TestStruct{Name: "test"})},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(rows))
	start := time.Now()

	handleInt := func(e *Event) (*Event, error) {
		assert.Equal(s.T(), RequestInt, e.Kind)
		duration := time.Since(start)
		assert.True(s.T(), duration < 1*time.Second)
		log.Info().Msgf("duration: %s", duration)
		var value int
		err := e.Export(&value)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), 42, value)
		wg.Done()
		return NewEvent(RequestInt, value), nil
	}

	handleStr := func(e *Event) (*Event, error) {
		assert.Equal(s.T(), RequestStr, e.Kind)
		duration := time.Since(start)
		assert.True(s.T(), duration < 1*time.Second)
		log.Info().Msgf("duration: %s", duration)
		var value string
		err := e.Export(&value)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "hello", value)
		wg.Done()
		return NewEvent(RequestStr, value), nil
	}

	handleMap := func(e *Event) (*Event, error) {
		assert.Equal(s.T(), RequestMap, e.Kind)
		duration := time.Since(start)
		assert.True(s.T(), duration < 1*time.Second)
		log.Info().Msgf("duration: %s", duration)
		var value map[string]any
		err := e.Export(&value)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), map[string]any{"name": "test"}, value)
		wg.Done()
		return NewEvent(RequestMap, value), nil
	}

	handleStruct := func(e *Event) (*Event, error) {
		assert.Equal(s.T(), RequestStruct, e.Kind)
		duration := time.Since(start)
		assert.True(s.T(), duration < 1*time.Second)
		log.Info().Msgf("duration: %s", duration)
		var value TestStruct
		err := e.Export(&value)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), TestStruct{Name: "test"}, value)
		wg.Done()
		return NewEvent(RequestStruct, value), nil
	}

	s.bus.Register(RequestInt, HandlerFunc(handleInt))
	s.bus.Register(RequestStr, HandlerFunc(handleStr))
	s.bus.Register(RequestMap, HandlerFunc(handleMap))
	s.bus.Register(RequestStruct, HandlerFunc(handleStruct))

	for _, row := range rows {
		s.T().Run(row.name, func(t *testing.T) {
			resp, err := s.bus.Request(row.e)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
		})
	}

	select {
	case <-time.After(1 * time.Second):
		assert.Fail(s.T(), "timeout")
	case <-wrapWG(&wg):
	}
}

func (s *EvtTestSuit) TestNatsEventBus_UnknownHandler() {
	s.T().Run("unknown", func(t *testing.T) {
		unknownEvent := NewEvent("unknown_type", nil)
		resp, err := s.bus.Request(unknownEvent)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "unknown_type", resp.Kind)
		assert.NotEmpty(t, resp.Error)
	})
}

func (s *EvtTestSuit) TestNatsEventBus_UnknownHandlerFunc() {
	s.T().Run("unknown", func(t *testing.T) {
		unknownEvent := NewEvent("unknown_type", nil)
		s.bus.Register("unknown_type", HandlerFunc(func(e *Event) (*Event, error) {
			return nil, nil
		}))
		resp, err := s.bus.Request(unknownEvent)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "unknown_type", resp.Kind)
		assert.NotEmpty(t, resp.Error)
	})
}

func (s *EvtTestSuit) TestNatsEventBus_Middleware() {
	s.T().Run("middleware", func(t *testing.T) {
		// setup middleware
		mw := func(e *Event) (*Event, error) {
			e.Set("middleware", "value")
			return e, nil
		}
		s.bus.Use(mw)

		// setup handler
		s.bus.Register("test", HandlerFunc(func(e *Event) (*Event, error) {
			assert.Equal(s.T(), "value", e.Get("middleware"))
			return nil, nil
		}))

		// publish event
		err := s.bus.Publish(NewEvent("test", nil))
		assert.NoError(s.T(), err)
	})
}
