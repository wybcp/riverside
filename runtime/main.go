package main

import (
	"github.com/pkg/profile"
	"log"
	"runtime"
)

var lastTotalFreed uint64
var m map[int]int

const count = 8192

func main() {
	defer profile.Start().Stop()
	printMemStats()
	log.Println(len(m))
	initMap()
	log.Println(len(m))
	printMemStats()
	runtime.GC()
	log.Println(len(m))
	printMemStats()
	for i:=range m{
		delete(m,i)
	}
	log.Println(len(m))
	runtime.GC()
	printMemStats()

	m = nil
	runtime.GC()
	printMemStats()
}

func initMap() {
	m = make(map[int]int, count)
	for i := 0; i < count; i++ {
		m[i] = i
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v  Just Freed = %v Sys = %v NumGC = %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, ((m.TotalAlloc-m.Alloc)-lastTotalFreed)/1024, m.Sys/1024, m.NumGC)
	lastTotalFreed = m.TotalAlloc - m.Alloc
}
