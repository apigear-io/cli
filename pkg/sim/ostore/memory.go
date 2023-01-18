package ostore

// MemoryStore implements IObjectStore
var _ IObjectStore = &MemoryStore{}

// MemoryStore is an in-memory implementation of IObjectStore
type MemoryStore struct {
	StoreWatcher
	objects map[string]map[string]any
}

// NewMemoryStore creates a new MemoryStore
func NewMemoryStore() IObjectStore {
	return &MemoryStore{
		objects: make(map[string]map[string]any),
	}
}

// Create an object in the store by id and properties
func (m *MemoryStore) Set(key string, value map[string]any) {
	m.objects[key] = value
	// notify watchers
	m.notify(StoreEvent{
		Type:  EventTypeCreate,
		Id:    key,
		Value: m.objects[key],
	})
}

// Update an object in the store by id and partial properties
func (m *MemoryStore) Update(id string, value map[string]any) {
	if !m.Has(id) {
		m.Set(id, value)
		return
	}
	for k, v := range value {
		m.objects[id][k] = v
	}
	// notify watchers
	m.notify(StoreEvent{
		Type:  EventTypeUpdate,
		Id:    id,
		Value: m.objects[id],
	})
}

// Delete an object in the store by id
func (m *MemoryStore) Delete(key string) {
	delete(m.objects, key)
	// notify watchers
	m.notify(StoreEvent{
		Type:  EventTypeDelete,
		Id:    key,
		Value: nil,
	})
}

// Get an object from the store by id
func (m *MemoryStore) Get(key string) map[string]any {
	// check if object exists
	if _, ok := m.objects[key]; !ok {
		return map[string]any{}
	}
	return m.objects[key]
}

// Has returns true if the object exists in the store
func (m *MemoryStore) Has(key string) bool {
	_, ok := m.objects[key]
	return ok
}

// Keys returns a list of keys
func (m *MemoryStore) Keys() []string {
	var keys []string
	for key := range m.objects {
		keys = append(keys, key)
	}
	return keys
}
