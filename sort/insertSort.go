package sort


//插入排序，原地排序，稳定，最好O(n),最坏O(n^2)，平均O(n^2)
func InsertSort(a []int,n int) {
	if n <= 1 {
		return
	}
	for i:=1;i<n ;i++  {
		value := a[i]
		j := i-1
		for ;j>=0;j--  {
			if a[j] > value {
				a[j+1] = a[j]
			}else {
				break
			}
		}
		a[j+1] = value
	}

}
