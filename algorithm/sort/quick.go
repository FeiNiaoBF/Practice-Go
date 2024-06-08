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
    // 以 nums[left] 为基准数
    med := medianThree(nums, left, (left+right)/2, right)
    // 将中位数交换至数组最左端
    nums[left], nums[med] = nums[med], nums[left]
    // 以 nums[left] 为基准数
    i, j := left, right
    for i < j {
        for i < j && nums[j] >= nums[left] {
            j-- //从右向左找首个小于基准数的元素
        }
        for i < j && nums[i] <= nums[left] {
            i++ //从左向右找首个大于基准数的元素
        }
        //元素交换
        nums[i], nums[j] = nums[j], nums[i]
    }
    //将基准数交换至两子数组的分界线
    nums[i], nums[left] = nums[left], nums[i]
    return i //返回基准数的索引
}

/* 选取三个候选元素的中位数 */
func medianThree(nums []int, left, mid, right int) int {
    l, m, r := nums[left], nums[mid], nums[right]
    if (l <= m && m <= r) || (r <= m && m <= l) {
        return mid // m 在 l 和 r 之间
    }
    if (m <= l && l <= r) || (r <= l && l <= m) {
        return left // l 在 m 和 r 之间
    }
    return right
}
