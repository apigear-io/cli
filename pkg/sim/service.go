package sim

import (
	"fmt"
	"reflect"

	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/dop251/goja"
)

type ObjectService struct {
	objectId        string
	properties      map[string]any
	propertyEmitter *Emitter[any]
	methods         map[string]goja.Callable
	signalEmitter   *Emitter[[]any]
	engine          *Engine
	source          *OLinkSource
	proxy           *goja.Object // Reference to the proxy object
}

func NewObjectService(engine *Engine, objectId string, properties map[string]any) *ObjectService {
	if properties == nil {
		properties = make(map[string]any)
	}

	s := &ObjectService{
		objectId:        objectId,
		properties:      properties,
		propertyEmitter: NewEmitter[any](),
		methods:         make(map[string]goja.Callable),
		signalEmitter:   NewEmitter[[]any](),
		engine:          engine,
		proxy:           nil, // Will be set after creation
	}
	s.source = NewOLinkSource(s)
	s.engine.registerSource(s.source)
	
	// Create the proxy for this service
	s.proxy = CreateServiceProxy(engine.rt, s)
	
	return s
}

// GetProxy returns the proxy object for this service
func (s *ObjectService) GetProxy() *goja.Object {
	return s.proxy
}

func (s *ObjectService) Close() {
	s.engine.unregisterSource(s.source)
	if s.source == nil {
		log.Warn().Msgf("ObjectService.Close: source is nil")
		return
	}
	s.source.Close()
}

func (s *ObjectService) ObjectId() string {
	return s.objectId
}

func (o *ObjectService) GetProperty(name string) any {
	return o.properties[name]
}

func (o *ObjectService) SetProperty(name string, value any) {
	o.setProperty(name, value)
}

func (o *ObjectService) setProperty(name string, value any) {
	log.Debug().Str("name", name).Interface("value", value).Msg("ObjectService.SetProperty")
	equals := reflect.DeepEqual(o.properties[name], value)
	if !equals {
		o.properties[name] = value
		o.propertyEmitter.Emit(name, value)
		if o.source == nil {
			log.Warn().Msgf("ObjectService.SetProperty: source is nil")
			return
		}
		o.source.NotifyPropertyChanged(name, value)
	}
}

func (o *ObjectService) OnProperty(name string, fn func(value any)) {
	o.propertyEmitter.Add(name, fn)
}

func (o *ObjectService) GetProperties() map[string]any {
	return o.properties
}

func (o *ObjectService) SetProperties(properties map[string]any) {
	for name, value := range properties {
		o.setProperty(name, value)
	}
}

// HasProperty
func (o *ObjectService) HasProperty(name string) bool {
	_, ok := o.properties[name]
	return ok
}

func (o *ObjectService) OnMethod(method string, v goja.Value) {
	fn, ok := goja.AssertFunction(v)
	if !ok {
		log.Warn().Msgf("ObjectService.OnMethod: value is not a function: %v", v)
		return
	}
	o.methods[method] = fn
}

func (o *ObjectService) CallMethod(method string, args ...any) (goja.Value, error) {
	log.Info().Str("method", method).Interface("args", args).Msg("ObjectService.CallMethod")
	fn, ok := o.methods[method]
	if !ok {
		log.Warn().Msgf("Method %s not found", method)
		return nil, fmt.Errorf("method %s not found", method)
	}
	jsArgs := make([]goja.Value, len(args))
	for i, arg := range args {
		jsArgs[i] = o.engine.rt.ToValue(arg)
	}
	// Use the proxy as 'this' context if available, otherwise use undefined
	thisContext := goja.Undefined()
	if o.proxy != nil {
		thisContext = o.engine.rt.ToValue(o.proxy)
	}
	return fn(thisContext, jsArgs...)
}

// GetMethod return method
func (o *ObjectService) GetMethod(method string) goja.Callable {
	return o.methods[method]
}

// HasMethod
func (o *ObjectService) HasMethod(method string) bool {
	_, ok := o.methods[method]
	return ok
}

// RemoveMethod removes a method from the service
func (o *ObjectService) RemoveMethod(method string) {
	delete(o.methods, method)
}

// RemoveProperty removes a property from the service
func (o *ObjectService) RemoveProperty(name string) {
	delete(o.properties, name)
}

func (o *ObjectService) EmitSignal(signal string, args ...any) {
	// Emit locally to JavaScript listeners
	o.signalEmitter.Emit(signal, args)
	
	// Also notify OLink clients if source is available
	if o.source != nil {
		o.source.NotifySignal(signal, core.Args(args))
	}
}

func (o *ObjectService) OnSignal(signal string, fn func(args ...any)) {
	o.signalEmitter.Add(signal, func(args []any) {
		fn(args...)
	})
}
