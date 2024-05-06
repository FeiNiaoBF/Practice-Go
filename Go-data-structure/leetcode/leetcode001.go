package leetcode

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 1. merge nums2 into nums1
	// 2. user sort & slice
	sli := append(nums1[:m], nums2[:n]...)
	sort.Ints(sli)
	copy(nums1, sli)

	// 1.sort

}
