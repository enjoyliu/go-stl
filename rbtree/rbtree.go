package rbtree

import (
	. "golang.org/x/exp/constraints"
)

func NewTree[K Ordered]() RbTreeI[K] {
	return RbTree[K]{}
}

// a red-black tree.
type RbTreeI[K Ordered] interface {
	// Find finds the node and return its value.
	Find(key K) any

	// FindIt finds the node and return it as an iterator.
	FindIt(key K) RbNode

	// Empty checks whether the rbtree is empty.
	Empty() bool

	// Iterator creates the rbtree's iterator that points to the minmum node.
	Iterator() RbNode

	// Size returns the size of the rbtree.
	Size() int

	// Insert inserts a new node into the rbtree.
	Insert(key K, value any) bool

	// Erase erases a node from the rbtree.
	Erase(key K) bool

	// Clear clears the rbtree.
	Clear()

	// Clone clones the rbtree.
	Clone() RbTreeI[K]

	// Contains checks whether the rbtree contains the key.
	Contains(key K) bool
}

type RbTree[K Ordered] struct{}

func (RbTree[K]) Find(key K) any {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K]) FindIt(key K) RbNode {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K]) Empty() bool {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K]) Iterator() RbNode {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K]) Insert(key K, value any) bool {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K]) Erase(key K) bool {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K]) Clone() RbTreeI[K] {
	//TODO implement me
	panic("implement me")
}

func (RbTree[K]) Contains(key K) bool {
	//TODO implement me
	panic("implement me")
}

// a red-black tree node.
type RbNode interface {
	// Key returns the key of the node.
	Key() Ordered

	// Value returns the value of the node.
	Value() any

	// Next returns the next node.
	Next() RbNode
}
