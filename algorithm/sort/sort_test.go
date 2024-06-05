package sort

import "testing"

func TestSelection(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2}
	Selection(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("Selection failed: %v", nums)
		}
	}
}

func TestBubble(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2}
	BubbleSort(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("Bubble failed: %v", nums)
		}
	}
}
