package main

import "fmt"

func main() {
	for i := 1; i < 6; i++ {
		fmt.Println(i, "步阶梯有：", walk(i), "走法！")
		fmt.Println(i, "步阶梯有：", walk2(i), "走法！")
	}
}

// 递归实现
func walk(n int) int {
	if n <= 0 {
		panic("阶梯大于0")
	}
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	return walk(n-1) + walk(n-2)

}

// 非递归
func walk2(n int) int {
	//int ret = 0;
	//int pre = 2;
	//int prePre = 1;
	//for (int i = 3; i <= n; ++i) {
	//	ret = pre + prePre;
	//	prePre = pre;
	//	pre = ret;
	//}
	//return ret;
	if n <= 0 {
		panic("阶梯大于0")
	}
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	next := 0
	pre := 2
	prePre := 1
	for i := 3; i <= n; i++ {
		next = pre + prePre
		pre, prePre = next, pre
	}
	return next

}
