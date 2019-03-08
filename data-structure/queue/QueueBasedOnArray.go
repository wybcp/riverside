package _9_queue

import "fmt"

type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}
//有问题
func (a *ArrayQueue) EnQueue(v interface{}) bool {
	if a.tail == a.capacity {
		return false
	}
	a.q[a.tail] = v
	a.tail++
	return true
}

func (a *ArrayQueue) DeQueue() interface{} {
	if a.head == a.tail {
		return nil
	}
	v := a.q[a.head]
	a.head++
	return v
}

func (a *ArrayQueue) String() string {
	if a.head == a.tail {
		return "empty queue"
	}
	result := "head"
	for i := a.head; i <= a.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", a.q[i])
	}
	result += "<-tail"
	return result
}
