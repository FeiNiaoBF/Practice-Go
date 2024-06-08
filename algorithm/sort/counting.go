package sort

func ConutSort(nums []int) {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	conut := make([]int, max+1)
	for _, num := range nums {
		conut[num]++
	}

	// prefix
	for i := 0; i < max; i++ {
		conut[i+1] += conut[i]
	}

	// sort
    n := len(nums)
    res := make([]int, n)
    for i := n - 1; i >= 0; i-- {
        num := nums[i]
        // 将 num 放置到对应索引处
        res[conut[num]-1] = num
        // 令前缀和自减 1 ，得到下次放置 num 的索引
        conut[num]--
    }
    // 使用结果数组 res 覆盖原数组 nums
    copy(nums, res)
}
