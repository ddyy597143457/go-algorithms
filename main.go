package main

import "ddyy/go-algorithms/tree"

func main() {
	tree23 := tree.NewTree23()
	for i:=1; i<=7; i++ {
		tree23.Tree23Insert(i)
	}
	tree23.LevelTravel()
}
