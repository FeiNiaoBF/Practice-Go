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

func TestInsertion(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2}
	Insertion(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("Insertion failed: %v", nums)
		}
	}
}

// benchmark

var NUMS = []int{
	34, 91, 13, 46, 28, 19, 85, 62,
	75, 22, 41, 67, 31, 98, 56, 11,
	82, 59, 38, 94, 26, 17, 63, 49,
	72, 88, 35, 44, 61, 29, 93, 15,
	58, 69, 23, 51, 39, 87, 65, 36,
	95, 21, 43, 70, 52, 97, 14, 60,
	33, 89, 25, 53, 78, 30, 54, 16,
	68, 42, 76, 90, 24, 55, 37, 92,
	18, 64, 71, 48, 83, 20, 56, 40,
	74, 27, 57, 11, 63, 45, 96, 321,
	31, 542, 1234, 11, 1, 234, 131, 31,
}

func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Selection(NUMS)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(NUMS)
	}
}

func BenchmarkInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Insertion(NUMS)
	}
}
