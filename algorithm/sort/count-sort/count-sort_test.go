package count

import "testing"

func TestCountSort(t *testing.T) {
	data := []int{1, 4, 0, 6, 5, 7, 3, 6, 2, 8, 1}
	t.Log(data)
	t.Log(CountSort(data))
}
