package sorting

func MaxInt(arr []int) int {
	maxint := arr[0]
	for i:=0; i<len(arr);i++ {
		if maxint < arr[i] {
			maxint = arr[i]
		}
	}
	return maxint
}

func MinInt(arr []int) int {
	minint := arr[0]
	for i:=0; i<len(arr);i++ {
		if minint > arr[i] {
			minint = arr[i]
		}
	}
	return minint
}
