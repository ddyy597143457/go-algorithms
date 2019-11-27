package sorting

import "fmt"

//堆是完全二叉树
//大顶堆：每个结点的值都大于或等于其左右孩子结点的值，即arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]
//对于节点i，父节点(i-1)/2，左子节点2i+1，右子节点2i+2

func HeadSort(arr []int)  {
	length := len(arr)
	//从最後一个父节点开始(切记)
	for i:=length/2-1;i>=0;i-- {
		HeadAdjust(arr,i,length-1)
	}
	//调整大顶堆
	for j:=length-1;j>0;j-- {
		//将堆顶元素(0)与末尾元素(j)进行交换,此时数组末尾是最大的了（类似于冒泡）
		Swap(arr,0,j)
		//余下的进行调整
		HeadAdjust(arr,0,j-1)
	}
}

func HeadAdjust(arr []int,start,length int)  {
	dad := start
	son := dad*2+1	//左son
	fmt.Println(dad,son)
	for son <= length {
		//如果右节点大，那么son取右节点
		if son+1 < length && arr[son] < arr[son+1] {
			son += 1
		}
		//如果父节点大了，说明满足大顶堆了，直接返回
		if arr[dad] > arr[son] {
			return
		} else {
			//父节点和最小的子节点交换
			Swap(arr,dad,son)
			dad = son
			son = dad*2+1
		}
	}
}