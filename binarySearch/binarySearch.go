package binarySearch

import (
	"fmt"
	"math"
	"strconv"
)

func BinarySearch(a []int, v int) int {
	low := 0
	high := len(a) - 1
	for {
		if low > high {
			break
		}
		middle := low + (high-low)/2
		if a[middle] == v {
			return middle
		} else if a[middle] > v {
			high = middle - 1
		} else if a[middle] < v {
			low = middle + 1
		}
	}
	return -1
}

//以下为二分法的变种问题:

//1.查找数组第一个等于给定元素的的值
func BinarySearchFirst(a []int, v int) int {
	n := len(a)
	low := 0
	high := n - 1
	for {
		if low > high {
			break
		}
		mid := low + (high-low)>>1
		if a[mid] > v {
			high = mid - 1
		} else if a[mid] < v {
			low = mid + 1
		} else {
			if a[mid] == 0 || a[mid-1] != v {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

//2.查找数组最后一个等于给定元素的的值
func BinarySearchLast(a []int, v int) int {
	n := len(a)
	low := 0
	high := n - 1
	for {
		if low > high {
			break
		}
		mid := low + (high-low)>>1
		if a[mid] > v {
			high = mid - 1
		} else if a[mid] < v {
			low = mid + 1
		} else {
			if mid == n-1 || a[mid+1] != v {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

//3.查找数组第一个大于等于给定元素的的值
func BinarySearchFirstEx(a []int, v int) int {
	n := len(a)
	low := 0
	high := n - 1
	for {
		if low > high {
			break
		}
		mid := low + (high-low)>>1
		if a[mid] >= v {
			if mid == 0 || a[mid-1] < v {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

//4.查找数组最后一个小于等于给定元素的的值
func BinarySearchLastEx(a []int, v int) int {
	n := len(a)
	low := 0
	high := n - 1
	for {
		if low > high {
			break
		}
		mid := low + (high-low)>>1
		if a[mid] <= v {
			if mid == n-1 || a[mid+1] > v {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}

func GetSquareRoot(s float64) float64 {
	if s <= 0 {
		return 0
	}
	var ret float64
	i := 0.0
	for {
		if i*i == s {
			ret = i
			goto transform
		}
		if i*i > s {
			break
		}
		i++
	}
	ret = getnumber(i-1, s, 0, 1)
transform:
	msgret := fmt.Sprintf("%.06f", ret)
	ret, err := strconv.ParseFloat(msgret, 64)
	if err != nil {
		fmt.Println("parse string to float64 error!")
		return 0
	}

	return ret
}

//精度保证在0.0000001
func getnumber(base, s, start, end float64) float64 {
	var ret float64
	mid := start + (end-start)/2
	num := base + mid
	sum := num * num
	if sum == s || math.Abs(s-sum) < 0.000001 {
		return num
	} else if sum > s {
		ret = getnumber(base, s, start, mid)
	} else if sum < s {
		ret = getnumber(base, s, mid, end)
	}
	return ret
}
