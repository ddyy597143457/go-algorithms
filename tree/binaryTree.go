package tree

import (
	"ddyy/go-algorithms/queue"
	stack2 "ddyy/go-algorithms/stack"
	"fmt"
	"math"
)

type treeNode struct {
	item interface{}
	left,right *treeNode
}

type BinaryTree struct {
	root *treeNode
}

func LevelCreateBinaryTree(arr []int) *BinaryTree {
	tree := new(BinaryTree)
	tree.root = levelCreateBinaryTree(arr,0,len(arr))
	return tree
}

func levelCreateBinaryTree(arr []int,start int, size int) *treeNode {
	curr := &treeNode{item:arr[start]}
	//这里为什么还要加1呢，因为是数组下标从0开始，而书面上的公式都是按1-n编码算的
	left := start*2 + 1
	right := start*2+1 + 1
	if left < size {
		curr.left = levelCreateBinaryTree(arr,left,size)
	}
	if right < size {
		curr.right = levelCreateBinaryTree(arr,right,size)
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

//输出最长路径
func (tree *BinaryTree)PrintMaxPath()  {
	stack := new(stack2.Stack)
	depth := tree.TreeDepth()
	printMaxPath(tree.root,stack,depth)
}
func printMaxPath(n *treeNode,stack *stack2.Stack,depth int) {
	if n == nil {
		return
	}
	stack.Push(n.item)
	if n.left == nil && n.right == nil {
		if stack.Depth == depth {
			stack.Print()
			fmt.Println()
		}
		stack.Pop()
		return
	}
	printMaxPath(n.left,stack,depth)
	printMaxPath(n.right,stack,depth)
	stack.Pop()
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
	stack.Pop()
}


//最近公共祖先
func (tree *BinaryTree)LCA(first,second int) *treeNode{
	return lca(tree.root,first,second)

}
func lca(curr *treeNode,first,second int) *treeNode{
	if curr == nil {
		return nil
	}
	if curr.item == first || curr.item == second {
		return curr
	}
	left := lca(curr.left,first,second)
	right := lca(curr.right,first,second)
	if left != nil && right != nil {
		return curr
	}else if left != nil {
		return left
	} else {
		return right
	}
}

//找出最大值
func (tree *BinaryTree)FindMaxBT() int {
	return findMaxBT(tree.root)
}
func findMaxBT(curr *treeNode) int {
	if curr == nil {
		return math.MinInt64
	}
	max := curr.item.(int)
	leftmax := findMaxBT(curr.left)
	rightmax := findMaxBT(curr.right)
	if leftmax > max {
		max = leftmax
	}
	if rightmax > max {
		max = rightmax
	}
	return max
}

//查找值
func (tree *BinaryTree)SearchBT(searchValue int) bool {
	return searchBT(tree.root,searchValue)
}
func searchBT(curr *treeNode,searchValue int) bool {
	if curr == nil {
		return false
	}
	if curr.item == searchValue {
		return true
	}
	left := searchBT(curr.left,searchValue)
	right := searchBT(curr.right,searchValue)
	return left || right
}

//构造二叉查找树（arr需要有序）
func CreateBinarySearchTree(arr []int) *BinaryTree {
	tree := new(BinaryTree)
	tree.root = createBinarySearchTree(arr,0,len(arr)-1)
	return tree
}

func createBinarySearchTree(arr []int,start,end int) *treeNode {
	if start > end {
		return nil
	}
	mid := (start+end)/2
	node := new(treeNode)
	node.item = arr[mid]
	node.left = createBinarySearchTree(arr,start,mid-1)
	node.right = createBinarySearchTree(arr,mid+1,end)
	return node
}

//二叉查找树的插入
func (tree *BinaryTree)BSTadd(item int)  {
	bstadd(tree.root,item)
}
func bstadd(curr *treeNode,item int) *treeNode{
	if curr == nil {
		curr = &treeNode{item:item}
		return curr
	}
	if curr.item.(int) > item {
		curr.left = bstadd(curr.left,item)
	} else if curr.item.(int) < item {
		curr.right = bstadd(curr.right,item)
	}
	return curr
}

//最大值
func FindMax(t *treeNode) (interface{},bool) {
	if t == nil {
		return nil,false
	}
	for t.right != nil {
		t = t.right
	}
	return t.item,true
}
//最小值
func FindMin(t *treeNode) (interface{},bool) {
	if t == nil {
		return nil,false
	}
	for t.left != nil {
		t = t.left
	}
	return t.item,true
}

//判断树是否是二叉查找树
func (tree *BinaryTree)IsBST() bool {
	return isBST(tree.root)
}
func isBST(curr *treeNode) bool {
	if curr == nil {
		return true
	}
	if curr.left != nil {
		max,_ := FindMax(curr.left)
		if max.(int) > curr.item.(int) {
			return false
		}
	}
	if curr.right != nil {
		min,_ := FindMin(curr.right)
		if min.(int) < curr.item.(int) {
			return false
		}
	}
	return isBST(curr.left) && isBST(curr.right)
}

func (tree *BinaryTree)DeleteNode(value int)  {
	tree.root = deleteNode(tree.root,value)
}
func deleteNode(curr *treeNode,value int) *treeNode {
	 if curr == nil {
	 	return nil
	 }
	 if curr.item == value {
	 	//目标删除节点是当前节点
	 	if curr.left == nil && curr.right == nil {
	 		//如果当前节点没有子节点了，删除自己，返回nil
	 		return nil
		}
	 	if curr.left == nil {
	 		//如果左子树是nil，右子树直接顶上
	 		return curr.right

		} else if curr.right == nil {
			//如果右子树是nil，左子树直接顶上
			return curr.left
		}
	 	lefttreedepth := treeDepth(curr.left)
	 	righttreedepth := treeDepth(curr.right)
	 	var t interface{}
	 	if lefttreedepth > righttreedepth {
	 		//如果左子树比较长，则从左子树选出一个最大的节点代替当前节点，然后删除那个最大节点
			t,_ = FindMax(curr.left)
			curr.item = t
			curr.left = deleteNode(curr.left,t.(int))
		} else {
			//如果右子树比较长，则从右子树选出一个最小的节点代替当前节点，然后删除那个最小节点
			t,_  = FindMin(curr.right)
			curr.item = t
			curr.right = deleteNode(curr.right,t.(int))
		}
	 } else if curr.item.(int) > value{
	 	//目标删除节点在左子树
	 	curr.left = deleteNode(curr.left,value)
	 } else {
	 	//目标删除节点在右子树
	 	curr.right = deleteNode(curr.right,value)
	 }
	 return curr
}
