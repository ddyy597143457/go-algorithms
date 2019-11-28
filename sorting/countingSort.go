package sorting

//计数排序
//非比较算法
//原理：假设待排序数组arr，最大值为maxValue，则利用一个长度为maxValue+1的辅助数组，它的key是数组arr的值
//填充完毕后直接输出就是有序的了
func CountingSort(arr []int)  {
	maxValue := MaxInt(arr)
	bucketlen := maxValue+1
	bucket := make([]int,bucketlen)	//构造长度bucketlen的slice
	length := len(arr)
	startIndex := 0
	for i:=0;i<length;i++ {
		bucket[arr[i]] += 1	//数组的值填充bucket的key
	}
	for j:=0;j<bucketlen;j++ {
		//bucket[j] > 0，有值，for是因为可能有些值是重复的
		for bucket[j] > 0 {
			arr[startIndex] = j
			startIndex++
			bucket[j] -= 1
		}
	}
}