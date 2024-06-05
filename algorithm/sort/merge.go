package sort

// 归并排序
func MergeSort(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}
	// 分
	mid := n / 2
	arrLeft := nums[:mid]
	arrRight := nums[mid:]

	// 并
	return merge(MergeSort(arrLeft), MergeSort(arrRight))
}

func merge(left, right []int) (arrTmp []int) {
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			arrTmp = append(arrTmp, left[0])
			left = left[1:]
		} else {
			arrTmp = append(arrTmp, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		arrTmp = append(arrTmp, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		arrTmp = append(arrTmp, right[0])
		right = right[1:]
	}

	return
}
