package sort

// time complexity: O(nlogn)
// space complexity: O(n)
func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	for left < right {
		pivot := partition(nums, left, right)
		if pivot-left < right-pivot {
			QuickSort(nums, left, pivot)
			left = pivot + 1
		} else {
			QuickSort(nums, pivot+1, right)
			right = pivot - 1
		}
	}
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
