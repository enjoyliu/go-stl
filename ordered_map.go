package ordered_map

import (
	. "golang.org/x/exp/constraints"
	"ordered_map/rbtree"
	"sort"
)

// OrderedMap is a map that maintains the order of insertion.
type OrderedMap interface {

	// Insert adds an element to the map. If the element already exists, it is not added again.
	Insert(key, value interface{}) bool

	// InsertOrReplace adds an element to the map. If the element already exists, it is replaced.
	InsertOrReplace(key, value interface{})

	// Remove removes an element from the map. If the element does not exist, it is a no-op.
	Remove(key interface{})

	// Contains returns true if the element exists in the map.
	Contains(key interface{}) bool

	// Merge merges the elements from another map into this map. If the element already exists, it is replaced.
	Merge(other OrderedMap)

	// Len returns the number of elements in the map.
	Len() int

	// Empty returns true if the map is empty.
	Empty() bool

	// Get returns the value associated with the key.
	Get(key interface{}) (interface{}, bool)

	// Keys returns the keys in the map.
	Keys() []interface{}
}

// how to ensure map's key is sortable
type I interface {
	// Insert1 key is sortable, but only base type
	Insert1(key Ordered, value any) bool

	// Insert2 key is sortable, but not support base type
	Insert2(key sort.Interface, value any) bool
}

type orderedMap[K Ordered] struct {
	// The map of keys to values.
	store rbtree.RbTreeI[K]
}

func (m *orderedMap[K]) Insert(key, value interface{}) bool {
	return m.store.Insert(key, value)
}
