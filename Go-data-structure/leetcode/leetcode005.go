package leetcode

func majorityElement(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	// // using map to store the nums
	// m := make(map[int]int)
	// for _, val := range nums {
	// 	m[val]++
	// }
	// // 大于 ⌊ n/2 ⌋ 的元素
	// for k, v := range m {
	// 	if v > l/2 {
	// 		return k
	// 	}s
	// }
	// return 0
	count := 0
	num := 0
	for _, v := range nums {
		// 若count为0，则将num初始化为当前元素
		if count == 0 {
			num, count = v, 1
		} else {
			if num == v {
				count++
			} else {
				count--
			}
		}
	}
	return num

}
