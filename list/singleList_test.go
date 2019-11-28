package list

import "testing"

func LessIntFunc(c1,c2 interface{}) bool {
	return c1.(int) < c2.(int)
}

func TestSingleList_HeadInsert(t *testing.T) {
	l := NewSingleList(LessIntFunc)
	for i:=1;i<=10;i++ {
		l.HeadInsert(i)
	}
	sp := l.Head
	i := 10
	for sp != nil {
		item := sp.Item
		if item != i {
			t.Error("TestSingleList_HeadInsert error")
		}
		sp = sp.Next
		i--
	}
}

func TestSingleList_TailInsert(t *testing.T) {
	l := NewSingleList(LessIntFunc)
	for i:=1;i<=10;i++ {
		l.TailInsert(i)
	}
	sp := l.Head
	i := 1
	for sp != nil {
		item := sp.Item
		if item != i {
			t.Error("TestSingleList_TailInsert error")
		}
		sp = sp.Next
		i++
	}
}

func TestSingleList_MidInsert(t *testing.T) {
	l := NewSingleList(LessIntFunc)
	l.HeadInsert(1)
	for i:=2 ;i<=10;i++ {
		l.MidInsert(i,i-1)
	}
	sp := l.Head
	i := 1
	for sp != nil {
		item := sp.Item
		if item != i {
			t.Error("TestSingleList_MidInsert error")
		}
		sp = sp.Next
		i++
	}
}

func TestSingleList_SortedInsert(t *testing.T) {
	l := NewSingleList(LessIntFunc)
	l.SortedInsert(8)
	l.SortedInsert(3)
	l.SortedInsert(4)
	l.SortedInsert(2)
	l.SortedInsert(5)
	l.SortedInsert(1)
	l.SortedInsert(6)
	l.SortedInsert(7)
	l.SortedInsert(9)
	sp := l.Head
	i := 1
	for sp != nil {
		item := sp.Item
		if item != i {
			t.Error("TestSingleList_SortedInsert error")
		}
		sp = sp.Next
		i++
	}
}

func TestSingleList_HeadDelete(t *testing.T) {
	l := NewSingleList(LessIntFunc)
	for i:=1;i<=10;i++ {
		l.TailInsert(i)
	}
	l.HeadDelete()
	sp := l.Head
	i := 2
	for sp != nil {
		item := sp.Item
		if item != i {
			t.Error("TestSingleList_HeadDelete error")
		}
		sp = sp.Next
		i++
	}
	if i != 11 {
		t.Error("TestSingleList_HeadDelete error")
	}
}

func TestSingleList_TailDelete(t *testing.T) {
	l := NewSingleList(LessIntFunc)
	for i:=1;i<=10;i++ {
		l.HeadInsert(i)
	}
	l.TailDelete()
	sp := l.Head
	i := 10
	for sp != nil {
		item := sp.Item
		if item != i {
			t.Error("TestSingleList_TailDelete error")
		}
		sp = sp.Next
		i--
	}
	if i != 1 {
		t.Error("TestSingleList_TailDelete error")
	}
}

func TestSingleList_MidDelete(t *testing.T) {
	l := NewSingleList(LessIntFunc)
	for i:=1;i<=10;i++ {
		l.TailInsert(i)
	}
	l.MidDelete(9)
	arr := []int{1,2,3,4,5,6,7,8,9}
	sp := l.Head
	for _,v := range arr {
		item := sp.Item
		if item != v {
			t.Error("TestSingleList_MidDelete error")
		}
		sp = sp.Next
	}
}

func TestSingleList_DeleteNode(t *testing.T) {
	l := NewSingleList(LessIntFunc)
	for i:=1;i<=10;i++ {
		l.TailInsert(i)
	}
	l.DeleteNode(1)
	sp := l.Head
	i := 2
	for sp != nil {
		item := sp.Item
		if item != i {
			t.Error("TestSingleList_DeleteNode error")
		}
		sp = sp.Next
		i++
	}
}

func TestSingleList_DeleteNodes(t *testing.T) {
	l := NewSingleList(LessIntFunc)
	for i:=1;i<=10;i++ {
		l.TailInsert(i)
	}
	l.TailInsert(1)
	l.HeadInsert(1)
	l.DeleteNodes(1)
	sp := l.Head
	i := 2
	for sp != nil {
		item := sp.Item
		if item != i {
			t.Error("TestSingleList_DeleteNodes error")
		}
		sp = sp.Next
		i++
	}
}

func TestSingleList_RemoveDuplicate(t *testing.T) {
	l := NewSingleList(LessIntFunc)
	for i:=1;i<=10;i++ {
		l.SortedInsert(i)
	}
	l.SortedInsert(8)
	l.SortedInsert(8)
	l.SortedInsert(9)
	l.RemoveDuplicate()
	sp := l.Head
	i := 1
	for sp != nil {
		item := sp.Item
		if item != i {
			t.Error("TestSingleList_RemoveDuplicate error")
		}
		sp = sp.Next
		i++
	}
}

func TestSingleList_CopySingleListReverse(t *testing.T) {
	l := NewSingleList(LessIntFunc)
	for i:=1;i<=10;i++ {
		l.TailInsert(i)
	}
	l2 := l.CopySingleListReverse()
	sp := l2.Head
	i := 10
	for sp != nil {
		item := sp.Item
		if item != i {
			t.Error("TestSingleList_CopySingleListReverse error")
		}
		sp = sp.Next
		i--
	}
}

func TestSingleList_ReverseSingleList(t *testing.T) {
	l := NewSingleList(LessIntFunc)
	for i:=1;i<=10;i++ {
		l.TailInsert(i)
	}
	l.ReverseSingleList()
	sp := l.Head
	i := 10
	for sp != nil {
		item := sp.Item
		if item != i {
			t.Error("TestSingleList_ReverseSingleList error")
		}
		sp = sp.Next
		i--
	}
}

func TestFindIntersection(t *testing.T) {
	l := NewSingleList(LessIntFunc)
	for i:=1;i<=10;i++ {
		l.TailInsert(i)
	}
	sp := l.Head
	length := 0
	for sp.Item.(int) < 5 {
		sp = sp.Next
		length++
	}
	l1 := NewSingleList(LessIntFunc)
	for i:=3;i< 5;i++ {
		l1.TailInsert(i)
	}
	sp1 := l1.Head
	for sp1.Next != nil {
		sp1 = sp1.Next
	}
	sp1.Next = sp
	l1.Count = l1.Count + (l.Count-length)
	l2 := FindIntersection(l,l1)
	i := 5
	for l2 != nil {
		if l2.Item != i {
			t.Error("TestFindIntersection error")
		}
		l2 = l2.Next
		i++
	}
}