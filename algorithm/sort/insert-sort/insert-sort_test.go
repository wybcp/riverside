package insert

import "testing"

var values = []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}

func TestInsertIncSort(t *testing.T) {
	t.Log(values)
	InsertIncSort(values)
	t.Log(values)
}

func TestInsertDecSort(t *testing.T) {
	t.Log(values)
	InsertDecSort(values)
	t.Log(values)
}
