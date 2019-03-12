package doublylinkedlist

import "fmt"

// 双向链表

// DoublyLinkedList 结构体
type DoublyLinkedList struct {
	head, tail *Node
	len, cap   int
}

// Node 双向节点
type Node struct {
	value     int
	pre, next *Node
}

// New initialize your data structure here
func New(k int) DoublyLinkedList {
	return DoublyLinkedList{
		cap: k,
	}
}

// InsertFront adds an item
func (d *DoublyLinkedList) InsertFront(value int) bool {
	if d.len == d.cap {
		return false
	}
	node := &Node{
		value: value,
	}
	if d.len == 0 {
		d.head = node
		d.tail = node
	} else {
		node.next = d.head
		d.head.pre = node
		d.head = node
	}
	d.len++
	return true
}

// InsertLast adds an item
func (d *DoublyLinkedList) InsertLast(value int) bool {
	if d.len == d.cap {
		return false
	}
	node := &Node{
		value: value,
	}
	if d.len == 0 {
		d.head = node
		d.tail = node
	} else {
		node.pre = d.tail
		d.tail.next = node
		d.tail = node
	}
	d.len++
	return true
}

// DeleteFront deletes an item from the front
func (d *DoublyLinkedList) DeleteFront() bool {
	if d.len == 0 {
		return false
	}
	if d.len == 1 {
		d.head, d.tail = nil, nil
	} else {
		d.head = d.head.next
		d.head.pre = nil
	}
	d.len--
	return true
}

// DeleteLast deletes an item
func (d *DoublyLinkedList) DeleteLast() bool {
	if d.len == 0 {
		return false
	}
	if d.len == 1 {
		d.head, d.tail = nil, nil
	} else {
		d.tail = d.tail.pre
		d.tail.next = nil
	}
	d.len--
	return true
}

// GetFront
func (d *DoublyLinkedList) GetFront() int {
	if d.len == 0 {
		return -1
	}
	return d.head.value
}

// GetRear
func (d *DoublyLinkedList) GetRear() int {
	if d.len == 0 {
		return -1
	}
	return d.tail.value
}

// IsEmpty
func (d *DoublyLinkedList) IsEmpty() bool {
	return d.len == 0
}

// IsFull checks
func (d *DoublyLinkedList) IsFull() bool {
	return d.len == d.cap
}

func (d *DoublyLinkedList) String() (str string) {
	if d.IsEmpty() {
		str = "双向链表为空"
	} else {
		node := d.head
		str = fmt.Sprintf("%d", node.value)
		for node.next != nil {
			str += fmt.Sprintf("<->%d", node.next.value)
			node = node.next
		}
	}
	return
}
