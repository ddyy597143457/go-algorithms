package main

import (
	"ddyy/go-algorithms/tree"
	"fmt"
)

func main() {
	arr := []int{6,8,5,9,7,4,3}
	fmt.Println(arr)
	tree.HeapSort(arr)
	fmt.Println(arr)
}
