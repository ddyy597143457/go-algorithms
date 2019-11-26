package sorting

//原理：以哨兵为基准，两边交换数据
//优点：时间复杂度最好
//缺点：不稳定
func QuickSort(arr []int,left,right int)  {
	if left >= right {
		return
	}
	i, j := left, right
	guard := arr[i]	//哨兵
	for i < j {
		for i < j && arr[j] > guard {
			j--
		}
		arr[i] = arr[j]
		for i < j && arr[i] < guard {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = guard
	QuickSort(arr,left,i)
	QuickSort(arr,i+1,right)
}