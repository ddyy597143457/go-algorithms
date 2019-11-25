package sorting

//原理：i=0，再用一个遍历从j=0开始j=n-i-1结束，比较j和j+1，将最大的数放到最后
//优点：简单
//缺点：时间复杂度最差,O(n^2)
func BubbleSort(arr[]int)  {
	length := len(arr)
	for i:=0;i<length-1;i++ {
		for j:=0;j<length-i-1;j++ {
			if arr[j] > arr[j+1] {
				Swap(arr,j,j+1)
			}
		}
	}
}
