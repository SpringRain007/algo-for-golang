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
