package sort

// 归并排序
/* 合并左子数组和右子数组 */
func merge(nums []int, left, mid, right int) {
	// 左子数组区间为 [left, mid], 右子数组区间为 [mid+1, right]
	// 创建一个临时数组 tmp ，用于存放合并后的结果
	tmp := make([]int, right-left+1)
	// 初始化左子数组和右子数组的起始索引
	i, j, k := left, mid+1, 0
	// 当左右子数组都还有元素时，进行比较并将较小的元素复制到临时数组中
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	// 将左子数组和右子数组的剩余元素复制到临时数组中
	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}
	// 将临时数组 tmp 中的元素复制回原数组 nums 的对应区间
	copy(nums[left:right+1], tmp)
}

/* 归并排序 */
func MergeSort(nums []int, left, right int) {
	// 终止条件
	if left >= right {
		return
	}
	// 划分阶段
	mid := left + (right-left)/2
	MergeSort(nums, left, mid)
	MergeSort(nums, mid+1, right)
	// 合并阶段
	merge(nums, left, mid, right)
}
