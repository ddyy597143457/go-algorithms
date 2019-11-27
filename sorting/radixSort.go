package sorting

func RadixSort(arr []int)  {
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
		//将所有桶中记录依次收集到tmp中，注意顺序是从高到低的，高位的数大于低位的
		for j:=length-1; j>=0; j-- {
			k := (arr[j]/radix)%10
			temp[count[k]-1] = arr[j]
			count[k]--
		}
		for j:=0; j<length; j++ {
			arr[j] = temp[j]
		}
		radix *= 10
	}
}