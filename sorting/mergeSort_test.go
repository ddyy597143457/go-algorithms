package sorting

import "testing"

func TestMergeSort(t *testing.T) {
	arr := []int{8,6,9,7,1,10,5,4,3,2}
	temp := make([]int,10)
	MergeSort(arr,temp,0,len(arr) - 1)
	i := 1
	for _,v := range arr {
		if v != i {
			t.Error("TestMergeSort failed...")
		}
		i++
	}
}