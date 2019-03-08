package main

import "fmt"

func partition(a []int, low, high int) int {

	pivot := a[high]
	i := low - 1
	for j := low; j < high; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}
func quickSort(a []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(a, low, high)
	quickSort(a, low, p-1)
	quickSort(a, p+1, high)
}

func main() {
	arr01 := []int{34, 45, 3, 6, 76, 34, 46, 809, 92, 8,11}

	fmt.Print("排序前")
	fmt.Println(arr01)

	quickSort(arr01, 0, len(arr01)-1)

	fmt.Print("排序后")
	fmt.Println(arr01)
}
