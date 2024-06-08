package sort

import "sort"

func BucketSort(nums []int, bucketSize int) {
	if len(nums) == 0 {
		return
	}

	max, min := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	// 计算桶的数量
	bucketCount := (max-min)/bucketSize + 1

	// 初始化桶
	buckets := make([][]int, bucketCount)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}

	// 将元素分配到桶中
	for _, num := range nums {
		index := (num - min) / bucketSize
		buckets[index] = append(buckets[index], num)
	}
	// 对每个桶内的元素进行排序
	for _, bucket := range buckets {
		sort.Ints(bucket)
	}
    index := 0
    for _, bucket := range buckets {
        for _, num := range bucket {
            nums[index] = num
            index++
        }
    }
}
