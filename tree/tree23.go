package tree

import (
	"ddyy/go-algorithms/queue"
	"fmt"
)

var debug bool

type Tree23Node struct {
	item [3]int	//item[2]事实上是不允许的，这里用于暂存3节点
	num int	//item的长度
	leftChild,midChild,rightChild  *Tree23Node	//左子树，中子树，右子树
	parent *Tree23Node	//指向父节点
	//左子树已满3节点分裂时，父节点的leftChildtemp指向左子树的右节点
	//右子树已满3节点分裂时，父节点的rightChildtemp指向右子树的左节点
	leftChildtemp,rightChildtemp *Tree23Node
}
type Tree23 struct {
	root *Tree23Node
}

func NewTree23() *Tree23{
	return new(Tree23)
}

//创建一个23树节点
func NewTree23Node(value int) *Tree23Node {
	node23 := new(Tree23Node)
	node23.num = 1
	node23.item[0] = value
	return node23
}
//在一个已存在的未满的23树节点添加值
func NodeAddValue(currNode *Tree23Node,value int) *Tree23Node {
	currNode.item[currNode.num] = value
	currNode.num++
	//值排序
	InsertionSort(currNode.item)
	//排序的item右可能是0,0,4，5，6这样的，4，5，6需要移到索引0，1，2位置
	ResetItem(currNode)
	return currNode
}

func ResetItem(currNode *Tree23Node) {
	var temp [3]int
	i := 0
	for _,v := range currNode.item {
		if v != 0 {
			temp[i] = v
			i++
		}
	}
	currNode.item = temp
}

