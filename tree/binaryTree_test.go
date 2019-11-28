package tree

import "testing"

func compareInt(x interface{}, y interface{}) bool {
	return x.(int) < y.(int)
}

func TestBinaryTree_Search(t *testing.T) {
	binarytree := New(compareInt)
	var intarray = []int{5,10,3,7,2,11,4}
	for _,v := range intarray {
		binarytree.Insert(v)
	}

	findNode := binarytree.Search(7)
	if findNode.Node != 7 {
		t.Error("Search error")
	}

	findNilNode := binarytree.Search(70)
	if findNilNode != nil {
		t.Error("Search nil error")
	}
}

func TestBinaryTree_MaxMin(t *testing.T) {
	binarytree := New(compareInt)
	var intarray = []int{5,10,3,7,2,11,4}
	for _,v := range intarray {
		binarytree.Insert(v)
	}
	max := binarytree.Max()
	if max != 11 {
		t.Error("Max error")
	}

	min := binarytree.Min()
	if min != 2 {
		t.Error("min error")
	}
}