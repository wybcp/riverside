package quicksort

import (
	"fmt"
	"math/rand"
)

// 快排算法
// 两种方式差别不是很大，只是避免出现极端情况

// 分区
func partition(a []int, low, high int) int {
	// 直接选取
	pivot := a[high]
	i := low
	for ; low < high; low++ {
		if a[low] < pivot {
			a[low], a[i] = a[i], a[low]
			i++
		}
	}

	a[i], a[high] = pivot, a[i]

	return i
}
func partitionByRand(a []int, low, high int) int {
	// 可以随机选取，也可以前中后，三取一，避免极端情况，概率平均，应该在大量数据
	// 才能体会到这个优势，因为多了几步运算
	randInt := rand.Intn(high-low) + low
	// 先交换取得的随机数据和最后一味地位置
	a[randInt], a[high] = a[high], a[randInt]
	pivot := a[high]
	i := low - 1
	for ; low < high; low++ {
		if a[low] < pivot {
			i++
			a[low], a[i] = a[i], a[low]
		}
	}

	a[i+1], a[high] = pivot, a[i+1]

	return i + 1
}

//GetNth 无序数组中的第 K 大元素
func GetNth(a []int, low, high, k int) {
	if k > high {
		panic("wrong")
	}
	if low < high {
		p := partition(a, low, high)
		if k > p {
			GetNth(a, p+1, high, k)
		}
	}
}

// QuickSort 快排算法,直接选择最后的数据作为pivot
func QuickSort(a []int, low, high int) {
	if low < high {
		p := partition(a, low, high)
		fmt.Println(a, p)
		QuickSort(a, low, p-1)
		QuickSort(a, p+1, high)
	}
}

//QuickSortByRand 快排算法,随机选取pivot
func QuickSortByRand(a []int, low, high int) {
	if low < high {
		p := partitionByRand(a, low, high)
		QuickSortByRand(a, low, p-1)
		QuickSortByRand(a, p+1, high)
	}
}
