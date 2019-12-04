package main

import (
	"ddyy/go-algorithms/tree"
)

func main() {
	avl := tree.NewAvlTree()
	avl.AvlAdd(100)
	avl.AvlAdd(90)
	avl.AvlAdd(8)
	avl.AvlAdd(110)
	avl.AvlAdd(120)
	avl.AvlAdd(5)
	avl.AvlAdd(4)
	avl.LevelTravel()
	avl.AvlAdd(99)
	avl.LevelTravel()
}
