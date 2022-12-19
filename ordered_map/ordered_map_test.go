package ordered_map

import (
	"testing"
)

func TestNewOrderedMap(t *testing.T) {
	var tt OrderedMap[string, string]
	tt.Len()

	m := NewOrderedMap[int, string]()
	m.Insert(1, "sss")
	m.Insert(33, "s")
}
