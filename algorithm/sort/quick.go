package sort

// time complexity: O(nlogn)
// space complexity: O(n)
func QuickSort(nums []int, left, right int) {
	for left < right {
		pivot := partition(nums, left, right)
		if pivot-left < right-pivot {
			QuickSort(nums, left, pivot-1)
			left = pivot + 1
		} else {
			QuickSort(nums, pivot+1, right)
			right = pivot - 1
		}
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	i := left + 1
	j := right

	for {
		for i <= j && nums[i] < pivot {
			i++
		}
		for i <= j && nums[j] > pivot {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else {
			break
		}
	}

	nums[left], nums[j] = nums[j], nums[left]
	return j
}
