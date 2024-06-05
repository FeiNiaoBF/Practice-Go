package sort

// time complexity: O(n^2)
// space complexity: O(1)
func Selection(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}

	for i := 0; i < size; i++ {
		minI := i
		for j := i + 1; j < size; j++ {
			if nums[j] < nums[minI] {
				minI = j
			}
		}
		// swap
		nums[i], nums[minI] = nums[minI], nums[i]
	}
}
