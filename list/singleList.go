/**
单链表
 */
package list

import "fmt"

type Node struct {
	Item interface{}
	Next *Node
}

type SingleList struct {
	Root *Node
	Len int
}

func NewSingleList() *SingleList{
	return new(SingleList)
}

//头部插入
func (list *SingleList) HeadInsert(item interface{})  {
	node := &Node{Item:item}
	if list.Root == nil {
		 list.Root = node
	} else {
		 node.Next = list.Root
		 list.Root = node
	}
	list.Len++
}

//中间插入
func (list *SingleList) MidInsert(item interface{},index int)  {
	node := &Node{Item:item}
	sp := list.Root
	for i:=1 ;i<index;i++ {
		if sp.Next == nil {
			break
		}
		sp = sp.Next
	}
	if sp.Next != nil {
		node.Next = sp.Next
	}
	sp.Next = node
	list.Len++
}

//尾部插入
func (list *SingleList) TailInsert(item interface{})  {
	node := &Node{Item:item}
	if list.Root == nil {
		list.Root = node
	} else {
		sp := list.Root
		for sp.Next != nil {
			sp = sp.Next
		}
		sp.Next = node
	}
	list.Len++
}

//头部删除元素
func (list *SingleList) HeadDelete() interface{} {
	if list.Root == nil {
		return nil
	}
	item := list.Root.Item
	list.Root = list.Root.Next
	list.Len--
	return item
}

//尾部删除元素
func (list *SingleList) TailDelete() interface{} {
	if list.Root == nil {
		return nil
	}
	sp := list.Root
	prev := list.Root
	i := 0
	for sp.Next != nil {
		sp = sp.Next
		if i > 0 {
			prev = prev.Next
		}
		i++
	}
	item := sp.Item
	prev.Next = nil
	return item
}

//中间删除元素
func (list *SingleList) MidDelete(index int) interface{} {
	if list.Root == nil {
		return nil
	}
	sp := list.Root
	for i:=0;i<index-1;i++ {
		sp = sp.Next
	}
	item := sp.Next.Item
	sp.Next = sp.Next.Next
	return item
}

//有序插入
func (list *SingleList) SortedInsert(item interface{})  {
	if list.Root == nil {
		list.Root = &Node{Item:item}
		list.Len++
	} else {
		//如果插入的元素比头部元素小，直接插入头部
		if list.Root.Item.(int) >= item.(int) {
			list.HeadInsert(item)
		} else {
			//从中部i位置插入
			sp := list.Root
			i := 0
			for sp.Item.(int) < item.(int) {
				i++
				sp = sp.Next
				if sp == nil {
					break
				}
			}
			list.MidInsert(item,i)
		}
	}
}

//反转链表
func (list *SingleList) ReverseSingleList() {
	if list.Root == nil {
		return
	}
	prev := list.Root
	sp := list.Root.Next
	prev.Next = nil
	for sp != nil {
		node := sp.Next
		sp.Next = prev
		prev = sp
		sp = node
	}
	list.Root = prev
}

//输出链表
func (list *SingleList) Print()  {
	sp := list.Root
	for sp != nil {
		fmt.Println(sp.Item)
		sp = sp.Next
	}
}