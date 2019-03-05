package sort

func SelectSort(a []int,n int) {
	if n <= 1 {
		return
	}
	for i:=0;i<n-1 ;i++  {
		min := a[i]
		index := i
		for j:=i+1;j<n ;j++  {
			if a[j] < min {
				min = a[j]
				index = j
			}
		}
		a[i],a[index] = a[index],a[i]
	}
}
