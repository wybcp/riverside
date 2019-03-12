package simplylinkedlist

import (
	"fmt"
	"testing"
)

func TestList_ReversedList(t *testing.T) {
	var list = List{}
	for i := 0; i < 5; i++ {
		list.Add(Node{data: i})
	}
	fmt.Println(list.String())
	list.ReversedList()
	fmt.Println("reverse list：", list.String())

}
