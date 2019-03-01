package main

import "fmt"

func test1() (x int) {
	defer fmt.Printf("in test1 defer: x = %d\n", x)
	x = 7
	return 9
}

func test2() (x int) {
	x = 7
	defer fmt.Printf("in test2 defer: x = %d\n", x)
	return 9
}

func test3() (x int) {
	defer func() {
		fmt.Printf("in test3 defer: x = %d\n", x)
	}()

	x = 7
	return 9
}

func test4() (x int) {
	defer func(n int) {
		fmt.Printf("in test4 defer x as parameter: x = %d\n", n)
		fmt.Printf("in test4 defer x after return: x = %d\n", x)
	}(x)

	x = 7
	return 9
}

func main() {
	fmt.Println("test1")
	fmt.Printf("in main: x = %d\n", test1())
	fmt.Println("test2")
	fmt.Printf("in main: x = %d\n", test2())
	fmt.Println("test3")
	fmt.Printf("in main: x = %d\n", test3())
	fmt.Println("test4")
	fmt.Printf("in main: x = %d\n", test4())
	test5()
}

func test5() (r int) {
	r = 1
	defer func(r int) {
		r += 5
		println(r)
	}(r)
	println(r)
	return
}
