package sort

func HeapSort(nums *[]int) {
	size := len(*nums)

	// build max heap
	// 不从根节点开始
	for i := size/2 - 1; i >= 0; i-- {
		heapfiy(nums, size, i)
	}
	// sort
	for i := size-1; i >= 0; i-- {
		(*nums)[0], (*nums)[i] = (*nums)[i], (*nums)[0]
		heapfiy(nums, i, 0)
	}
}

func heapfiy(nums *[]int, size int, index int) {
	for {
		largest := index
		left := 2*index + 1
		right := 2*index + 2

		if left < size && (*nums)[left] > (*nums)[largest] {
			largest = left
		}
		if right < size && (*nums)[right] > (*nums)[largest] {
			largest = right
		}
		if largest == index {
			break
		}
		// swap
		(*nums)[index], (*nums)[largest] = (*nums)[largest], (*nums)[index]

		// continue to heapify
		index = largest
	}
}
