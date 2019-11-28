package searching

import (
	"math"
)

//跳跃查找
//原理：以步长step跳跃，跳啊跳，跳到大于key的数字的前一个区间,目标必定位于该区间或者不存在
func JumpSearch(arr []int,searchValue int) int {
	prev := 0
	sz := len(arr)
	step := int(math.Sqrt(float64(sz)))
	start := step

	//跳跃到目标区间
	for arr[int(math.Min(float64(start),float64(sz))) - 1] < searchValue {
		prev = start
		start += step
		if prev >= sz {
			return -1
		}
	}
	//在目标区间中线性前进
	for arr[prev] < searchValue {
		prev++
		if prev >= sz {
			return -1
		}
	}
	if arr[prev] == searchValue {
		return  prev
	}
	return -1
}
