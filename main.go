package main

import "ddyy/go-algorithms/tree"

func main() {
	avl := tree.NewAvlTree()
	avl.AvlAdd(100)
	avl.AvlAdd(90)
	avl.AvlAdd(8)
	avl.AvlAdd(110)
	avl.AvlAdd(120)
	avl.AvlAdd(7)
	avl.AvlAdd(130)
	avl.AvlRemove(7)
	avl.AvlAdd(140)
	avl.LevelTravel()
}
