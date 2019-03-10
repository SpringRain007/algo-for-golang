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
	for i := 0; i < 100; i++ {
		rand.NewSource(time.Now().UnixNano())
		a[i] = rand.Intn(1000)
	}
	sort.QuickSort(a[:], len(a))
	ret := BinarySearch(a[:], 90)
	if ret == -1 {
		fmt.Println("not found!")
		return
	}
	fmt.Println(a)
	fmt.Printf("found a[%d]=%d \n", ret, a[ret])
}

func TestBinarySearchFirst(t *testing.T) {
	a := []int{1, 2, 3, 4, 8, 8, 8, 8, 9, 10, 12, 13}
	ret := BinarySearchFirst(a[:], 8)
	if ret == -1 {
		fmt.Println("not found!")
		return
	}
	fmt.Println(a)
	fmt.Printf("found a[%d]=%d \n", ret, a[ret])
}

func TestBinarySearchLast(t *testing.T) {
	a := []int{1, 2, 3, 4, 8, 8, 8, 8, 9, 10, 12, 13}
	ret := BinarySearchLast(a[:], 8)
	if ret == -1 {
		fmt.Println("not found!")
		return
	}
	fmt.Println(a)
	fmt.Printf("found a[%d]=%d \n", ret, a[ret])
}

func TestBinarySearchFirstEx(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 9, 9, 11, 13}
	ret := BinarySearchFirstEx(a[:], 8)
	if ret == -1 {
		fmt.Println("not found!")
		return
	}
	fmt.Println(a)
	fmt.Printf("found a[%d]=%d \n", ret, a[ret])
}
func TestBinarySearchLastEx(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 9, 9, 11, 13}
	ret := BinarySearchLastEx(a[:], 10)
	if ret == -1 {
		fmt.Println("not found!")
		return
	}
	fmt.Println(a)
	fmt.Printf("found a[%d]=%d \n", ret, a[ret])
}
func TestGetSquareRoot(t *testing.T) {
	fmt.Println(GetSquareRoot(101.101))
}
