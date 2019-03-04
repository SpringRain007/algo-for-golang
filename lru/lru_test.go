package lru

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLruCache_Print(t *testing.T) {
	cache,_ := NewCache(10)
	i := 0
	for ;i<10;i++ {
		cache.Set(strconv.Itoa(i),strconv.Itoa(i*i))
	}
	cache.Print()
	v,ok := cache.Get("1")
	if ok {
		fmt.Printf("%v\n",v)
	}
	cache.Print()

	v,ok = cache.Get("5")
	if ok {
		fmt.Printf("%v\n",v)
	}
	cache.Print()
	v,ok = cache.Get("10")
	if ok {
		fmt.Printf("%v\n",v)
	}else {
		fmt.Printf("not found !\n")
		cache.Set("10",10*10)
	}
	cache.Print()
}
