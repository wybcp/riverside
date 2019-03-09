package main

import (
	"testing"
)

func BenchmarkThreeSum(b *testing.B) {
	b.ResetTimer()
	a := []int{-1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3, 7, 8, 9, 10, -8, -6, -7, -1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3, 7, 8, 9, 10, -8, -6, -7}
	for i := 0; i < b.N; i++ {
		ThreeSum(a)
	}
}
func BenchmarkThreeSum2(b *testing.B) {
	b.ResetTimer()
	a := []int{-1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3, 7, 8, 9, 10, -8, -6, -7, -1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3, 7, 8, 9, 10, -8, -6, -7}
	for i := 0; i < b.N; i++ {
		threeSum(a)
	}
}

// goos: darwin
// goarch: amd64
// pkg: riverside/letcode
// BenchmarkThreeSum-8   	  100000	     11337 ns/op	    1536 B/op	      22 allocs/op
// PASS
// ok  	riverside/letcode	1.268s

//func Test_threeSum(t *testing.T) {
//	type args struct {
//		nums []int
//	}
//	tests := []struct {
//		name string
//		args args
//		want [][]int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ThreeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_threeSumZero(t *testing.T) {
//	type args struct {
//		wg   *sync.WaitGroup
//		nums []int
//		i    int
//		n    int
//		in   chan []int
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			threeSumZero(tt.args.wg, tt.args.nums, tt.args.i, tt.args.n, tt.args.in)
//		})
//	}
//}
