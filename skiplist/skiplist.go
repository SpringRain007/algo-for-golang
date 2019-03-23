package skiplist

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

const MAX_LEVEL = 16

type skipListNode struct {
	score int
	value interface{}
	level int
	next []*skipListNode
}

//
func NewNode(score int,v interface{},level int) *skipListNode {
	return &skipListNode{
		score:score,
		value:v,
		level:level,
		next:make([]*skipListNode,level,level),
	}
}

type skipList struct {
	head *skipListNode
	level int
	len int
}

func NewSkipList() *skipList {
	head := NewNode(math.MinInt32,nil,MAX_LEVEL)
	return &skipList{
		head:head,
		level:1,
		len:0,
	}
}

func (this *skipList) Length() int {
	return this.len
}

func (this *skipList) Level() int {
	return this.level
}

//
func (this *skipList) Insert(score int,v interface{}) error {
	if v == nil {
		return errors.New("invalid param!")
	}

	current := this.head
	//此处需要记录每层索引的位置
	update := [MAX_LEVEL]*skipListNode{}
	for i:= MAX_LEVEL-1;i>=0;i-- {
		//从最上层开始寻找
		for current.next[i] != nil {
			if current.next[i].value == v {
				return errors.New("element already exists!")
			}
			if current.next[i].score > score {
				update[i] = current
				break
			}
			current = current.next[i]
		}
		if current.next[i] == nil {
			update[i] = current
		}
	}

	//随机一个层数
	level := 1
	for i:=1;i<MAX_LEVEL;i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	//
	newNode := NewNode(score,v,level)
	for i:=0;i<level;i++ {
		next := update[i].next[i]
		update[i].next[i] = newNode
		newNode.next[i] = next
	}
	if level > this.level {
		this.level = level
	}
	this.len++
	return nil
}

func (this *skipList) Find(score int,v interface{}) *skipListNode{
	if v == nil || this.len == 0 {
		return nil
	}
	current := this.head
	for i:=this.level-1;i>=0;i-- {
		for current.next[i] != nil {
			if current.next[i].score == score && current.next[i].value == v {
				return current.next[i]
			}
			if current.next[i].score > score {
				break
			}
			current = current.next[i]
		}
	}
	return nil
}

func (this *skipList) Delete(score int,v interface{}) error {
	if v == nil {
		return errors.New("invalid param!")
	}
	current := this.head
	//保存前驱节点
	update := [MAX_LEVEL]*skipListNode{}
	for i:=this.level-1;i>=0;i-- {
		update[i] = nil
		for current.next[i] != nil {
			if current.next[i].score == score && current.next[i].value == v {
				update[i] = current
				break
			}
			if current.next[i].score > score {
				break
			}
			current = current.next[i]
		}
	}
	current = update[0].next[0]
	for i:= current.level-1;i>=0;i-- {
		if update[i] == nil {
			continue
		}
		if update[i] == this.head && current.next[i] == nil {
			this.level = i
		}
		if update[i].next[i] == nil {
			update[i].next[i] = nil
		}else {
			update[i].next[i] = update[i].next[i].next[i]
		}
	}
	this.len--
	return nil
}

func (this *skipList) Print() {
	current := this.head.next[0]
	for current != nil {
		fmt.Printf("score[%v],value[%v]\n",current.score,current.value)
		current = current.next[0]
	}
	fmt.Printf("len[%v],level[%d]\n",this.len,this.level)
}