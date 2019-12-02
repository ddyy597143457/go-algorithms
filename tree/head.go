package tree

import (
	"fmt"
)
//堆是完全二叉树

//完全二叉树的性质
//1) 具有n个节点的完全二叉树的深度为[log2n]+1，其中[log2n]是向下取整。
//2）若对含 n 个结点的完全二叉树从上到下且从左至右进行 1 至 n 的编号，则任意一个编号为 i 的结点有如下特性：
//(1) 若 i=1，则该结点是二叉树的根，无双亲, 否则，编号为 [i/2] 的结点为其双亲结点;
//(2) 若 2i>n，则该结点无左孩子， 否则，编号为 2i 的结点为其左孩子结点；
//(3) 若 2i+1>n，则该结点无右孩子结点， 否则，编号为2i+1 的结点为其右孩子结点。
type Heap struct {
	arr []int
	size int
}

//构建小顶堆
func NewHeap(arrInput []int) *Heap {
	length := len(arrInput)
	//契合公式,0号位置不用
	arr := []int{0}
	arr = append(arr,arrInput...)
	head := &Heap{arr:arr,size:length}
	//从第一个父节点开始（0号位置不用，第一个父节点是length/2）
	for i:=length/2;i>0;i-- {
		head.ProclateDown(i)
	}
	return head
}

//向下调整小顶堆
func (h *Heap)ProclateDown(parent int)  {
	lchild := parent*2
	rchild := lchild+1
	if lchild > h.size {
		return
	}
	small := lchild
	if rchild <= h.size && h.arr[rchild] < h.arr[small] {
		small = rchild
	}
	//如果父节点大于子节点
	if h.arr[parent] > h.arr[small] {
		//交换
		h.swap(parent,small)
		//继续向下比较
		h.ProclateDown(small)
	}
}

//向上调整小顶堆
func (h *Heap)ProclateUp(child int)  {
	parant := child/2
	if parant == 0 {
		//0号是不用的，所以只能是1，1已经是堆顶了
		return
	}
	if h.arr[parant] > h.arr[child] {
		h.swap(parant,child)
		h.ProclateUp(parant)
	}
}

func (h *Heap)Print()  {
	fmt.Println(h.arr[1:])
}

//添加元素
func (h *Heap)Add(value int)  {
	//先加到结尾
	h.arr = append(h.arr,value)
	h.size++
	//调整堆
	h.ProclateUp(h.size)
}

func (h *Heap)Remove() int {
	//堆顶元素出列
	value := h.arr[1]
	//末尾元素复制给堆顶
	h.arr[1] = h.arr[h.size]
	//末尾元素复制给堆顶后它这个位置不需要参与调整了
	h.size--
	//从栈顶开始调整
	h.ProclateDown(1)
	//调整完成后末尾那个元素彻底离开arr，剩下的数组复制给arr
	//这里为什么需要size+1,因为上面调整前已经size--了，而且[0:7]实际上是arr[0]...arr[6]
	h.arr = h.arr[0:(h.size+1)]
	return value
}

func HeapSort(arr []int) {
	t := NewHeap(arr)
	length := len(arr)
	for i:=0; i<length; i++ {
		arr[i] = t.Remove()
	}
}

func (h *Heap)lessFunc(x,y int) bool {
	return h.arr[x] < h.arr[y]
}


func (h *Heap)swap(x,y int)  {
	h.arr[x],h.arr[y] = h.arr[y],h.arr[x]
}