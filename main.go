package main

import (
	"fmt"
	"github.com/luomingshun/algo-for-golang/sort"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	var a [100000]int
	for i:=0;i<100000 ;i++  {
		a[i] = rand.Intn(100000)
	}
	b := a
	c := a
	d := a
	start := time.Now()
	sort.BubbleSort(a[:],len(a))
	fmt.Printf("bubble sort cost %q\n",time.Now().Sub(start))
	start1 := time.Now()
	sort.InsertSort(b[:],len(b))
	fmt.Printf("insert sort cost %q\n",time.Now().Sub(start1))
	start2 := time.Now()
	sort.SelectSort(c[:],len(c))
	fmt.Printf("select sort cost %q\n",time.Now().Sub(start2))
	start3 := time.Now()
	sort.QuickSort(d[:],len(d))
	fmt.Printf("quick sort cost %q\n",time.Now().Sub(start3))
}
