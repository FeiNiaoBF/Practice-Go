package sort

// tims complexity: O(n^2)
// space complexity: O(1)
func Insertion(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}
	// lift is sorted, right is unsorted
	for i := 1; i < size; i++ {
		base := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > base {
			// move to right
			// j+1 == i
			nums[j+1] = nums[j]
			j--
		}
		// swap
		nums[j+1] = base
	}
}
