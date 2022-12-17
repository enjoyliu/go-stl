package rbtree

import (
	. "golang.org/x/exp/constraints"
)

// a red-black tree.
type RbTreeI[K Ordered, V any] interface {
	// Find finds the node and return its value.
	Find(key K) V

	// FindIt finds the node and return it as an iterator.
	FindIt(key K) RbNode[K, V]

	// Empty checks whether the rbtree is empty.
	Empty() bool

	// Iterator creates the rbtree's iterator that points to the minmum node.
	Iterator() RbNode[K, V]

	// Size returns the size of the rbtree.
	Size() int

	// Insert inserts a new node into the rbtree.
	Insert(key K, value V) bool

	// Erase erases a node from the rbtree.
	Erase(key K)

	// Clear clears the rbtree.
	Clear()

	// Clone clones the rbtree.
	Clone() RbTreeI[K, V]

	// Contains checks whether the rbtree contains the key.
	Contains(key K) bool
}

type RbTree[K Ordered, V any] struct{}

func (RbTree[K, V]) Find(key K) V {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K, V]) FindIt(key K) RbNode[K, V] {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K, V]) Empty() bool {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K, V]) Iterator() RbNode[K, V] {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K, V]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K, V]) Insert(key K, value V) bool {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K, V]) Erase(key K) bool {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K, V]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K, V]) Clone() RbTreeI[K, V] {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K, V]) Contains(key K) bool {
	//TODO implement me
	panic("implement me")
}

// a red-black tree node.
type RbNode[K Ordered, V any] interface {
	// GetKey returns the key of the node.
	GetKey() K

	// GetValue returns the value of the node.
	GetValue() V

	// Next returns the next node.
	Next() RbNode[K, V]
}
