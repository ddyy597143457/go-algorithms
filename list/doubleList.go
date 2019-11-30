package list

import "fmt"

const LRULENGTH = 10

//双向链表
type doubleNode struct {
	item interface{}
	next *doubleNode
	prev *doubleNode
}
type doubleList struct {
	head *doubleNode
	tail *doubleNode
	size int
	lessFunc Comparable
}

func NewDoubleList(comparable Comparable) *doubleList {
	return &doubleList{
		lessFunc:comparable,
	}
}

func (list *doubleList)HeadInsert(item interface{})  {
	node := &doubleNode{item:item}
	if list.size == 0 {
		list.head = node
		list.tail = list.head
		list.size++
		return
	}
	list.head.prev = node
	node.next = list.head
	list.head = node
	list.size++
	return
}

func (list *doubleList)TailInsert(item interface{})  {
	node := &doubleNode{item:item}
	if list.size == 0 {
		list.head = node
		list.tail = list.head
		list.size++
		return
	}
	node.prev = list.tail
	list.tail.next = node
	list.tail = node
	list.size++
}

func (list *doubleList)HeadDelete() interface{} {
	if list.size == 0 {
		return nil
	}
	if list.size == 1 {
		it := list.head.item
		list.tail = nil
		list.head = nil
		list.size--
		return it
	}
	it := list.head.item
	list.head = list.head.next
	list.head.prev = nil
	list.size--
	return it
}

func (list *doubleList)TailDelete() interface{}{
	if list.size == 0 {
		return nil
	}
	if list.size == 1 {
		it := list.head.item
		list.tail = nil
		list.head = nil
		list.size--
		return it
	}
	it := list.tail
	list.tail = list.tail.prev
	list.tail.next = nil
	list.size--
	return it
}

func (list *doubleList)HeadPrint()  {
	if list.size == 0 {
		return
	}
	sp := list.head
	for sp != nil {
		fmt.Print(sp.item," ")
		sp = sp.next
	}
	fmt.Println()
}

func (list *doubleList)TailPrint()  {
	if list.size == 0 {
		return
	}
	sp := list.tail
	for sp != nil {
		fmt.Print(sp.item," ")
		sp = sp.prev
	}
	fmt.Println()
}

//用链表来实现LRU缓存淘汰策略
//思路：维护一个有序单链表，越靠近链表尾部的结点是越早之前访问的。当有一个新的数据被访问时，我们从链表头部开
//顺序遍历链表。
//如果此数据之前已经被缓存在链表中了，我们遍历得到这个数据的对应结点，并将其从原来的位置删除，并插入到链表头部。
//如果此数据没在缓存链表中，又可以分为两种情况处理：
//如果此时缓存未满，可直接在链表头部插入新节点存储此数据；
//如果此时缓存已满，则删除链表尾部节点，再在链表头部插入新节点。

func (list *doubleList)LRUVisit(item interface{})  {
	if list.size <= 1 {
		list.HeadInsert(item)
		return
	}
	var found bool = false
	sp := list.head
	for sp.next != nil {
		if sp.item == item {
			found = true
			break
		}
		sp = sp.next
	}
	//item已经存在了，把它移到链表头部
	if found {
		if sp == list.head {
			return
		}
		sp.next.prev = sp.prev
		sp.prev.next = sp.next
		list.HeadInsert(item)
		return
	} else {
		if list.size >=  LRULENGTH {
			list.TailDelete()
			list.HeadInsert(item)
		} else {
			list.HeadInsert(item)
		}
	}
}
