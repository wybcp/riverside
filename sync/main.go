package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	go father(&wg)

	wg.Wait()
	log.Printf("main: father and all chindren exit")
}

func father(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	fmt.Println("father")
	for i := 0; i < 10; i++ {
		go child(wg, i)
	}
}

func child(wg *sync.WaitGroup, i int) {
	wg.Add(1)
	defer wg.Done()

	fmt.Printf("child [%d]\n", i)
}
