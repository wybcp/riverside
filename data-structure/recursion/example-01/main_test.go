package main

import "testing"

func Benchmark_walk(b *testing.B)  {
	for i:=0;i<b.N ; i++ {
		walk(50)
	}

}
func Benchmark_walk2(b *testing.B)  {
	for i:=0;i<b.N ; i++ {
		walk2(50)
	}

}
//10
//Benchmark_walk-8    	 5000000	       238 ns/op
//Benchmark_walk2-8   	200000000	         6.26 ns/op

//50
//Benchmark_walk-8    	       1	52961117525 ns/op
//Benchmark_walk2-8   	100000000	        21.8 ns/op
