package sort

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelection(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2, 2, 3}
	Selection(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("Selection failed: %v", nums)
		}
	}
	require.Equal(t, []int{1, 2, 2, 3, 3, 4, 5}, nums)
}

func TestBubble(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2, 2, 3}
	BubbleSort(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("Bubble failed: %v", nums)
		}
	}
	require.Equal(t, []int{1, 2, 2, 3, 3, 4, 5}, nums)
}

func TestInsertion(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2, 2, 3}
	Insertion(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("Insertion failed: %v", nums)
		}
	}
	require.Equal(t, []int{1, 2, 2, 3, 3, 4, 5}, nums)
}

func TestQuickSort(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2, 2, 3}
	QuickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("QuickSort failed: %v", nums)
		}
	}
	require.Equal(t, []int{1, 2, 2, 3, 3, 4, 5}, nums)
}

func TestMergeSort(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2, 2, 3}
	MergeSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("MergeSort failed: %v", nums)
		}
	}
	require.Equal(t, []int{1, 2, 2, 3, 3, 4, 5}, nums)
}

func TestHeapSort(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2, 2, 3}
	HeapSort(&nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("HeapSort failed: %v", nums)
		}
	}
	require.Equal(t, []int{1, 2, 2, 3, 3, 4, 5}, nums)
}

func TestBucketSort(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2, 2, 3}
	BucketSort(nums, 2)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("BucketSort failed: %v", nums)
		}
	}
	require.Equal(t, []int{1, 2, 2, 3, 3, 4, 5}, nums)
}

func TestConutSort(t *testing.T) {
	nums := []int{5, 3, 4, 1, 2, 2, 3}
	CountSort(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Errorf("ConutSort failed: %v", nums)
		}
	}
	require.Equal(t, []int{1, 2, 2, 3, 3, 4, 5}, nums)
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
		QuickSort(NUMS, 0, 9999)
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(NUMS, 0, 9999)
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
		CountSort(NUMS)
	}
}
