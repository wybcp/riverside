package _9_queue

import "fmt"

type CircularQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{make([]interface{}, n), n, 0, 0}
}

/*
栈空条件：head==tail为true
*/
func (c *CircularQueue) IsEmpty() bool {
	if c.head == c.tail {
		return true
	}
	return false
}

/*
栈满条件：(tail+1)%capacity==head为true
*/
func (c *CircularQueue) IsFull() bool {
	if c.head == (c.tail+1)%c.capacity {
		return true
	}
	return false
}

func (c *CircularQueue) EnQueue(v interface{}) bool {
	if c.IsFull() {
		return false
	}
	c.q[c.tail] = v
	c.tail = (c.tail + 1) % c.capacity
	return true
}

func (c *CircularQueue) DeQueue() interface{} {
	if c.IsEmpty() {
		return nil
	}
	v := c.q[c.head]
	c.head = (c.head + 1) % c.capacity
	return v
}

func (c *CircularQueue) String() string {
	if c.IsEmpty() {
		return "empty queue"
	}
	result := "head"
	var i = c.head
	for true {
		result += fmt.Sprintf("<-%+v", c.q[i])
		i = (i + 1) % c.capacity
		if i == c.tail {
			break
		}
	}
	result += "<-tail"
	return result
}