//插入节点
func (tree *Tree23)Tree23Insert(value int) {
	if tree.root == nil {
		//根节点为空，直接生成新节点插入
		tree.root = NewTree23Node(value)
	} else {
		tree.root = tree23Insert(tree.root,value)
	}
}
//从根开始遍历，直到找到合适的位置插入
func tree23Insert(currNode *Tree23Node,value int) *Tree23Node {
	if currNode == nil {
		currNode = NewTree23Node(value)
		return currNode
	}
	if currNode.num == 1 {
		if value < currNode.item[0] {
			if currNode.leftChild != nil {
				currNode.leftChild = tree23Insert(currNode.leftChild,value)
			} else {
				currNode = NodeAddValue(currNode,value)
			}
		} else if value > currNode.item[0] {
			if currNode.rightChild != nil {
				currNode.rightChild = tree23Insert(currNode.rightChild,value)
			} else {
				currNode = NodeAddValue(currNode,value)
			}
		}
	} else if currNode.num == 2 {
		 if value < currNode.item[0] {
		 	if currNode.leftChild != nil {
				currNode.leftChild = tree23Insert(currNode.leftChild,value)
			} else {
				currNode = NodeAddValue(currNode,value)
			}
		 } else if value > currNode.item[1] {
			 if currNode.rightChild != nil {
				 currNode.rightChild = tree23Insert(currNode.rightChild,value)
			 } else {
				 currNode = NodeAddValue(currNode,value)
			 }
		 } else {
			 if currNode.midChild != nil {
				 currNode.midChild = tree23Insert(currNode.midChild,value)
			 } else {
				 currNode = NodeAddValue(currNode,value)
			 }
		 }
	}
	if currNode.num == 3 {
		//如果当前节点插入新节点后成为了3节点，那么就需要分裂
		currNode = Division(currNode)
	}
	return currNode
}
//节点分裂
func Division(currNode *Tree23Node) *Tree23Node{
	if currNode.parent == nil {
		//如果当前节点的父节点为nil，那么它是根节点
		leftNode := GetDivisionNode(currNode,0)	//当前节点的左节点
		midNode := GetDivisionNode(currNode,1)	//当前节点的中节点
		rightNode := GetDivisionNode(currNode,2)	//当前节点的右节点
		//如果leftChild或者rightChild分裂，重新整合
		makeDivisionNode(currNode,leftNode,rightNode)
		//midNode成为leftNode和rightNode的父节点
		midNode.leftChild = leftNode
		leftNode.parent = midNode
		midNode.rightChild = rightNode
		rightNode.parent = midNode
		return midNode
	}
	//如果父节点是1节点，直接插入，无需再分裂
	//如果父节点是2节点，需要分裂
	if currNode.parent.num == 1 || currNode.parent.num == 2{
		leftNode := GetDivisionNode(currNode,0)	//当前节点的左节点
		rightNode := GetDivisionNode(currNode,2)	//当前节点的右节点
		//如果leftChild或者rightChild处于分裂状态，整合之
		makeDivisionNode(currNode,leftNode,rightNode)
		//父节点的最大值
		parentmax := maxItem(currNode.parent.item)
		var temp *Tree23Node
		//父节点是1节点，无需分裂，midNode上移给父节点
		//并且leftNode或rightNode将成为父节点的midChild
		if currNode.parent.num == 1 {
			//如果当前节点的最小值大于父节点最大值，则当前必定是父节点的右子树
			if currNode.item[0] > parentmax {
				currNode.parent.midChild = leftNode
				temp = rightNode
			} else {
				currNode.parent.midChild = rightNode
				temp = leftNode
			}
		} else if currNode.parent.num == 2{
			//父节点是2节点，节点需要标记分裂
			if currNode.item[0] > parentmax {
				currNode.parent.rightChildtemp = leftNode
				currNode.parent.rightChild = rightNode
				temp = rightNode
			} else {
				currNode.parent.leftChild = leftNode
				currNode.parent.leftChildtemp = rightNode
				temp = leftNode
			}
		}
		leftNode.parent = currNode.parent
		rightNode.parent = currNode.parent
		currNode.parent = NodeAddValue(currNode.parent,currNode.item[1])
		currNode =  temp
	}
	return currNode
}
//如果左子树或者右子树是分裂的,需要把分裂的子树节点重新整合
func makeDivisionNode(currNode,leftNode,rightNode *Tree23Node)  {
	if currNode.leftChildtemp != nil {
		//如果是左子树分裂,
		//leftChild和leftChildtemp成为leftNode的左右子树
		leftNode.leftChild = currNode.leftChild
		leftNode.rightChild = currNode.leftChildtemp
		leftNode.leftChild.parent = leftNode
		leftNode.rightChild.parent = leftNode
		//midChild和rightChild成为rightNode的左右子树
		rightNode.leftChild = currNode.midChild
		rightNode.rightChild = currNode.rightChild
		rightNode.leftChild.parent = rightNode
		rightNode.rightChild.parent = rightNode
	} else if currNode.rightChildtemp != nil {
		//如果是右子树分裂
		//leftChild和midChild成为leftNode的左右子树
		leftNode.leftChild = currNode.leftChild
		leftNode.rightChild = currNode.midChild
		leftNode.leftChild.parent = leftNode
		leftNode.rightChild.parent = leftNode
		//rightChildtemp和rightChild成为rightNode的左右子树
		rightNode.leftChild = currNode.rightChildtemp
		rightNode.rightChild = currNode.rightChild
		rightNode.leftChild.parent = rightNode
		rightNode.rightChild.parent = rightNode
	}
}

//获取当前节点的左中右节点
func GetDivisionNode(currNode *Tree23Node,index int) *Tree23Node {
	node := new(Tree23Node)
	node.num = 1
	node.item[0] = currNode.item[index]
	return node
}

func maxItem(item [3]int) int {
	max := item[0]
	for i:=0; i<len(item); i++ {
		if max < item[i] {
			max = item[i]
		}
	}
	return max
}

func (tree *Tree23) LevelTravel() {
	if tree.root == nil {
		return
	}
	q := queue.New()
	q.EnQueue(tree.root)
	for !q.IsEmpty() {
		it,_ := (q.DeQueue()).(*Tree23Node)
		fmt.Println("item:",it.item,"num:",it.num,"parent",it.parent)
		if it.leftChild != nil {
			q.EnQueue(it.leftChild)
		}
		if it.midChild != nil {
			q.EnQueue(it.midChild)
		}
		if it.rightChild != nil {
			q.EnQueue(it.rightChild)
		}
	}
	fmt.Println()
}

func InsertionSort(arr [3]int)  {
	length := len(arr)
	for i:=1;i<length;i++ {
		t := arr[i]
		j := i-1
		for j >= 0 && t < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = t
	}
}