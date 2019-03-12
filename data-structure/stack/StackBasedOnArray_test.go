package stack

import (
	"fmt"
	"testing"
)

func TestArrayStack_Push(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	t.Log(s.Pop())
	s.Push(3)
	// t.Log(s.Pop())
	// t.Log(s.Pop())
	s.Push(4)
	// t.Log(s.Pop())
	// s.String()
	// t.Log(s)
	fmt.Println(s)
}

func TestArrayStack_Pop(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
}

func TestArrayStack_Top(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
}
