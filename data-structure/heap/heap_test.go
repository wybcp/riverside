package heap

import "testing"

func TestHeap_Insert(t *testing.T) {
	h := NewHeap(5)
	t.Log(h)
	h.Insert(5)
	t.Log(h)
	h.Insert(6)
	t.Log(h)
	h.Insert(4)
	t.Log(h)
	h.Insert(3)
	t.Log(h)
	h.Insert(7)
	t.Log(h)
	// h.RemoveMax()
	// t.Log(h)
	h.Sort()
	t.Log(h)
}
