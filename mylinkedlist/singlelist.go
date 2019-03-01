package mylinkedlist

import "fmt"

type LinkNode struct {
	next *LinkNode
	value interface{}
}

func newNode(v interface{}) *LinkNode {
	return &LinkNode{
		next:nil,
		value:v,
	}
}

type LinkList struct {
	head *LinkNode
	length uint32
}

func NewLinkList() *LinkList {
	return &LinkList{
		head:nil,
		length:0,
	}
}

func (this *LinkList)GetLength() uint32 {
	return this.length
}

func (this *LinkList)GetTail() *LinkNode {
	if this.head == nil {
		return nil
	}
	tmp := this.head
	for {
		if tmp.next == nil {
			return tmp
		}
		tmp = tmp.next
	}
}

func (this *LinkList)PushBack(v interface{}) bool {
	tail := this.GetTail()
	if tail == nil {
		this.head = newNode(v)
		this.length += 1
		return true
	}
	newnode := newNode(v)
	tail.next = newnode
	this.length += 1
	return true
}

func (this *LinkList)PushFront(v interface{}) bool {
	if this.head == nil {
		this.head = newNode(v)
		this.length += 1
		return true
	}
	newnode := newNode(v)
	newnode.next = this.head
	this.head = newnode
	this.length += 1
	return true
}

func (this *LinkList)PrintList() {
	if this.length == 0 || this.head == nil {
		return
	}
	tmp := this.head
	for {
		if tmp == nil {
			break
		}
		fmt.Printf("%v ",tmp.value)
		tmp = tmp.next
	}
}

func (this *LinkList)Reverse() {
	if this.length == 0 || this.length ==1 {
		return
	}
	oldhead := this.head
	prev := this.head
	cur := this.head.next
	tmp := cur.next
	for {
		if cur == nil {
			break
		}
		tmp = cur.next
		cur.next = prev
		prev = cur
		cur = tmp
	}
	this.head = prev
	oldhead.next = nil
}