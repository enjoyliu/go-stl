package ordered_map

import (
	"sort"

	. "golang.org/x/exp/constraints"

	"go-stl/rbtree"
)

// OrderedMapI is a map that maintains the order of insertion.
type OrderedMapI interface {

	// Insert adds an element to the map. If the element already exists, it is not added again.
	Insert(key, value interface{}) bool

	// InsertOrReplace adds an element to the map. If the element already exists, it is replaced.
	InsertOrReplace(key, value interface{})

	// Remove removes an element from the map. If the element does not exist, it is a no-op.
	Remove(key interface{})

	// Contains returns true if the element exists in the map.
	Contains(key interface{}) bool

	// Merge merges the elements from another map into this map. If the element already exists, it is replaced.
	Merge(other OrderedMapI)

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
type I[K Ordered, V any] interface {
	// Insert1 key is sortable, but only base type
	Insert1(key K, value V) bool

	// Insert2 key is sortable, but not support base type
	Insert2(key sort.Interface, value V) bool
}

type OrderedMap[K Ordered, V any] struct {
	// The map of keys to values.
	store rbtree.RbTreeI[K, V]
}

func (m *OrderedMap[K, V]) InsertOrReplace(key K, value V) {
	//TODO implement me
	panic("implement me")
}

func (m *OrderedMap[K, V]) Remove(key K) {
	//TODO implement me
	panic("implement me")
}

func (m *OrderedMap[K, V]) Contains(key K) bool {
	//TODO implement me
	panic("implement me")
}

func (m *OrderedMap[K, V]) Merge(other OrderedMap[K, V]) {
	//TODO implement me
	panic("implement me")
}

func (m *OrderedMap[K, V]) Len() int {
	//TODO implement me
	panic("implement me")
}

func (m *OrderedMap[K, V]) Empty() bool {
	//TODO implement me
	panic("implement me")
}

func (m *OrderedMap[K, V]) Get(key K) (interface{}, bool) {
	//TODO implement me
	panic("implement me")
}

func (m *OrderedMap[K, V]) Keys() []K {
	//TODO implement me
	panic("implement me")
}

func (m *OrderedMap[K, V]) Insert(key K, value V) bool {
	return m.store.Insert(key, value)
}

func NewOrderedMap[K Ordered, V any]() *OrderedMap[K, V] {

	return &OrderedMap[K, V]{store: rbtree.NewTree[K, V]()}
}
