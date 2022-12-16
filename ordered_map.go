package ordered_map

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
