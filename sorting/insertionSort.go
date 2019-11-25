package sorting

//原理：i=1开始,j从i-1开始,往前遍历，比自己大的元素均往后挪，然后自己插入
//优点：简单
//缺点：时间复杂度比简单选择好点,O(n^2)
//直接插入可以改进为了折半插入法，但是只是比较次数少了，移动没少，复杂度依然为O(n^2)
func InsertionSort(arr []int)  {
	length := len(arr)
	for i:=1;i<length;i++ {
		t := arr[i]
		j := i-1
		for j >= 0 && t < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = t
	}
}