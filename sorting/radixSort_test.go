package sorting

import "testing"

func TestRadixSort(t *testing.T) {
	arr := []int{29,25,1,49,9,37,21,43,17,14}
	tar := []int{1,9,14,17,21,25,29,37,43,49}
	RadixSort(arr)
	i := 0
	for _,v := range arr {
		if v != tar[i] {
			t.Error("TestRadixSort failed...")
		}
		i++
	}
}