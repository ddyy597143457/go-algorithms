package sorting

import (
	"math"
)

//桶排序，计数排序的升级版
//原理：将数组分成n个桶，遍历数组将每个元素映射到对应的桶,每个桶再执行计数排序，然后合并
func BucketSort(arr []int,bucketSize int){
	length := len(arr)
	min := MinInt(arr)
	max := MaxInt(arr)
	bucketCount := (int)(math.Floor(float64((max-min)/bucketSize)))+1
	buckets := make([][]int,bucketCount)
	for i:=0; i<length; i++ {
		index := (int)(math.Floor(float64((arr[i]-min)/bucketSize)))
		buckets[index] = append(buckets[index],arr[i])
	}
	startIndex := 0
	for i:=0; i<len(buckets); i++ {
		if len(buckets[i]) > 0 {
			CountingSort(buckets[i])
			for j:=0;j<len(buckets[i]); j++ {
				arr[startIndex] = buckets[i][j]
				startIndex++
			}
		}
	}
}