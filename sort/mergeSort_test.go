package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	var a []int
	rand.NewSource(time.Now().UnixNano())
	for i:=0;i<20 ;i++  {
		a = append(a,rand.Intn(100))
	}
	MergeSort(a,len(a))
	fmt.Println(a)
}

func BenchmarkMergeSort(b *testing.B) {
	rand.NewSource(time.Now().UnixNano())
	var a []int
	for i:=0;i<b.N ;i++  {
		a = append(a,rand.Intn(100000))
	}
	//fmt.Println(a)
	MergeSort(a,len(a))
	//fmt.Println(a)
}
