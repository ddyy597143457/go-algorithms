package sorting

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{8,6,9,7,1,5,4,3,2}
	BubbleSort(arr)
	i := 1
	for _,v := range arr {
		if v != i {
			t.Error("TestBubbleSort failed...")
		}
		i++
	}
}
