package list

import "testing"

func TestSingleList_BeforeInert(t *testing.T) {
	l := NewSingleList()
	for i:=1;i<=10;i++ {
		l.BeforeInert(i)
	}
	sp := l.Root
	i := 10
	for sp != nil {
		item := sp.Item
		if item != i {
			t.Error("TestSingleList_BeforeInert error")
		}
		sp = sp.Next
		i--
	}
}

func TestSingleList_AfterInsert(t *testing.T) {
	l := NewSingleList()
	for i:=1;i<=10;i++ {
		l.AfterInsert(i)
	}
	sp := l.Root
	i := 1
	for sp != nil {
		item := sp.Item
		if item != i {
			t.Error("TestSingleList_AfterInsert error")
		}
		sp = sp.Next
		i++
	}
}

func TestSingleList_MidInsert(t *testing.T) {
	l := NewSingleList()
	l.BeforeInert(1)
	for i:=2 ;i<=10;i++ {
		l.MidInsert(i,i-1)
	}
	sp := l.Root
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
	l := NewSingleList()
	l.SortedInsert(8)
	l.SortedInsert(3)
	l.SortedInsert(4)
	l.SortedInsert(2)
	l.SortedInsert(5)
	l.SortedInsert(1)
	l.SortedInsert(6)
	l.SortedInsert(7)
	l.SortedInsert(9)
	sp := l.Root
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
	l := NewSingleList()
	for i:=1;i<=10;i++ {
		l.AfterInsert(i)
	}
	l.HeadDelete()
	sp := l.Root
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
	l := NewSingleList()
	for i:=1;i<=10;i++ {
		l.BeforeInert(i)
	}
	l.TailDelete()
	sp := l.Root
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
	l := NewSingleList()
	for i:=1;i<=10;i++ {
		l.AfterInsert(i)
	}
	l.MidDelete(9)
	arr := []int{1,2,3,4,5,6,7,8,9}
	sp := l.Root
	for _,v := range arr {
		item := sp.Item
		if item != v {
			t.Error("TestSingleList_MidDelete error")
		}
		sp = sp.Next
	}
}