package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	// "time"
)

// func main() {
// 	// var a=[]int{1,2,3}
// 	// // 在开头添加一个元素
// 	// a=append([]int{0},a...)
// 	// // 在开头添加一个切片
// 	// a=append([]int{-3,-2,-1},a...)

// 	// // 在第i的位置插入x
// 	// i:=3
// 	// x:=9
// 	// a=append(a[:i],append([]int{x},a[i:]...)...)
// 	// // 切片扩展一个空间
// 	// a=append(a,0)
// 	// // a[i:]后移一个位置
// 	// copy(a[i+1:],a[i:])
// 	// a[i]=x
// 	// xSlice:=[]int{2,8,33}
// 	// a=append(a,xSlice...)
// 	// copy(a[i+len(xSlice):],a[i:])
// 	// copy(a[i:],xSlice)
// 	// fmt.Println(a)
// 	// fmt.Println(a[:0])
// 	for i := 0; i < 3; i++ {
// 		// defer func() { println(i) }()
// 		defer func(i int){ println(i) } (i)
// 	}

// }

// 生产者: 生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
		spew.Dump(v)
	}
}
func main() {
	//ch := make(chan int, 64) // 成果队列
	//
	//go Producer(3, ch) // 生成 3 的倍数的序列
	//go Producer(5, ch) // 生成 5 的倍数的序列
	//go Consumer(ch)    // 消费 生成的队列
	//
	//// 运行一定时间后退出
	//// time.Sleep(5 * time.Second)
	//select {}

	//// nil 切片
	//var s1 []int
	//// 空切片
	//var s2 []int = []int{}
	//spew.Dump(s1,s2)
	//slice := []int{1, 2, 3, 4, 5}
	////长度为2-1=1，容量为3-1=2
	//newSlice := slice[1:2:3]
	//spew.Dump(slice,newSlice)
	////内置的append也是一个可变参数的函数，可以同时追加好几个值。
	//newSlice=append(newSlice,10,20,30)
	//spew.Dump(newSlice)

	//通过...操作符，把一个切片追加到另一个切片里。
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[1:2:3]
	newSlice = append(newSlice, slice...)
	spew.Dump(newSlice)
	//fmt.Println(newSlice)
	//fmt.Println(slice)
}
