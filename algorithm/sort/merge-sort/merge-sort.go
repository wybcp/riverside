package mergesort

import (
	"math"
)

// 归并排序,把数组从中间分成前后两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起，

// MergeSort 归并排序
func MergeSort(arr []int, low, high int) {
	sentinel := math.MaxInt64
	if low < high {
		mid := (low + high) / 2
		MergeSort(arr, low, mid)
		MergeSort(arr, mid+1, high)
		merge(arr, low, mid, high, sentinel)

	}
}

func merge(arr []int, low, mid, high, sentinel int) {
	leftLen := mid - low + 1
	rightLen := high - mid

	arrLeft := make([]int, leftLen+1)
	for i := 0; i < leftLen; i++ {
		arrLeft[i] = arr[low+i]
	}
	arrLeft[leftLen] = sentinel //哨兵牌,比其中的所有数据都大
	// fmt.Println(arrLeft)

	arrRight := make([]int, rightLen+1)
	for j := 0; j < rightLen; j++ {
		arrRight[j] = arr[mid+j+1]
	}
	arrRight[rightLen] = sentinel //哨兵牌
	// fmt.Println(arrRight)
	i, j := 0, 0
	for k := low; k <= high; k++ {
		if arrLeft[i] <= arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}
}
