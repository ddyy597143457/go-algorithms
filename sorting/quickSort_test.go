package sorting

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{8,6,9,7,1,10,5,4,3,2}
	InsertionSort(arr)
	i := 1
	for _,v := range arr {
		if v != i {
			t.Error("TestQuickSort failed...")
		}
		i++
	}
}