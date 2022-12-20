package ordered_map

import (
	"sort"
)

// OrderedMap is a map with ordered keys, the key's type cloud be any struct implements sort.Interface
type OrderedMap[K sort.Interface, V any] struct {
}
