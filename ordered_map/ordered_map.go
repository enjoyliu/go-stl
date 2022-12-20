package ordered_map

import (
	. "golang.org/x/exp/constraints"

	"go-stl/container"
	"go-stl/rbtree"
)

type LessFunc[K any] func(K, K) bool

type OrderedMap[K any, V any] struct {
	// The map of keys to values.
	store *rbtree.Tree[K, V]

	less LessFunc[K]
}

func (m *OrderedMap[K, V]) Insert(key K, value V) bool {
	return m.store.Insert(key, value)
}

func (m *OrderedMap[K, V]) InsertOrReplace(key K, value V) {
	m.store.Insert(key, value)
}

func (m *OrderedMap[K, V]) Remove(key K) {
	m.store.Erase(key)
}

func (m *OrderedMap[K, V]) Contains(key K) bool {
	return m.store.Contains(key)
}

func (m *OrderedMap[K, V]) Merge(other container.OrderedMap[K, V]) {
	//TODO implement me
	panic("implement me")
}

func (m *OrderedMap[K, V]) Len() int {
	return m.store.Size()
}

func (m *OrderedMap[K, V]) Empty() bool {
	return m.store.Empty()
}

func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	return m.store.Find(key)
}

func (m *OrderedMap[K, V]) Iterator() container.MapIterator[K, V] {
	return m.store.Iterator()
}

func (m *OrderedMap[K, V]) Find(key K) container.MapIterator[K, V] {
	return m.store.FindIt(key)
}

// Clear removes all elements from the map.
func (m *OrderedMap[K, V]) Clear() {
	m.store.Clear()
}

func NewOrderedMap[K Ordered, V any]() container.OrderedMap[K, V] {
	cmp := DefaultCmp[K]()
	return &OrderedMap[K, V]{
		store: rbtree.NewTree[K, V](cmp),
		less:  cmp,
	}
}

func DefaultCmp[K Ordered]() LessFunc[K] {
	return func(a, b K) bool {
		return a < b
	}
}

func WithCompareFunc[K any, V any](cmp func(K, K) bool) container.OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		store: rbtree.NewTree[K, V](cmp),
		less:  cmp,
	}
}
