package count

import "testing"

func TestCountSort(t *testing.T) {
	data := []int{34, 30, 3, 6, 5, 34, 35, 6, 22, 8, 11}
	t.Log(data)
	t.Log(CountSort(data))
}
