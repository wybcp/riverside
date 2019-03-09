package mergesort

import "testing"

var data = []int{34, 45, 3, 6, 76, 34, 46, 809, 92, 8, 11}

func TestMergeSort(t *testing.T) {
	t.Log(data)
	MergeSort(data, 0, len(data)-1)
	t.Log(data)
}

func BenchmarkMergeSort(b *testing.B) {
	for index := 0; index < b.N; index++ {
		MergeSort(data, 0, len(data)-1)
	}
}
