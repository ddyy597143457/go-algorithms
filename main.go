package main

import "ddyy/go-algorithms/list"

func main() {
	l := list.NewSingleList()
	l.SortedInsert(8)
	l.SortedInsert(3)
	l.SortedInsert(4)
	l.SortedInsert(2)
	l.SortedInsert(5)
	l.SortedInsert(1)
	l.SortedInsert(6)
	l.SortedInsert(7)
	l.SortedInsert(9)
	l.Print()
}