package doublylinkedlist

import (
	"fmt"
	"testing"
)

func TestDoublyLinkedList_InsertFront(t *testing.T) {
	list := New(10)
	list.InsertFront(2)
	list.InsertFront(23)
	list.InsertFront(3)
	fmt.Println(list.String())
	// t.Log(list)
}
