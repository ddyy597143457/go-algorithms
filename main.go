package main

import (
	"ddyy/go-algorithms/sorting"
	"fmt"
)

func main() {
	arr := []int{29,25,1,49,9,37,21,43,17,14}
	sorting.RadixSort(arr)
	fmt.Println(arr)
}