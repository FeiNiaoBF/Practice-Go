package sort

func ConutSort(nums []int) {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	// 申请一个数组，下标为0-max，值为出现的次数
	// 空间不好说
	// 当我有10个，{9，8，7，100...}
	// 申请一个数组，长度为100，其他都是0
	conut := make([]int, max+1)
	for _, num := range nums {
		conut[num]++
	}

	for i, num := 0, 0; i<max+1; i++ {
		for j := 0; j < conut[i]; j++ {
			nums[num] = i
			num++
		}
	}
}
