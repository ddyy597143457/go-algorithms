package searching

//二分查找
func BinarySearch(arr []int,searchValue int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left+right)/2
		if arr[mid] ==  searchValue {
			return mid
		}
		if arr[mid] < searchValue {
			left = mid+1
		} else {
			right = mid-1
		}
	}
	return -1
}