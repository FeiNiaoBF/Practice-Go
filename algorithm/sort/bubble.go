package sort

// time complexity: O(n^2)
// space complexity: O(1)
func BubbleSort(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if nums[j] > nums[j+1] {
				// swap
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
