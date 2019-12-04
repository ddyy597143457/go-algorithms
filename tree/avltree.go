package tree

import (
	"ddyy/go-algorithms/queue"
	"fmt"
)

//二叉平衡查找树

type Elem interface {}

type AvlNode struct {
	data Elem
	balance int
	lchild,rchild *AvlNode
}

type AvlTree struct {
	root *AvlNode
}

func NewAvlNode(value Elem) *AvlNode {
	return &AvlNode{data:value}
}

func NewAvlTree() *AvlTree {
	return new(AvlTree)
}

func (avl *AvlTree)AvlAdd(value Elem) {
	avl.root = avlAdd(avl.root,value)
}

func avlAdd(curr *AvlNode,value Elem) *AvlNode {
	if curr == nil {
		curr = NewAvlNode(value)
		return curr
	}
	if less(value,curr.data) {
		curr.lchild = avlAdd(curr.lchild,value)
		//如果失去平衡
		if abs(heighet(curr.lchild)-heighet(curr.rchild)) > 1 {
			if less(value,curr.lchild.data) {
				//LL型，顺时针旋转
				curr = left_left_ratation(curr)
			} else {
				//LR型
				curr = left_right_ratation(curr)
			}
		}
	} else {
		curr.rchild = avlAdd(curr.rchild,value)
		//如果失去平衡
		if abs(heighet(curr.rchild)-heighet(curr.lchild)) > 1 {
			if less(value,curr.rchild.data) {
				//RL型
				curr = right_left_ratation(curr)
			} else {
				//RR型，逆时针旋转
				curr = right_right_ratation(curr)
			}
		}
	}
	balance(curr)
	return curr
}

//LL树型，顺时针旋转
func left_left_ratation(curr *AvlNode) *AvlNode {
	var lchild *AvlNode
	lchild = curr.lchild
	curr.lchild = lchild.rchild
	lchild.rchild = curr
	balance(curr)
	balance(lchild)
	return lchild
}
//RR树型，逆时针旋转
func right_right_ratation(curr *AvlNode) *AvlNode {
	var rchild *AvlNode
	rchild = curr.rchild
	curr.rchild = rchild.lchild
	rchild.lchild = curr
	balance(curr)
	balance(rchild)
	return rchild
}
//LR型，先逆时针旋再顺时针旋
func left_right_ratation(curr *AvlNode) *AvlNode {
	curr.lchild = right_right_ratation(curr.rchild)
	return left_left_ratation(curr)
}
//RL型，先顺时针旋再逆时针旋
func right_left_ratation(curr *AvlNode) *AvlNode {
	curr.rchild = left_left_ratation(curr.rchild)
	return right_right_ratation(curr)
}

func balance(a *AvlNode) {
	if a != nil  {
		a.balance = abs(heighet(a.lchild)-heighet(a.rchild))
	}
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func heighet(curr *AvlNode) int {
	if curr == nil {
		return 0
	}
	lh := heighet(curr.lchild)
	rh := heighet(curr.rchild)
	if lh > rh {
		return lh + 1
	}
	return  rh + 1
}

func abs(t int) int {
	if t < 0 {
		return -t
	}
	return t
}

func less(x,y Elem) bool {
	if x.(int) < y.(int) {
		return true
	}
	return false
}

func (tree *AvlTree) LevelTravel() {
	if tree.root == nil {
		return
	}
	q := queue.New()
	q.EnQueue(tree.root)
	for !q.IsEmpty() {
		it,_ := (q.DeQueue()).(*AvlNode)
		fmt.Print(it," ")
		if it.lchild != nil {
			q.EnQueue(it.lchild)
		}
		if it.rchild != nil {
			q.EnQueue(it.rchild)
		}
	}
	fmt.Println()
}