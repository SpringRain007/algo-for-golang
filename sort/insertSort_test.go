package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInsertSort(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())
	var a []int
	for i := 0; i < 20; i++ {
		a = append(a, rand.Intn(1000))
	}
	fmt.Println(a)
	InsertSort(a, len(a))
	fmt.Println(a)
}

func BenchmarkInsertSort(b *testing.B) {
	rand.NewSource(time.Now().UnixNano())
	var a []int
	for i := 0; i < b.N; i++ {
		a = append(a, rand.Intn(100000))
	}
	//fmt.Println(a)
	InsertSort(a, len(a))
	//fmt.Println(a)
}
