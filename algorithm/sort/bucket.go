package sort

func BucketSort(nums []int) {
	if len(nums) == 0 {
		return
	}

	max, min := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

 // 计算桶的数量
 k := max - min + 1

 // 初始化桶
 bucket := make([][]int, k)
 for i := range bucket {
	 bucket[i] = make([]int, 0)
 }

 // 将元素分配到桶中
 for _, num := range nums {
	 index := num - min
	 bucket[index] = append(bucket[index], num)
 }

 // 对每个桶内的元素进行排序
 for _, b := range bucket {
	 Insertion(b)
 }

 // 合并桶中的元素
 index := 0
 for _, b := range bucket {
	 for _, num := range b {
		 nums[index] = num
		 index++
	 }
 }
}
