package searching

import (
	"math"
)

func JumpSearch(arr []int,key int) int {
	prev := 0
	sz := len(arr)
	step := int(math.Sqrt(float64(sz)))
	start := step

	//find the block
	for arr[int(math.Min(float64(start),float64(sz))) - 1] < key {
		prev = start
		start += step
		if prev >= sz {
			return -1
		}
	}

	for arr[prev] < key {
		prev++
		if prev >= sz {
			return -1
		}
	}
	if arr[prev] == key {
		return  prev
	}
	return -1
}
