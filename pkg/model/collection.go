package model

import "fmt"

type NamedElement interface {
	Name() string
	Kind() Kind
}

type Collection struct {
	items []NamedElement
	keys  map[string]bool
}

func NewCollection() *Collection {
	return &Collection{
		items: make([]NamedElement, 0),
		keys:  make(map[string]bool),
	}
}

func (c *Collection) Items() []NamedElement {
	return c.items
}

func (c *Collection) Filter(t Kind) []NamedElement {
	r := make([]NamedElement, 0)
	for _, i := range c.items {
		if i.Kind() == t {
			r = append(r, i)
		}
	}
	return r
}

func (c *Collection) Add(n NamedElement) error {
	if c.keys[n.Name()] {
		return fmt.Errorf("duplicate name: %s", n.Name())
	}
	c.items = append(c.items, n)
	c.keys[n.Name()] = true
	return nil
}

func (c *Collection) Lookup(name string) NamedElement {
	for _, i := range c.items {
		if i.Name() == name {
			return i
		}
	}
	return nil
}

// NestedCollection is a collection of collections.
type NestedCollection struct {
	collections []*Collection
}

// NewNestedCollection returns a new NestedCollection.
func NewNestedCollection() *NestedCollection {
	return &NestedCollection{
		collections: make([]*Collection, 0),
	}
}

// Items returns all items in all collections.
func (c *NestedCollection) Items() []NamedElement {
	r := make([]NamedElement, 0)
	for _, i := range c.collections {
		r = append(r, i.Items()...)
	}
	return r
}

// Add adds a collection to the nested collection.
func (c *NestedCollection) Add(n *Collection) {
	c.collections = append(c.collections, n)
}

// Filter returns all items with the given kind.
func (c *NestedCollection) Filter(t Kind) []NamedElement {
	r := make([]NamedElement, 0)
	for _, i := range c.collections {
		r = append(r, i.Filter(t)...)
	}
	return r
}

// Lookup returns the first item with the given name.
func (c *NestedCollection) Lookup(name string) NamedElement {
	for _, i := range c.collections {
		if r := i.Lookup(name); r != nil {
			return r
		}
	}
	return nil
}
