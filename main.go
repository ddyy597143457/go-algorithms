package main

import (
	"ddyy/go-algorithms/tree"
	"fmt"
)

func LessIntFunc(c1,c2 interface{}) bool {
	return c1.(int) < c2.(int)
}

func main() {
	arr := []int{4,5,6,7,8,9,10,11}
	tree1 := tree.LevelOrderBinaryTree(arr)
	tree1.PreOrder()
	fmt.Println()
	tree1.PrintAllPath()
}
