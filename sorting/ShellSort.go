package sorting

//原理：直接插入排序的改进，每次把数组分为若干个区间，每个区间进行直接插入排序
//优点：第一个突破O(n^2)的算法，具体时间复杂度难以衡量取决于gap，最好约为O(n^1.3)
//缺点：
func ShellSort(arr []int)  {
	length := len(arr)
	gap := length/2
	for gap > 0 {
		for i:=gap;i<length;i++ {
			 t := arr[i]
			 j := i-gap
			 for j>=0 && t<arr[j] {
			 	arr[j+gap] = arr[j]
			 	j -= gap
			 }
			arr[j+gap] = t
		}
		gap /= 2
	}
}