package rbtree

import (
	"go-stl/container"
)

// color of node
const (
	RED   = 0
	BLACK = 1
)

// Tree is a struct of red-black tree.
type Tree[K any, V any] struct {
	root     *node[K, V]
	size     int
	lessFunc func(K, K) bool
}

// NewTree creates a new rbtree.
func NewTree[K any, V any](cmp func(K, K) bool) *Tree[K, V] {
	return &Tree[K, V]{
		lessFunc: cmp,
	}
}

// Find finds the node and return its value.
func (t *Tree[K, V]) Find(key K) (V, bool) {
	n := t.findnode(key)
	if n != nil {
		return n.value, true
	}
	var result V
	return result, false
}

// Contains checks whether the rbtree contains the key.
func (t *Tree[K, V]) Contains(key K) bool {
	if t.findnode(key) == nil {
		return false
	}
	return true
}

// FindIt finds the node and return it as an iterator.
func (t *Tree[K, V]) FindIt(key K) container.MapIterator[K, V] {
	return t.findnode(key)
}

// Empty checks whether the rbtree is empty.
func (t *Tree[K, V]) Empty() bool {
	if t.root == nil {
		return true
	}
	return false
}

// Iterator creates the rbtree's iterator that points to the minmum node.
func (t *Tree[K, V]) Iterator() container.MapIterator[K, V] {
	return minimum(t.root)
}

// Size returns the size of the rbtree.
func (t *Tree[K, V]) Size() int {
	return t.size
}

// Clear destroys the rbtree.
func (t *Tree[K, V]) Clear() {
	t.root = nil
	t.size = 0
}

// Insert inserts the key-value pair into the rbtree.
func (t *Tree[K, V]) Insert(key K, value V) bool {
	// TODO need return if the key has already existed
	x := t.root
	var y *node[K, V]

	for x != nil {
		y = x
		if t.lessFunc(key, x.key) {
			x = x.left
		} else if t.lessFunc(x.key, key) {
			x = x.right
		} else {
			return false
		}
	}

	z := &node[K, V]{parent: y, color: RED, key: key, value: value}
	t.size++

	if y == nil {
		// if z has no parent, it is the root
		z.color = BLACK
		t.root = z
		return true
	} else if t.lessFunc(z.key, y.key) {
		y.left = z
	} else {
		y.right = z
	}
	t.rbInsertFixup(z)
	return true
}

// Erase deletes the node by key
func (t *Tree[K, V]) Erase(key K) {
	z := t.findnode(key)
	if z == nil {
		return
	}

	var x, y *node[K, V]
	if z.left != nil && z.right != nil {
		y = successor(z)
	} else {
		y = z
	}

	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}

	xparent := y.parent
	if x != nil {
		x.parent = xparent
	}
	if y.parent == nil {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	if y != z {
		z.key = y.key
		z.value = y.value
	}

	if y.color == BLACK {
		t.rbDeleteFixup(x, xparent)
	}
	t.size--
}

// Clone creates a new rbtree that is a clone of the original rbtree.
func (t *Tree[K, V]) Clone() *Tree[K, V] {
	// TODO: need to implement

	// iterate the tree
	var newTree *Tree[K, V]
	newTree = NewTree[K, V](t.lessFunc)
	*newTree.root = *t.root

	return nil
}
