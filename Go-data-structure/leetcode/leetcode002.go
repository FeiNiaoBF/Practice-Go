package leetcode


func removeElement(nums []int, val int) int {
	l := len(nums)
	i := 0
	for i < l {
		if nums[i] == val {
			// 重组数组
			nums = append(nums[:i], nums[i+1:]...)
			l--
		} else {
			i++
		}
	}
	// fmt.Println(nums)
	return len(nums)  // return the length of the new array
}
