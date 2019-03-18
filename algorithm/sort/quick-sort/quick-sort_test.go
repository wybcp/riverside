package quicksort

import (
	"math/rand"
	"testing"
	"time"
)

var data = []int{34, 45,10, 3, 6, 76, 34, 46, 809, 92, 8, 11}

func TestQuickSort(t *testing.T) {
	t.Log(data)
	QuickSort(data, 0, len(data)-1)
	t.Log(data)
}

// 无序数组中的第 K 大元素。比如，4， 2， 5， 12， 3 这样一组数据，第 3 大元素就是 4。

// 我们选择数组区间 A[0…n-1] 的最后一个元素 A[n-1] 作为 pivot，对数组 A[0…n-1] 原地分区，这样数组就分成了三部分，A[0…p-1]、A[p]、A[p+1…n-1]。

// 如果 p+1=K，那 A[p] 就是要求解的元素；如果 K>p+1, 说明第 K 大元素出现在 A[p+1…n-1] 区间，我们再按照上面的思路递归地在 A[p+1…n-1] 这个区间内查找。同理，如果 K<p+1，那我们就在 A[0…p-1] 区间查找。
func TestGetNth(t *testing.T) {
	t.Log(data)
	rand.Seed(time.Now().UnixNano())
	k := (rand.Int())%len(data) + 1
	GetNth(data, 0, len(data)-1, k)
	t.Log(data)
	t.Log(k, "->", data[k-1])
}

func makeLargeSlice(n int) {
	// b.Log(data)
	for index := 0; index < n; index++ {
		data = append(data, rand.Int())
	}
	// b.Log(data)
}
func BenchmarkQuickSort(b *testing.B) {
	makeLargeSlice(10)
	for index := 0; index < b.N; index++ {
		QuickSort(data, 0, len(data)-1)
	}
	// b.Log(data)
}
func BenchmarkQuickSortByRand(b *testing.B) {

	for index := 0; index < b.N; index++ {
		QuickSortByRand(data, 0, len(data)-1)
	}
	// b.Log(data)
}

// BenchmarkQuickSort-8   	10000000	       135 ns/op	       0 B/op	       0 allocs/op
// BenchmarkQuickSort-8   	     100	 196570134 ns/op	  120689 B/op
// BenchmarkQuickSort-8   	    1000	   1521732 ns/op	   19690 B/op
