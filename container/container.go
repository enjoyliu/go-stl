package container

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Sortable interface {
	sort.Interface
}

type Ordered interface {
	Sortable | constraints.Ordered
}

type Container interface {
	// Len returns the number of elements in the container.
	Len() int
	// Empty returns true if the container is empty.
	Empty() bool
	// Clear removes all elements from the container.
	Clear()
}

// Queue is a container adapter that gives the programmer the functionality of a queue.
type Queue[T any] interface {
	Container
	// Push adds an element to the back of the queue.
	Push(value T)
	// Pop removes an element from the front of the queue.
	Pop() T
	// Front returns the element at the front of the queue.
	Front() T
}

// Stack is a container adapter that gives the programmer the functionality of a stack.
type Stack[T any] interface {
	Container
	// Push adds an element to the top of the stack.
	Push(value T)
	// Pop removes an element from the top of the stack.
	Pop() T
	// Top returns the element at the top of the stack.
	Top() T
}

// Deque is a container adapter that gives the programmer the functionality of a deque.
type Deque[T any] interface {
	Container
	// PushFront adds an element to the front of the deque.
	PushFront(value T)
	// PushBack adds an element to the back of the deque.
	PushBack(value T)
	// PopFront removes an element from the front of the deque.
	PopFront() T
	// PopBack removes an element from the back of the deque.
	PopBack() T
	// Front returns the element at the front of the deque.
	Front() T
	// Back returns the element at the back of the deque.
	Back() T
}

// PriorityQueue is a container adapter that gives the programmer the functionality of a priority queue.
type PriorityQueue[T any] interface {
	Container
	Queue[T]
}

// Map is a container adapter that gives the programmer the functionality of a map.
type Map[K any, V any] interface {
	Container
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
}

// OrderedMap is a map that maintains the order of insertion.
type OrderedMap[K Ordered, V any] interface {
	Map[K, V]

	// Find returns the value associated with the key.
	Find(key K) MapIterator[K, V]

	// Iterator returns the value associated with the key.
	Iterator() MapIterator[K, V]
}
