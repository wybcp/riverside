package stack

import "fmt"

/*
基于数组实现的栈
*/

// ArrayStack 数组实现的栈的结构
type ArrayStack struct {
	//数据
	data []interface{}
	//栈顶指针
	top int
}
// NewArrayStack 新建一个数组实现的栈
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

// IsEmpty 判断是否为空
func (a *ArrayStack) IsEmpty() bool {
	isEmpty:=false
	if a.top < 0 {
		isEmpty=true
	}
	return isEmpty
}

// Push 入栈
func (a *ArrayStack) Push(v interface{}) {
	if a.top < 0 {
		a.top = 0
	} else {
		a.top++
	}

	if a.top > len(a.data)-1 {
		a.data = append(a.data, v)
	} else {
		a.data[a.top] = v
	}
}

// Pop 出栈
func (a *ArrayStack) Pop() interface{} {
	if a.IsEmpty() {
		return nil
	}
	v := a.data[a.top]
	a.top--
	return v
}

// Top 栈顶
func (a *ArrayStack) Top() interface{} {
	if a.IsEmpty() {
		return nil
	}
	return a.data[a.top]
}

// Flush 重置
func (a *ArrayStack) Flush() {
	a.top = -1
}

// String 字符串输出
func (a *ArrayStack) String() (str string) {
	if a.IsEmpty() {
		str = "empty stack"
	} else {
		for i := a.top; i >= 0; i-- {
			// str += fmt.PrintF(i, "层:", a.data[i])
			fmt.Printf("%d 层:%v;", i, a.data[i])
		}
	}
	return
}
