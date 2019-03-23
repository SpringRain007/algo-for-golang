package myqueue

import "fmt"

type Myqueue struct {
	array []interface{}
	size  int
	head  int
	tail  int
}

func NewQueue(n int) *Myqueue {
	return &Myqueue{
		array: make([]interface{}, n),
		size:  n,
		head:  0,
		tail:  0,
	}
}

func (this *Myqueue) Enqueue(v interface{}) bool {
	if this.size == 0 {
		fmt.Println("size of queue is 0!")
		return false
	}
	if this.tail == this.size {
		//fmt.Println("queue is full!")
		this.array = append(this.array, v)
		this.tail++
		return true
	}
	this.tail++
	this.array[this.tail-1] = v
	return true
}

func (this *Myqueue) Dequeue() interface{} {
	if this.head > this.tail {
		fmt.Println("the queue is empty!")
		return nil
	}
	tmp := this.array[this.head]
	this.array[this.head] = nil
	this.head++
	return tmp
}

func (this *Myqueue) Print() {
	for _, v := range this.array {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
