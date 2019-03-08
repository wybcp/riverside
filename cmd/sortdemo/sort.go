package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 4, 7, 9, 23, 2, 37}
	sort.Ints(a)
	for i, v := range a {
		fmt.Println(i, v)
	}
}
