package sort

/*
逆序度 = 满有序度 - 有序度，即需要交换的次数
满有序度 = 完全有序的数组的有序对数（n*(n-1)/2）
有序度 = 初始状态数组的有序对的个数

*/



//冒泡排序，原地排序，稳定，最好情况O(n),最坏情况O(n^2)，平均时间复杂度O(n^2)
func BubbleSort(a []int,n int) {
	if n <= 1 {
		return ;
	}
	for i:=0;i<n;i++ {
		flag := false
		for j:=0;j<n-i-1 ;j++  {
			if a[j] > a[j+1] {
				a[j],a[j+1] = a[j+1],a[j]
				flag = true
			}
		}
		if !flag { //上一次冒泡没有发生数据交换，直接退出
			break
		}
	}
}