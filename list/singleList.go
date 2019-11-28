/**
单链表
 */
package list

import (
	"fmt"
)

//节点值比较函数
type Comparable func(c1 interface{},c2 interface{}) bool

type Node struct {
	Item interface{}
	Next *Node
}

type SingleList struct {
	Head *Node
	Count int
	LessFunc Comparable
}

func NewSingleList(comparableFunc Comparable) *SingleList{
	return &SingleList{
		LessFunc:comparableFunc,
	}
}

//链表长度
func (list *SingleList)Size() int {
	return list.Count
}

func (list *SingleList)IsEmpty() bool {
	return list.Count == 0
}

func (list *SingleList)FreeList()  {
	list.Head = nil
	list.Count = 0
}

//头部插入
func (list *SingleList) HeadInsert(item interface{})  {
	node := &Node{Item:item}
	if list.Head == nil {
		 list.Head = node
	} else {
		 node.Next = list.Head
		 list.Head = node
	}
	list.Count++
}

//中间插入
func (list *SingleList) MidInsert(item interface{},index int)  {
	node := &Node{Item:item}
	sp := list.Head
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
	list.Count++
}

//尾部插入
func (list *SingleList) TailInsert(item interface{})  {
	node := &Node{Item:item}
	if list.Head == nil {
		list.Head = node
	} else {
		sp := list.Head
		for sp.Next != nil {
			sp = sp.Next
		}
		sp.Next = node
	}
	list.Count++
}

//头部删除元素
func (list *SingleList) HeadDelete() interface{} {
	if list.Head == nil {
		return nil
	}
	item := list.Head.Item
	list.Head = list.Head.Next
	list.Count--
	return item
}

//尾部删除元素
func (list *SingleList) TailDelete() interface{} {
	if list.Head == nil {
		return nil
	}
	sp := list.Head
	prev := list.Head
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
	if list.Head == nil {
		return nil
	}
	sp := list.Head
	for i:=0;i<index-1;i++ {
		sp = sp.Next
	}
	item := sp.Next.Item
	sp.Next = sp.Next.Next
	return item
}

//删除一个目标元素
func (list *SingleList)DeleteNode(delValue interface{}) bool {
	if list.Count == 0 {
		return false
	}
	//如果头元素刚好是key，则删除
	if list.Head != nil && list.Head.Item == delValue {
		list.Head = list.Head.Next
		list.Count--
		return true
	}
	curr := list.Head
	//由于上面的处理，第一个curr确定不是key了，遍历下一个
	for curr.Next != nil {
		if curr.Next.Item == delValue {
			curr.Next = curr.Next.Next
			list.Count--
			return true
		} else {
			curr = curr.Next
		}
	}
	return false
}

//删除全部目标元素
func (list *SingleList)DeleteNodes(delValue interface{}) bool {
	if list.Count == 0 {
		return false
	}
	//如果头元素刚好是key，则先删除，如果前面几个都是key，循环删除
	for list.Head != nil && list.Head.Item == delValue {
		list.Head = list.Head.Next
		list.Count--
	}
	var isDel bool = false
	curr := list.Head
	//由于上面的处理，第一个curr确定不是key了，遍历下一个
	for curr.Next != nil {
		if curr.Next.Item == delValue {
			curr.Next = curr.Next.Next
			list.Count--
			isDel = true
		} else {
			curr = curr.Next
		}
	}
	return isDel
}

//有序插入
func (list *SingleList) SortedInsert(item interface{})  {
	newNode := &Node{Item:item}
	curr := list.Head
	if curr == nil || list.LessFunc(item,curr.Item) {
		newNode.Next = list.Head
		list.Head = newNode
		list.Count++
		return
	}
	for curr.Next != nil && list.LessFunc(curr.Next.Item,item) {
		curr = curr.Next
	}
	newNode.Next = curr.Next
	curr.Next = newNode
	list.Count++
}

