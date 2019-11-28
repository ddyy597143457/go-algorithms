package sorting

//基数排序
//原理：从个位开始，十位，百位一直到最高位对每个元素排序
func RadixSort(arr []int) {
	radix  := 1
	length := len(arr)
	max := MaxInt(arr)
	temp := make([]int,length)
	var count [10]int
	for max/radix > 0{
		//count每次初始化为0
		for j:=0; j<10; j++ {
			count[j] = 0
		}
		//当前位数下每个数值的数量
		for j:=0; j<length; j++ {
			k := (arr[j]/radix)%10
			count[k]++
		}
		//前缀和,理解这里是最关键的，知道count[9]是多少吗?count[9]=length
		for j:=1; j<10; j++ {
			count[j] += count[j-1]
		}
		//将所有桶中记录依次收集到tmp中,这里要从length-1开始，因为高位排序要以低一位的排序结果为基础去排序
		for j:=length-1; j>=0; j-- {
			k := (arr[j]/radix)%10	//算出位数
			temp[count[k]-1] = arr[j]	//结合前缀和，根据位数把数据放进temp
			count[k]--
		}
		//temp重新放回arr
		for j:=0; j<length; j++ {
			arr[j] = temp[j]
		}
		//位数*10
		radix *= 10
	}
}