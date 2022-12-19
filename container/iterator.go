package container

type Iterator[V any] interface {
	HasNext() bool
	Value() V
}

type MapIterator[K any, V any] interface {
	Iterator[V]
	Key() K
	Next() MapIterator[K, V]
}
