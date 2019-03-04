package lru

import "fmt"

type DataNode struct {
	next *DataNode
	key string
	value interface{}
}

type LruCache struct {
	head *DataNode
	length int
	size int
}
//
func NewCache(len int) (*LruCache,bool) {
	if len == 0 {
		return nil,false
	}
	return &LruCache{
		head:nil,
		length:0,
		size:len,
	},true
}

func newDataNode(k string,v interface{}) *DataNode {
	return &DataNode{
		next:nil,
		key:k,
		value:v,
	}
}

func (this *LruCache)Set(k string,v interface{}) {
 	if this.length <= 1{
		newhead := newDataNode(k,v)
		newhead.next = this.head
		this.head = newhead
		this.length++
	} else if this.length == this.size {
		//find tail
		tmp := this.head
		for {
			if tmp.next.next == nil {
				break
			}
			tmp = tmp.next
		}
		tmp.next = tmp.next.next
		newhead := newDataNode(k,v)
		newhead.next = this.head
		this.head = newhead
	} else {
		newhead := newDataNode(k,v)
		newhead.next = this.head
		this.head = newhead
		this.length++
	}
}

func (this *LruCache)Get(k string) (interface{},bool) {
	if this.length == 0 {
		return nil,false
	}
	if this.head.key == k {
		return this.head.value,true
	}
	prev := this.head
	tmp := prev.next
	for {
		if tmp == nil {
			goto case1
		}
		if tmp.key == k {
			goto case2
		}
		prev = tmp
		tmp = tmp.next
	}
	case1:
		return nil,false
	case2:
		newhead := tmp
		prev.next = tmp.next
		newhead.next = this.head
		this.head = newhead
		return this.head.value,true
}

func (this *LruCache)Print() {
	tmp := this.head
	for {
		if tmp == nil {
			break
		}
		fmt.Printf("k[%q],v[%q],len[%d]\n",tmp.key,tmp.value,this.length)
		tmp = tmp.next
	}
}