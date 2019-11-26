package sorting

//原理：该算法是采用分治法（Divide and Conquer）的一个非常典型的应用
//优点：两个字，稳定！
//缺点：空间复杂度较差

//治
func Merge(arr,temp []int,left,mid,right int)  {
	i,j := left,mid+1
	k := left
	for i < mid+1 && j < right+1 {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	for i < mid + 1 {
		temp[k] = arr[i]
		i++
		k++
	}
	for j < right+1 {
		temp[k] = arr[j]
		j++
		k++
	}
	for i:=left; i<= right; i++ {
		arr[i] = temp[i]
	}
}
//分
func MergeSort(arr []int,temp []int,left,right int)  {
	if left >= right {
		return
	}
	mid := left+(right-left)/2
	MergeSort(arr,temp,left,mid)
	MergeSort(arr,temp,mid+1,right)
	Merge(arr,temp,left,mid,right)
}