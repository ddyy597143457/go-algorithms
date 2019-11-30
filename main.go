package main

import "ddyy/go-algorithms/list"

func LessIntFunc(c1,c2 interface{}) bool {
	return c1.(int) < c2.(int)
}

func main() {
	l := list.NewDoubleList(LessIntFunc)
	for i:=1;i<=10;i++ {
		l.LRUVisit(i)
	}
	l.LRUVisit(11)
	l.LRUVisit(12)
	l.LRUVisit(9)
	l.HeadPrint()
	l.TailPrint()
}