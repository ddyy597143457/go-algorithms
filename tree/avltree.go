package tree

import (
	"ddyy/go-algorithms/queue"
	"fmt"
)

//二叉平衡查找树

type Elem interface {}

type AvlNode struct {
	item Elem
	balance int
	left,right *AvlNode
}

type AvlTree struct {
	root *AvlNode
}

func NewAvlNode(value Elem) *AvlNode {
	return &AvlNode{item:value}
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
	if less(value,curr.item) {
		curr.left = avlAdd(curr.left,value)
		//如果失去平衡
		if abs(heighet(curr.left)-heighet(curr.right)) > 1 {
			if less(value,curr.left.item) {
				//LL型，顺时针旋转
				curr = left_left_ratation(curr)
			} else {
				//LR型
				curr = left_right_ratation(curr)
			}
		}
	} else {
		curr.right = avlAdd(curr.right,value)
		//如果失去平衡
		if abs(heighet(curr.right)-heighet(curr.left)) > 1 {
			if less(value,curr.right.item) {
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

func (avl *AvlTree)AvlRemove(value Elem) {
	avl.root = avlRemove(avl.root,value)
}
func avlRemove(curr *AvlNode,value Elem) *AvlNode {
	if curr == nil {
		return nil
	}
	if less(value,curr.item) {
		//目标元素在左子树
		curr.left = avlRemove(curr.left,value)
		if abs(heighet(curr.left)-heighet(curr.right)) > 1 {
			 //左子树删除了一个，如果失衡肯定是右子树长，属于RR型
			 curr = right_right_ratation(curr)
		}
	} else if less(curr.item,value) {
		//目标元素在右子树
		curr.right = avlRemove(curr.right,value)
		if abs(heighet(curr.left)-heighet(curr.right)) > 1 {
			//右子树删除了一个，如果失衡肯定左右子树长，属于LL型
			curr = left_left_ratation(curr)
		}
	} else {
		//目标元素在此，删除
		if curr.left == nil && curr.right == nil  {
			//左右孩子都没有，直接返回nil
			return nil
		} else if curr.left == nil {
			//左孩子为空，返回右孩子
			return curr.right
		} else if curr.right == nil {
			//右孩子为空，返回左孩子
			return curr.left
		} else {
			//左右孩子都非空
			//从左子树选出一个最大的节点代替当前节点，然后删除那个最大节点
			//或者从右子树选出一个最小的节点代替当前节点，然后删除那个最小节点
			t,_  := AvlFindMin(curr.right)
			curr.item = t
			curr.right = avlRemove(curr.right,t.(int))
		}
	}
	balance(curr)
	return curr
}

//LL树型，顺时针旋转
func left_left_ratation(curr *AvlNode) *AvlNode {
	var left *AvlNode
	left = curr.left
	curr.left = left.right
	left.right = curr
	balance(curr)
	balance(left)
	return left
}
//RR树型，逆时针旋转
func right_right_ratation(curr *AvlNode) *AvlNode {
	var right *AvlNode
	right = curr.right
	curr.right = right.left
	right.left = curr
	balance(curr)
	balance(right)
	return right
}
//LR型，左子树先逆时针旋，curr再顺时针旋
func left_right_ratation(curr *AvlNode) *AvlNode {
	curr.left = right_right_ratation(curr.left)
	return left_left_ratation(curr)
}
//RL型，右子树先顺时针旋，curr再逆时针旋
func right_left_ratation(curr *AvlNode) *AvlNode {
	curr.right = left_left_ratation(curr.right)
	return right_right_ratation(curr)
}

func balance(a *AvlNode) {
	if a != nil  {
		a.balance = abs(heighet(a.left)-heighet(a.right))
	}
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

//最大值
func AvlFindMax(t *AvlNode) (interface{},bool) {
	if t == nil {
		return nil,false
	}
	for t.right != nil {
		t = t.right
	}
	return t.right,true
}
//最小值
func AvlFindMin(t *AvlNode) (interface{},bool) {
	if t == nil {
		return nil,false
	}
	for t.left != nil {
		t = t.left
	}
	return t.item,true
}

func heighet(curr *AvlNode) int {
	if curr == nil {
		return 0
	}
	lh := heighet(curr.left)
	rh := heighet(curr.right)
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
		if it.left != nil {
			q.EnQueue(it.left)
		}
		if it.right != nil {
			q.EnQueue(it.right)
		}
	}
	fmt.Println()
}