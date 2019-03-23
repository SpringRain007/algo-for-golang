package sort

func merge(a []int,p,q,r int) {
	n := r-p+1
	tmp := make([]int,n)
	i := p
	j := q+1
	k := 0
	for {
		if i > q || j > r {
			break
		}
		if a[i] <= a[j] {
			tmp[k] = a[i]
			k++
			i++
		}else {
				tmp[k] = a[j]
				k++
				j++
		}
	}
	start := i
	end := q
	if j <= r {
		start = j
		end = r
	}
	for ;start <= end ;start++  {
		tmp[k] = a[start]
		k++
	}
	k = 0
	for x:=p;x<=r;x++ {
		a[x] = tmp[k]
		k++
	}
}

func mergeSort_c(a []int,p,r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	mergeSort_c(a,p,q)
	mergeSort_c(a,q+1,r)
	merge(a,p,q,r)
}

func MergeSort(a []int, n int) {
	mergeSort_c(a,0,n-1)
}
