package sorting

func Swap(arr []int, x,y int)  {
	arr[x],arr[y] = arr[y],arr[x]
}

//原理 ：i=0设定为最小，然后从i+1开始到后面这段找到一个最小的元素，和i交换
//优点，直观
//缺点，时间复杂度比冒泡略好一点,O(n^2)
func SimpleChoiceSort(arr []int)  {
	length := len(arr)
	for i:=0;i<length-1;i++ {
		min := i
		for j:=i+1;j<length;j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i && arr[min] < arr[i] {
			Swap(arr,i,min)
		}
	}
}