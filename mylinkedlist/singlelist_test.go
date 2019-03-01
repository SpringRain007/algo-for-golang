package mylinkedlist

import (
	"fmt"
	"testing"
)

func TestLinkList_PrintfList(t *testing.T) {
	l := NewLinkList()
	l.PrintList()
	fmt.Println("length: ",l.GetLength())
	fmt.Println()
	l.PushBack(1)
	l.PrintList()
	fmt.Println("length: ",l.GetLength())
	fmt.Println()
	l.PushBack("aaa")
	l.PrintList()
	fmt.Println("length: ",l.GetLength())
	fmt.Println()
	l.PushFront([3]int{1,2,3})
	l.PrintList()
	fmt.Println("length: ",l.GetLength())
}

func TestLinkList_Reverse(t *testing.T) {
	l := NewLinkList()
	l.PushBack(1)
	l.Reverse()
	l.PrintList()
	l.PushBack(2)
	l.Reverse()
	l.PrintList()
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PrintList()
	l.Reverse()
	l.PrintList()
}
