package main

import (
	"fmt"
	"log"
)

// var where = log.Print

func main() {
	// test1()
	// test2()
	test3()
}

func test1() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
	where()

}
func test2() {
	for i := 0; ; i++ {
		fmt.Println("Value of i is now:", i)
	}
}
func test3() {
	where()
	for i := 0; i < 3; i++ {
		fmt.Println("Value of i:", i)
	}
}

func where() {
	log.SetFlags(log.Llongfile)
	log.Print("")
}
