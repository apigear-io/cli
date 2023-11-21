// Package ostore provides a simple object store for the simulation.
// A object-store is a key-value store for objects. The objects are identified by a unique key.
// The objects are maps of properties (kwargs) in form of `map[string]any`.
// The object store is used to store the state of the simulation for each object.
// A watcher can be registered to get notified on changes to the store.
package ostore
