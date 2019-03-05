package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	var a []int
	rand.NewSource(time.Now().UnixNano())
	for i:=0;i<20 ;i++  {
		a = append(a,rand.Intn(100))
	}
	QuickSort(a,len(a))
	fmt.Println(a)
}

func BenchmarkQuickSort(b *testing.B) {
	rand.NewSource(time.Now().UnixNano())
	var a []int
	for i:=0;i<b.N ;i++  {
		a = append(a,rand.Intn(100000))
	}
	//fmt.Println(a)
	QuickSort(a,len(a))
	//fmt.Println(a)
}