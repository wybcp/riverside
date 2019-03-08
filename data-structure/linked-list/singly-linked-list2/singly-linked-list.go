package main

import (
	"fmt"
)

type Object interface{}

//节点
type Node struct {
	data Object
	next *Node
}

//单向链表
type List struct {
	head *Node
}

//从list头部添加
func (l *List) Add(node Node) {
	node.next, l.head = l.head, &node
}
func (l *List) IsEmpty() bool {
	return l.head == nil
}

func (l *List) ReversedList() {
	// 中间变量
	var nodeNew *Node = nil

	for !l.IsEmpty() {
		//取list头处理
		node := *(l.head)
		nodeNext := node.next
		//移动list头
		l.head = nodeNext

		//取出来的头结点处理，改变指针
		node.next = nodeNew

		nodeNew = &node
	}
	//将新的变量值赋予就得list头部
	l.head = nodeNew
}

func (l *List) String() string {
	if l.IsEmpty() {
		return "链表为空"
	}
	node := l.head
	result := ""
	for node.next != nil {
		result += fmt.Sprintf("%+v->", node.data)
		node = node.next
	}
	//处理尾节点
	result += fmt.Sprintf("%+v", node.data)
	return result

}

func main() {
	var list = List{}
	for i := 0; i < 5; i++ {
		list.Add(Node{data: i})
	}
	fmt.Println(list.String())
	list.ReversedList()
	fmt.Println("reverse list：", list.String())
}
