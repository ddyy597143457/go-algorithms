/**
二叉树
 */
package tree

import "fmt"

//节点值比较函数
type Comparable func(c1 interface{},c2 interface{}) bool

type BinaryTree struct {
	Node interface{}
	Left *BinaryTree
	Right *BinaryTree
	LessFunc Comparable
}

func New(comparafunc Comparable) *BinaryTree{
	return &BinaryTree{
		LessFunc:comparafunc,
	}
}

//插入元素
func (binarytree *BinaryTree) Insert(item interface{})  {
	if binarytree.Node == nil {
		binarytree.Node = item
		binarytree.Left = New(binarytree.LessFunc)
		binarytree.Right = New(binarytree.LessFunc)
		return
	}
	if binarytree.LessFunc(item,binarytree.Node) {
		binarytree.Left.Insert(item)
	} else {
		binarytree.Right.Insert(item)
	}
}

//查找元素
func (binarytree *BinaryTree) Search(item interface{}) *BinaryTree{
	if binarytree.Node == nil {
		return nil
	}
	if binarytree.Node == item {
		return binarytree
	}
	if binarytree.LessFunc(item,binarytree.Node) {
		return binarytree.Left.Search(item)
	} else {
		return binarytree.Right.Search(item)
	}
}

//二叉树最大值，也就是最右子节点值
func (binarytree *BinaryTree) Max() interface{} {
	if binarytree.Node == nil || binarytree.Right.Node == nil{
		return binarytree.Node
	}
	return binarytree.Right.Max()
}

//二叉树最小值，也就是最左子节点
func (binarytree *BinaryTree) Min() interface{} {
	if binarytree.Node == nil || binarytree.Left.Node == nil{
		return binarytree.Node
	}
	return binarytree.Left.Min()
}

//前序遍历
func (binarytree *BinaryTree) PreTravel()  {
	if binarytree.Node != nil {
		fmt.Println(binarytree.Node)
		binarytree.Left.PreTravel()
		binarytree.Right.PreTravel()
	}
}

//中序遍历
func (binarytree *BinaryTree) MidTravel()  {
	if binarytree.Node != nil {
		binarytree.Left.MidTravel()
		fmt.Println(binarytree.Node)
		binarytree.Right.MidTravel()
	}
}

//后序遍历
func (binarytree *BinaryTree) PostOrderTravel()  {
	if binarytree.Node != nil {
		binarytree.Left.PostOrderTravel()
		binarytree.Right.PostOrderTravel()
		fmt.Println(binarytree.Node)
	}
}

//层次遍历
func (binarytree *BinaryTree) LevelTravel() {
	queue := make([]*BinaryTree,0)
	queue = append(queue,binarytree)
	for len(queue) > 0 {
		item := queue[0]
		fmt.Println(item.Node)
		queue = queue[1:]
		if item.Left.Node != nil {
			queue = append(queue,item.Left)
		}
		if item.Right.Node != nil {
			queue = append(queue,item.Right)
		}
	}
}
