package skiplist

import (
	"fmt"
	"github.com/luomingshun/algo-for-golang/public"
	"math/rand"
	"testing"
	"time"
)

func TestSkipList_Print(t *testing.T) {
	sl := NewSkipList()
	sl.Insert(23,"hello")
	rand.NewSource(time.Now().UnixNano())
	for i:=0;i<20;i++ {
		sl.Insert(rand.Intn(50),public.GetRandomStr(5))
	}
	sl.Print()
	node := sl.Find(23,"hello")
	if node != nil {
		fmt.Println("find the one!")
	}
	sl.Delete(23,"hello")
	sl.Print()

}
