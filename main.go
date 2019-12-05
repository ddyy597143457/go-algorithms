package main

import "ddyy/go-algorithms/tree"

func main() {
	avl := tree.NewAvlTree()
	avl.AvlAdd(10)
	avl.AvlAdd(5)
	avl.AvlAdd(4)
	avl.AvlAdd(100)
	avl.AvlAdd(6)
	avl.AvlAdd(7)
	avl.LevelTravel()
}
