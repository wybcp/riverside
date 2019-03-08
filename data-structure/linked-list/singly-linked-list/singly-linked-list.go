
package main

import (
	"fmt"
)

type Object interface {}

//节点
type Node struct {
	data Object
	next *Node
}
//单向链表
type List struct {
	head *Node
	tail *Node
	size uint64
}


//初始化
func(l *List) Init(){
	(*l).size = 0   // 此时链表是空的
	(*l).head = nil // 没有头
	(*l).tail = nil // 没有尾
}

//向尾部添加数据
func (l *List) Append(node *Node) bool {
	if node == nil {
		return false
	}

	(*node).next = nil
	// 将新元素放入单链表中
	if (*l).size == 0 {
		(*l).head = node
	} else {
		(*l).tail.next = node
	}

	// 调整尾部位置及链表元素数量
	(*l).tail = node // node成为新的尾部
	(*l).size++      // 元素数量增加

	return true
}
//插入数据
func (l *List) Insert(i uint64,node *Node) bool {
	// 空的节点、索引超出范围和空链表都无法做插入操作
	if node == nil || i > (*l).size || (*l).size == 0 {
		return false
	}

	if i == 0 { // 直接排第一
		(*node).next = (*l).head
		(*l).head = node
	} else {
		// 找到前一个元素
		preItem := (*l).head
		for j := 1 ; uint64(j) < i; j++ { // 数前面i个元素
			preItem = (*preItem).next
		}
		// 原来元素放到新元素后面,新元素放到前一个元素后面
		(*node).next = (*preItem).next
		(*preItem).next = preItem
	}

	(*l).size++

	return true
}

//删除元素
func (l *List) Remove(i uint64, node *Node) bool {
	if i >= (*l).size {
		return false
	}

	if i == 0 { // 删除头部
		node = (*l).head
		(*l).head = (*node).next
		if (*l).size == 1 { // 如果只有一个元素，那尾部也要调整
			(*l).tail = nil
		}
	} else {
		preItem := (*l).head
		for  j  := 1; uint64(j) < i; j++ {
			preItem = (*preItem).next
		}

		node = (*preItem).next
		(*preItem).next = (*node).next

		if i == ((*l).size - 1) { // 若删除的尾部，尾部指针需要调整
			(*l).tail = preItem
		}
	}
	(*l).size--
	return true
}
//获取元素
func (l *List) Get(i uint64) *Node {
	if i >= (*l).size {
		return nil
	}

	item := (*l).head
	for j := 0; uint64(j) < i ; j++ {    // 从head数i个
		item = (*item).next
	}

	return item
}
func (l *List)String() string {

	if l.size==0 {
		return "链表为空"
	}else {
		var i uint64
		node:=l.head
		result:=""
		for ;i<l.size ;i++  {
			//spew.Dump(node)
			//spew.Dump(node.data)
			//result += fmt.Sprintf("<-%+v", a.q[i])
			if l.tail==node {
				result += fmt.Sprintf("%+v", node.data)
			}else {
				result += fmt.Sprintf("%+v->", node.data)
			}
			node=node.next
		}
		return result
	}
}
//直接新建一个reverse list
func reversedList(list List)List  {
	i :=list.size
	revList:=List{}
	for ;i>0 ;  {
		revList.Append(list.Get(i))
		i--
	}
	revList.Append(list.head)
	return revList
}
////把数据先存到数组或栈里面
//func (l *List)reverseListByStack()  {
//	var i uint64
//	nodeL:=(*l).head
//	for ;i<(*l).size;i++{
//
//
//
//		//node=node.next
//
//	}
//}

func (l *List)ReverseList(){
	nodeLeft:=l.head

	for i:=1;i<3 ;i++  {
		nodeLeft.next=l.head
		node=node.next
	}
}
//func reverseListByArray(list List)  {
//	arr:=new([list.size]interface{})
//	var i uint64
//	node:=list.head
//
//	for ; i<list.size;i++  {
//		arr[i]=node.data
//		node=node.next
//	}
//	for ; ;  {
//
//	}
//	list.printList(list)
//}
////直接新建一个reverse list
//func reversedList2(list List)List  {
//	i :=list.size
//
//	for ;i>0 ;  {
//		headNode:=list.head
//		list.tail.next=headNode
//		list.tail=headNode
//		list.head=list.head.next
//		i--
//		list.printList(list)
//	}
//	//revList.Append(list.head)
//	return list
//}
func main(){
	var list= List{}
	list.Init()
	for i:=1;i<=40 ;i++  {
		node:=Node{data:i}
		list.Append(&node)
	}
	// 输出链表
	fmt.Println(list.String())

	//list.printList(reversedList(list))
	//list.printList(reversedList2(list))
	//var node=list.Get(35)
	//if node!=nil {
	//	fmt.Printf("\n当前节点位置：%+q ,数据：%d \n",node,node.data)
	//}

	//var deleteNode=&Node{}
	//var result=list.Remove(35,deleteNode)
	//fmt.Printf("删除是否成功 %+q \n",result)
	//
	//var node2=list.Get(35)
	//fmt.Printf("当前节点位置：%+q \n,数据：%d \n",node2,node2.data)
}