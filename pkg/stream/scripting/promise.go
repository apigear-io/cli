package scripting

import (
	"sync"

	"github.com/dop251/goja"
)

// Promise provides a simple Promise implementation for Goja.
// It supports .then(onFulfilled, onRejected) and .catch(onRejected) chaining.
type Promise struct {
	engine *Engine

	mu        sync.Mutex
	state     string // "pending", "fulfilled", "rejected"
	value     interface{}
	reason    interface{}
	onFulfill []goja.Callable
	onReject  []goja.Callable
}

// NewPromise creates a new Promise.
func NewPromise(engine *Engine) *Promise {
	return &Promise{
		engine: engine,
		state:  "pending",
	}
}

// Resolve fulfills the promise with a value.
func (p *Promise) Resolve(value interface{}) {
	p.mu.Lock()
	if p.state != "pending" {
		p.mu.Unlock()
		return
	}
	p.state = "fulfilled"
	p.value = value
	handlers := make([]goja.Callable, len(p.onFulfill))
	copy(handlers, p.onFulfill)
	p.mu.Unlock()

	// Call handlers asynchronously
	if len(handlers) > 0 {
		p.engine.ScheduleCallback(func(vm *goja.Runtime) {
			for _, handler := range handlers {
				p.engine.CallHandler(handler, "promise.then", vm.ToValue(value))
			}
		})
	}
}

// Reject rejects the promise with a reason.
func (p *Promise) Reject(reason interface{}) {
	p.mu.Lock()
	if p.state != "pending" {
		p.mu.Unlock()
		return
	}
	p.state = "rejected"
	p.reason = reason
	handlers := make([]goja.Callable, len(p.onReject))
	copy(handlers, p.onReject)
	p.mu.Unlock()

	// Call handlers asynchronously
	if len(handlers) > 0 {
		p.engine.ScheduleCallback(func(vm *goja.Runtime) {
			for _, handler := range handlers {
				p.engine.CallHandler(handler, "promise.catch", vm.ToValue(reason))
			}
		})
	}
}

// ToValue converts the Promise to a JavaScript object with .then() and .catch().
func (p *Promise) ToValue(vm *goja.Runtime) goja.Value {
	obj := vm.NewObject()

	// then(onFulfilled, onRejected) - returns new Promise for chaining
	_ = obj.Set("then", func(call goja.FunctionCall) goja.Value {
		var onFulfilled, onRejected goja.Callable

		if len(call.Arguments) >= 1 {
			if fn, ok := goja.AssertFunction(call.Arguments[0]); ok {
				onFulfilled = fn
			}
		}
		if len(call.Arguments) >= 2 {
			if fn, ok := goja.AssertFunction(call.Arguments[1]); ok {
				onRejected = fn
			}
		}

		// Create a new promise for chaining
		nextPromise := NewPromise(p.engine)

		p.mu.Lock()
		state := p.state
		value := p.value
		reason := p.reason

		if state == "pending" {
			// Register handlers for later
			if onFulfilled != nil {
				p.onFulfill = append(p.onFulfill, wrapHandler(vm, onFulfilled, nextPromise))
			} else {
				// Pass through value if no handler
				p.onFulfill = append(p.onFulfill, func(_ goja.Value, args ...goja.Value) (goja.Value, error) {
					if len(args) > 0 {
						nextPromise.Resolve(args[0].Export())
					} else {
						nextPromise.Resolve(nil)
					}
					return goja.Undefined(), nil
				})
			}
			if onRejected != nil {
				p.onReject = append(p.onReject, wrapHandler(vm, onRejected, nextPromise))
			} else {
				// Pass through rejection if no handler
				p.onReject = append(p.onReject, func(_ goja.Value, args ...goja.Value) (goja.Value, error) {
					if len(args) > 0 {
						nextPromise.Reject(args[0].Export())
					} else {
						nextPromise.Reject(nil)
					}
					return goja.Undefined(), nil
				})
			}
			p.mu.Unlock()
		} else {
			p.mu.Unlock()

			// Already settled, call handler immediately
			p.engine.ScheduleCallback(func(vm *goja.Runtime) {
				if state == "fulfilled" {
					if onFulfilled != nil {
						result, err := onFulfilled(goja.Undefined(), vm.ToValue(value))
						if err != nil {
							nextPromise.Reject(err.Error())
						} else {
							nextPromise.Resolve(result.Export())
						}
					} else {
						nextPromise.Resolve(value)
					}
				} else {
					if onRejected != nil {
						result, err := onRejected(goja.Undefined(), vm.ToValue(reason))
						if err != nil {
							nextPromise.Reject(err.Error())
						} else {
							nextPromise.Resolve(result.Export())
						}
					} else {
						nextPromise.Reject(reason)
					}
				}
			})
		}

		return nextPromise.ToValue(vm)
	})

	// catch(onRejected) - shorthand for then(undefined, onRejected)
	_ = obj.Set("catch", func(call goja.FunctionCall) goja.Value {
		var onRejected goja.Callable
		if len(call.Arguments) >= 1 {
			if fn, ok := goja.AssertFunction(call.Arguments[0]); ok {
				onRejected = fn
			}
		}

		nextPromise := NewPromise(p.engine)

		p.mu.Lock()
		state := p.state
		value := p.value
		reason := p.reason

		if state == "pending" {
			// Pass through fulfillment
			p.onFulfill = append(p.onFulfill, func(_ goja.Value, args ...goja.Value) (goja.Value, error) {
				if len(args) > 0 {
					nextPromise.Resolve(args[0].Export())
				} else {
					nextPromise.Resolve(nil)
				}
				return goja.Undefined(), nil
			})
			if onRejected != nil {
				p.onReject = append(p.onReject, wrapHandler(vm, onRejected, nextPromise))
			} else {
				p.onReject = append(p.onReject, func(_ goja.Value, args ...goja.Value) (goja.Value, error) {
					if len(args) > 0 {
						nextPromise.Reject(args[0].Export())
					} else {
						nextPromise.Reject(nil)
					}
					return goja.Undefined(), nil
				})
			}
			p.mu.Unlock()
		} else {
			p.mu.Unlock()

			p.engine.ScheduleCallback(func(vm *goja.Runtime) {
				if state == "fulfilled" {
					nextPromise.Resolve(value)
				} else {
					if onRejected != nil {
						result, err := onRejected(goja.Undefined(), vm.ToValue(reason))
						if err != nil {
							nextPromise.Reject(err.Error())
						} else {
							nextPromise.Resolve(result.Export())
						}
					} else {
						nextPromise.Reject(reason)
					}
				}
			})
		}

		return nextPromise.ToValue(vm)
	})

	return obj
}

// wrapHandler wraps a handler to propagate results to the next promise.
func wrapHandler(vm *goja.Runtime, handler goja.Callable, next *Promise) goja.Callable {
	return func(_ goja.Value, args ...goja.Value) (goja.Value, error) {
		result, err := handler(goja.Undefined(), args...)
		if err != nil {
			next.Reject(err.Error())
		} else {
			next.Resolve(result.Export())
		}
		return goja.Undefined(), nil
	}
}
