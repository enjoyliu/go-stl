package container

type Container interface {
	// Len returns the number of elements in the container.
	Len() int
	// Empty returns true if the container is empty.
	Empty() bool
	// Clear removes all elements from the container.
	Clear()
}

type Map[K any, V any] interface {
	Container
}

// OrderedMap is a map that maintains the order of insertion.
type OrderedMap[K any, V any] interface {
	Map[K, V]

	// Insert adds an element to the map. If the element already exists, it is not added again.
	Insert(key K, value V) bool

	// InsertOrReplace adds an element to the map. If the element already exists, it is replaced.
	InsertOrReplace(key K, value V)

	// Remove removes an element from the map. If the element does not exist, it is a no-op.
	Remove(key K)

	// Contains returns true if the element exists in the map.
	Contains(key K) bool

	// Merge merges the elements from another map into this map. If the element already exists, it is replaced.
	Merge(other OrderedMap[K, V])

	// Len returns the number of elements in the map.
	Len() int

	// Empty returns true if the map is empty.
	Empty() bool

	// Get returns the value associated with the key.
	Get(key K) (V, bool)

	// Find returns the value associated with the key.
	Find(key K) MapIterator[K, V]

	// Iterator returns the value associated with the key.
	Iterator() MapIterator[K, V]
}
