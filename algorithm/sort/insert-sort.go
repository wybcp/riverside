package main

import "fmt"

func main() {
	numbers := []int{6, 2, 7, 3, 8, 5}

	//insertIncSort(numbers)
	insertDecSort2(numbers)

}

//插入排序

func insertIncSort(s []int) {
	for j := 1; j < len(s); j++ {
		key := s[j]
		i := j - 1
		for i >= 0 && s[i] > key {
			s[i+1] = s[i]
			i--
		}
		s[i+1] = key
	}
	fmt.Println(s)
}
func insertDecSort2(s []int) {
	for j := 1; j < len(s); j++ {
		key := s[j]
		i := j - 1
		for i >= 0 && s[i] < key {
			s[i+1] = s[i]
			i--
		}
		s[i+1] = key
	}
	fmt.Println(s)
}
