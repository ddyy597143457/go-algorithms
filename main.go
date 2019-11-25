package main

import (
	"ddyy/go-algorithms/list"
)

func main() {
	l := list.NewSingleList()
	for i:=1;i<=10;i++ {
		l.TailInsert(i)
	}
	l.Print()
	l.ReverseSingleList()
	l.Print()
}