package ostore

import "errors"

var (
	// ErrObjectNotFound is returned when an object is not found
	ErrObjectNotFound = errors.New("object not found")
	// ErrObjectExists is returned when an object already exists
	ErrObjectExists = errors.New("object already exists")
	// ErrPropertyNotFound is returned when a property is not found
	ErrPropertyNotFound = errors.New("property not found")
)

type IObjectStore interface {
	// Set properties by key
	Set(key string, value map[string]any)
	// Update properties by key
	Update(key string, value map[string]any)
	// Get properties by key
	Get(key string) map[string]any
	// Delete properties by key
	Delete(key string)
	// Has returns true if the store has the key
	Has(key string) bool
	// Keys returns a list of keys
	Keys() []string
	// OnEvent for changes to the store
	OnEvent(fn StoreNotifyFunc) StoreUnWatch
}
