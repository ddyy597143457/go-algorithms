package sorting

import "testing"

func TestInsertionSort(t *testing.T) {
	arr := []int{8,6,9,7,1,5,4,3,2}
	InsertionSort(arr)
	i := 1
	for _,v := range arr {
		if v != i {
			t.Error("TestInsertionSort failed...")
		}
		i++
	}
}