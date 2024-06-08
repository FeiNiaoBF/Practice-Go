package sort

func CountSort(nums []int) {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	counter := make([]int, max+1)
	for _, num := range nums {
		counter[num]++
	}

	// prefix
	for i := 0; i < max; i++ {
		counter[i+1] += counter[i]
	}
	// sort
	n := len(nums)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		res[counter[nums[i]]-1] = nums[i]
		counter[nums[i]]--
	}

	copy(nums, res)
}
