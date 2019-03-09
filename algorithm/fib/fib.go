package fib

// Fib 费布拉切函数
func Fib(n uint) uint {
	if n == 0 {
		panic("n 大于 0 的整数")
	}
	if n < 3 {
		return 1
	}
	var n1 uint = 1
	var n2 uint = 1
	for index := 2; uint(index) < n; index++ {
		n2, n1 = n1+n2, n2
	}
	return n2
}
