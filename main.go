package main

import (
	"ddyy/go-algorithms/sorting"
	"fmt"
)

func main() {
	arr := []int{8,6,9,7,1,10,5,4,3,2}
	sorting.HeadSort(arr)
	fmt.Println(arr)
}