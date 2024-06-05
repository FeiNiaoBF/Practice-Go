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

func TestQuickSort(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2}
	QuickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("QuickSort failed: %v", nums)
		}
	}
}

func TestMergeSort(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2}
	nums = MergeSort(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("MergeSort failed: %v", nums)
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
	74, 27, 57, 11, 63, 45, 96, 32,
	31, 54, 14, 11, 11, 24, 11, 90,
	31, 14, 92, 67, 38, 25, 81, 59,
	46, 13, 85, 21, 68, 39, 56, 11,
	74, 29, 63, 48, 95, 17, 82, 31,
	60, 43, 89, 22, 58, 35, 77, 26,
	99, 18, 64, 41, 72, 53, 97, 15,
	84, 30, 62, 49, 88, 23, 55, 37,
	91, 19, 66, 44, 78, 27, 61, 50,
	93, 20, 1, 36, 94, 16, 65, 51,
	83, 28, 5, 40, 76, 32, 90, 24,
	59, 17, 71, 45, 98, 13, 86, 33,
	69, 52, 96, 62, 17, 93, 29, 56,
	34, 91, 3, 46, 28, 19, 85, 62,
	75, 22, 41, 67, 31, 98, 56, 11,
	82, 59, 38, 4, 26, 17, 63, 49,
	72, 88, 35, 44, 61, 29, 93, 15,
	58, 69, 23, 51, 39, 87, 65, 36,
	84, 30, 62, 49, 88, 23, 55, 37,
	91, 19, 66, 4, 78, 27, 61, 50,
	93, 20, 79, 36, 94, 16, 65, 51,
	83, 28, 58, 40, 76, 32, 90, 24,
	59, 17, 71, 45, 98, 13, 86, 33,
	69, 52, 96, 62, 17, 93, 29, 56,
	34, 91, 13, 46, 2, 19, 85, 62,
	75, 22, 41, 67, 31, 98, 56, 11,
	82, 59, 38, 94, 26, 17, 63, 49,
	72, 88, 35, 44, 1, 29, 93, 15,
	58, 69, 23, 51, 39, 87, 65, 36,
	95, 21, 43, 70, 52, 97, 14, 60,
	33, 89, 25, 53, 78, 30, 54, 16,
	68, 42, 76, 0, 24, 55, 37, 92,
	18, 64, 71, 48, 83, 20, 56, 40,
	74, 27, 57, 11, 63, 45, 96, 32,
	31, 54, 14, 11, 11, 24, 11, 90,
	31, 14, 92, 67, 38, 25, 81, 59,
	46, 13, 85, 21, 68, 39, 56, 11,
	74, 29, 63, 48, 95, 17, 82, 31,
	60, 43, 89, 22, 58, 35, 77, 26,
	99, 18, 64, 41, 72, 53, 97, 15,
	84, 30, 62, 49, 88, 23, 55, 37,
	91, 19, 66, 44, 78, 27, 61, 50,
	93, 20, 79, 36, 94, 16, 65, 51,
	83, 28, 58, 40, 76, 32, 90, 24,
	59, 17, 71, 45, 98, 13, 86, 33,
	69, 52, 96, 62, 17, 93, 29, 56,
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

func BenchmarkQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(NUMS, 0, len(NUMS)-1)
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(NUMS)
	}
}
