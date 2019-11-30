package tree

import (
	"ddyy/go-algorithms/queue"
	stack2 "ddyy/go-algorithms/stack"
	"fmt"
)

type treeNode struct {
	item interface{}
	left,right *treeNode
}

type BinaryTree struct {
	root *treeNode
}

func LevelOrderBinaryTree(arr []int) *BinaryTree {
	tree := new(BinaryTree)
	tree.root = levelOrderBinaryTree(arr,0,len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int,start int, size int) *treeNode {
	curr := &treeNode{item:arr[start]}
	//这里为什么还要加1呢，因为是数组下标从0开始，而书面上的公式都是按1-n编码算的
	left := start*2 + 1
	right := start*2+1 + 1
	if left < size {
		curr.left = levelOrderBinaryTree(arr,left,size)
	}
	if right < size {
		curr.right = levelOrderBinaryTree(arr,right,size)
	}
	return curr
}

func (tree *BinaryTree)PreOrder() {
	preOrder(tree.root)
}

func preOrder(node *treeNode)  {
	if node == nil {
		return
	}
	fmt.Print(node.item," ")
	preOrder(node.left)
	preOrder(node.right)
}

func (tree *BinaryTree)MidOrder() {
	midOrder(tree.root)
}

func midOrder(node *treeNode)  {
	if node == nil {
		return
	}
	midOrder(node.left)
	fmt.Print(node.item," ")
	midOrder(node.right)
}

func (tree *BinaryTree)PostOrder() {
	postOrder(tree.root)
}

func postOrder(node *treeNode)  {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Print(node.item," ")
}

//层次遍历（宽度优先）（利用slice实现）
func (tree *BinaryTree) LevelTravel() {
	if tree.root == nil {
		return
	}
	queue := make([]*treeNode,0)
	queue = append(queue,tree.root)
	for len(queue) > 0 {
		it := queue[0]
		fmt.Print(it.item," ")
		queue = queue[1:]
		if it.left != nil {
			queue = append(queue,it.left)
		}
		if it.right != nil {
			queue = append(queue,it.right)
		}
	}
}

//层次遍历（宽度优先）（利用队列实现）
func (tree *BinaryTree) LevelTravel1() {
	if tree.root == nil {
		return
	}
	q := queue.New()
	q.EnQueue(tree.root)
	for !q.IsEmpty() {
		it,_ := (q.DeQueue()).(*treeNode)
		fmt.Print(it.item," ")
		if it.left != nil {
			q.EnQueue(it.left)
		}
		if it.right != nil {
			q.EnQueue(it.right)
		}
	}
}

//计算树的深度
func (tree *BinaryTree)TreeDepth() int {
	return treeDepth(tree.root)
}
func treeDepth(n *treeNode) int {
	if n == nil {
		return 0
	}
	leftDepth := treeDepth(n.left) + 1
	rightDepth := treeDepth(n.right) + 1
	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}

//复制树
func (tree *BinaryTree)CopyTree() *BinaryTree{
	tree2 := new(BinaryTree)
	tree2.root = copyTree(tree.root)
	return tree2
}

func copyTree(n *treeNode) *treeNode {
	if n == nil {
		return nil
	}
	temp := new(treeNode)
	temp.item = n.item
	temp.left = copyTree(n.left)
	temp.right = copyTree(n.right)
	return temp
}

//复制树镜像
func (tree *BinaryTree)CopyMirrorTree() *BinaryTree{
	tree2 := new(BinaryTree)
	tree2.root = copyMirrorTree(tree.root)
	return tree2
}

func copyMirrorTree(n *treeNode) *treeNode {
	if n == nil {
		return nil
	}
	temp := new(treeNode)
	temp.item = n.item
	temp.right = copyTree(n.left)
	temp.left = copyTree(n.right)
	return temp
}

//树的节点数量
func (tree *BinaryTree)NumNodes() int {
	return numNodes(tree.root)
}

func numNodes(n *treeNode) int {
	if n == nil {
		return 0
	}
	return 1+numNodes(n.left)+numNodes(n.right)
}

//树的叶子节点数量
func (tree *BinaryTree)NumLeafNodes() int {
	return numLeafNodes(tree.root)
}

func numLeafNodes(n *treeNode) int {
	if n == nil {
		return 0
	}
	if n.left == nil && n.right == nil {
		return 1
	}
	return numLeafNodes(n.left)+numLeafNodes(n.right)
}

//判断两棵树是否相等
func IsEqual(tree,tree2 *BinaryTree) bool {
	return isEqual(tree.root,tree2.root)
}

func isEqual(n,n2 *treeNode) bool {
	if n == nil && n2 == nil {
		return true
	}else if n == nil || n2 == nil {
		return false
	} else {
		return (n.item == n2.item) && isEqual(n.left,n2.left) && isEqual(n.right,n2.right)
	}
}

//输出所有路径
func (tree *BinaryTree)PrintAllPath()  {
	stack := new(stack2.Stack)
	printAllPath(tree.root,stack)
}
func printAllPath(n *treeNode,stack *stack2.Stack) {
	if n == nil {
		return
	}
	stack.Push(n.item)
	if n.left == nil && n.right == nil {
		stack.Print()
		stack.Pop()
		fmt.Println()
		return
	}
	printAllPath(n.left,stack)
	printAllPath(n.right,stack)
	//这个pop有点讲究
	stack.Pop()
}
