package main

import (
	"ddyy/go-algorithms/tree"
	"fmt"
)

func LessIntFunc(c1,c2 interface{}) bool {
	return c1.(int) < c2.(int)
}

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9}
	tree1 := tree.CreateBinarySearchTree(arr)
	tree1.LevelTravel()
	tree1.DeleteNode(5)
	fmt.Println()
	tree1.LevelTravel()
}
