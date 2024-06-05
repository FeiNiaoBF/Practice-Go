package sort

func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	pivot := partition(nums, left, right)
	QuickSort(nums, left, pivot)
	QuickSort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	leftInd := left
	// partition nums with pivot
	for i := left + 1; i <= right; i++ {
		if nums[i] < pivot {
			nums[i], nums[leftInd] = nums[leftInd], nums[i]
			leftInd++
		}
	}
	// swap pivot
	nums[leftInd] = pivot
	return leftInd
}
