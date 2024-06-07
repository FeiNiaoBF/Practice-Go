package sort

func BucketSort(nums []int) {
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

	bucket := make([]int, max-min+1)
	for i := 0; i < len(nums); i++ {
		bucket[nums[i]-min]++
	}

	index := 0
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			nums[index] = i + min
			index++
			bucket[i]--
		}
	}
}
