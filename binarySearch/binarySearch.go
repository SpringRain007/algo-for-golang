package binarySearch

func BinarySearch(a []int,v int) int{
	low := 0
	high := len(a)-1
	for {
		if low > high {
			break
		}
		middle := low + (high-low)/2
		if a[middle] == v {
			return middle
		}else if a[middle] > v {
			high = middle-1
		}else if a[middle] < v {
			low = middle + 1
		}
	}
	return -1
}
