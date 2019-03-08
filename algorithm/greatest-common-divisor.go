package main

import "fmt"

func main() {
	fmt.Println(gcdx(867, 1122))
	fmt.Println(gcd(867, 1122))
}

/*
*辗转相除法：最大公约数
*递归写法，进入运算是x和y都不为0
 */
func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}

/*
*辗转相除法：最大公约数
*非递归写法
 */
func gcdx(x, y int) int {
	for {
		tmp := x % y
		if tmp > 0 {
			x ,y= y,tmp
		} else {
			return y
		}
	}
}
