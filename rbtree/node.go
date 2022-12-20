package rbtree

import (
	"go-stl/container"
)

type node[K any, V any] struct {
	left, right, parent *node[K, V]
	color               int
	key                 K
	value               V
}

// Next returns the node's successor as an iterator.
func (n *node[K, V]) Next() container.MapIterator[K, V] {
	return successor(n)
}

// HasNext return
func (n *node[K, V]) HasNext() bool {
	return successor(n).parent != nil
}

// Key returns the key of the node.
func (n *node[K, V]) Key() K {
	return n.key
}

// Value returns the value of the node.
func (n *node[K, V]) Value() V {
	return n.value
}

// Left returns the left-child of the node.
func (n *node[K, V]) Left() *node[K, V] {
	return n.left
}

// Right
func (n *node[K, V]) Right() *node[K, V] {
	return n.right
}

// Parent
func (n *node[K, V]) Parent() *node[K, V] {
	return n.parent
}

// Clone returns a copy of the node.
func (n *node[K, V]) Clone() *node[K, V] {
	return &node[K, V]{
		key:   n.key,
		value: n.value,
		color: n.color,
	}
}
