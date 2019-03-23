package sort

func partition(a []int, p, r int) int {
	pviot := a[r]
	i := p
	for j := p; j <= r-1; j++ {
		if a[j] < pviot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	return i
}
func quickSort_c(a []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(a, p, r)
	quickSort_c(a, p, q-1)
	quickSort_c(a, q+1, r)
}
func QuickSort(a []int, n int) {
	quickSort_c(a, 0, n-1)
}

func partitionRe(a []int,p,r int) int{
	pviot := a[r]
	i := p
	for j:=p;j<=r-1 ;j++  {
		if a[j] > pviot {
			a[i],a[j] = a[j],a[i]
			i++
		}
	}
	a[i],a[r] = a[r],a[i]
	return i
}
func SelectK(a []int,k int) int {
	n := len(a)
	if  n == 0 || k > n || k < 0 {
		panic("k > len(a)")
	}
	return selectK(a,0,n-1,k)
}

func selectK(a []int,p,r,k int) int{
	var ret int
	q := partitionRe(a,p,r)
	if k == q+1 {
		return a[q]
	}else if k > q+1 {
		ret = selectK(a,q+1,r,k)
	}else if k < q+1 {
		ret = selectK(a,p,q-1,k)
	}
	return ret
}
