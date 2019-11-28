package main

import (
	"ddyy/go-algorithms/list"
)

func LessIntFunc(c1,c2 interface{}) bool {
	return c1.(int) < c2.(int)
}

func main() {
	l := list.NewSingleList(LessIntFunc)
	for i:=1;i<=10;i++ {
		l.TailInsert(i)
	}
	sp := l.Head
	length := 0
	for sp.Item.(int) < 5 {
		sp = sp.Next
		length++
	}
	l1 := list.NewSingleList(LessIntFunc)
	for i:=3;i< 5;i++ {
		l1.TailInsert(i)
	}
	sp1 := l1.Head
	for sp1.Next != nil {
		sp1 = sp1.Next
	}
	sp1.Next = sp
	l1.Count = l1.Count + (l.Count-length)
	l2 := list.FindIntersection(l,l1)
	list.PrintNodes(l2)
}