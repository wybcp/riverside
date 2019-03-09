package fib

import (
	"testing"
)

func TestFib(t *testing.T) {
	var i uint = 1
	for ; i < 6; i++ {
		t.Log(i, "的 fib is ", Fib(i))
	}
}
