package simplylinkedlist

import "fmt"

//Node 节点
type Node struct {
	data interface{}
	next *Node
}

//List 单向链表
type List struct {
	head *Node
}

//Add 从list头部添加
func (l *List) Add(node Node) {
	node.next, l.head = l.head, &node
}

// IsEmpty 是否为空
func (l *List) IsEmpty() bool {
	return l.head == nil
}

// ReversedList 翻转链表
func (l *List) ReversedList() {
	// 中间变量
	var nodeNew *Node

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
