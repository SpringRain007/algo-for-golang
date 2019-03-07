package binarySearch

import (
	"fmt"
	"github.com/luomingshun/algo-for-golang/sort"
	"math/rand"
	"testing"
	"time"
)

func TestBinarySearch(t *testing.T) {
	//rand.NewSource(time.Now().UnixNano())
	var a [100]int
	for i:=0;i<100 ;i++  {
		rand.NewSource(time.Now().UnixNano())
		a[i] = rand.Intn(1000)
	}
	sort.QuickSort(a[:],len(a))
	ret := BinarySearch(a[:],90)
	if ret == -1 {
		fmt.Println("not found!")
		return
	}
	fmt.Println(a)
	fmt.Printf("found a[%d]=%d \n",ret,a[ret])
}

func TestGetSquareRoot(t *testing.T) {
	fmt.Println(GetSquareRoot(1.11111))
}