//删除重复元素（以有序链表为基础）
func (list *SingleList)RemoveDuplicate()  {
	curr := list.Head
	for curr != nil {
		if curr.Next != nil && curr.Item == curr.Next.Item {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
}

//查找元素
func (list *SingleList)IsPresent(key interface{}) bool {
	curr := list.Head
	if curr == nil {
		return false
	}
	for curr != nil {
		if curr.Item == key {
			return true
		}
		curr = curr.Next
	}
	return false
}

//反转链表
func (list *SingleList) ReverseSingleList() {
	curr := list.Head
	var prev,next *Node
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	list.Head = prev
}

//反转链表（递归）
func (list *SingleList)ReverseRecureSingleList()  {
	list.Head = list.ReverseRecure(list.Head)
}
//反转链表（递归）
func (list *SingleList)ReverseRecure(h *Node) *Node {
	//如果h为空直接返回
	//如果h没有后置节点了，返回它本身,此时末尾元素成了递归栈的最上面一个，从它开始反向每一个节点
	if h == nil || h.Next == nil {
		return h
	}
	var newNode = list.ReverseRecure(h.Next)
	h.Next.Next = h //反转每个节点的指向,注意是从倒数第二个开始反转的
	h.Next = nil
	return newNode
}

//合并链表
func (list *SingleList)MergeList(list1 *SingleList) {
	if list1.Head == nil {
		return
	}
	if list.Head == nil {
		list.Head = list1.Head
		return
	}
	sp := list.Head
	for sp.Next != nil {
		sp = sp.Next
	}
	sp.Next = list1.Head
	list.Count = list.Count+list1.Count
}

//有序合并链表
func MergeSortList(list,list1 *SingleList) *SingleList{
	if list == nil {
		return list1
	}
	if list1 == nil {
		return list
	}
	newList := NewSingleList(list.LessFunc)
	newHead := list.MergeReverseSortList(list.Head,list1.Head)
	newList.Head = newHead
	newList.Count = list.Count+list1.Count
	return newList
}
//递归合并有序链表
func (list *SingleList)MergeReverseSortList(l,l1 *Node) *Node {
	if l == nil {
		return l1
	}
	if l1 == nil {
		return l
	}
	var res *Node
	if list.LessFunc(l.Item,l1.Item) {
		res = l
		res.Next = list.MergeReverseSortList(l.Next,l1)
	} else {
		res = l1
		res.Next = list.MergeReverseSortList(l,l1.Next)
	}
	return res
}

//复制链表
func (list *SingleList)CopySingleListReverse() *SingleList{
	curr := list.Head
	var t1,t2 *Node
	for curr != nil {
		t1 = &Node{Item:curr.Item,Next:t2}
		t2 = t1
		curr = curr.Next
	}
	ll2 := NewSingleList(list.LessFunc)
	ll2.Head = t1
	ll2.Count = list.Count
	return ll2
}

//判断是否链表是否存在环
func (list *SingleList) LoopDetect() bool {
	if list.Count <= 2 {
		return false
	}
	slowPrt := list.Head
	fastPtr := list.Head
	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		slowPrt = slowPrt.Next
		fastPtr = fastPtr.Next.Next
		if slowPrt == fastPtr {
			return true
		}
	}
	return false
}

//判断环类型 0无环 1圆环 2普通环
func (list *SingleList) LoopTypeDetect() int {
	if !list.LoopDetect() {
		return 0
	}
	slowPrt := list.Head
	fastPtr := list.Head
	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		if list.Head == fastPtr.Next || list.Head == fastPtr.Next.Next {
			return 1
		}
		slowPrt = slowPrt.Next
		fastPtr = fastPtr.Next.Next
		if slowPrt == fastPtr {
			return 2
		}
	}
	return 0
}

//判断是否是圆环
func (list *SingleList)IsCircleList() bool {
	if list.Count <= 2 {
		return false
	}
	isCircle := false
	tempHead := list.Head
	list.ReverseSingleList()
	if tempHead == list.Head {
		isCircle = true
	}
	list.ReverseSingleList()
	return isCircle
}

//查找两条链表的交点，不相交则返回nil
func FindIntersection(s,s2 *SingleList) *Node{
	l := s.Count
	head := s.Head
	l2 := s2.Count
	head2 := s2.Head
	var diff int
	if l < l2 {
		//保证head是长的那个
		temp := head
		head = head2
		head2 = temp
		diff = l2 -l
	} else {
		diff = l - l2
	}
	for ;diff>0;diff-- {
		head = head.Next
	}
	for head != head2 {
		if head == nil || head2 == nil {
			return nil
		}
		head = head.Next
		head2 = head2.Next
	}
	return head
}

//给定一个节点，输出到结尾
func PrintNodes(sp *Node)  {
	for sp != nil {
		fmt.Print(sp.Item," ")
		sp = sp.Next
	}
	fmt.Println()
}