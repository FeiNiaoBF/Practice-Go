package sort

import (
	"sort"
	"testing"
)

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
	MergeSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("MergeSort failed: %v", nums)
		}
	}
}

func TestHeapSort(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2}
	HeapSort(&nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("HeapSort failed: %v", nums)
		}
	}
}

func TestBucketSort(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2}
	BucketSort(nums, 2)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("HeapSort failed: %v", nums)
		}
	}
}

func TestConutSort(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2}
	ConutSort(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("ConutSort failed: %v", nums)
		}
	}
}

// benchmark

var NUMS []int

func init(){
	NUMS = RandNums(10000)
}

func BenchmarkGOSORT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(NUMS)
	}
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
		MergeSort(NUMS, 0, len(NUMS)-1)
	}
}

func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort(&NUMS)
	}
}

func BenchmarkBucket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BucketSort(NUMS, 5)
	}
}

func BenchmarkConut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConutSort(NUMS)
	}
}
