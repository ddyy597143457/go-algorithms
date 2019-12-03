package timer

import (
	"fmt"
	"time"
)

//堆实现优先级队列

type PQueue struct {
	arr []*Pitem
	size int
	isMin bool
}

type Pitem struct {
	time time.Time
	jobs interface{}
	done int
}

func NewPitem(t time.Time,jobs interface{}) *Pitem {
	return &Pitem{
		time:t,
		jobs:jobs,
	}
}

func NewPQueue(arrInput []*Pitem,ismin bool) *PQueue {
	length := len(arrInput)
	//契合公式,0号位置不用
	arr := []*Pitem{new(Pitem)}
	arr = append(arr,arrInput...)
	pq := &PQueue{arr:arr,size:length,isMin:ismin}
	//从第一个父节点开始（0号位置不用，第一个父节点是length/2）
	for i:=length/2;i>0;i-- {
		pq.ProclateDown(i)
	}
	return pq
}

//向下调整
func (pq *PQueue)ProclateDown(parent int)  {
	lchild := parent*2
	rchild := lchild+1
	if lchild > pq.size {
		return
	}
	small := lchild
	if rchild <= pq.size && pq.comparable(small,rchild) {
		small = rchild
	}
	//比较父节点和子节点
	if pq.comparable(parent,small) {
		//交换
		pq.swap(parent,small)
		//继续向下比较
		pq.ProclateDown(small)
	}
}

//向上调整
func (pq *PQueue)ProclateUp(child int)  {
	parent := child/2
	if parent == 0 {
		//0号是不用的，所以只能是1，1已经是堆顶了
		return
	}
	if pq.comparable(parent,child) {
		pq.swap(parent,child)
		pq.ProclateUp(parent)
	}
}

func (pq *PQueue)Remove() *Pitem {
	//堆顶元素出列
	value := pq.arr[1]
	//末尾元素复制给堆顶
	pq.arr[1] = pq.arr[pq.size]
	//末尾元素复制给堆顶后它这个位置不需要参与调整了
	pq.size--
	//从栈顶开始调整
	pq.ProclateDown(1)
	//调整完成后末尾那个元素彻底离开arr，剩下的数组复制给arr
	//这里为什么需要size+1,因为上面调整前已经size--了，而且[0:7]实际上是arr[0]...arr[6]
	pq.arr = pq.arr[0:(pq.size+1)]
	return value
}

//添加元素
func (pq *PQueue)Add(value *Pitem)  {
	//先加到结尾
	pq.arr = append(pq.arr,value)
	pq.size++
	//调整堆
	pq.ProclateUp(pq.size)
}


func (pq *PQueue)comparable(x,y int) bool {
	if pq.isMin {
		return pq.arr[x].time.Unix() > pq.arr[y].time.Unix()
	}
	return pq.arr[x].time.Unix() < pq.arr[y].time.Unix()
}

func (pq *PQueue)swap(x,y int)  {
	pq.arr[x],pq.arr[y] = pq.arr[y],pq.arr[x]
}

func (pq *PQueue)Peek() *Pitem {
	if pq.size < 1 {
		return nil
	}
	return pq.arr[1]
}

func (pq *PQueue)JobStart() *Pitem{
	for {
		item := pq.Peek()
		if item != nil {
			if item.time.Unix() < time.Now().Unix() && item.done == 0 {
				if item.done == 0 {
					fmt.Println(item," Be overdue,do it")
				}
				pq.Remove()
				continue
			}
			if (item.time.Unix() - time.Now().Unix()) <= 1 && item.done == 0 {
				pq.Remove()
				fmt.Println(item,"called do it")
			}
		} else {
			fmt.Println("wait for jobs...")
			time.Sleep(time.Second)
		}
	}
}

func (pq *PQueue)Print()  {
	arr := pq.arr[1:]
	for _,v := range arr  {
		fmt.Print(v," ")
	}
	fmt.Println()
}