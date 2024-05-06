package leetcode

func removeDuplicates(nums []int) int {
	l := len(nums)
	rel := 0
	if l == 0 {
		return rel
	}
	j := 0
	for i := 1; i < l; i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
