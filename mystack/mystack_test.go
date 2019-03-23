package mystack

import (
	"fmt"
	"testing"
)

func TestMystack_Push(t *testing.T) {
	s, ok := NewStack(0)
	if !ok {
		fmt.Println("bad n number!")
		return
	}
	fmt.Println(s.Size(), s.Count())
	s.Push("a")
	fmt.Println(s.Size(), s.Count())
	s.Push("b")
	fmt.Println(s.Size(), s.Count())
	s.Push("c")
	fmt.Println(s.Size(), s.Count())
	s.Push("d")
	fmt.Println(s.Size(), s.Count())
	fmt.Println(s.Pop())
	fmt.Println(s.Size(), s.Count())
	fmt.Println(s.Pop())
	fmt.Println(s.Size(), s.Count())
	fmt.Println(s.Pop())
	fmt.Println(s.Size(), s.Count())
	fmt.Println(s.Pop())
	fmt.Println(s.Size(), s.Count())
}
