package _8_stack

import "fmt"

/*
基于链表实现的栈
*/
type node struct {
	next *node
	val  interface{}
}

type LinkedListStack struct {
	//栈顶节点
	topNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (l *LinkedListStack) IsEmpty() bool {
	if l.topNode == nil {
		return true
	}
	return false
}

func (l *LinkedListStack) Push(v interface{}) {
	l.topNode = &node{next: l.topNode, val: v}
}

func (l *LinkedListStack) Pop() interface{} {
	if l.IsEmpty() {
		return nil
	}
	v := l.topNode.val
	l.topNode = l.topNode.next
	return v
}

func (l *LinkedListStack) Top() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return l.topNode.val
}

func (l *LinkedListStack) Flush() {
	l.topNode = nil
}

func (l *LinkedListStack) Print() {
	if l.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := l.topNode
		for nil != cur {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
