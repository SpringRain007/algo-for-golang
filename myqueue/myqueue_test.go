package myqueue

import "testing"

func TestMyqueue_Enqueue(t *testing.T) {
	q := NewQueue(5)
	q.Enqueue("a")
	q.Print()
	q.Enqueue("b")
	q.Print()
	q.Enqueue("c")
	q.Print()
	q.Enqueue("d")
	q.Print()
	q.Enqueue("e")
	q.Print()
	q.Enqueue("123")
	q.Print()
}

func TestMyqueue_Dequeue(t *testing.T) {
	q := NewQueue(5)
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")
	q.Dequeue()
	q.Print()
	q.Enqueue("d")
	q.Enqueue("e")
	q.Dequeue()
	q.Print()
	q.Enqueue("123")
	q.Dequeue()
	q.Print()
}
