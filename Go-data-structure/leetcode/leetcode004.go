package leetcode

func removeTwoDuplicates(nums []int) int {
	i := 0
	for _, v := range nums {
		println("i=", i, "v=", v)
		if i < 2 || nums[i-2] != v {
			println("i=", i, "v=", v)
			// println("nums[i-2]=", nums[i-2])
			// println("nums[i]=", nums[i])
			nums[i] = v
			i++
		}
	}
	return i
}
