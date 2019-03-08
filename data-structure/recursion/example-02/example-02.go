package main
import "fmt"

func recursiveCall(product int, num int, ch chan int)  {
	product += num

	if num == 1 {
		ch <- product
		return
	}

	go recursiveCall(product, num - 1, ch)
}

func main()  {
	ch := make(chan int)
	go recursiveCall(0, 4, ch)
	product := <-ch
	fmt.Printf("Product is %d\n", product)
}